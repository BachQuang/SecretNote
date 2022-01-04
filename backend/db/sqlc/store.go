package db

import (
	"database/sql"
)

type Store interface {
	Querier
}

type SQLStore struct {
	*Queries
	db *sql.DB
}

// Create NewStore

func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}
