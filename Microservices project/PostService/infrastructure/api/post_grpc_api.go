package api

import (
	"PostService/application"
	"context"
	"fmt"
	pb "github.com/dislinked/common/proto/post_service"
	"google.golang.org/grpc/status"
)

type PostHandler struct {
	pb.UnimplementedPostServiceServer
	service *application.PostService
}

func NewPostHandler(service *application.PostService) *PostHandler {
	return &PostHandler{service: service}
}

func (handler *PostHandler) GetAll(ctx context.Context, request *pb.Empty) (*pb.GetMultipleResponse, error) {
	posts, err := handler.service.GetAll()
	if err != nil {
		return nil, err
	}
	response := &pb.GetMultipleResponse{Posts: []*pb.Post{}}
	for _, post := range posts {
		current := mapPost(post)
		response.Posts = append(response.Posts, current)
	}
	return response, nil
}

func (handler *PostHandler) CreatePost(ctx context.Context, request *pb.NewPost) (*pb.NewPost, error) {
	fmt.Println((*request).Post)
	post := mapNewPost(request.Post)
	fmt.Println(post)

	newPost, err := handler.service.Insert(post)
	if err != nil {
		return nil, status.Error(400, err.Error())
	}

	response := &pb.NewPost{
		Post: inversePostMap(newPost),
	}

	return response, nil
}
