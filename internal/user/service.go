package user

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

type UserServiceInterface interface {
	HashPassword(password string) (string, error)
	CheckPassword(hash, password string) bool
	CreateUser(user *User) error
	FindUserByUsername(username string) (*User, error)
	InitializeUserFlashcards(userID primitive.ObjectID) error
}

type UserService struct {
	Repo *UserRepository
}

func NewUserService(repo *UserRepository) *UserService {
	return &UserService{Repo: repo}
}

// HashPassword hashes a plain text password using bcrypt
func (s *UserService) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Failed to hash password: %v", err)
		return "", err
	}
	return string(hashedPassword), nil
}

// CheckPassword compares a plain text password with a hashed password
func (s *UserService) CheckPassword(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// CreateUser creates a new user in the database
func (s *UserService) CreateUser(user *User) error {
	return s.Repo.CreateUser(context.Background(), user)
}

// FindUserByUsername finds a user by their username
func (s *UserService) FindUserByUsername(username string) (*User, error) {
	return s.Repo.FindUserByUsername(context.Background(), username)
}

// InitializeUserFlashcards initializes flashcards for a new user
func (s *UserService) InitializeUserFlashcards(userID primitive.ObjectID) error {
	ctx := context.Background()

	// Fetch all global words from the repository
	words, err := s.Repo.GetAllWords(ctx)
	if err != nil {
		return err
	}

	// Prepare user-specific flashcards
	var flashcards []interface{}
	for _, word := range words {
		flashcards = append(flashcards, bson.M{
			"user_id":     userID,
			"word_id":     word["_id"],
			"interval":    1,
			"next_review": time.Now().Unix(),
			"status":      "unknown",
		})
	}

	// Insert the flashcards into the user's collection
	err = s.Repo.InsertUserFlashcards(ctx, flashcards)
	if err != nil {
		return err
	}

	log.Println("User flashcards initialized successfully")
	return nil
}
