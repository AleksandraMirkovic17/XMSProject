package api

import (
	"PostService/application"
	"PostService/infrastructure/orchestrators"
	"context"
	"fmt"
	pb "github.com/dislinked/common/proto/post_service"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/status"
)

type PostHandler struct {
	pb.UnimplementedPostServiceServer
	service                              *application.PostService
	friendPostedNotificationOrchestrator *orchestrators.FriendPostedNotificationOrchestrator
}

func NewPostHandler(service *application.PostService, orchestrator *orchestrators.FriendPostedNotificationOrchestrator) *PostHandler {
	return &PostHandler{
		service:                              service,
		friendPostedNotificationOrchestrator: orchestrator,
	}
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

	//pozivanje sage za slanje notifikacija
	handler.friendPostedNotificationOrchestrator.Start(MapDomainNotificationToEventNotification(newPost))

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

func (handler *PostHandler) ReactToPost(ctx context.Context, request *pb.ReactionRequest) (*pb.ReactionResponse, error) {
	id := request.Reaction.PostId
	println("Id in react+", id)
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	post1, err := handler.service.Get(objectId)
	if err != nil {
		return nil, err
	}
	println("Post that is gotten: ", post1.PostText)
	postWithReaction, err := handler.service.ReactToPost(post1, request.Reaction.Username, mapPbReactionTypeToDomain(request.Reaction.ReactionType))
	if err != nil {
		panic(fmt.Errorf("Invalid reaction"))
	}

	reactionResponse := pb.ReactionResponse{
		Post: mapPostFromDomainToPb(postWithReaction),
	}

	return &reactionResponse, nil

}

func (handler *PostHandler) DeleteReaction(ctx context.Context, request *pb.DeleteReactionRequest) (*pb.ReactionResponse, error) {
	id := request.PostId
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	post, err := handler.service.Get(objectId)
	if err != nil {
		return nil, err
	}
	post, err = handler.service.DeleteReaction(post, request.Username)
	if err != nil {
		return nil, err
	}
	postPb := mapPostFromDomainToPb(post)
	response := &pb.ReactionResponse{
		Post: postPb,
	}
	return response, nil
}

func (handler *PostHandler) GetLikes(ctx context.Context, request *pb.GetRequest) (*pb.MultipleReactionsResponse, error) {
	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	likes, err := handler.service.GetLikes(objectId)
	response := &pb.MultipleReactionsResponse{}
	for _, like := range likes {
		response.Owner = append(response.Owner, &pb.Owner{
			Username: like.Ahuthor.Username,
			Name:     like.Ahuthor.Name,
			Surname:  like.Ahuthor.Surname,
		})
	}
	return response, nil
}

func (handler *PostHandler) GetDislikes(ctx context.Context, request *pb.GetRequest) (*pb.MultipleReactionsResponse, error) {
	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	dislikes, err := handler.service.GetDislikes(objectId)
	response := &pb.MultipleReactionsResponse{}
	for _, like := range dislikes {
		response.Owner = append(response.Owner, &pb.Owner{
			Username: like.Ahuthor.Username,
			Name:     like.Ahuthor.Name,
			Surname:  like.Ahuthor.Surname,
		})
	}
	return response, nil
}

func (handler *PostHandler) CreateCommentOnPost(ctx context.Context, request *pb.CommentRequest) (*pb.CommentResponse, error) {
	id := request.PostId
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	post, err := handler.service.Get(objectId)
	if err != nil {
		return nil, err
	}

	postD, err := handler.service.CreateComment(post, mapNewPbCommentToDomainComment(request.Comment))
	if err != nil {
		return nil, err
	}
	response := &pb.CommentResponse{Post: mapPostFromDomainToPb(postD)}
	return response, nil
}

func (handler *PostHandler) GetComments(ctx context.Context, request *pb.GetRequest) (*pb.MultipleCommentsResponse, error) {
	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	comments, error := handler.service.GetComments(objectId)
	if error != nil {
		return nil, err
	}
	response := mapDomainCommentsToPbComments(comments)
	return response, nil

}
