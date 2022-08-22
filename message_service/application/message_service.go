package application

import (
	"MessageService/domain"
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MessageService struct {
	store domain.MessageStore
}

func NewMessageService(store domain.MessageStore) *MessageService {
	return &MessageService{
		store: store,
	}
}

func (service *MessageService) Get(id primitive.ObjectID) (*domain.Message, error) {
	//mongoId.(primitive.ObjectID).Hex()
	print("Id primitive", id.Hex())
	return service.store.Get(id)
}

func (service *MessageService) GetAll() ([]*domain.Message, error) {
	return service.store.GetAll()
}

func (service *MessageService) Insert(message *domain.Message) (*domain.Message, error) {
	(*message).Id = primitive.NewObjectID()
	err := service.store.Insert(message)
	if err != nil {
		err := errors.New("Error occured during creating new message")
		return nil, err
	}

	return message, nil
}

func (service *MessageService) GetAllByUser(uuid string) ([]*domain.Message, error) {
	return service.store.GetAllByUser(uuid)
}

func (service *MessageService) DeleteAll() {
	service.store.DeleteAll()
}
