package config

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"

	"github.com/Jake-Schuler/ORC-MatchMaker/models"
)

func InitDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("data/ORC.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.QualsMatch{})
	return db
}
