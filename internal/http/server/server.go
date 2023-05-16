package server

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httprate"
	"github.com/z3orc/dynamic-rpc/internal/http/middleware"
	"github.com/z3orc/dynamic-rpc/internal/http/routes"
)

type HttpServer struct {
	router *chi.Mux
	port   string
}

func New(port string) *HttpServer {
	return &HttpServer{
		router: chi.NewRouter(),
		port:   port,
	}
}

func (server *HttpServer) Start() {
	//Middleware
	server.router.Use(middleware.Recover)
	server.router.Use(middleware.Logger)
	server.router.Use(httprate.LimitByIP(
		240,
		60*time.Second,
	))

	//Routes
	routes.Init(server.router)

	//Init listener
	log.Print("| Server listening on ", server.port, " 🚀")
	log.Fatal(http.ListenAndServe(server.port, server.router))
}
