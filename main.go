package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/izaakdale/service-product/app"
	db "github.com/izaakdale/service-product/db/sqlc"
	"github.com/kelseyhightower/envconfig"
)

type SvcSpec struct {
	Port string
}

type DBSpec struct {
	Host     string
	Port     string
	User     string
	Password string
	Table    string
}

func main() {
	var dbs DBSpec
	err := envconfig.Process("db", &dbs)
	if err != nil {
		log.Fatal(err.Error())
	}
	// TODO remove hard code
	err = db.OpenClientConnection(dbs.Host, dbs.Port, dbs.User, dbs.Password, dbs.Table)
	if err != nil {
		panic(err)
	}
	var s SvcSpec
	err = envconfig.Process("service", &s)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", s.Port), app.Router()))
}
