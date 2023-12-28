package controllers

import (
	"log"
	"net/http"
	"poker-game/internal/initializers"
	"poker-game/internal/models"

	"github.com/gin-gonic/gin"
)

func CreatePlayer(c *gin.Context) {
	var player models.Player

	// Bind JSON to the player struct
	if err := c.BindJSON(&player); err != nil {
		// Return a 400 Bad Request response if JSON binding fails
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create the player record in the database
	result := initializers.DB.Create(&player)
	if result.Error != nil {
		// Log the error and return a 500 Internal Server Error response
		log.Printf("Error creating player: %v\n", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	// Log the created player
	log.Printf("Created player: %+v\n", player)

	// Return a 201 Created response with success message
	c.IndentedJSON(http.StatusCreated, gin.H{"msg": "Success", "player": player})
}
