package middleware

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/z3orc/dynamic-rpc/internal/database"
	"github.com/z3orc/dynamic-rpc/internal/http/recorder"
	"github.com/z3orc/dynamic-rpc/internal/models"
	"github.com/z3orc/dynamic-rpc/internal/util"
)

func Cache(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		cachedResult, err := getFromDatabase(r)
		if err != nil {
			w.Header().Add("cached", "False")

			c := &recorder.ResponseRecorder{
				ResponseWriter: w,
				StatusCode: http.StatusOK,
				Body: []byte{},
			}
			next.ServeHTTP(c, r)
			go pushToDatabase(c, r)
		} else {
			w.Header().Add("cached", "True")
			util.ReturnJson(w, r, cachedResult)
		}

	})
}

func getFromDatabase(r *http.Request) (models.Version, error){
	values := strings.Split(r.RequestURI, "/")

	identifier := fmt.Sprint(values[1], "-", values[2])

	client := database.Connect()
	val, err := client.HGetAll(database.RedisCtx, identifier).Result()
	if err != nil{
		log.Print("| Could not fetch from database")
		return models.Version{}, err
	}

	defer client.Close()

	verified := verifyResult(val)
	if verified {
		version := models.Version{
			Url: val["url"],
			ChecksumType: val["checksumtype"],
			Checksum: val["checksum"],
			Version: val["version"],
		}
		return version, nil
	}
	
	return models.Version{}, errors.New("invalid result")

}

func pushToDatabase(c *recorder.ResponseRecorder, r *http.Request) {
	if c.StatusCode == http.StatusOK{
		values := strings.Split(r.RequestURI, "/")

		identifier := fmt.Sprint(values[1], "-", values[2])

		version := models.Version{}

		json.Unmarshal(c.Body, &version)

		client := database.Connect()
		defer client.Close()
		client.HSet(database.RedisCtx, identifier, "url", version.Url, "checksumtype", version.ChecksumType, "checksum", version.Checksum, "version", version.Version)
		client.Expire(database.RedisCtx, identifier, 48*time.Hour)
	}
}

// func pushToDatabase(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		c := httptest.NewRecorder()
// 		next.ServeHTTP(c, r)
		
// 		for k, v := range c.Header() {
//             w.Header()[k] = v
//         }

//         w.WriteHeader(c.Code)
//         c.Body.WriteTo(w)

// 		fmt.Println(r.Response.StatusCode)

// 		if r.Response.StatusCode == http.StatusOK{
// 			values := strings.Split(r.RequestURI, "/")

// 			identifier := fmt.Sprint(values[1], "-", values[2])

// 			version := models.Version{}

// 			defer r.Body.Close()
// 			body, err := io.ReadAll(r.Body)
// 			if err != nil {
// 				log.Print("Could not cache result")
// 			}

// 			json.Unmarshal(body, &version)

// 			client := database.Connect()
// 			client.HSet(database.RedisCtx,identifier, version)
// 			client.Close()
// 		}
// 	})
// }

func verifyResult(result map[string]string) bool {

	if len(result) != 4 {
		return false
	}

	for _, element := range result {
		if element == "" {
			return false
		}
	}

	return true
}