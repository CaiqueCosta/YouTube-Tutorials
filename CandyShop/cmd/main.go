package main

import (
	"HexArch/pkg/adding"
	"HexArch/pkg/http/rest"
	"HexArch/pkg/reading"
	"HexArch/pkg/storage/postgresql"
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
