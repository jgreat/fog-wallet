package models

import (
	"log"

	"gorm.io/gorm"
)

func DoMigrations(db *gorm.DB) {
	log.Println("Start DB migrations.")

	err := db.Debug().AutoMigrate(&Account{})
	if err != nil {
		log.Panic("Problem with migration")
	}

	log.Println("DB up to date.")
}
