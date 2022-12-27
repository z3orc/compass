package middleware

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c := httptest.NewRecorder()
		next.ServeHTTP(c, r)
		
		for k, v := range c.Header() {
            w.Header()[k] = v
        }
        w.WriteHeader(c.Code)
        c.Body.WriteTo(w)

		log.SetOutput(os.Stdout)
		log.Print("| ", r.Method, " | ",  r.RemoteAddr, " | ", r.RequestURI,  " | ", c.Result().Status,  " | ", c.Header().Get("cached"))
	})
}