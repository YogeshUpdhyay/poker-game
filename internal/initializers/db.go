package initializers

import (
	"log"
	"os"
	"poker-game/internal/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitializeDatabase() {
	dsn := os.Getenv("DB_URL")
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("There was a error connecting to the data base")
	}

	// running migrtions
	DB.AutoMigrate(&models.Table{}, &models.Game{}, &models.Player{})
}
