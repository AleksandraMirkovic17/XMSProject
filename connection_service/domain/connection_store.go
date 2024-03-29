package domain

import (
	"context"

	pb "github.com/dislinked/common/proto/connection_service"
)

type ConnectionStore interface {
	Register(userID string, username string, isPublic bool) (*pb.ActionResult, error)
	GetFriends(id string) ([]UserConn, error)
	GetBlockeds(userID string) ([]UserConn, error)
	AddFriend(userIDa, userIDb string) (*pb.ActionResult, error)
	AddBlockUser(userIDa, userIDb string) (*pb.ActionResult, error)
	RemoveFriend(userIDa, userIDb string) (*pb.ActionResult, error)
	UnblockUser(userIDa, userIDb string) (*pb.ActionResult, error)
	GetRecommendation(userID string) ([]*UserRecommendation, error)
	Init()
	SendFriendRequest(userIDa, userIDb string) (*pb.ActionResult, error)
	UnsendFriendRequest(userIDa, userIDb string) (*pb.ActionResult, error)
	DeclineFriendRequest(userIDa string, userIDb string) (*pb.ActionResult, error)
	GetConnectionDetail(userIDa, userIDb string) (*pb.ConnectionDetail, error)
	GetFriendRequests(userID string) ([]UserConn, error)
	Update(user UserConn) (*pb.ActionResult, error)
	GetMyContacts(ctx context.Context, request *pb.GetMyContactsRequest) (*pb.ContactsResponse, error)
}
