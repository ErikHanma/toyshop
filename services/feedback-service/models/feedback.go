package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Feedback struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	UserID    string             `bson:"user_id"`
	Email     string             `bson:"email"`
	Message   string             `bson:"message"`
	CreatedAt time.Time          `bson:"created_at"`
}
