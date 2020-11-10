package rest

import (
	"encoding/json"
	"net/http"

	"github.com/CaiqueCosta/YouTube-Tutorials/CandyShop/pkg/reading"
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
