package utils

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"os"
	"store/models"
)

// Database connection
func InitDB() *gorm.DB {
	db_name := os.Getenv("DB_NAME")
	db, err := gorm.Open(sqlite.Open(db_name), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}

	// Auto migrate the schema
	if err := db.AutoMigrate(&models.Product{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	return db
}
