package handler

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/z3orc/dynamic-rpc/internal/client/purpur"
	"github.com/z3orc/dynamic-rpc/internal/util"
)

func Purpur(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	uri := r.RequestURI

	url, err := purpur.GetDownloadUrl(id)
	if err != nil {
		util.Error(w, err)
	}

    var isRedirect bool = strings.Contains(uri, "download")

    switch isRedirect {
    case true:
        http.Redirect(w, r, url, http.StatusTemporaryRedirect)
    case false:
        version, err := purpur.GetFormatted(id)
        
        if err != nil {
            util.Error(w, err)
        }
        util.ReturnJson(w, r, version)
    }

}
