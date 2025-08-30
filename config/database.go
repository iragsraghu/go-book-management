package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/iragsraghu/go-book-management/models" // <-- import your models package
)

var DB *gorm.DB

func ConnectDatabase() {
	// Load .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// Connect without DB first
	tempDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8mb4&parseTime=True&loc=Local",
		user, pass, host, port)

	tempDB, err := gorm.Open(mysql.Open(tempDSN), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Failed to connect to MySQL server:", err)
	}

	// Create database if not exists
	createDB := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;", dbName)
	if err := tempDB.Exec(createDB).Error; err != nil {
		log.Fatal("❌ Failed to create database:", err)
	}
	fmt.Println("✅ Database ensured:", dbName)

	// Now connect with the actual DB
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, pass, host, port, dbName)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Failed to connect to database:", err)
	}

	fmt.Println("✅ Connected to database:", dbName)

	// Run AutoMigrate
	err = DB.AutoMigrate(&models.Book{})
	if err != nil {
		log.Fatal("❌ AutoMigrate failed:", err)
	}
	fmt.Println("✅ AutoMigrate completed")
}
