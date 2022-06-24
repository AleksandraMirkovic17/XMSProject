package application

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type JwtManager struct {
	secretKey string
}

func NewJwtManager() *JwtManager {
	return &JwtManager{
		secretKey: "123456",
	}
}

func (manager *JwtManager) GenerateHashPassword(password string) (string, error) {
	return password, nil
}

func (manager *JwtManager) GenerateJWT(username, role string) (string, error) {
	var mySigningKey = []byte(manager.secretKey)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["username"] = username
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Errorf("something Went Wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}

func (manager *JwtManager) CheckPasswordHash(password, hash string) bool {
	return password == hash
}
