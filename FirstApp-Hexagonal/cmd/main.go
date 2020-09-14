package main

import (
	"YouTube-Tutorials/FirstApp-Hexagonal/pkg/adding"
	"YouTube-Tutorials/FirstApp-Hexagonal/pkg/rest"
	"YouTube-Tutorials/FirstApp-Hexagonal/pkg/storage"
	"fmt"
	"log"
	"net/http"
)

func main() {
	r, err := storage.SetupStorage()
	if err != nil {
		fmt.Println(err)
	}

	as := adding.NewService(r)

	router := rest.Handler(as)

	fmt.Println("Starting server on port 8080.....")
	log.Fatal(http.ListenAndServe(":8080", router))
}
