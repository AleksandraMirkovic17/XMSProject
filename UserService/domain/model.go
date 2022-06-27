package domain

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Role int

const (
	Regular Role = iota
	Admin
	Agent
)

type Gender int

const (
	Other = iota
	MALE
	FEMALE
)

type User struct {
	Id          uuid.UUID `gorm:"index:idx_name,unique"`
	Name        string    `bson:"name"`
	Surname     string    `bson:"surname"`
	Username    string    `bson:"username"`
	Email       string    `bson:"email"`
	Password    string    `bson:"password"`
	Phone       string    `bson:"phone"`
	Gender      Gender    `bson:"gender"`
	Role        Role      `bson:"role"`
	DateOfBirth time.Time `bson:"date"`
	Public      bool      `bson:"isPublic"`
}
