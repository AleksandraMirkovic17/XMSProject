package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserStore interface {
	Create(user *User) error
	GetAll() ([]*User, error)
	GetById(id primitive.ObjectID) (*User, error)
	DeleteAll()
	GetByUsername(username string) (*User, error)
}
