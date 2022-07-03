package api

import (
	"PostService/application"
	"context"
	"fmt"
	pb "github.com/dislinked/common/proto/post_service"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
		current := mapPostFromDomainToPb(post)
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
		Post: mapPostFromDomainToPb(newPost),
	}

	return response, nil
}

func (handler *PostHandler) Get(ctx context.Context, request *pb.GetRequest) (*pb.GetResponse, error) {
	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	post, err := handler.service.Get(objectId)
	if err != nil {
		return nil, err
	}
	postPb := mapPostFromDomainToPb(post)
	response := &pb.GetResponse{
		Post: postPb,
	}
	return response, nil
}

func (handler *PostHandler) GetByUser(ctx context.Context, request *pb.GetByUserRequest) (*pb.GetAllResponse, error) {
	uuidUser := request.Uuid
	posts, err := handler.service.GetAllByUser(uuidUser)
	if err != nil {
		return nil, err
	}
	response := &pb.GetAllResponse{
		Posts: []*pb.Post{},
	}
	for _, post := range posts {
		current := mapPostFromDomainToPb(post)
		response.Posts = append(response.Posts, current)
	}
	return response, nil
}
