package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dislinked/api_gateway/domain"
	handler "github.com/dislinked/api_gateway/infrastructure/api"
	clients "github.com/dislinked/api_gateway/infrastructure/services"
	"github.com/dislinked/api_gateway/startup/config"
	postPb "github.com/dislinked/common/proto/post_service"
	userPb "github.com/dislinked/common/proto/user_service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"strings"
)

type PostHandler struct {
	userServiceAddress string
	postServiceAddress string
	config             *config.Config
}

func NewPostHandler(c *config.Config) handler.Handler {
	return &PostHandler{
		userServiceAddress: fmt.Sprintf("%s:%s", c.UserHost, c.UserPort),
		postServiceAddress: fmt.Sprintf("%s:%s", c.PostHost, c.PostPort),
		config:             c,
	}

}

func (p PostHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("GET", "/post/username/{idPost}", p.GetPostWithUsername)
	if err != nil {
		panic(err)
	}

}

func (p PostHandler) GetPostWithUsername(w http.ResponseWriter, r *http.Request, params map[string]string) {
	postId := params["idPost"]
	println("id post ", postId)
	if postId == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userService := clients.NewUserClient(p.userServiceAddress)
	postService := clients.NewPostClient(p.postServiceAddress)

	post, err := postService.Get(context.TODO(), &postPb.GetRequest{Id: postId})
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		println("get post by id error")
		return
	}
	println("UserAuthentication id is ", strings.Split(post.Post.User, "\"")[1])
	user, err := userService.Get(context.TODO(), &userPb.GetUserRequest{Id: strings.Split(post.Post.User, "\"")[1]})
	if err != nil {
		println("find by id user error")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	var postMapped = mapPbPostToDomainPost(post.Post, user.User)
	postJSON, _ := json.Marshal(postMapped)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(postJSON)
	if err != nil {
		println("writing response error")
		return
	}

}

func mapPbPostToDomainPost(postPb *postPb.Post, userPb *userPb.User) domain.Post {
	id, _ := primitive.ObjectIDFromHex(postPb.Id)
	var post = &domain.Post{
		Id:        id,
		User:      userPb.Username,
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
	return *post

}
