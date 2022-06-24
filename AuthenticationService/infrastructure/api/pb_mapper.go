package api

import "AuthenticationService/domain"
import pb "github.com/dislinked/common/proto/authentication_service"

func mapUserToPB(user *domain.User) *pb.User {
	userPb := &pb.User{
		Username: user.Username,
		Password: user.Password,
		Role:     user.Role,
	}
	return userPb
}

func mapUserToDomain(user *pb.User) *domain.User {
	userDomain := &domain.User{
		Username: (*user).Username,
		Password: (*user).Password,
		Role:     (*user).Role,
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
