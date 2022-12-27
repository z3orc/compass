package handler

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/z3orc/dynamic-rpc/internal/client/purpur"
	"github.com/z3orc/dynamic-rpc/internal/util"
)

func PurpurAsJson(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

    version, err := purpur.GetFormatted(id)
    if err != nil {
        util.Error(w, err)
    } else {
        util.ReturnJson(w, r, version)
    }
    
}

func PurpurAsRedirect(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

    url, err := purpur.GetDownloadUrl(id)
    if err != nil {
        util.Error(w, err)
    } else {
        http.Redirect(w, r, url, http.StatusTemporaryRedirect)
    }

}
