package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/z3orc/dynamic-rpc/internal/database"
	"github.com/z3orc/dynamic-rpc/internal/http/middleware"
	"github.com/z3orc/dynamic-rpc/internal/http/routes"
	"github.com/z3orc/dynamic-rpc/internal/util"
)

var port string = util.GetPort()

func main() {

	client := database.Connect()
	fmt.Println(database.Check(client))

	router := mux.NewRouter()

	//Middleware
	router.Use(middleware.Recover)
	router.Use(middleware.Logger)
	router.Use(middleware.Cache)

	//Routes
	routes.Init(router)
	
	//ASCII-banner on launch
	util.Banner("DynamicRPC")
	log.Fatal(http.ListenAndServe(port, router))
}