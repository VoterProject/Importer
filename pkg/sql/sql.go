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
	//db.LogMode(true)
	db.DB().SetMaxIdleConns(10)
	return &VoterDB{DB: db}
}
