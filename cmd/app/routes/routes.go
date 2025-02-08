package routes

import (
	"github.com/epic-labs-org/thinktokbackend/internal/flashcard"
	"github.com/epic-labs-org/thinktokbackend/internal/platform/middleware"
	"github.com/epic-labs-org/thinktokbackend/internal/user"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, userHandler *user.UserHandler, flashcardHandler *flashcard.FlashcardHandler) {
	api := router.Group("/api/v1")
	{
		// Health check
		api.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{"status": "OK"})
		})

		// User routes
		api.POST("/register", userHandler.RegisterUser)
		api.POST("/login", userHandler.LoginUser)

		// Protected routes
		flashcards := api.Group("/flashcards", middleware.JWTAuthMiddleware())
		{
			flashcards.GET("/", flashcardHandler.ListFlashcards)
			flashcards.POST("/", flashcardHandler.AddFlashcard)         // Add flashcard
			flashcards.DELETE("/:id", flashcardHandler.RemoveFlashcard) // Remove flashcard
			flashcards.POST("/:id/memorize", flashcardHandler.MarkFlashcard)
		}
	}
}
