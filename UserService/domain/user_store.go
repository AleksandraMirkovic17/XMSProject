package domain

import (
	"go.mongodb.org/mongo-driver/x/mongo/driver/uuid"
)

type UserStore interface {
	Insert(user *User) error
	Update(user *User) error
	GetAll() (*[]User, error)
	FindByID(uuid uuid.UUID) (*User, error)
	FindByUsername(username string) (*User, error)
	Search(searchText string) (*[]User, error)
	Delete(user *User) error
	DeleteAll() error
}
