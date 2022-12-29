package handler

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

const baseURL = "https://api.compass.z3orc.com"

func Redirect(w http.ResponseWriter, r *http.Request) {
	flavour := chi.URLParam(r, "flavour")
	id := chi.URLParam(r, "id")
    
    newURL := fmt.Sprintf("%s/%s/%s", baseURL, flavour, id)

	http.Redirect(w, r, newURL, http.StatusPermanentRedirect)
}
