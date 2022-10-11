package main

import (
	"log"
	"net/http"

	"github.com/izaakdale/service-product/app"
	db "github.com/izaakdale/service-product/db/sqlc"
)

func main() {
	// TODO remove hard code
	err := db.OpenClientConnection("localhost", "admin", "password", "ordering-app")
	if err != nil {
		panic(err)
	}
	log.Fatal(http.ListenAndServe("localhost:8082", app.Router()))
}
