package main

import (
	"CandyShop/pkg/adding"
	"CandyShop/pkg/http/rest"
	"CandyShop/pkg/reading"
	"CandyShop/pkg/storage/postgresql"
	"fmt"
	"log"
	"net/http"
)

func main() {
	r, err := postgresql.SetupStorage()
	if err != nil {
		log.Fatalln("error while setting up storage:", err)
	}

	rs := reading.NewService(r)
	as := adding.NewService(r)
	fmt.Println("starting server on port 8080...")
	router := rest.InitHandlers(rs, as)
	log.Fatal(http.ListenAndServe(":8080", router))
}


/*
-------GitHub Tag Commands:-------
git tag v2.0
git tag
git push origin v2.0
*/