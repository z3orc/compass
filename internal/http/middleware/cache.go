package middleware

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/z3orc/dynamic-rpc/internal/database"
	"github.com/z3orc/dynamic-rpc/internal/models"
	"github.com/z3orc/dynamic-rpc/internal/util"
)

func Cache(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		cachedResult, err := getFromDatabase(r)
		if err != nil {
			fmt.Println(err.Error())
			w.Header().Add("cached", "False")
			next.ServeHTTP(w, r)
			pushToDatabase(next)
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
		log.Print("Could not fetch from database")
		return models.Version{}, err
	}

	client.Close()

	verified := verifyResult(val)
	fmt.Println(verified)
	if verified {
		version := models.Version{
			Url: val["url"],
			ChecksumType: val["checksumtype"],
			Checksum: val["checksum"],
			Version: val["version"],
		}
		return version, nil
	} else {
		return models.Version{}, errors.New("invalid result")
	}

}

func pushToDatabase(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c := httptest.NewRecorder()
		next.ServeHTTP(c, r)
		
		for k, v := range c.Header() {
            w.Header()[k] = v
        }

        w.WriteHeader(c.Code)
        c.Body.WriteTo(w)

		fmt.Println(r.Response.StatusCode)

		if r.Response.StatusCode == http.StatusOK{
			values := strings.Split(r.RequestURI, "/")

			identifier := fmt.Sprint(values[1], "-", values[2])

			version := models.Version{}

			defer r.Body.Close()
			body, err := io.ReadAll(r.Body)
			if err != nil {
				log.Print("Could not cache result")
			}

			json.Unmarshal(body, &version)

			client := database.Connect()
			client.HSet(database.RedisCtx,identifier, version)
			client.Close()
		}
	})
}

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