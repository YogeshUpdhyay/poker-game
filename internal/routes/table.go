package routes

import (
	"poker-game/internal/controllers"

	"github.com/gin-gonic/gin"
)

func SetupTableRouter(router *gin.Engine) {
	router.GET("/tables", controllers.GetTables)
	router.POST("/table", controllers.CreateTable)
}
