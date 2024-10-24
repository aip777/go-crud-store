package utils

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"store/models"
)

// Database connection
func InitDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("store.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}

	// Auto migrate the schema
	if err := db.AutoMigrate(&models.Product{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	return db
}
