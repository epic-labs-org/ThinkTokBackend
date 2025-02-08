package main

import (
	"github.com/epic-labs-org/thinktokbackend/cmd/app/routes"
	"github.com/epic-labs-org/thinktokbackend/config"
	"github.com/epic-labs-org/thinktokbackend/internal/flashcard"
	"github.com/epic-labs-org/thinktokbackend/internal/platform/mongodb"
	"github.com/epic-labs-org/thinktokbackend/internal/user"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	// Load configuration
	config.LoadConfig("config/config.yaml")

	// Connect to MongoDB using loaded configuration
	dbClient := mongodb.NewMongoClient(config.AppConfig.Database.URI, config.AppConfig.Database.Name)
	defer dbClient.Disconnect()

	// Initialize user module
	userRepo := user.NewUserRepository(dbClient.DB)
	userService := user.NewUserService(userRepo)
	userHandler := user.NewUserHandler(userService)

	// Initialize flashcard module
	flashcardRepo := flashcard.NewFlashcardRepository(dbClient.DB)
	flashcardService := flashcard.NewFlashcardService(flashcardRepo)
	flashcardHandler := flashcard.NewFlashcardHandler(flashcardService)

	// Create a Gin router
	r := gin.Default()

	// Register routes
	routes.RegisterRoutes(r, userHandler, flashcardHandler)

	// Start the server
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
