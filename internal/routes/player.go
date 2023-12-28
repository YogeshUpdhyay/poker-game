package routes

import (
	"poker-game/internal/controllers"

	"github.com/gin-gonic/gin"
)

func SetupPlayerRouter(router *gin.Engine) {
	router.POST("/player", controllers.CreatePlayer)
}
