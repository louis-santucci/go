package database

import (
	"go-go.com/go-back/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase(db_filename string) {
	db, err := gorm.Open(sqlite.Open(db_filename), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to DB")
	}

	err = db.AutoMigrate(&models.User{}, &models.Redirection{})
	if err != nil {
		panic(err.Error())
	}

	DB = db
}
