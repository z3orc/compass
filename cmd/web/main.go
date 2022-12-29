package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/z3orc/dynamic-rpc/internal/env"
	"github.com/z3orc/dynamic-rpc/internal/http/handler"
	"github.com/z3orc/dynamic-rpc/internal/http/middleware"
	"github.com/z3orc/dynamic-rpc/internal/util"
)

var port string = env.ListenerPort()

func main() {
	//ASCII-banner on launch
	util.Banner("CompassWeb")

	router := chi.NewRouter()

	//Middleware
	router.Use(middleware.Recover)
	router.Use(middleware.Logger)

	//Static index
	router.Handle("/", http.FileServer(http.Dir("./static")))
	router.HandleFunc("/{flavour}/{id}", handler.Redirect)
	

	log.Fatal(http.ListenAndServe(port, router))
}
