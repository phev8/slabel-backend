package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // use as DB connector
)

// DB is a pointer to the database
var DB *gorm.DB

// OpenDB connects to the database
func OpenDB() {
	var err error
	DB, err = gorm.Open("sqlite3", "data.DB")

	if err != nil {
		log.Printf("Error from gorm.Open: %s\n", err)
	}

	log.Println("You connected to your database.")
}

// GetCorsSetting for setting api properties
func GetCorsSetting() gin.HandlerFunc {
	return cors.Middleware(cors.Config{
		Origins:         "*",
		RequestHeaders:  "Authorization, Origin, Content-Type",
		Methods:         "GET, POST, PUT, DELETE",
		Credentials:     true,
		ValidateHeaders: false,
		MaxAge:          10 * time.Minute,
	})
}
