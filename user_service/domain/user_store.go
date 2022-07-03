package domain

import (
	uuid "github.com/satori/go.uuid"
)

type UserStore interface {
	Insert(user *User) error
	Update(user *User) error
	GetAll() (*[]User, error)
	FindByID(uuid uuid.UUID) (*User, error)
	FindByUsername(username string) (*User, error)
	Search(searchText string) (*[]User, error)
	Delete(user *User) error
	AddWorkExperience(experience *WorkExperience, uuid uuid.UUID) error
	AddEducationExperience(education *EducationExperience, uuuid uuid.UUID) error
	AddSkill(skill string, uuid uuid.UUID) error
	RemoveSkill(skill string, uuid uuid.UUID) error
	AddInterest(companyId uuid.UUID, userId uuid.UUID) error
	DeleteWorkExperience(experienceId uuid.UUID, userId uuid.UUID) error
	DeleteEducationExperience(educationId uuid.UUID, userId uuid.UUID) error
	RemoveInterest(companyId uuid.UUID, userId uuid.UUID) error
}
