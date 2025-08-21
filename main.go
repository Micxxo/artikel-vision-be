package main

import (
	"log"
	"os"

	"github.com/Micxxo/artikel-vision-be/databases"
	"github.com/Micxxo/artikel-vision-be/migrations"
	"github.com/Micxxo/artikel-vision-be/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// database
	databasePath := os.Getenv("DB_PATH")
	db, err := databases.ConnectDB(databasePath)
	if err != nil {
		log.Fatal("Error connecting to database")
	}

	// database migration
	if err := migrations.RunMigrations(db); err != nil {
		log.Fatal("Database migration error")
	}

	app := fiber.New()

	// CORS middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS,PATCH",
	}))

	// routes
	api := app.Group("/api")
	routes.ArticleRoutes(api)

	serverPort := os.Getenv("APP_PORT")
	app.Listen(":" + serverPort)
}
