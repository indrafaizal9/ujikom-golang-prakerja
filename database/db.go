package database

import (
	"fmt"
	"log"
	"os"

	"ujikom/config"
	"ujikom/pkg/models"

	//gorm
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Init(config *config.Config) {
	var err error

	dsn := fmt.Sprintf("host=%s user=%s password=%s  dbname=%s port=%s sslmode=disable",
		config.DBHost, config.DBUserName, config.DBUserPassword, config.DBName, config.DBPort)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to Database! \n", err.Error())
		os.Exit(1)
	}

	DB.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running Migrations...")
	AutoMigrate()
	log.Println("Migrations Completed!")
}

func AutoMigrate() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Recipe{})
	DB.AutoMigrate(&models.Ingredient{})
}
