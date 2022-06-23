package application

import (
	"UserService/domain"
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService struct {
	store domain.UserStore
}

func NewUserService(store domain.UserStore) *UserService {
	return &UserService{
		store: store,
	}
}

func (service *UserService) Get(id primitive.ObjectID) (*domain.User, error) {
	return service.store.Get(id)
}

func (service *UserService) GetAll() ([]*domain.User, error) {
	return service.store.GetAll()
}

func (service *UserService) Insert(post *domain.User) (*domain.User, error) {
	(*post).Id = primitive.NewObjectID()
	err := service.store.Insert(post)
	if err != nil {
		err := errors.New("Error occured during creating new user")
		return nil, err
	}

	return post, nil
}

func (service *UserService) DeleteAll() {
	service.store.DeleteAll()
}
