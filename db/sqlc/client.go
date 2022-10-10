package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var client *Client

type Client struct {
	*Queries
	// used for tx blocks, unused for now.
	db *sql.DB
}

func OpenClientConnection(host, user, password string) error {
	db, err := sql.Open("postgres", "postgresql://admin:password@localhost:5432/ordering-app?sslmode=disable")
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
