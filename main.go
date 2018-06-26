package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func init() {
	log.Println("Start initialisation")
	OpenDB()

	log.Println("Initialisation finished")
}

func main() {
	log.Println("main")
	session := Session{
		Name: "test",
	}
	log.Println(session)

	router := gin.Default()
	router.HandleMethodNotAllowed = true
	router.Use(GetCorsSetting())

	api1 := router.Group("/api/v1")
	api1.Use(GetCorsSetting())
	{
		api1.GET("", nil)
	}

	defer DB.Close()

	router.Run(":65432")
}
