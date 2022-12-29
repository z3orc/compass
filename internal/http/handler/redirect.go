package handler

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Redirect(w http.ResponseWriter, r *http.Request) {
	flavour := chi.URLParam(r, "flavour")
	id := chi.URLParam(r, "id")

	http.Redirect(w, r, fmt.Sprint("https://api.compass.z3orc.com/", flavour, "/", id), http.StatusPermanentRedirect)
}