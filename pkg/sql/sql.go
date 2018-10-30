package sql

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

type VoterDB struct {
	DB *sql.DB
}

func NewSQL(connectionString string) *VoterDB {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	//db.LogMode(true)
	return &VoterDB{DB: db}
}
