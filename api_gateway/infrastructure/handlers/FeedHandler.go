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
	postPb "github.com/dislinked/common/proto/post_service"
	userPb "github.com/dislinked/common/proto/user_service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

type FeedHandler struct {
	userServiceAddress       string
	postServiceAddress       string
	connectionServiceAddress string
	config                   *config.Config
}

func (f FeedHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("GET", "/user/{username}/feed", f.GetFeedPostsForUser)
	if err != nil {
		panic(err)
	}
}

func (f FeedHandler) GetFeedPostsForUser(w http.ResponseWriter, r *http.Request, params map[string]string) {
	username := params["username"]
	if username == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	userService := clients.NewUserClient(f.userServiceAddress)
	postsService := clients.NewPostClient(f.postServiceAddress)
	connectionService := clients.NewConnectionClient(f.connectionServiceAddress)

	user, err := userService.FindByUsername(context.TODO(), &userPb.GetUserByUsernameRequest{Username: username})
	if err != nil {
		println("find by username error")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	println("user id is", user.User.Id)
	id := user.User.Id
	connections, err := connectionService.GetFriends(context.TODO(), &connectionPb.GetRequest{UserID: id})
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		println("get friend error")
		return
	}

	var posts []domain.Post
	println("Number of friends: ", len(connections.Users))
	for _, friend := range connections.Users {
		println("friend user id:", friend.UserID, friend.Username)
		postsPbs, err := postsService.GetByUser(context.TODO(), &postPb.GetByUserRequest{Uuid: friend.UserID})
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			println("Get by user")
			return
		}
		println("Ukupno postova prijatelja ", len(postsPbs.Posts))
		posts = append(posts, mapPbsToDomain(postsPbs.Posts, friend.Username)...)
	}

	println("Ukupno postova:", len(posts))

	postsDto := FeedPostsResponseDto{
		Feed: posts,
	}

	postsDtoJson, _ := json.Marshal(postsDto)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(postsDtoJson)
	if err != nil {
		println("writing response error")
		return
	}
}

type FeedPostsResponseDto struct {
	Feed []domain.Post
}

func mapPbsToDomain(postsPb []*postPb.Post, username string) []domain.Post {
	var posts []domain.Post

	for _, postPb := range postsPb {
		id, _ := primitive.ObjectIDFromHex(postPb.Id)
		var post = &domain.Post{
			Id:        id,
			User:      username,
			UserID:    postPb.User,
			PostText:  postPb.Posttext,
			Image:     postPb.Image,
			Date:      *postPb.Date,
			IsDeleted: postPb.Isdeleted,
		}
		for _, link := range postPb.Links {
			post.Links = append(post.Links, link)
		}
		for _, reaction := range postPb.Reactions {
			postId, _ := primitive.ObjectIDFromHex(reaction.PostId)
			post.Reactions = append(post.Reactions, &domain.Reaction{
				Id:       postId,
				User:     reaction.Username,
				Reaction: mapReactionTypeFromPbToDomain(reaction.ReactionType),
			})
		}
		for _, comment := range postPb.Comments {
			commentId, _ := primitive.ObjectIDFromHex(comment.Id)
			post.Comments = append(post.Comments, &domain.Comment{
				Id:      commentId,
				User:    comment.Username,
				Content: comment.Content,
				Date:    *comment.Date,
			})
		}
		posts = append(posts, *post)

	}
	return posts
}

func mapReactionTypeFromPbToDomain(reactionType postPb.Reaction_ReactionType) domain.ReactionType {
	switch reactionType {
	case postPb.Reaction_LIKE:
		return domain.LIKE
	case postPb.Reaction_DISLIKE:
		return domain.DISLIKED
	}
	return domain.LIKE

}

func NewFeedHandler(c *config.Config) handler.Handler {
	return &FeedHandler{
		config:                   c,
		userServiceAddress:       fmt.Sprintf("%s:%s", c.UserHost, c.UserPort),
		postServiceAddress:       fmt.Sprintf("%s:%s", c.PostHost, c.PostPort),
		connectionServiceAddress: fmt.Sprintf("%s:%s", c.ConnHost, c.ConnPort),
	}
}
