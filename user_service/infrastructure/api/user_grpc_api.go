package api

import (
	"UserService/application"
	"context"
	"fmt"

	pb "github.com/dislinked/common/proto/user_service"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserHandler struct {
	pb.UnimplementedUserServiceServer
	service *application.UserService
}

func NewUserHandler(service *application.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (handler *UserHandler) GetAll(ctx context.Context, request *pb.GetUserBySearchParamsRequest) (*pb.GetAllUserResponse, error) {
	users, err := handler.service.GetAllByUsernameAndNameAndSurname(request.Username, request.Username, request.Username)
	if err != nil {
		return nil, err
	}
	response := &pb.GetAllUserResponse{Users: []*pb.User{}}
	for _, user := range users {
		current := mapUser(user)
		response.Users = append(response.Users, current)
	}
	return response, nil
}

func (handler *UserHandler) Get(ctx context.Context, request *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	println(request.Id)
	userId, _ := primitive.ObjectIDFromHex(request.Id)
	user, err := handler.service.GetOne(userId)
	if err != nil || user == nil {
		return nil, err
	}
	response := &pb.GetUserResponse{
		User: mapUser(user),
	}
	return response, nil
}

func (handler *UserHandler) Insert(ctx context.Context, request *pb.RegisterUserRequest) (*pb.User, error) {
	user := mapUserPbToDomain(request.User)
	fmt.Println("mapper zavrsio")

	err := handler.service.Insert(user)
	fmt.Println("Register zavrsio")
	if err != nil {
		return nil, err
	}

	return mapUser(user), nil
}

func (handler *UserHandler) Update(ctx context.Context, request *pb.UpdateUserRequest) (*pb.User, error) {
	user := mapUserPbToDomain(request.User)
	foundUser, findErr := handler.service.FindByUsername((*user).Username)
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

/*func (handler *UserHandler) SearchProfile(ctx context.Context, request *pb.SearchProfileRequest) (*pb.GetAllUserResponse, error) {

	users, err := handler.service.Search(request.GetParam())
	if err != nil {
		return nil, err
	}
	response := &pb.GetAllUserResponse{
		Users: []*pb.User{},
	}
	println("Name comming")
	for _, user := range *users {
		current := mapUser(&user)
		response.Users = append(response.Users, current)
	}
	return response, nil
}*/

func (handler *UserHandler) FindByUsername(ctx context.Context, request *pb.GetUserByUsernameRequest) (*pb.GetUserResponse, error) {
	print("Searching fo user by username")
	println("Searching for: " + request.Username)
	user, err := handler.service.FindByUsername(request.Username)
	if err != nil || user == nil {
		return nil, err
	}
	response := &pb.GetUserResponse{
		User: mapUser(user),
	}
	return response, nil
}
