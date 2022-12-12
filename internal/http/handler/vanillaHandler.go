package handler

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/z3orc/dynamic-rpc/internal/client/piston"
	"github.com/z3orc/dynamic-rpc/internal/models"
	"github.com/z3orc/dynamic-rpc/internal/util"
)

func Vanilla(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	uri := r.RequestURI

	url, err := piston.GetDownloadUrl(id)
	if err != nil {
		util.Error(w, err)
	} else {
		var isRedirect bool = strings.Contains(uri, "download")

		switch isRedirect {
		case true:
			http.Redirect(w, r, url, http.StatusTemporaryRedirect)
		case false:
			version := &models.Version{
				Url: url,
			}
			util.ReturnJson(w, r, *version)
		}

	}
}
