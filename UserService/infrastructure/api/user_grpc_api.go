package api

import (
	"UserService/application"
	"UserService/domain"
	"context"
	"fmt"

	pb "github.com/dislinked/common/proto/user_service"
	uuid "github.com/satori/go.uuid"
)

type UserHandler struct {
	pb.UnimplementedUserServiceServer
	service *application.UserService
}

func NewUserHandler(service *application.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (handler *UserHandler) GetAll(ctx context.Context, request *pb.EmptyUser) (*pb.GetAllUserResponse, error) {
	// span := tracer.StartSpanFromContextMetadata(ctx, "GetAllAPI")
	// defer span.Finish()

	// ctx = tracer.ContextWithSpan(context.Background(), span)
	// users, err := handler.service.GetAll(ctx)
	var users *[]domain.User
	var err error
	users, err = handler.service.GetAll()
	if err != nil || *users == nil {
		return nil, err
	}
	response := &pb.GetAllUserResponse{
		Users: []*pb.User{},
	}
	for _, user := range *users {
		current := mapUser(&user)
		response.Users = append(response.Users, current)
	}
	return response, nil
}

func (handler *UserHandler) Get(ctx context.Context, request *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	// span := tracer.StartSpanFromContextMetadata(ctx, "GetAllAPI")
	// defer span.Finish()

	// ctx = tracer.ContextWithSpan(context.Background(), span)
	// users, err := handler.service.GetAll(ctx)
	println(request.Id)
	parsedUUID, err := uuid.FromString(request.Id)
	user, err := handler.service.GetOne(parsedUUID)
	if err != nil || user == nil {
		return nil, err
	}
	response := &pb.GetUserResponse{
		User: mapUser(user),
	}
	return response, nil
}

func (handler *UserHandler) Insert(ctx context.Context, request *pb.RegisterUserRequest) (*pb.User, error) {
	user := mapNewUserPbToDomain(request.User)
	fmt.Println("mapper zavrsio")

	err := handler.service.Insert(user)
	fmt.Println("Register zavrsio")
	if err != nil {
		return nil, err
	}

	return mapUser(user), nil
}

func (handler *UserHandler) Update(ctx context.Context, request *pb.UpdateUserRequest) (*pb.User, error) {
	user := mapNewUserPbToDomain(request.User)
	foundUser, findErr := handler.service.FindByUsername(*((*user).Username))
	if findErr != nil {
		return nil, findErr
	}
	if foundUser == nil {
		return nil, findErr
	}
	user.Id = foundUser.Id

	updateErr := handler.service.Update(user.Id, user)
	if updateErr != nil {
		return nil, updateErr
	}

	return mapUser(user), nil
}

func (handler *UserHandler) SearchProfile(request *pb.SearchProfileRequest) (*pb.GetAllUserResponse, error) {

	users, err := handler.service.Search(request.GetParam())
	if err != nil {
		return nil, err
	}
	response := &pb.GetAllUserResponse{
		Users: []*pb.User{},
	}
	for _, user := range users {
		current := mapUser(user)
		response.Users = append(response.Users, current)
	}
	return response, nil
}
