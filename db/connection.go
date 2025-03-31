package db

import (
	"fmt"
	"log"
	"os"

	"github.com/cameronbibb/backlog-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	host := os.Getenv("")
	user := os.Getenv("")
	password := os.Getenv("")
	dbname := os.Getenv("")
	port := os.Getenv("")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disabled TimeZone=UTC", host, user, password, dbname, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}

	DB = db

	db.AutoMigrate(&models.User{}, &models.BacklogItem{}, &models.Game{}, &models.Playlist{})
}
