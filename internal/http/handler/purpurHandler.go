package handler

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/z3orc/dynamic-rpc/internal/client/purpur"
	"github.com/z3orc/dynamic-rpc/internal/models"
	"github.com/z3orc/dynamic-rpc/internal/util"
)

func Purpur(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	uri := r.RequestURI

	url, err := purpur.GetDownloadUrl(id)
	if err != nil {
		util.Error(w, err)
	} else {
		if !strings.Contains(uri, "download"){

			// build, err := paper.GetLatestBuild(id)
			// if err != nil {
			// 	util.Error(w, err)
			// }
	
			version := & models.Version{
				Url: url,
			}
			util.ReturnJson(w, r, *version)
		}
	
		http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	}
}