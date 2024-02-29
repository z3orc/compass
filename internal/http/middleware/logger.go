package middleware

import (
	"log"
	"net/http"
	"os"

	"github.com/z3orc/compass/internal/http/recorder"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		c := &recorder.ResponseRecorder{
			ResponseWriter: w,
			StatusCode:     http.StatusOK,
		}

		next.ServeHTTP(c, r)

		log.SetOutput(os.Stdout)
		log.Print("| ", r.Method, " | ", r.RemoteAddr, " | ", r.RequestURI, " | ", c.StatusCode, " | ", c.Header().Get("cached"), " |")
	})
}
