package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/z3orc/compass/internal/http/handler"
	"github.com/z3orc/compass/internal/http/middleware"
)

func Init(router *chi.Mux) {

	router.Get("/", handler.Home)

	router.Route("/vanilla", func(r chi.Router) {
		r.With(middleware.Cache).Get("/{id}", handler.VanillaAsJson)
		r.With(middleware.Cache).Get("/{id}/download", handler.VanillaAsRedirect)
	})

	router.Route("/paper", func(r chi.Router) {
		r.With(middleware.Cache).Get("/{id}", handler.PaperAsJson)
		r.With(middleware.Cache).Get("/{id}/download", handler.PaperAsRedirect)
	})

	router.Route("/purpur", func(r chi.Router) {
		r.With(middleware.Cache).Get("/{id}", handler.PurpurAsJson)
		r.With(middleware.Cache).Get("/{id}/download", handler.PurpurAsRedirect)
	})
}
