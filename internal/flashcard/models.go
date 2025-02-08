package flashcard

import "go.mongodb.org/mongo-driver/bson/primitive"

type ReviewEntry struct {
	Timestamp int64  `bson:"timestamp" json:"timestamp"` // Time of the review
	Status    string `bson:"status" json:"status"`       // "known" or "unknown"
	Interval  int    `bson:"interval" json:"interval"`   // Interval at the time of review
}

type Flashcard struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Word          string             `bson:"word" json:"word"`
	Translation   string             `bson:"translation" json:"translation"`
	Language      string             `bson:"language" json:"language"`
	UserID        primitive.ObjectID `bson:"user_id" json:"user_id"`
	Interval      int                `bson:"interval" json:"interval"`
	NextReview    int64              `bson:"next_review" json:"next_review"`
	Status        string             `bson:"status" json:"status"`
	ReviewHistory []ReviewEntry      `bson:"review_history" json:"review_history"` // New field for tracking history
}
