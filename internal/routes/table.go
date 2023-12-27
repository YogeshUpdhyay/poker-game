package routes

import (
	contrtollers "poker-game/internal/controllers"

	"github.com/gin-gonic/gin"
)

func SetupTableRouter(router *gin.Engine) {
	router.GET("/tables", contrtollers.GetTables)
	router.POST("/tables", contrtollers.JoinATable)
}
