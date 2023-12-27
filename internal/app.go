package internal

import (
	"poker-game/internal/routes"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func GetApp() *gin.Engine {
	router = gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	routes.SetupTableRouter(router)
	return router
}
