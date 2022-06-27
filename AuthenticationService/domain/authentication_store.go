package domain

import uuid "github.com/jackc/pgtype/ext/gofrs-uuid"

type UserStore interface {
	Create(user *User) (*User, error)
	GetAll() (*[]User, error)
	GetById(id uuid.UUID) (*User, error)
	DeleteAll()
	GetByUsername(username string) (*User, error)
}
