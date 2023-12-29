package routes

import (
	"poker-game/internal/controllers"

	"github.com/gin-gonic/gin"
)

func SetupGameRouter(router *gin.Engine) {
	router.POST("/joingame", controllers.JoinAGame)
	router.GET("/game/:gameId", controllers.GetGame)
}
