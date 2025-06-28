package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"ctf-toolkit-backend/internal/api/routes"
	"ctf-toolkit-backend/internal/config"
	"ctf-toolkit-backend/internal/database"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	//  MongoDB
	database.Connect(cfg.MongoURI)

	// Create uploads directory if it doesn't exist
	if err := os.MkdirAll("uploads", 0755); err != nil {
		log.Fatal("Failed to create uploads directory:", err)
	}

	// Fiber app init
	app := fiber.New(fiber.Config{
		BodyLimit: 100 * 1024 * 1024, // 100MB limit
	})

	//middleware
	app.Use(cors.New())
	app.Use(logger.New())

	// Setup routes
	routes.SetupRoutes(app)

	// Start server
	port := cfg.Port
	if port == "" {
		port = "8080"
	}
	log.Printf("Starting server on port %s...", port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}