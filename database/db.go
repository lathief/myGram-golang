package database

import (
	"fmt"
	"os"
	"log"
	"myGram/entities"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host       = os.Getenv("DB_HOST")
	user       = os.Getenv("DB_USER")
	password   = os.Getenv("DB_PASSWORD")
	dbPort     = os.Getenv("DB_PORT")
	dbName     = os.Getenv("DB_NAME")
	DEBUG_MODE = true
	db         *gorm.DB
	err        error
)

func StartDB() (*gorm.DB, error) {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbName, dbPort)
	dsn := config
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to database :", err)
	}

	fmt.Println("Connection database success")
	err = db.Debug().AutoMigrate(entities.User{}, entities.Photo{}, entities.Comment{}, entities.SocialMedia{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func GetDB() *gorm.DB {
	if DEBUG_MODE {
		return db.Debug()
	}
	return db
}
