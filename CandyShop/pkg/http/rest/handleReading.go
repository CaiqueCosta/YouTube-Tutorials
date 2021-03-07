package rest

import (
	"CandyShop/pkg/reading"
	"encoding/json"
	"net/http"
)

func welcomeHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("Welcome to our candy shop!")
	}
}

func getAllCandiesHandler(rs reading.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		cs, err := rs.GetAllCandyNames()
		if err != nil {
			http.Error(w, "Cannot process your request at this time...", http.StatusInternalServerError)
		}
		json.NewEncoder(w).Encode(cs)
	}
}
