package application

import (
	"UserService/domain"

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

func (service *UserService) Register(user *domain.User) error {
	// span := tracer.StartSpanFromContext(ctx, "Register-Service")
	// defer span.Finish()
	//
	// newCtx := tracer.ContextWithSpan(context.Background(), span)

	return nil
}

func (service *UserService) Insert(user *domain.User) error {
	// span := tracer.StartSpanFromContext(ctx, "Register-Service")
	// defer span.Finish()
	//
	// newCtx := tracer.ContextWithSpan(context.Background(), span)
	err := service.store.Insert(user)
	return err
}
func (service *UserService) Update(uuid primitive.ObjectID, user *domain.User) error {
	service.store.Update(user)
	return nil
}

func (service *UserService) PatchUser(updatePaths []string, requestUser *domain.User,
	uuid primitive.ObjectID) (*domain.User, error) {
	// span := tracer.StartSpanFromContext(ctx, "Update-Service")
	// defer span.Finish()
	//
	// newCtx := tracer.ContextWithSpan(context.Background(), span)
	foundUser, err := service.store.FindByID(uuid)
	if err != nil {
		return nil, err
	}

	updatedUser, err := updateField(updatePaths, foundUser, requestUser)
	if err != nil {
		return nil, err
	}
	err = service.store.Update(updatedUser)
	return updatedUser, nil
}

func updateField(paths []string, user *domain.User, requestUser *domain.User) (*domain.User, error) {
	return user, nil
}

func (service *UserService) GetAll() ([]*domain.User, error) {
	// span := tracer.StartSpanFromContext(ctx, "GetAll-Service")
	// defer span.Finish()
	//
	// newCtx := tracer.ContextWithSpan(context.Background(), span)
	return service.store.GetAll()
}

func (service *UserService) Search(searchText string) (*[]domain.User, error) {
	return service.store.Search(searchText)
}

func (service *UserService) GetOne(uuid primitive.ObjectID) (*domain.User, error) {
	return service.store.FindByID(uuid)
}

func (service *UserService) FindByUsername(username string) (*domain.User, error) {
	println("Finding by username in user_service.go")
	return service.store.FindByUsername(username)
}

func (service *UserService) Delete(user *domain.User) interface{} {
	return service.store.Delete(user)
}
