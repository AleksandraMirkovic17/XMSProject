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
	Id          uuid.UUID `gorm:"primaryKey"`
	Name        string
	Surname     string
	Username    *string `gorm:"unique"`
	Email       *string `gorm:"unique"`
	Password    string
	Phone       string
	Gender      Gender
	Role        Role
	DateOfBirth time.Time
	Public      bool
}
