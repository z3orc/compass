package routes

import (
	"github.com/gorilla/mux"
	"github.com/z3orc/dynamic-rpc/internal/http/handler"
)

// func Init(router *chi.Mux) {
// 	router.Route("/vanilla", func(r chi.Router){
// 		r.Get("/{id}", handler.VanillaAsJson)
// 		r.Get("/{id}/download", handler.VanillaAsRedirect)
// 	})
// 	// s := router.PathPrefix("/vanilla").Subrouter()
// 	// s.HandleFunc("/{id}", handler.VanillaAsJson)
// 	// s.HandleFunc("/{id}/download", handler.VanillaAsRedirect)

// 	// s = router.PathPrefix("/paper").Subrouter()
// 	// s.HandleFunc("/{id}", handler.PaperAsJson)
// 	// s.HandleFunc("/{id}/download", handler.PaperAsRedirect)

// 	// s = router.PathPrefix("/purpur").Subrouter()
// 	// s.HandleFunc("/{id}", handler.PurpurAsJson)
// 	// s.HandleFunc("/{id}/download", handler.PurpurAsRedirect)
// }

func Init(router *mux.Router) {
	s := router.PathPrefix("/vanilla").Subrouter()
	s.HandleFunc("/{id}", handler.VanillaAsJson)
	s.HandleFunc("/{id}/download", handler.VanillaAsRedirect)

	s = router.PathPrefix("/paper").Subrouter()
	s.HandleFunc("/{id}", handler.PaperAsJson)
	s.HandleFunc("/{id}/download", handler.PaperAsRedirect)

	s = router.PathPrefix("/purpur").Subrouter()
	s.HandleFunc("/{id}", handler.PurpurAsJson)
	s.HandleFunc("/{id}/download", handler.PurpurAsRedirect)
}