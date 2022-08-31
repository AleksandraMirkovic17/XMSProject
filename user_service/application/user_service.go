package application

import (
	"UserService/domain"
	"UserService/infrastructure/orchestrators"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type UserService struct {
	store              domain.UserStore
	orchestrator       *orchestrators.UserOrchestrator
	orchestratorUpdate *orchestrators.UpdateUserOrchestrator
}

func NewUserService(store domain.UserStore, orchestrator *orchestrators.UserOrchestrator, orchestratorUpdate *orchestrators.UpdateUserOrchestrator) *UserService {
	return &UserService{
		store:              store,
		orchestrator:       orchestrator,
		orchestratorUpdate: orchestratorUpdate,
	}
}

func (service *UserService) Register(user *domain.User) error {
	// span := tracer.StartSpanFromContext(ctx, "Register-Service")
	// defer span.Finish()
	//
	// newCtx := tracer.ContextWithSpan(context.Background(), span)
	println("Registracija usera user_service.go")
	service.store.Insert(user)
	return nil
}

func (service *UserService) Insert(user *domain.User) error {
	// span := tracer.StartSpanFromContext(ctx, "Register-Service")
	// defer span.Finish()
	//
	// newCtx := tracer.ContextWithSpan(context.Background(), span)
	println("Insertovanje usera user_service.go")
	err := service.store.Insert(user)
	if err != nil {
		println("Greska se desila prilikom inserta")
		return err
	}
	err = service.orchestrator.CreateUser(user)
	if err != nil {
		println("Greska se desila prilikom pozivanja orkestratora 1!")
		return err
	}
	println("Nije se desila greska prilikom pozivanja orkestratora 1")

	/*err = service.orchestrator.CreateConnectionUser(user)
	if err != nil {
		println("Greska se desila prilikom pozivanja orkestratora 2!")
		return err
	}*/
	println("posle publishovanja")

	return err
}
func (service *UserService) Update(uuid primitive.ObjectID, user *domain.User) error {
	service.store.Update(uuid, user)
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
	err = service.store.Update(updatedUser.Id, updatedUser)
	return updatedUser, nil
}

func updateField(paths []string, user *domain.User, requestUser *domain.User) (*domain.User, error) {
	return user, nil
}

func (service *UserService) GetOne(uuid primitive.ObjectID) (*domain.User, error) {
	return service.store.FindByID(uuid)
}

func (service *UserService) GetAll() ([]*domain.User, error) {
	return service.store.GetAll()
}

func (service *UserService) GetAllByUsernameAndNameAndSurname(username string, name string, surname string) ([]*domain.User, error) {
	return service.store.FindByUsernameAndNameAndSurname(username, name, surname)
}

func (service *UserService) Search(searchText string) (*[]domain.User, error) {
	return service.store.Search(searchText)
}

func (service *UserService) FindByUsername(username string) (*domain.User, error) {
	println("Finding by username in user_service.go")
	return service.store.FindByUsername(username)
}

func (service *UserService) Delete(user *domain.User) interface{} {
	return service.store.Delete(user)
}

func (service *UserService) GenerateApiToken(user *domain.User) (string, error) {
	var claims = &APIToken{}
	claims.Username = user.Username
	claims.Method = "ShareJobOffer"

	userRoles := "Agent"

	claims.Roles = append(claims.Roles, userRoles)
	var tokenCreationTime = time.Now().UTC()
	var tokenExpirationTime = tokenCreationTime.Add(time.Duration(720) * time.Hour)

	token, err := generateToken(claims, tokenExpirationTime)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (service *UserService) CheckAccess(token string) (bool, string, error) {
	claims, err := verifyToken(token)
	if err != nil {
		println("Error verifying token!")
		return false, "", err
	}
	if claims.valid() != nil {
		return false, "", claims.valid()
	}

	return true, claims.Username, nil
}

func (token APIToken) valid() error {
	now := time.Now().UTC().Unix()

	if token.VerifyExpiresAt(now, false) == false {
		diff := time.Unix(now, 0).Sub(time.Unix(token.ExpiresAt, 0))
		return fmt.Errorf("token is expired by %v", diff)
	}
	return nil
}

func generateToken(claims *APIToken, expirationTime time.Time) (string, error) {
	claims.ExpiresAt = expirationTime.Unix()
	claims.IssuedAt = time.Now().UTC().Unix()
	claims.Issuer = "Dislinkt"

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	mySigningKey := []byte("secretString")
	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func verifyToken(token string) (*APIToken, error) {
	claims := &APIToken{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("secretString"), nil
	})
	if err != nil {
		fmt.Println("Error parsing claims:", err.Error())
		return nil, err
	}

	return claims, nil
}

type APIToken struct {
	Username string
	Method   string
	Roles    []string
	jwt.StandardClaims
}
