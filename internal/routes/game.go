package routes

import (
	"github.com/gin-gonic/gin"
)

func SetupGameRouter(router *gin.Engine) {
	router.POST("/joingame")
}
