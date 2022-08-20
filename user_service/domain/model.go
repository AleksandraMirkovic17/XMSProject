package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Role int

const (
	Regular Role = iota
	Admin
	Agent
)

///////////
// ENUMS //
///////////

type Gender int

const (
	Other = iota
	MALE
	FEMALE
)

type EducationType int

const (
	PRIMARY_EDUCATION = iota
	SECONDARY_EDUCATION
	COLLEGE_EDUCATION
)

/////////////
// CLASSES //
/////////////

type Skill struct {
	Id   primitive.ObjectID `bson:"_id"`
	Name string             `bson:"name"`
}

type Interest struct {
	Id   primitive.ObjectID `bson:"_id"`
	Name string             `bson:"name"`
}

type EducationExperience struct {
	Id              primitive.ObjectID `bson:"_id"`
	InstitutionName string             `bson:"institution_name"`
	Type            EducationType      `bson:"type"`
	StartDate       time.Time          `bson:"start_date"`
	EndDate         time.Time          `bson:"end_date"`
}

type WorkExperience struct {
	Id               primitive.ObjectID `bson:"_id"`
	OrganizationName string             `bson:"organization_name"`
	PositionName     string             `bson:"position_name"`
	StartDate        time.Time          `bson:"start_date"`
	EndDate          time.Time          `bson:"end_date"`
}

type User struct {
	Id                   primitive.ObjectID    `bson:"_id"`
	Name                 string                `bson:"name"`
	Surname              string                `bson:"surname"`
	Username             string                `bson:"username"`
	Email                string                `bson:"email"`
	Password             string                `bson:"password"`
	Phone                string                `bson:"phone"`
	Gender               Gender                `bson:"gender"`
	Role                 Role                  `bson:"role"`
	DateOfBirth          string                `bson:"date_of_birth"`
	Public               bool                  `bson:"public"`
	Skills               []Skill               `bson:"skills"`
	Interests            []Interest            `bson:"interests"`
	EducationExperiences []EducationExperience `bson:"education_experience"`
	WorkExperiences      []WorkExperience      `bson:"work_experience"`
	Biography            string                `bson:"biography"`
}
