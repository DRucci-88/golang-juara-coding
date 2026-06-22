package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"mini-hris/config"
	"mini-hris/routes"
)

func main() {
	if err := config.ConnectDatabase(); err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	if err := config.SeedData(); err != nil {
		log.Fatalf("failed to seed database: %v", err)
	}

	router := gin.Default()
	routes.SetupRoutes(router)

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
