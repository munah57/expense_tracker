package database

import (
	"fmt"
	"log"
	"os"
	"tracker/models"
   
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

   )
   

   /* always make sure to put the following in the terminal: go get gorm.io/gorm
     go get gorm.io/driver/postgres */
	 
   var DB *gorm.DB
   
   func ConnectDB() {
	// Connection string 
	connStr := fmt.Sprintf(
	 "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
	 os.Getenv("DB_HOST"),
	 os.Getenv("DB_USER"),
	 os.Getenv("DB_PASSWORD"),
	 os.Getenv("DB_NAME"),
	 os.Getenv("DB_PORT"),
	)
   
	var err error // Initialize error here
	// Open database connection 
	DB, err = gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
	 log.Fatalf("failed to connect to database: %v", err)
	}
	
	// Migrate models to form tables in the database 
	err = DB.AutoMigrate(&models.User{}, &models.Transaction{}, &models.Budget{})
	if err != nil {
	 log.Fatalf("failed to migrate models: %v", err)
	}
	fmt.Println("Connected to the database successfully")
   }
   
