package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
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
	Id          primitive.ObjectID `bson:"_id"`
	Name        string             `bson:"name"`
	Surname     string             `bson:"surname"`
	Username    string             `bson:"username"`
	Email       string             `bson:"email"`
	Password    string             `bson:"password"`
	Gender      Gender             `bson:"gender"`
	Role        Role               `bson:"role"`
	DateOfBirth time.Time          `bson:"date"`
}
