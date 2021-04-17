package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/leozz37/gin-app-template/routes"
	"github.com/leozz37/gin-app-template/services/db"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load(".env.example")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Start DB
	db.ConnectMySQL(os.Getenv("DATABASE_TYPE"), os.Getenv("DATABASE_DSN"))

	// Start Gin
	routes.InitRoutes()
}
