package main

import (
	"FirstApp-Hexagonal/pkg/adding"
	"FirstApp-Hexagonal/pkg/rest"
	"FirstApp-Hexagonal/pkg/storage"
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
