package middleware

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/z3orc/compass/internal/database"
	"github.com/z3orc/compass/internal/http/recorder"
	"github.com/z3orc/compass/internal/models"
	"github.com/z3orc/compass/internal/util"
)

func Cache(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		URI := r.RequestURI

		//Skips cache
		if strings.Contains(URI, "download") {
			next.ServeHTTP(w, r)

		}

		cachedResult, err := getFromDatabase(r)
		if err != nil {
			w.Header().Add("cached", "False")

			c := &recorder.ResponseRecorder{
				ResponseWriter: w,
				StatusCode:     http.StatusOK,
				Body:           []byte{},
			}
			next.ServeHTTP(c, r)
			go pushToDatabase(c, r)
		} else {
			w.Header().Add("cached", "True")
			util.ReturnJson(w, r, cachedResult)
		}

	})
}

func getFromDatabase(r *http.Request) (models.Version, error) {
	values := strings.Split(r.RequestURI, "/")

	identifier := fmt.Sprint(values[1], "-", values[2])

	client := database.Connect()
	defer client.Close()

	val, err := client.HGetAll(database.RedisCtx, identifier).Result()
	if err != nil {
		// log.Print("| Could not fetch from database")
		return models.Version{}, err
	}

	verified := verifyResult(val)
	if verified {
		version := models.Version{
			Url:          val["url"],
			ChecksumType: val["checksumtype"],
			Checksum:     val["checksum"],
			Version:      val["version"],
		}
		return version, nil
	}

	return models.Version{}, errors.New("invalid result")

}

func pushToDatabase(c *recorder.ResponseRecorder, r *http.Request) {
	if c.StatusCode == http.StatusOK {
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
