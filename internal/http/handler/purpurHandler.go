package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/z3orc/compass/internal/client/purpur"
	"github.com/z3orc/compass/internal/util"
)

func PurpurAsJson(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	version, err := purpur.GetFormatted(id)
	if err != nil {
		util.Error(w, err)
	} else {
		util.ReturnJson(w, r, version)
	}

}

func PurpurAsRedirect(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	url, err := purpur.GetDownloadUrl(id)
	if err != nil {
		util.Error(w, err)
	} else {
		http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	}

}
