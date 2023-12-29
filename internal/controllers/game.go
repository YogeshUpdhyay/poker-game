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
	result := initializers.DB.Where("table_id = ? AND available_seats > ?", payload.TableId, 0).Order("id").Preload("Players").First(&game)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		// Fetching the table data
		var table models.Table
		if result := initializers.DB.Find(&table, payload.TableId); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating new game"})
			return
		}

		// fetching the player to be added to the game
		var player models.Player
		if result := initializers.DB.Find(&player, payload.PlayerId); result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid Player ID recieved"})
				return
			} else {
				// some other random error is recieved
				c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "There was an error fetching the player"})
				return
			}
		}
		// update the player with the buy in amount for the game
		player.BuyInAmount = &payload.BuyInAmount
		initializers.DB.Save(&player)

		// No game found, create a new game and add the player to it
		game = models.Game{
			TableId:        payload.TableId,
			AvailableSeats: table.MaxNoOfSeats - 1,
			MaxNoOfSeats:   table.MaxNoOfSeats,
			Players:        []models.Player{player},
		}

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
		// fetching the player to be added to the game
		var player models.Player
		if result := initializers.DB.Find(&player, payload.PlayerId); result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid Player ID recieved"})
				return
			} else {
				// some other random error is recieved
				c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "There was an error fetching the player"})
				return
			}
		}
		// update the player with the buy in amount for the game
		player.BuyInAmount = &payload.BuyInAmount
		initializers.DB.Save(&player)

		game.AvailableSeats--
		game.Players = append(game.Players, player)
		if result := initializers.DB.Save(&game); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating game"})
			return
		}
	}

	// Return game data in the response
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "Success",
		"game":    game,
	})
}

func GetGame(c *gin.Context) {
	var game models.Game
	if result := initializers.DB.Preload("Players").First(&game, c.Param("gameId")); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Game not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "Success",
		"game":    game,
	})
}
