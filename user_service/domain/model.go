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
	Id     uuid.UUID `gorm:"primaryKey"`
	Name   string
	UserId uuid.UUID
}

type EducationExperience struct {
	Id              uuid.UUID `gorm:"primaryKey"`
	InstitutionName string
	Type            EducationType
	StartDate       time.Time
	EndDate         time.Time
	UserId          uuid.UUID
}

type WorkExperience struct {
	Id               uuid.UUID `gorm:"primaryKey"`
	OrganizationName string
	PositionName     string
	StartDate        time.Time
	EndDate          time.Time
	UserId           uuid.UUID
}

type User struct {
	Id                   uuid.UUID `gorm:"primaryKey"`
	Name                 string
	Surname              string
	Username             *string `gorm:"unique"`
	Email                *string `gorm:"unique"`
	Password             string
	Phone                string
	Gender               Gender
	Role                 Role
	DateOfBirth          time.Time
	Public               bool
	Skills               []Skill               `gorm:"foreignKey:UserId"`
	EducationExperiences []EducationExperience `gorm:"foreignKey:UserId"`
	WorkExperience       []WorkExperience      `gorm:"foreignKey:UserId"`
	Biography            string
}
