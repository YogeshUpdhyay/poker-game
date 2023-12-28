package internal

import (
	"log"
	"os"
	"poker-game/internal/models"
	"poker-game/internal/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var router *gin.Engine
var db *gorm.DB

func InitializeEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("There was an error setting up the app environment")
	}
}

func InitializeDatabase() {
	dsn := os.Getenv("DB_URL")
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("There was a error connecting to the data base")
	}

	// running migrtions
	db.AutoMigrate(&models.Table{})
}

func GetDB() *gorm.DB {
	return db
}

func GetApp() *gin.Engine {
	// initialize environment
	InitializeEnv()

	// intializing the database
	InitializeDatabase()

	router = gin.Default()
	routes.SetupTableRouter(router)
	return router
}
