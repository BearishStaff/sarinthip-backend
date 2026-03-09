package database

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	// 1. Load from .env or Environment Variables
	// Ensure your string ends with sslmode=require
	dsn := os.Getenv("DATABASE_URL")

	// 2. Open Connection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to Supabase:", err)
	}

	// 3. Optional: Auto-Migrate (Creates tables based on your Go structs)
	// err = db.AutoMigrate(&models.Branch{}, &models.Bill{}, &models.Expense{})
	// if err != nil {
	// 	log.Fatal("Migration Failed:", err)
	// }

	log.Println("✅ Database connected via GORM")
	DB = db
}
