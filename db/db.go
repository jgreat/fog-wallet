package db

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("fog-wallet.db"), &gorm.Config{})
	if err != nil {
		log.Panic("Database connection failed.")
	} else {
		log.Println("Database connection established.")
	}

	return db
}
