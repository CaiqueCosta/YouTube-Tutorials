package rest

import (
	"github.com/CaiqueCosta/YouTube-Tutorials/FirstApp-Hexagonal/FirstApp-Hexagonal/pkg/adding"
	"github.com/go-chi/chi"
)

func Handler(as adding.Service) *chi.Mux {
	router := chi.NewRouter()

	router.Get("/api/users", handleAdding(as))

	return router
}
