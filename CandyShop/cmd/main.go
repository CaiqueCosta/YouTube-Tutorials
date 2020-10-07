package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/CaiqueCosta/YouTube-Tutorials/CandyShop/pkg/http/rest"
)

func main() {
	fmt.Println("starting server on port 8080...")
	router := rest.InitHandlers()
	log.Fatal(http.ListenAndServe(":8080", router))
}
