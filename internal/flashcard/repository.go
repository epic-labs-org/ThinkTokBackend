package flashcard

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type FlashcardRepositoryInterface interface {
	InsertFlashcard(ctx context.Context, flashcard *Flashcard) error
	DeleteFlashcard(ctx context.Context, cardID primitive.ObjectID) error
	ListFlashcards(ctx context.Context, userID primitive.ObjectID) ([]Flashcard, error)
	UpdateFlashcardStatus(ctx context.Context, cardID primitive.ObjectID, status string, interval int, nextReview int64) error
	FindFlashcardByID(ctx context.Context, cardID primitive.ObjectID) (*Flashcard, error)
}

type FlashcardRepository struct {
	Collection *mongo.Collection
}

func NewFlashcardRepository(db *mongo.Database) *FlashcardRepository {
	return &FlashcardRepository{
		Collection: db.Collection("flashcards"),
	}
}

func (r *FlashcardRepository) InsertFlashcard(ctx context.Context, flashcard *Flashcard) error {
	_, err := r.Collection.InsertOne(ctx, flashcard)
	return err
}

func (r *FlashcardRepository) DeleteFlashcard(ctx context.Context, cardID primitive.ObjectID) error {
	_, err := r.Collection.DeleteOne(ctx, bson.M{"_id": cardID})
	return err
}

func (r *FlashcardRepository) ListFlashcards(ctx context.Context, userID primitive.ObjectID) ([]Flashcard, error) {
	var flashcards []Flashcard
	cursor, err := r.Collection.Find(ctx, bson.M{"user_id": userID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &flashcards); err != nil {
		return nil, err
	}
	return flashcards, nil
}

func (r *FlashcardRepository) UpdateFlashcardStatus(ctx context.Context, cardID primitive.ObjectID, status string, interval int, nextReview int64) error {
	reviewEntry := ReviewEntry{
		Timestamp: time.Now().Unix(),
		Status:    status,
		Interval:  interval,
	}

	_, err := r.Collection.UpdateOne(ctx, bson.M{"_id": cardID}, bson.M{
		"$set":  bson.M{"status": status, "interval": interval, "next_review": nextReview},
		"$push": bson.M{"review_history": reviewEntry}, // Append to review history
	})
	return err
}

func (r *FlashcardRepository) FindFlashcardByID(ctx context.Context, cardID primitive.ObjectID) (*Flashcard, error) {
	var flashcard Flashcard
	err := r.Collection.FindOne(ctx, bson.M{"_id": cardID}).Decode(&flashcard)
	if err != nil {
		return nil, err
	}
	return &flashcard, nil
}
