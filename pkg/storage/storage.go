package storage

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// DumpsterData
type DumpsterData interface {
	SaveObject(context.Context, *Payload) error
	DeleteObject(context.Context, string) error
	ListObjects(context.Context) ([]*Payload, error)
	GetByID(context.Context, string) (*Payload, error)
	GetByTwID(context.Context, string) (*Payload, error)
}

// Payload - Dumpster Payload Object
type Payload struct {
	ID             primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Date           time.Time          `json:"date" bson:"date"`
	Text           []byte             `json:"text" bson:"text"`
	TwID           string             `json:"twitter_id" bson:"twitter_id"`
	TwDate         time.Time          `json:"twitter_date" bson:"twitter_date"`
	UserName       string             `json:"twitter_user_name" bson:"twitter_user_name"`
	UserScreenName string             `json:"twitter_user_screen_name" bson:"twitter_user_screen"`
	TextSentiment  int                `json:"text_sentiment,omitempty" bson:"text_sentiment,omitempty"`
}

// Response - Result of Querying
type Response struct {
	ID      primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Date    time.Time          `json:"date" bson:"date"`
	Payload Payload            `json:"payload" bson:"payload"`
}
