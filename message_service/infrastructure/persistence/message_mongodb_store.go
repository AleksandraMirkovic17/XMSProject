package persistence

import (
	"MessageService/domain"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE   = "message"
	COLLECTION = "message"
)

type MessageMongoDBStore struct {
	messages *mongo.Collection
}

func NewMessageMongoDBStore(client *mongo.Client) domain.MessageStore {
	messages := client.Database(DATABASE).Collection(COLLECTION)
	return &MessageMongoDBStore{messages: messages}
}

func (store *MessageMongoDBStore) Get(id primitive.ObjectID) (*domain.Message, error) {
	filter := bson.M{"_id": id}
	return store.filterOne(filter)
}

func (store *MessageMongoDBStore) GetAll() ([]*domain.Message, error) {
	filter := bson.D{}
	return store.filter(filter)
}

func (store *MessageMongoDBStore) Insert(message *domain.Message) error {
	result, err := store.messages.InsertOne(context.TODO(), message)
	if err != nil {
		return err
	}
	message.Id = result.InsertedID.(primitive.ObjectID)

	return nil
}

func (store *MessageMongoDBStore) Update(message *domain.Message) error {
	_, err := store.messages.UpdateOne(context.TODO(), bson.M{"_id": message.Id}, bson.M{"$set": message})
	if err != nil {
		return err
	}
	return nil
}

func (store *MessageMongoDBStore) GetAllByUser(uuid string) ([]*domain.Message, error) {
	filter := bson.M{"user_id": uuid}
	return store.filter(filter)
}

func (store *MessageMongoDBStore) DeleteAll() {
	//TODO implement me
	panic("implement me")
}

func (store *MessageMongoDBStore) filter(filter interface{}) ([]*domain.Message, error) {
	cursor, err := store.messages.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}

	return decode(cursor)
}

func (store *MessageMongoDBStore) filterOne(filter interface{}) (message *domain.Message, err error) {
	result := store.messages.FindOne(context.TODO(), filter)
	err = result.Decode(&message)
	return
}

func decode(cursor *mongo.Cursor) (messages []*domain.Message, err error) {
	for cursor.Next(context.TODO()) {
		var message domain.Message
		err = cursor.Decode(&message)
		if err != nil {
			return
		}
		messages = append(messages, &message)
	}
	err = cursor.Err()
	return
}
