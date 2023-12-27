package controllers

import (
	"net/http"
	"poker-game/internal/serializers"

	"github.com/gin-gonic/gin"
)

func GetTables(c *gin.Context) {
	tables := []serializers.Table{
		{Name: "LA", MaxBuyIn: 10000},
		{Name: "Miami", MaxBuyIn: 50000},
		{Name: "Las Vegas", MaxBuyIn: 100000},
	}
	c.JSON(200, gin.H{"tables": tables})
}

func JoinATable(c *gin.Context) {
	var body struct {
		Name string `json:"name"`
	}
	if err := c.BindJSON(&body); err != nil {
		// Handle the error if the JSON is not as expected
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, body)
}
