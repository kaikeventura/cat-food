package database

import (
	"github.com/kaikeventura/cat-food/ms-onboarding/configuration/database/migrations"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var database *gorm.DB

func RunDatabase()  {
	dsn := "root:root@tcp(127.0.0.1:3306)/onboarding?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error: ", err)
	}

	database = db

	migrations.ExecuteMigrations(database)
}

func CloseConnection() error {
	config, err := database.DB()
	if err != nil {
		return err
	}

	err = config.Close()
	if err != nil {
		return err
	}

	return nil
}

func GetDatabase() *gorm.DB {
	return database
}
