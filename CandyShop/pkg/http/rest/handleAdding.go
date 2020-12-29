package rest

import (
	"encoding/json"
	"net/http"

	"github.com/CaiqueCosta/YouTube-Tutorials/CandyShop/pkg/adding"
)

func addCandy(as adding.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var nc adding.Candy
		if err := json.NewDecoder(r.Body).Decode(&nc); err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		id, err := as.AddCandy(nc)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		nc.Id = id
		json.NewEncoder(w).Encode(nc)
	}
}
