package database

import (
	"fmt"
	"log"
	"myGram/entities"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host       = "localhost"
	user       = "admin"
	password   = "admin"
	dbPort     = "5432"
	dbName     = "fp-go"
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
