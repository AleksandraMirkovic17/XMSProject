package domain

import uuid "github.com/jackc/pgtype/ext/gofrs-uuid"

type User struct {
	ID       uuid.UUID `gorm:"index:idx_name,unique"`
	Username string    `gorm:"unique"`
	Password string
	Role     string
}
