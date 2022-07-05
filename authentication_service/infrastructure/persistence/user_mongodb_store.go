package persistence

import (
	"AuthenticationService/domain"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE   = "user"
	COLLECTION = "user"
)

type AuthenticationMongoDBStore struct {
	users *mongo.Collection
}

func NewAuthenticationMongoDBStore(client *mongo.Client) domain.UserStore {
	users := client.Database(DATABASE).Collection(COLLECTION)
	return &AuthenticationMongoDBStore{users: users}
}

func (store *AuthenticationMongoDBStore) Create(user *domain.User) error {
	_, err := store.users.InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}
	return nil
}

func (store *AuthenticationMongoDBStore) GetAll() ([]*domain.User, error) {
	filter := bson.D{}
	return store.filter(filter)
}

func (store *AuthenticationMongoDBStore) GetById(id primitive.ObjectID) (*domain.User, error) {
	filter := bson.M{"_id": id}
	return store.filterOne(filter)
}

func (store *AuthenticationMongoDBStore) GetByUsername(username string) (user *domain.User, err error) {
	filter := bson.M{"username": username}
	return store.filterOne(filter)
}

func (store *AuthenticationMongoDBStore) DeleteById(id primitive.ObjectID) error {
	_, err := store.users.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		return err
	}
	return nil
}

func (store *AuthenticationMongoDBStore) DeleteAll() {
	panic("implement me")
}

func (store *AuthenticationMongoDBStore) filter(filter interface{}) ([]*domain.User, error) {
	cursor, err := store.users.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}

	return decode(cursor)
}

func (store *AuthenticationMongoDBStore) filterOne(filter interface{}) (post *domain.User, err error) {
	result := store.users.FindOne(context.TODO(), filter)
	err = result.Decode(&post)
	return
}

func decode(cursor *mongo.Cursor) (posts []*domain.User, err error) {
	for cursor.Next(context.TODO()) {
		var post domain.User
		err = cursor.Decode(&post)
		if err != nil {
			return
		}
		posts = append(posts, &post)
	}
	err = cursor.Err()
	return
}
