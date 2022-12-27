package handler

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/z3orc/dynamic-rpc/internal/client/paper"
	"github.com/z3orc/dynamic-rpc/internal/util"
)

func PaperAsJson(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

    version, err := paper.GetFormatted(id)
    if err != nil {
        util.Error(w, err)
    } else {
        util.ReturnJson(w, r, version)
    }
    
}

func PaperAsRedirect(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

    url, err := paper.GetDownloadUrl(id)
    if err != nil {
        util.Error(w, err)
    } else {
        http.Redirect(w, r, url, http.StatusTemporaryRedirect)
    }

}
