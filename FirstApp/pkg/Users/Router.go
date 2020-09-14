package Users

import "github.com/go-chi/chi"

func UsersRoutes(s Service) *chi.Mux {
	router := chi.NewRouter()
	router.Get("/getUser", getHandler(s))
	router.Post("/updateUser", postHandler(s))
	return router
}
