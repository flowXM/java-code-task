package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"java-code-task/configs"
)

func Connect() (*sql.DB, error) {
	connStr := configs.GetConnectionString()
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	return db, nil
}
