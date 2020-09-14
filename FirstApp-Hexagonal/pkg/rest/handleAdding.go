package rest

import (
	"net/http"

	"github.com/CaiqueCosta/YouTube-Tutorials/FirstApp-Hexagonal/FirstApp-Hexagonal/pkg/adding"
)

func handleAdding(as adding.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
