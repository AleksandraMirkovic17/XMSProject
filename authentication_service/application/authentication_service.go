package application

import (
	"AuthenticationService/domain"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthenticationService struct {
	store      domain.UserStore
	jwtManager JwtManager
}

func NewAuthenticationService(store domain.UserStore) *AuthenticationService {
	return &AuthenticationService{
		store:      store,
		jwtManager: *NewJwtManager(),
	}
}

func (service *AuthenticationService) Login(credentials *domain.Credentials) (*domain.Token, error) {
	dbUser, _ := service.store.GetByUsername((*credentials).Username)
	if (*dbUser).Username == "" {
		err := errors.New("bad credentials")
		return nil, err
	}

	isPasswordCorrect := service.jwtManager.CheckPasswordHash((*credentials).Password, (*dbUser).Password)
	if !isPasswordCorrect {
		err := errors.New("bad credentials")
		return nil, err
	}

	validToken, err := service.jwtManager.GenerateJWT((*dbUser).Username, (*dbUser).Role)
	if err != nil {
		err := errors.New("failed to generate token")
		return nil, err
	}

	var token domain.Token
	token.Username = (*dbUser).Username
	token.Role = (*dbUser).Role
	token.TokenString = validToken

	return &token, nil
}

func (service *AuthenticationService) Register(user *domain.User) (primitive.ObjectID, error) {
	dbUser, _ := service.store.GetByUsername((*user).Username)
	if (*dbUser).Username != "" {
		err := errors.New("username already exists")
		return primitive.ObjectID{}, err
	}
	_, err := service.store.Create(user)
	if err != nil {
		return primitive.ObjectID{}, err
	}
	return user.ID, err
}

func (service *AuthenticationService) DeleteById(id primitive.ObjectID) error {
	err := service.store.DeleteById(id)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func (service *AuthenticationService) IsAuthorized(token *domain.Token) {
	//service.store.Create()
}
