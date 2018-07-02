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
		authorized := api1.Group("/")
		authorized.Use(KeyRequired())

		authorized.GET("/test", TestAPI)
		authorized.GET("/labelset", GetLabelSetsHandl)
		authorized.POST("/labelset", CreateLabelSetHandl)
		authorized.PUT("/labelset", UpdateLabelSetHandl)
		authorized.DELETE("/labelset", DeleteLabelSetHandl)

		authorized.GET("/labelset/labels", GetSingleLabelSetHandl)
		authorized.POST("/labelset/label", CreateLabelTemplateHandl)
		authorized.PUT("/labelset/label", UpdateLabelTemplateHandl)
		authorized.DELETE("/labelset/label", DeleteLabelTemplateHandl)

		authorized.GET("/session", GetSessionsHandl)
		authorized.POST("/session", CreateSessionHandl)
		authorized.PUT("/session", UpdateSessionHandl)
		authorized.DELETE("/session", DeleteSessionHandl)

		authorized.GET("/session/labels", GetSingleSessionHandl)
		authorized.POST("/session/label", CreateLabelHandl)
		authorized.PUT("/session/label", UpdateLabelHandl)
		authorized.DELETE("/session/label", DeleteLabelHandl)

	}

	defer DB.Close()

	router.Run(":65432")
}
