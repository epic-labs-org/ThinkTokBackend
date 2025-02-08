package user

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Username       string             `bson:"username" json:"username"`
	PasswordHash   string             `bson:"password_hash" json:"-"`
	NativeLanguage string             `bson:"native_language" json:"native_language"`
}
