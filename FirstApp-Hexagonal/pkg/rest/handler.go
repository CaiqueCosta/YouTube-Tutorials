package rest

import (
	"FirstApp-Hexagonal/pkg/adding"

	"github.com/go-chi/chi"
)

func Handler(as adding.Service) *chi.Mux {
	router := chi.NewRouter()

	router.Post("/api/users", handleAdding(as))
	router.Get("/api/users", handleReading(as))

	return router
}
