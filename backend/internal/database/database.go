package database

import "database/sql"

var DB *sql.DB

func Connect(path string) (*sql.DB, error) {
	return nil, nil
}

func Close() {}