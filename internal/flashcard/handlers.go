package flashcard

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FlashcardHandler struct {
	Service FlashcardServiceInterface
}

func NewFlashcardHandler(service FlashcardServiceInterface) *FlashcardHandler {
	return &FlashcardHandler{Service: service}
}

func (h *FlashcardHandler) ListFlashcards(c *gin.Context) {
	userID := c.GetString("user_id")
	objectID, _ := primitive.ObjectIDFromHex(userID)

	flashcards, err := h.Service.ListFlashcards(context.Background(), objectID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve flashcards"})
		return
	}

	c.JSON(http.StatusOK, flashcards)
}

func (h *FlashcardHandler) MarkFlashcard(c *gin.Context) {
	cardID := c.Param("id")
	objectID, _ := primitive.ObjectIDFromHex(cardID)

	var request struct {
		Knows bool `json:"knows"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := h.Service.MarkFlashcard(context.Background(), objectID, request.Knows); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update flashcard"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Flashcard updated successfully"})
}

func (h *FlashcardHandler) AddFlashcard(c *gin.Context) {
	var flashcard Flashcard
	if err := c.ShouldBindJSON(&flashcard); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	userID := c.GetString("user_id")
	flashcard.UserID, _ = primitive.ObjectIDFromHex(userID)

	if err := h.Service.AddFlashcard(context.Background(), &flashcard); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add flashcard"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Flashcard added successfully", "id": flashcard.ID.Hex()})
}

func (h *FlashcardHandler) RemoveFlashcard(c *gin.Context) {
	cardID := c.Param("id")
	objectID, _ := primitive.ObjectIDFromHex(cardID)

	if err := h.Service.RemoveFlashcard(context.Background(), objectID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove flashcard"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Flashcard removed successfully"})
}
