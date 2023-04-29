package config

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBConnect(dbDsn string) {
	db, err := gorm.Open(postgres.Open(dbDsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database", err.Error())
	} else {
		fmt.Println("connected to db")
		DB = db
	}
}
