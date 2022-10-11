package main

import (
	"log"
	"net/http"

	"github.com/izaakdale/service-product/app"
	db "github.com/izaakdale/service-product/db/sqlc"
	"github.com/kelseyhightower/envconfig"
)

type DBSpec struct {
	Host     string
	Port     string
	User     string
	Password string
	Table    string
}

func main() {
	var s DBSpec
	err := envconfig.Process("db", &s)
	if err != nil {
		log.Fatal(err.Error())
	}
	// TODO remove hard code
	err = db.OpenClientConnection(s.Host, s.Port, s.User, s.Password, s.Table)
	if err != nil {
		panic(err)
	}
	log.Fatal(http.ListenAndServe("localhost:8082", app.Router()))
}
