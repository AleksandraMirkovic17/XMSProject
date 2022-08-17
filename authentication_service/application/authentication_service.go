package application

import (
	"AuthenticationService/domain"
	orchestrators "AuthenticationService/infrastructure/orchestrator"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthenticationService struct {
	store        domain.UserStore
	orchestrator *orchestrators.RegisterUserOrchestrator
	jwtManager   JwtManager
}

func NewAuthenticationService(store domain.UserStore, orchestrator *orchestrators.RegisterUserOrchestrator) *AuthenticationService {
	return &AuthenticationService{
		store:        store,
		jwtManager:   *NewJwtManager(),
		orchestrator: orchestrator,
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

func (service *AuthenticationService) Register(user *domain.UserAuthentication) (primitive.ObjectID, error) {
	dbUser, _ := service.store.GetByUsername((*user).Username)
	if dbUser != nil && (*dbUser).Username != "" {
		err := errors.New("username already exists")
		println("Desio se error: username already exists")
		return primitive.ObjectID{}, err
	}
	_, err := service.store.Create(user)
	if err != nil {
		println("desio se neki error")
		return primitive.ObjectID{}, err
	}
	println("Sve je ok proslo u metodu register")
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

func (service *AuthenticationService) GetOne(id primitive.ObjectID) (*domain.UserAuthentication, error) {
	return service.store.GetById(id)

}

func (service *AuthenticationService) Update(user *domain.UserAuthentication) error {
	return service.store.Update(user)

}
