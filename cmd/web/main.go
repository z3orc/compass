package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/z3orc/dynamic-rpc/internal/http/middleware"
	"github.com/z3orc/dynamic-rpc/internal/util"
)

var port string = util.GetPort()

func main() {
	router := chi.NewRouter()

	//Middleware
	router.Use(middleware.Recover)
	router.Use(middleware.Logger)

	//Static index
	router.Handle("/*", http.FileServer(http.Dir("./static")))
	
	//ASCII-banner on launch
	util.Banner("DynamicWeb")
	log.Fatal(http.ListenAndServe(port, router))
}