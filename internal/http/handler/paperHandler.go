package handler

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/z3orc/dynamic-rpc/internal/client/paper"
	"github.com/z3orc/dynamic-rpc/internal/util"
)

func Paper(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	uri := r.RequestURI
    asRedirect := strings.Contains(uri, "download")

    switch asRedirect {
    case true:
        url, err := paper.GetDownloadUrl(id)
        if err != nil {
            util.Error(w, err)
        }

        http.Redirect(w, r, url, http.StatusTemporaryRedirect)
    case false:
        version, err := paper.GetFormatted(id)
        if err != nil {
            util.Error(w, err)
        } 
        util.ReturnJson(w, r, version)
    }
}
