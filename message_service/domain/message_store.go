package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MessageStore interface {
	Get(id primitive.ObjectID) (*Message, error)
	GetAll() ([]*Message, error)
	Insert(message *Message) error
	Update(message *Message) error
	GetAllByUser(uuid string) ([]*Message, error)
	DeleteAll()
}
