package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/z3orc/dynamic-rpc/internal/http/middleware"
	"github.com/z3orc/dynamic-rpc/internal/http/routes"
	"github.com/z3orc/dynamic-rpc/internal/util"
)

var port string = util.GetPort()

func main() {

	//ASCII-banner on launch
	util.Banner("DynamicRPC")

	router := mux.NewRouter()

	//Middleware
	router.Use(middleware.Recover)
	router.Use(middleware.Logger)
	router.Use(middleware.Cache)

	//Routes
	routes.Init(router)
	
	//Init server
	log.Fatal(http.ListenAndServe(port, router))
}