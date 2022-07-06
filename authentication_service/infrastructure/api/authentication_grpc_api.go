package api

import (
	"AuthenticationService/application"
	"context"
	"net/http"
	"time"

	pb "github.com/dislinked/common/proto/authentication_service"
	events "github.com/dislinked/common/saga/create_order"
	"google.golang.org/grpc/status"
)

type AuthenticationHandler struct {
	pb.UnimplementedAuthenticationServiceServer
	service                  *application.AuthenticationService
	RegisterUserOrchestrator *application.RegisterUserOrchestrator
}

func NewAuthenticationHandler(service *application.AuthenticationService) *AuthenticationHandler {
	return &AuthenticationHandler{
		service: service,
	}
}

func (handler *AuthenticationHandler) Login(ctx context.Context, request *pb.LoginRequest) (*pb.Token, error) {
	credentials := mapCredentialsToDomain(request.Credentials)
	token, err := handler.service.Login(credentials)

	if err != nil {
		return nil, status.Error(401, "Bad credentials!")
	}

	tokenPB := mapTokenToPB(token)
	return tokenPB, nil
}

func (handler *AuthenticationHandler) Register(ctx context.Context, request *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	user := mapUserToDomain(request.User)
	_, err := handler.service.Register(user) //registruje se prosledjeni user
	if err != nil {
		return nil, status.Error(http.StatusBadRequest, "Username already exists!")
	}
	dateOfBirth, _ := time.Parse("2006-01-02T15:04", request.User.DateOfBirth)
	handler.RegisterUserOrchestrator.Start(events.UserDetails{
		Id:          request.User.Id,
		Name:        request.User.Name,
		Surname:     request.User.Surname,
		Username:    request.User.Username,
		Password:    request.User.Username,
		Email:       request.User.Email,
		Birthday:    dateOfBirth,
		Gender:      mapAuthGenderToCreateOrderGender(request.User.Gender),
		Role:        mapAuthRoleToCreateOrderRole(request.User.Role),
		PhoneNumber: request.User.ContactPhone,
		IsPublic:    request.User.Public,
	})

	response := &pb.RegisterResponse{
		Username: request.User.Username,
	}
	return response, nil
}

func (handler *AuthenticationHandler) IsAuthorized(ctx context.Context, request *pb.AuthorizationRequest) (*pb.AuthorizationResponse, error) {

	return nil, nil
}
