package user

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	UserCollection          *mongo.Collection
	UserFlashcardCollection *mongo.Collection
	WordCollection          *mongo.Collection
}

func NewUserRepository(db *mongo.Database) *UserRepository {
	return &UserRepository{
		UserCollection:          db.Collection("users"),
		UserFlashcardCollection: db.Collection("user_flashcards"),
		WordCollection:          db.Collection("words"),
	}
}

func (r *UserRepository) CreateUser(ctx context.Context, user *User) error {
	user.ID = primitive.NewObjectID()
	_, err := r.UserCollection.InsertOne(ctx, user)
	return err
}

func (r *UserRepository) FindUserByUsername(ctx context.Context, username string) (*User, error) {
	var user User
	err := r.UserCollection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetAllWords(ctx context.Context) ([]bson.M, error) {
	cursor, err := r.WordCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var words []bson.M
	if err := cursor.All(ctx, &words); err != nil {
		return nil, err
	}
	return words, nil
}

func (r *UserRepository) InsertUserFlashcards(ctx context.Context, flashcards []interface{}) error {
	_, err := r.UserFlashcardCollection.InsertMany(ctx, flashcards)
	return err
}
