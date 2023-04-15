package main

import (
	"fmt"
	"log"
	"myGram/database"
	"os"

	"myGram/router"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := database.StartDB()
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()
	r.Use(cors.Default())
	router.StartRouter(r, db)

	r.Use(gin.Recovery())

	port := os.Getenv("PORT")

	if len(port) == 0 {
		port = "8080"
	}

	r.Run(fmt.Sprintf(":%s", port))
}
