package handler

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func RedirectDownload(w http.ResponseWriter, r *http.Request) {
	flavour := chi.URLParam(r, "flavour")
	id := chi.URLParam(r, "id")
    
    newURL := fmt.Sprintf("%s/%s/%s/%s", baseURL, flavour, id, "/download")

	http.Redirect(w, r, newURL, http.StatusPermanentRedirect)
}
