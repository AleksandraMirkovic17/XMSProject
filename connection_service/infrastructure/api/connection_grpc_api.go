package api

import (
	"ConnectionService/application"
	"context"
	"fmt"
	pb "github.com/dislinked/common/proto/connection_service"
	events "github.com/dislinked/common/saga/connection_notification"
)

type ConnectionHandler struct {
	pb.UnimplementedConnectionServiceServer
	service *application.ConnectionService
}

func NewConnectionHandler(service *application.ConnectionService) *ConnectionHandler {
	return &ConnectionHandler{
		UnimplementedConnectionServiceServer: pb.UnimplementedConnectionServiceServer{},
		service:                              service,
	}
}

func (handler *ConnectionHandler) Register(ctx context.Context, request *pb.RegisterRequest) (*pb.ActionResult, error) {
	fmt.Println("[ConnectionHandler]:Register")
	userID := request.User.UserID
	isPublic := !request.User.IsPrivate
	username := request.User.Username
	return handler.service.Register(userID, username, isPublic)
}

func (handler *ConnectionHandler) GetFriends(ctx context.Context, request *pb.GetRequest) (*pb.Users, error) {

	fmt.Println("[ConnectionHandler]:GetFriends")

	id := request.UserID
	friends, err := handler.service.GetFriends(id)
	if err != nil {
		return nil, err
	}
	response := &pb.Users{}
	for _, user := range friends {
		fmt.Println("User", id, "is friend with", user.UserID, user.Username)
		response.Users = append(response.Users, mapUserConn(user))
	}
	fmt.Println("Returning response")
	return response, nil
}

func (handler *ConnectionHandler) GetBlockeds(ctx context.Context, request *pb.GetRequest) (*pb.Users, error) {
	fmt.Println("[ConnectionHandler]:GetBlockeds")

	id := request.UserID
	blockedUsers, err := handler.service.GetBlockeds(id)
	if err != nil {
		return nil, err
	}
	response := &pb.Users{}
	for _, user := range blockedUsers {
		fmt.Println("User", id, "block", user.UserID)
		response.Users = append(response.Users, mapUserConn(user))
	}
	return response, nil
}

func (handler *ConnectionHandler) GetFriendRequests(ctx context.Context, request *pb.GetRequest) (*pb.Users, error) {
	fmt.Println("[ConnectionHandler]:GetFriendRequests")

	id := request.UserID
	friendRequests, err := handler.service.GetFriendRequests(id)
	if err != nil {
		return nil, err
	}
	response := &pb.Users{}
	for _, user := range friendRequests {
		fmt.Println("User", id, "friend reqest to", user.UserID)
		response.Users = append(response.Users, mapUserConn(&user))
	}
	return response, nil
}

func (handler *ConnectionHandler) AddFriend(ctx context.Context, request *pb.AddFriendRequest) (*pb.ActionResult, error) {
	fmt.Println("[ConnectionHandler]:AddFriend")
	userIDa := request.AddFriendDTO.UserIDa
	userIDb := request.AddFriendDTO.UserIDb

	actionResult, err := handler.service.AddFriend(userIDa, userIDb)
	if actionResult.Status == 201 {
		handler.service.Orchestrator.Start(&events.ConnectionNotification{
			Content:    "",
			SenderId:   userIDa,
			ReceiverId: userIDb,
			Sender:     "",
			Receiver:   "",
			Request:    false,
		})
	}
	return actionResult, err
}
func (handler *ConnectionHandler) AddBlockUser(ctx context.Context, request *pb.AddBlockUserRequest) (*pb.ActionResult, error) {
	fmt.Println("[ConnectionHandler]:AddBlockUser")
	userIDa := request.AddBlockUserDTO.UserIDa
	userIDb := request.AddBlockUserDTO.UserIDb
	return handler.service.AddBlockUser(userIDa, userIDb)
}

func (handler *ConnectionHandler) RemoveFriend(ctx context.Context, request *pb.RemoveFriendRequest) (*pb.ActionResult, error) {
	fmt.Println("[ConnectionHandler]:RemoveFriend")
	userIDa := request.RemoveFriendDTO.UserIDa
	userIDb := request.RemoveFriendDTO.UserIDb
	return handler.service.RemoveFriend(userIDa, userIDb)
}

func (handler *ConnectionHandler) UnblockUser(ctx context.Context, request *pb.UnblockUserRequest) (*pb.ActionResult, error) {
	fmt.Println("[ConnectionHandler]:UnblockUser")
	userIDa := request.UnblockUserDTO.UserIDa
	userIDb := request.UnblockUserDTO.UserIDb
	return handler.service.UnblockUser(userIDa, userIDb)
}

func (handler *ConnectionHandler) GetRecommendation(ctx context.Context, request *pb.GetRequest) (*pb.RecommendedUsers, error) {
	fmt.Println("[ConnectionHandler]:GetRecommendation")

	id := request.UserID
	recommendation, err := handler.service.GetRecommendation(id)
	if err != nil {
		return nil, err
	}
	response := &pb.RecommendedUsers{}
	for _, user := range recommendation {
		response.Users = append(response.Users, mapDomainUserToPbRecommendedUser(user))
	}
	return response, nil
}

func (handler *ConnectionHandler) SendFriendRequest(ctx context.Context, request *pb.SendFriendRequestRequest) (*pb.ActionResult, error) {
	fmt.Println("[ConnectionHandler]:SendFriendRequest")
	userIDa := request.SendFriendRequestRequestDTO.UserIDa
	userIDb := request.SendFriendRequestRequestDTO.UserIDb

	actionResult, err := handler.service.SendFriendRequest(userIDa, userIDb)
	println("action result")
	println(actionResult.Status)
	if actionResult.Status == 201 {
		handler.service.Orchestrator.Start(&events.ConnectionNotification{
			Content:    "",
			SenderId:   userIDa,
			ReceiverId: userIDb,
			Sender:     "",
			Receiver:   "",
			Request:    true,
		})

	}
	return actionResult, err
}
func (handler *ConnectionHandler) UnsendFriendRequest(ctx context.Context, request *pb.UnsendFriendRequestRequest) (*pb.ActionResult, error) {
	fmt.Println("[ConnectionHandler]:UnsendFriendRequest")
	userIDa := request.UnsendFriendRequestRequestDTO.UserIDa
	userIDb := request.UnsendFriendRequestRequestDTO.UserIDb
	return handler.service.UnsendFriendRequest(userIDa, userIDb)
}

func (handler *ConnectionHandler) DeclineFriendRequest(ctx context.Context, request *pb.DeclineFriendRequestRequest) (*pb.ActionResult, error) {
	fmt.Println("[ConnectionHandler]:DeclineFriendRequest")
	userIDa := request.DeclineRequestDTO.UserIDa
	userIDb := request.DeclineRequestDTO.UserIDb
	return handler.service.DeclineFriendRequest(userIDa, userIDb)
}

func (handler *ConnectionHandler) GetConnectionDetail(ctx context.Context, request *pb.GetConnectionDetailRequest) (*pb.ConnectionDetail, error) {
	fmt.Println("[ConnectionHandler]:GetConnectionDetail")
	userIDa := request.UserIDa
	userIDb := request.UserIDb
	return handler.service.GetConnectionDetail(userIDa, userIDb)
}

func (handler *ConnectionHandler) ChangePrivacy(ctx context.Context, request *pb.ChangePrivacyRequest) (*pb.ActionResult, error) {
	panic("Not implemented")
}

func (handler *ConnectionHandler) GetMyContacts(ctx context.Context, request *pb.GetMyContactsRequest) (*pb.ContactsResponse, error) {
	return handler.service.GetMyContacts(ctx, request)
}
