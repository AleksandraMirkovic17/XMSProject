package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Message struct {
	Id       primitive.ObjectID `bson:"_id"`
	Content  string             `bson:"content"`
	Date     time.Time          `bson:"date"`
	FromUser string             `bson:"user"`
	ToUser   string             `bson:"user"`
}
