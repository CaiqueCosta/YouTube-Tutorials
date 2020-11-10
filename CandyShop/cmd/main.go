package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/CaiqueCosta/YouTube-Tutorials/CandyShop/pkg/http/rest"
	"github.com/CaiqueCosta/YouTube-Tutorials/CandyShop/pkg/reading"
	"github.com/CaiqueCosta/YouTube-Tutorials/CandyShop/pkg/storage"
)

func main() {

	r, err := storage.SetupStorage()
	if err != nil {
		log.Fatalln("error while setting up storage:", err)
	}

	rs := reading.NewService(r)

	fmt.Println("starting server on port 8080...")
	router := rest.InitHandlers(rs)
	log.Fatal(http.ListenAndServe(":8080", router))
}

/*
-------GitHub Tag Commands:-------
git tag v2.0
git tag
git push origin v2.0
*/
