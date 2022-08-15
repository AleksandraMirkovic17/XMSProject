package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserStore interface {
	Create(user *UserAuthentication) (primitive.ObjectID, error)
	GetAll() ([]*UserAuthentication, error)
	GetById(id primitive.ObjectID) (*UserAuthentication, error)
	DeleteById(id primitive.ObjectID) error
	DeleteAll()
	GetByUsername(username string) (*UserAuthentication, error)
}
