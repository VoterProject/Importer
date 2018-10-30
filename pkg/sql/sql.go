package sql

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

type VoterDB struct {
	DB *gorm.DB
}

func NewSQL(connectionString string) *VoterDB {
	db, err := gorm.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	return &VoterDB{DB: db}
}
