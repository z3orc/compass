package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/z3orc/compass/internal/client/piston"
	"github.com/z3orc/compass/internal/util"
)

func VanillaAsJson(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	version, err := piston.GetFormatted(id)
	if err != nil {
		util.Error(w, err)
	} else {
		util.ReturnJson(w, r, version)
	}

}

func VanillaAsRedirect(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	url, err := piston.GetDownloadUrl(id)
	if err != nil {
		util.Error(w, err)
	} else {
		http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	}

}
