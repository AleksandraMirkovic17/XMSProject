package api

import (
	"ConnectionService/domain"

	pb "github.com/dislinked/common/proto/connection_service"
)

func mapUserConn(userConn *domain.UserConn) *pb.User {
	userConnPb := &pb.User{
		UserID:    userConn.UserID,
		IsPrivate: !userConn.IsPublic,
	}

	return userConnPb
}

func mapDomainUserToPbRecommendedUser(user *domain.UserRecommendation) *pb.RecommendedUser {
	recommendedUser := &pb.RecommendedUser{
		UserID:    user.UserID,
		IsPrivate: !user.IsPublic,
		Username:  user.Username,
		IsMutual:  user.IsMutual,
		Mutual:    int32(user.Mutual),
	}
	return recommendedUser

}
