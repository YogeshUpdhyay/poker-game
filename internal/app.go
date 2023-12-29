package internal

import (
	"log"
	"poker-game/internal/initializers"
	"poker-game/internal/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var router *gin.Engine

func InitializeEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("There was an error setting up the app environment")
	}
}

func GetApp() *gin.Engine {
	// initialize environment
	InitializeEnv()

	// intializing the database
	initializers.InitializeDatabase()

	router = gin.Default()
	routes.SetupTableRouter(router)
	routes.SetupPlayerRouter(router)
	routes.SetupGameRouter(router)

	return router
}
