package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Gender int

const (
	MALE = iota
	FEMALE
)

type User struct {
	Id          primitive.ObjectID `bson:"_id"`
	Username    string             `bson:"username"`
	Email       string             `bson:"email"`
	Password    string             `bson:"password"`
	Gender      Gender             `bson:"gender"`
	DateOfBirth time.Time          `bson:"date"`
}
