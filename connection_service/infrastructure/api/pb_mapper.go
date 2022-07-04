package api

import (
	"ConnectionService/domain"

	pb "github.com/dislinked/common/proto/connection_service"
)

func mapUserConn(userConn *domain.UserConn) *pb.User {
	userConnPb := &pb.User{
		UserID:    userConn.UserID,
		IsPrivate: userConn.IsPrivate,
	}

	return userConnPb
}
