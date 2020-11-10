package rest

import (
	"github.com/CaiqueCosta/YouTube-Tutorials/CandyShop/pkg/reading"
	"github.com/gorilla/mux"
)

func InitHandlers(rs reading.Service) *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/", welcomeHandler()).Methods("GET")
	router.HandleFunc("/api/candies", getAllCandiesHandler(rs)).Methods("GET")
	return router
}
