package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type PostStore interface {
	Get(id primitive.ObjectID) (*Post, error)
	GetAll() ([]*Post, error)
	Insert(post *Post) error
	GetAllByUser(uuid string) ([]*Post, error)
	GetAllByConnections(uuids []string) ([]*Post, error)
	CreateComment(post *Post, comment *Comment) error
	ReactToPost(post *Post, username string, reaction *Reaction) error
	DeleteAll()
}
