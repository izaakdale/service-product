package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var client *Client

type Client struct {
	*Queries
	// used for tx blocks, unused for now.
	db *sql.DB
}

func OpenClientConnection(host, port, user, password, tableName string) error {
	db, err := sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, tableName))
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}
	client = &Client{
		db:      db,
		Queries: New(db),
	}
	return nil
}

func ClientQueries() *Queries {
	return client.Queries
}
