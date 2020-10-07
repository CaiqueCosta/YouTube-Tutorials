package rest

import "github.com/gorilla/mux"

func InitHandlers() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/", welcomeHandler()).Methods("GET")

	return router
}
