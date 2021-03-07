package rest

import (
	"HexArch/pkg/adding"
	"encoding/json"
	"net/http"
)

func addCandy(as adding.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var nc adding.Candy
		if err := json.NewDecoder(r.Body).Decode(&nc); err != nil {
			http.Error(w, "internal Server Error", http.StatusInternalServerError)
			return
		}
		id, err := as.AddCandy(nc)
		if err != nil {
			http.Error(w, "internal Server Error", http.StatusInternalServerError)
			return
		}
		nc.Id = id
		json.NewEncoder(w).Encode(nc)
	}
}
