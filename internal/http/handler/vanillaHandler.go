package handler

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/z3orc/dynamic-rpc/internal/client/piston"
	"github.com/z3orc/dynamic-rpc/internal/util"
)

func VanillaAsJson(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

    version, err := piston.GetFormatted(id)
    if err != nil {
        util.Error(w, err)
    } else {
        util.ReturnJson(w, r, version)
    }
    
}

func VanillaAsRedirect(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

    url, err := piston.GetDownloadUrl(id)
    if err != nil {
        util.Error(w, err)
    } else {
        http.Redirect(w, r, url, http.StatusTemporaryRedirect)
    }

}
