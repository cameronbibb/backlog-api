package handlers

import (
	"net/http"

	"github.com/cameronbibb/backlog-api/db"
	"github.com/cameronbibb/backlog-api/models"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error})
		return
	}

	// Here you would typically check for duplicate username/email, hash the password, etc.

	if err := db.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":           user.ID,
		"username":     user.Username,
		"email":        user.Email,
		"playlists":    user.Playlists,
		"backlogitems": user.BacklogItems,
		"created_at":   user.CreatedAt,
		"updated_at":   user.UpdatedAt,
	})
}

/*
JSON from the client payload -> create a user model struct -> do some checks to make sure the username and email doesn't already exist in the system -> create the user -> send information back to the frontend (beyond initial MVP probably will do some sort of verification of email and use cookies for session)

What am I getting from gin? I need to check the gin docs for router and what to expect here.
In main.go do I import handlers and call user.CreateUser in a POST function?

*/
