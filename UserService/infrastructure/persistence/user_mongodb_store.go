package persistence

import (
	"UserService/domain"
	"strings"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type UserPostgresStore struct {
	db *gorm.DB
}

func (store *UserPostgresStore) AddWorkExperience(experience *domain.WorkExperience, uuid uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (store *UserPostgresStore) AddEducationExperience(education *domain.EducationExperience, uuuid uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (store *UserPostgresStore) AddSkill(skill string, uuid uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (store *UserPostgresStore) RemoveSkill(skill string, uuid uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (store *UserPostgresStore) AddInterest(companyId uuid.UUID, userId uuid.UUID) error {
	
}

func (store *UserPostgresStore) DeleteWorkExperience(experienceId uuid.UUID, userId uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (store *UserPostgresStore) DeleteEducationExperience(educationId uuid.UUID, userId uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (store *UserPostgresStore) RemoveInterest(companyId uuid.UUID, userId uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func NewUserPostgresStore(db *gorm.DB) (domain.UserStore, error) {
	err := db.AutoMigrate(&domain.User{}, &domain.WorkExperience{}, &domain.EducationExperience{}, &domain.LanguageSkill{})
	if err != nil {
		return nil, err
	}
	return &UserPostgresStore{db: db}, nil
}

func (store *UserPostgresStore) Insert(user *domain.User) error {
	// span := tracer.StartSpanFromContext(ctx, "Register-DB")
	// defer span.Finish()
	result := store.db.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (store *UserPostgresStore) Update(user *domain.User) error {
	// span := tracer.StartSpanFromContext(ctx, "Update-DB")
	// defer span.Finish()
	if result := store.db.Save(&user); result.Error != nil {
		return result.Error
	}
	return nil
}

func (store *UserPostgresStore) GetAll() (*[]domain.User, error) {
	// span := tracer.StartSpanFromContext(ctx, "GetAll-DB")
	// defer span.Finish()
	var users []domain.User
	result := store.db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return &users, nil
}

func (store *UserPostgresStore) FindByID(uuid uuid.UUID) (user *domain.User, err error) {
	foundUser := domain.User{}
	if result := store.db.First(&foundUser, uuid); result.Error != nil {
		return nil, result.Error
	}
	return &foundUser, nil
}

func (store *UserPostgresStore) FindByUsername(username string) (user *domain.User, err error) {
	foundUser := domain.User{}
	if result := store.db.Model(domain.User{Username: &username}).First(&foundUser); result.Error != nil {
		return nil, result.Error
	}
	return &foundUser, nil
}

func (store *UserPostgresStore) Search(searchText string) (*[]domain.User, error) {
	var users *[]domain.User
	args := strings.TrimSpace(searchText)
	splitArgs := strings.Split(args, " ")
	allUsers, _ := store.GetAll()
	for _, cuser := range *allUsers {
		nameSearch := strings.ToLower(cuser.Name)
		surnameSearch := strings.ToLower(cuser.Surname)
		usernameSearch := strings.ToLower(*cuser.Username)
		for _, searchParam := range splitArgs {
			searchParamLower := strings.ToLower(searchParam)
			if !(strings.Contains(nameSearch, searchParamLower) || strings.Contains(surnameSearch, searchParamLower) || strings.Contains(usernameSearch, searchParamLower)) {
				break
			}
			*users = append(*users, cuser)
		}
	}

	return users, nil
}

func (store *UserPostgresStore) Delete(user *domain.User) error {
	result := store.db.Delete(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
