package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// Storage keeps data in postgres db
type Storage struct {
	db *sql.DB
}

// NewStorage returns a new Postgres storage
func NewStorage(host, port, user, password, dbName string) (*Storage, error) {
	connect := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbName)

	db, err := sql.Open("postgres", connect)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &Storage{db}, nil
}
