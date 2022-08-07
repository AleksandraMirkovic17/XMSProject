package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type ReactionDetailsDTO struct {
	Ahuthor *Author
	PostId  primitive.ObjectID
}

type Author struct {
	UserId   primitive.ObjectID
	Username string
	Name     string
	Surname  string
}

type FeedPostsResponseDto struct {
	Feed []Post
}
