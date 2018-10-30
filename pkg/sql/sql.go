package sql

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

type voterdb struct {
	DB *sql.DB
}

func NewSQL(connectionString string) *voterdb {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	return &voterdb{DB: db}
}
