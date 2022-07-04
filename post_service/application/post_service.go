package application

import (
	"PostService/domain"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PostService struct {
	store domain.PostStore
}

func NewPostService(store domain.PostStore) *PostService {
	return &PostService{
		store: store,
	}
}

func (service *PostService) Get(id primitive.ObjectID) (*domain.Post, error) {
	//mongoId.(primitive.ObjectID).Hex()
	print("Id primitive", id.Hex())
	return service.store.Get(id)
}

func (service *PostService) GetAll() ([]*domain.Post, error) {
	return service.store.GetAll()
}

func (service *PostService) Insert(post *domain.Post) (*domain.Post, error) {
	(*post).Id = primitive.NewObjectID()
	err := service.store.Insert(post)
	if err != nil {
		err := errors.New("Error occured during creating new post")
		return nil, err
	}

	return post, nil
}

func (service *PostService) GetAllByUser(uuid string) ([]*domain.Post, error) {
	return service.store.GetAllByUser(uuid)
}

func (service *PostService) GetAllByConnections(uuids []string) ([]*domain.Post, error) {
	return service.store.GetAllByConnections(uuids)
}

func (service *PostService) CreateComment(post *domain.Post, comment *domain.Comment) error {
	return service.store.CreateComment(post, comment)
}
func (service *PostService) ReactToPost(post *domain.Post, user string, reaction domain.ReactionType) (*domain.Post, error) {
	return service.store.ReactToPost(post, user, reaction)
}

func (service *PostService) DeleteReaction(post *domain.Post, user string) (*domain.Post, error) {
	return service.store.DeleteReaction(post, user)
}

func (service *PostService) GetLikes(id primitive.ObjectID) ([]*domain.ReactionDetailsDTO, error) {
	return service.store.GetAllLikes(id)
}

func (service *PostService) GetDislikes(id primitive.ObjectID) ([]*domain.ReactionDetailsDTO, error) {
	return service.store.GetAllDislikes(id)
}

func (service *PostService) DeleteAll() {
	service.store.DeleteAll()
}
