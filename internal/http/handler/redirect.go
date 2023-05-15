package handler

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/z3orc/dynamic-rpc/internal/env"
)

var baseURL string = env.APIURL()

func Redirect(w http.ResponseWriter, r *http.Request) {
	flavour := chi.URLParam(r, "flavour")
	id := chi.URLParam(r, "id")

	newURL := fmt.Sprintf("%s/%s/%s", baseURL, flavour, id)

	http.Redirect(w, r, newURL, http.StatusPermanentRedirect)
}
