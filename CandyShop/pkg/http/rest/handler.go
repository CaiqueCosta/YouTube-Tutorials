package rest

import (
	"HexArch/pkg/adding"
	"HexArch/pkg/reading"
	"github.com/gorilla/mux"
)

func InitHandlers(rs reading.Service, as adding.Service) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/", welcomeHandler()).Methods("GET")
	router.HandleFunc("/api/candies", getAllCandiesHandler(rs)).Methods("GET")

	// Adding
	router.HandleFunc("/api/candy", addCandy(as)).Methods("POST")
	return router
}
