package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/z3orc/compass/internal/client/paper"
	"github.com/z3orc/compass/internal/util"
)

func PaperAsJson(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	version, err := paper.GetFormatted(id)
	if err != nil {
		util.Error(w, err)
	} else {
		util.ReturnJson(w, r, version)
	}

}

func PaperAsRedirect(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	url, err := paper.GetDownloadUrl(id)
	if err != nil {
		util.Error(w, err)
	} else {
		http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	}

}
