package controllers

import (
	"errors"
	"net/http"
	"poker-game/internal/initializers"
	"poker-game/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func JoinAGame(c *gin.Context) {
	// Get payload data
	var payload struct {
		PlayerId    int `json:"playerId"`
		TableId     int `json:"tableId"`
		BuyInAmount int `json:"buyInAmount"`
	}

	if err := c.BindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Look for an existing game with available seats
	var game models.Game
	result := initializers.DB.Where("table_id = ? AND available_seats > ?", payload.TableId, 0).Order("id").First(&game)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		// Fetching the table data
		var table models.Table
		if result := initializers.DB.Find(&table, payload.TableId); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating new game"})
			return
		}

		// No game found, create a new game and add the player to it
		game = models.Game{
			TableId:        payload.TableId,
			AvailableSeats: table.MaxNoOfSeats - 1,
			MaxNoOfSeats:   table.MaxNoOfSeats,
			// Initialize other fields as needed
		}
		// Add logic to create a new player and add to game.Players
		// ...

		if result := initializers.DB.Create(&game); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating new game"})
			return
		}
	} else if result.Error != nil {
		// Handle other errors
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error finding game"})
		return
	} else {
		// Game found, add the player to this game
		// Add logic to add the player to game.Players
		// ...

		game.AvailableSeats--
		if result := initializers.DB.Save(&game); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating game"})
			return
		}
	}

	// Return game data in the response
	c.IndentedJSON(http.StatusOK, gin.H{
		"gameId":         game.ID,
		"tableId":        game.TableId,
		"availableSeats": game.AvailableSeats,
		"maxNoOfSeats":   game.MaxNoOfSeats,
		// Include other game data as needed
	})
}
