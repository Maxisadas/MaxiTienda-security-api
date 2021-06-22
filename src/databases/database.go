package databases

import (
	"api-rest/src/configuration"
	"api-rest/src/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DBConn *gorm.DB
)

func InitDB() {
	var err error
	config := configuration.GetConfig()
	dbURL := fmt.Sprintf("host=%s user=%s dbname=%s password=%s sslmode=disable", config.Database_host, config.Database_user, config.Database_dbname, config.Database_password)
	DBConn, err = gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	fmt.Println("Database connection successfully opened")
	//Migrations
	DBConn.AutoMigrate(&models.User{})
	fmt.Println("Migrations completed")
}
