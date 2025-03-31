package main

import (
	"log"
	"os"

	"github.com/cameronbibb/backlog-api/db"
	"github.com/cameronbibb/backlog-api/handlers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if os.Getenv("GO_ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Println("Error loading .env file:", err)
		} else {
			log.Println("Environment variables loaded from .env")
		}
	} else {
		log.Println("Running in production mode; skipping .env load")
	}

	db.Connect()

	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.POST("/users", handlers.CreateUser)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
