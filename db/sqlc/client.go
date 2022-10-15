package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/izaakdale/utils-go/logger"
	_ "github.com/lib/pq"
)

var client *Client

type Client struct {
	*Queries
	// used for tx blocks, unused for now.
	db *sql.DB
}

func connectToDb(host, port, user, password, tableName string) (*sql.DB, error) {
	db, err := sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, tableName))
	if err != nil {
		logger.Error("error opening db")
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		logger.Error("error pinging db")
		return nil, err
	}
	return db, err
}

func OpenClientConnection(host, port, user, password, tableName string) error {

	var counts int
	for {
		conn, err := connectToDb(host, port, user, password, tableName)
		if err == nil {
			logger.Info("Connected to the DB!")
			client = &Client{
				db:      conn,
				Queries: New(conn),
			}
			break
		} else {
			logger.Error("Could not connect to db, attempting reconnect")
			counts++
		}
		if counts > 10 {
			logger.Error("Connection attempts surpassed 10, giving up...")
			return err
		}
		log.Println("Back off for two seconds")
		time.Sleep(time.Second * 2)
		continue
	}
	return nil
}

func ClientQueries() *Queries {
	return client.Queries
}
