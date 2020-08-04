package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// NewStorage returns a new Postgres connection
func NewStorage(host, port, user, password, dbName string) (*sql.DB, error) {
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

	return db, nil
}
