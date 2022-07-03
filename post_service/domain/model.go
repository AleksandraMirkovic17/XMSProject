package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Comment struct {
	Id      primitive.ObjectID `bson:"_id"`
	Content string             `bson:"content"`
	Date    time.Time          `bson:"date"`
	User    string             `bson:"user"`
}

type Reaction struct {
	User     string       `bson:"user"`
	Reaction ReactionType `bson:"reaction_type"`
}

type ReactionType int

const (
	LIKE = iota
	DISLIKED
)

type Post struct {
	Id        primitive.ObjectID `bson:"_id"`
	User      string             `bson:"user_id"`
	PostText  string             `bson:"post_text"`
	Image     string             `bson:"images"`
	Links     []string           `bson:"links"`
	Date      time.Time          `bson:"date"`
	Reactions []Reaction         `bson:"reactions"`
	Comments  []Comment          `bson:"comments"`
	IsDeleted bool               `bson:"is_deleted"`
}
