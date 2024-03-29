package application

import (
	"ConnectionService/infrastructure/orchestrator"
	"context"
	"fmt"

	"ConnectionService/domain"
	"ConnectionService/startup/config"

	pb "github.com/dislinked/common/proto/connection_service"
	userService "github.com/dislinked/common/proto/user_service"
)

type ConnectionService struct {
	store        domain.ConnectionStore
	UserClient   userService.UserServiceClient
	Orchestrator *orchestrator.ConnectionNotificationOrchestrator
}

func NewConnectionService(store domain.ConnectionStore, c *config.Config, orchestrator *orchestrator.ConnectionNotificationOrchestrator) *ConnectionService {
	return &ConnectionService{
		store:        store,
		Orchestrator: orchestrator,
		UserClient:   NewUserClient(fmt.Sprintf("%s:%s", c.UserHost, c.UserPort)),
	}
}

func (service *ConnectionService) Register(userID string, username string, isPublic bool) (*pb.ActionResult, error) {
	println("Pozivanje registracije!")
	return service.store.Register(userID, username, isPublic)
}

func (service *ConnectionService) GetFriends(id string) ([]*domain.UserConn, error) {

	var friendsRetVal []*domain.UserConn

	friends, err := service.store.GetFriends(id)
	if err != nil {
		return nil, err
	}
	for _, s := range friends {
		friendsRetVal = append(friendsRetVal, &domain.UserConn{UserID: s.UserID, IsPublic: s.IsPublic, Username: s.Username})
	}
	return friendsRetVal, nil
}

func (service *ConnectionService) GetBlockeds(id string) ([]*domain.UserConn, error) {

	var friendsRetVal []*domain.UserConn

	friends, err := service.store.GetBlockeds(id)
	if err != nil {
		return nil, nil
	}
	for _, s := range friends {
		friendsRetVal = append(friendsRetVal, &domain.UserConn{UserID: s.UserID, IsPublic: s.IsPublic})
	}
	return friendsRetVal, nil
}

func (service *ConnectionService) GetFriendRequests(userID string) ([]domain.UserConn, error) {
	return service.store.GetFriendRequests(userID)
}

func (service *ConnectionService) AddFriend(userIDa, userIDb string) (*pb.ActionResult, error) {
	service.UserClient.Get(context.TODO(), &userService.GetUserRequest{Id: userIDa})
	return service.store.AddFriend(userIDa, userIDb)
}

func (service *ConnectionService) AddBlockUser(userIDa, userIDb string) (*pb.ActionResult, error) {
	return service.store.AddBlockUser(userIDa, userIDb)
}

func (service *ConnectionService) RemoveFriend(userIDa, userIDb string) (*pb.ActionResult, error) {
	return service.store.RemoveFriend(userIDa, userIDb)
}

func (service *ConnectionService) UnblockUser(userIDa, userIDb string) (*pb.ActionResult, error) {
	return service.store.UnblockUser(userIDa, userIDb)
}

func (service *ConnectionService) GetRecommendation(userID string) ([]*domain.UserRecommendation, error) {
	return service.store.GetRecommendation(userID)
}

func (service *ConnectionService) SendFriendRequest(userIDa, userIDb string) (*pb.ActionResult, error) {
	service.UserClient.Get(context.TODO(), &userService.GetUserRequest{Id: userIDa})
	return service.store.SendFriendRequest(userIDa, userIDb)
}

func (service *ConnectionService) UnsendFriendRequest(userIDa, userIDb string) (*pb.ActionResult, error) {
	return service.store.UnsendFriendRequest(userIDa, userIDb)
}

func (service *ConnectionService) DeclineFriendRequest(userIDa, userIDb string) (*pb.ActionResult, error) {
	return service.store.DeclineFriendRequest(userIDa, userIDb)
}

func (service *ConnectionService) GetConnectionDetail(userIDa, userIDb string) (*pb.ConnectionDetail, error) {
	return service.store.GetConnectionDetail(userIDa, userIDb)
}

func (service *ConnectionService) UpdateUser(user domain.UserConn) (*pb.ActionResult, error) {
	return service.store.Update(user)
}

func (service *ConnectionService) GetMyContacts(ctx context.Context, request *pb.GetMyContactsRequest) (*pb.ContactsResponse, error) {
	return service.store.GetMyContacts(ctx, request)
}
