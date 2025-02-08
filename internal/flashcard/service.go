package flashcard

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type FlashcardServiceInterface interface {
	AddFlashcard(ctx context.Context, flashcard *Flashcard) error
	RemoveFlashcard(ctx context.Context, cardID primitive.ObjectID) error
	MarkFlashcard(ctx context.Context, cardID primitive.ObjectID, knows bool) error
	ListFlashcards(ctx context.Context, userID primitive.ObjectID) ([]Flashcard, error)
}
type FlashcardService struct {
	Repo FlashcardRepositoryInterface
}

func NewFlashcardService(repo FlashcardRepositoryInterface) *FlashcardService {
	return &FlashcardService{Repo: repo}
}

func (s *FlashcardService) AddFlashcard(ctx context.Context, flashcard *Flashcard) error {
	flashcard.ID = primitive.NewObjectID()
	flashcard.Interval = 1                   // Initial interval is 1 day
	flashcard.NextReview = time.Now().Unix() // Set the first review to the current time
	flashcard.Status = "unknown"             // Default status is "unknown"

	return s.Repo.InsertFlashcard(ctx, flashcard)
}

func (s *FlashcardService) RemoveFlashcard(ctx context.Context, cardID primitive.ObjectID) error {
	return s.Repo.DeleteFlashcard(ctx, cardID)
}

func (s *FlashcardService) ListFlashcards(ctx context.Context, userID primitive.ObjectID) ([]Flashcard, error) {
	return s.Repo.ListFlashcards(ctx, userID)
}

// MarkFlashcard updates the flashcard's review interval using an enhanced SM2 algorithm
func (s *FlashcardService) MarkFlashcard(ctx context.Context, cardID primitive.ObjectID, knows bool) error {
	flashcard, err := s.Repo.FindFlashcardByID(ctx, cardID)
	if err != nil {
		return err
	}

	if knows {
		flashcard.Interval = s.calculateNewInterval(flashcard.Interval)
	} else {
		flashcard.Interval = 1
	}

	flashcard.NextReview = time.Now().AddDate(0, 0, flashcard.Interval).Unix()

	// Track the review with history
	status := "known"
	if !knows {
		status = "unknown"
	}

	return s.Repo.UpdateFlashcardStatus(ctx, flashcard.ID, status, flashcard.Interval, flashcard.NextReview)
}

func (s *FlashcardService) calculateNewInterval(currentInterval int) int {
	if currentInterval < 8 {
		return int(float64(currentInterval) * 1.5) // Smaller intervals grow slower
	} else if currentInterval < 30 {
		return currentInterval * 2 // Medium intervals grow moderately
	} else {
		return currentInterval * 3 // Larger intervals grow faster
	}
}
