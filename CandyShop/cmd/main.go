package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/CaiqueCosta/YouTube-Tutorials/CandyShop/pkg/http/rest"
	"github.com/CaiqueCosta/YouTube-Tutorials/CandyShop/pkg/storage"
)

func main() {

	r, err := storage.SetupStorage()
	if err != nil {
		log.Fatalln("error while setting up storage:", err)
	}

	c, err := r.GetAllCandyNames()
	if err != nil {
		log.Fatalln("error while getting candies in storage:", err)
	}
	log.Println(c)

	fmt.Println("starting server on port 8080...")
	router := rest.InitHandlers()
	log.Fatal(http.ListenAndServe(":8080", router))
}
