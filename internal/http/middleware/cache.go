package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/z3orc/dynamic-rpc/internal/models"
	"github.com/z3orc/dynamic-rpc/internal/util"
)

func Cache(next http.Handler) http.Handler {

	cache := make(map[string]string)
	cache["vanilla-1.19.2"] = "https://piston-data.mojang.com/v1/objects/f69c284232d7c7580bd89a5a4931c3581eae1378/server.jar"
	cache["paper-1.19.2"] = "https://api.papermc.io/v2/projects/paper/versions/1.19.2/builds/307/downloads/paper-1.19.2-307.jar"

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		values := strings.Split(r.RequestURI, "/")

		identifier := fmt.Sprint(values[1], "-", values[2])

		val, ok := cache[identifier]

		if ok {
			version := models.Version{
				Url: val,
			}
			w.Header().Add("cached", "True")
			util.ReturnJson(w, r, version)
		} else {
			w.Header().Add("cached", "False")
			next.ServeHTTP(w, r)
		}
	})
}