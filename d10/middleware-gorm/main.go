package main

import (
	"log"
	"materi-middleware-gorm/config"
	"materi-middleware-gorm/routes"
)

func main() {
	// Initialize database, migrations, and seeding
	config.InitDB()

	// Setup routes and inject database connection
	r := routes.SetupRouter(config.DB)

	log.Println("Server is running on port :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
