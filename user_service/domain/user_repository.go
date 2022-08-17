package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserStore interface {
	Insert(user *User) error
	Update(uuid primitive.ObjectID, user *User) error
	GetAll() ([]*User, error)
	FindByID(uuid primitive.ObjectID) (*User, error)
	FindByUsername(username string) (*User, error)
	FindByUsernameAndNameAndSurname(username string, name string, surname string) ([]*User, error)
	Search(searchText string) (*[]User, error)
	Delete(user *User) error
}
