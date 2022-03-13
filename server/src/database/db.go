package database

import (
	"fmt"
	"log"
	"new_go_app/src/models"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error

	dsn := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo",
		"postgres", "postgres", "gorm", "5432")

	log.Print("Connecting to PostgreSQL DB...")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}
	log.Println("connected")

	log.Print("Running the migrations...")
}

func AutoMigrate() {
	DB.AutoMigrate(&models.User{}, &models.Post{}, &models.Vote{}, &models.Reply{})
}
