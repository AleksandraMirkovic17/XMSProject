package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dislinked/api_gateway/domain"
	handler "github.com/dislinked/api_gateway/infrastructure/api"
	clients "github.com/dislinked/api_gateway/infrastructure/services"
	"github.com/dislinked/api_gateway/startup/config"
	connectionPb "github.com/dislinked/common/proto/connection_service"
	userPb "github.com/dislinked/common/proto/user_service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"strings"
)

type SearchHandler struct {
	userServiceAddress       string
	conncetionServiceAddress string
	config                   *config.Config
}

func (s SearchHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("GET", "/user/search/{userID}/{params}", s.getUsersBySearch)
	if err != nil {
		panic(err)
	}
}

func (s SearchHandler) getUsersBySearch(w http.ResponseWriter, r *http.Request, params map[string]string) {
	fmt.Println("[ApiGateway]:Searching for")
	searchParams := params["params"]
	userID := params["userID"]
	println("params: " + searchParams + ", searching by user: " + userID)
	if searchParams == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userService := clients.NewUserClient(s.userServiceAddress)
	connectionService := clients.NewConnectionClient(s.conncetionServiceAddress)

	users, err := userService.GetAll(context.TODO(), &userPb.GetUserBySearchParamsRequest{
		Username: searchParams,
		Name:     searchParams,
		Surname:  searchParams,
	})
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		println("get users by search")
		return
	}
	friends, err1 := connectionService.GetFriends(context.TODO(), &connectionPb.GetRequest{UserID: userID})
	if err1 != nil {
		w.WriteHeader(http.StatusNotFound)
		println("get friends by id err:", err1.Error())
		return
	}
	blockeds, err := connectionService.GetBlockeds(context.TODO(), &connectionPb.GetRequest{UserID: userID})
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		println("get blockeds by id err")
		return
	}
	var searchedUsersPublic []domain.User
	println("Number of searched users: ", len(users.Users))
	for _, usearch := range users.Users {
		searchedUsersPublic = append(searchedUsersPublic, mapPbUserToDomainUser(usearch))
	}
	//dodati privatne prijatelje
	for _, friend := range friends.Users {
		if friend.IsPrivate {
			friendUser, err := userService.Get(context.TODO(), &userPb.GetUserRequest{Id: friend.UserID})
			if err != nil {
				println("friend exists only in connection servis ", friend.UserID)
				continue
			}
			if strings.Contains(friendUser.User.Username, searchParams) || strings.Contains(friendUser.User.Name, searchParams) || strings.Contains(friendUser.User.Surname, searchParams) {
				searchedUsersPublic = append(searchedUsersPublic, mapPbUserToDomainUser(friendUser.User))
			}

		}
	}
	//ukloniti blokirane iz pretrage
	for _, blockedUser := range blockeds.Users {
		if !blockedUser.IsPrivate {
			for index, searchedUser := range searchedUsersPublic {
				if searchedUser.Public && searchedUser.Id.Hex() == blockedUser.UserID {
					searchedUsersPublic = append(searchedUsersPublic[:index], searchedUsersPublic[(index+1):]...)
					break
				}

			}
		}
	}

	usersDTO := SearchedUserResponseDTO{
		Users: searchedUsersPublic,
	}

	usersDTOJson, _ := json.Marshal(usersDTO)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(usersDTOJson)
	if err != nil {
		println("writing response error")
		return
	}

}

type SearchedUserResponseDTO struct {
	Users []domain.User
}

func mapPbUserToDomainUser(usearch *userPb.User) domain.User {
	id, _ := primitive.ObjectIDFromHex((*usearch).Id)
	var user = domain.User{
		Id:       id,
		Name:     (*usearch).Name,
		Surname:  (*usearch).Surname,
		Username: (*usearch).Username,
		Public:   (*usearch).Public,
	}
	return user

}

func NewSearchHandlerc(c *config.Config) handler.Handler {
	return &SearchHandler{
		userServiceAddress:       fmt.Sprintf("%s:%s", c.UserHost, c.UserPort),
		conncetionServiceAddress: fmt.Sprintf("%s:%s", c.ConnHost, c.ConnPort),
		config:                   nil,
	}

}
