package middleware

import (
	"log"
	"net/http"
)

func Recover(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		defer func() {
			err := recover()
			if err != nil {
				log.Print(err)

				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}

		}()

		next.ServeHTTP(w, r)
	})
}