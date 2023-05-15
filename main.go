package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httprate"
	"github.com/z3orc/dynamic-rpc/internal/env"
	"github.com/z3orc/dynamic-rpc/internal/http/middleware"
	"github.com/z3orc/dynamic-rpc/internal/http/routes"
	"github.com/z3orc/dynamic-rpc/internal/util"
)

var Version string = ""
var Build string = ""
var port string = env.ListenerPort()

func main() {

	//ASCII-banner on launch
	util.Banner("Compass", Version, Build)

	//Init router
	router := chi.NewRouter()

	//Middleware
	router.Use(middleware.Recover)
	router.Use(middleware.Logger)
	router.Use(httprate.LimitByIP(
		240,
		60*time.Second,
	))

	//Routes
	routes.Init(router)

	//Init listener
	log.Print("| Server listening on ", port, " 🚀")
	log.Fatal(http.ListenAndServe(port, router))
}