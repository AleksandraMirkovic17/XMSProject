package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type PostStore interface {
	Get(id primitive.ObjectID) (*Post, error)
	GetAll() ([]*Post, error)
	Insert(post *Post) error
	Update(post *Post) error
	GetAllByUser(uuid string) ([]*Post, error)
	GetAllByConnections(uuids []string) ([]*Post, error)
	CreateComment(post *Post, comment *Comment) error
	ReactToPost(post *Post, username string, reaction ReactionType) (*Post, error)
	DeleteReaction(post *Post, username string) (*Post, error)
	GetAllLikes(id primitive.ObjectID) ([]*ReactionDetailsDTO, error)
	GetAllDislikes(id primitive.ObjectID) ([]*ReactionDetailsDTO, error)
	DeleteAll()
}
