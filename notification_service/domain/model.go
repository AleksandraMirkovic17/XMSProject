package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Notification struct {
	Id      primitive.ObjectID `bson:"_id"`
	User    string             `bson:"user"`
	Content string             `bson:"content"`
	Url     string             `bson:"url"`
	Seen    bool               `bson:"seen"`
	Date    time.Time          `bson:"date"`
}
