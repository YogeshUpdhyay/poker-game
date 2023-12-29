package controllers

import (
	"net/http"
	"poker-game/internal/initializers"
	"poker-game/internal/models"

	"github.com/gin-gonic/gin"
)

func GetTables(c *gin.Context) {
	var tables []models.Table
	if result := initializers.DB.Find(&tables); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "There was an error fetching the tables"})
	}
	c.JSON(200, gin.H{"tables": tables})
}

func CreateTable(c *gin.Context) {
	var table models.Table

	if err := c.BindJSON(&table); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if result := initializers.DB.Create(&table); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "There was an error adding the table to the database"})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{
		"message": "Created",
		"table":   table,
	})
}
