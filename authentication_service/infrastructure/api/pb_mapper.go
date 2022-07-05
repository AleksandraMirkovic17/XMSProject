package api

import (
	"AuthenticationService/domain"

	pb "github.com/dislinked/common/proto/authentication_service"
)

func mapUserToPB(user *domain.User) *pb.User {
	userPb := &pb.User{
		Username: user.Username,
		Password: user.Password,
		Role:     mapRoleToPb(user.Role),
	}
	return userPb
}

func mapUserToDomain(user *pb.User) *domain.User {
	userDomain := &domain.User{
		Username: (*user).Username,
		Password: (*user).Password,
		Role:     (*user).Role.String(),
	}
	return userDomain
}

func mapTokenToPB(token *domain.Token) *pb.Token {
	tokenPb := &pb.Token{
		Role:     token.Role,
		Username: token.Username,
		Token:    token.TokenString,
	}
	return tokenPb
}

func mapCredentialsToDomain(credentials *pb.Credentials) *domain.Credentials {
	credentialsDomain := &domain.Credentials{
		Username: (*credentials).Username,
		Password: (*credentials).Password,
	}
	return credentialsDomain
}

func mapRoleToPb(role string) pb.UserRole {
	switch role {
	case "REGULAR":
		return pb.UserRole_Regular
	case "AGENT":
		return pb.UserRole_Agent
	case "ADMIN":
		return pb.UserRole_Admin
	}
	return pb.UserRole_Regular
}
