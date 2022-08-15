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
	COLLECTION = "user_authentication"
)

type AuthenticationMongoDBStore struct {
	users *mongo.Collection
}

func NewAuthenticationMongoDBStore(client *mongo.Client) *AuthenticationMongoDBStore {
	users := client.Database(DATABASE).Collection(COLLECTION)
	return &AuthenticationMongoDBStore{
		users: users,
	}
}

func (store *AuthenticationMongoDBStore) Create(user *domain.UserAuthentication) (primitive.ObjectID, error) {
	println("Kreiranje usera sa username", user.Username)
	_, err := store.users.InsertOne(context.TODO(), user)
	if err != nil {
		return primitive.ObjectID{}, err
	}
	return user.ID, nil
}

func (store *AuthenticationMongoDBStore) GetAll() ([]*domain.UserAuthentication, error) {
	filter := bson.D{}
	return store.filter(filter)
}

func (store *AuthenticationMongoDBStore) GetById(id primitive.ObjectID) (*domain.UserAuthentication, error) {
	filter := bson.M{"_id": id}
	return store.filterOne(filter)
}

func (store *AuthenticationMongoDBStore) GetByUsername(username string) (user *domain.UserAuthentication, err error) {
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

func (store *AuthenticationMongoDBStore) filter(filter interface{}) ([]*domain.UserAuthentication, error) {
	cursor, err := store.users.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}

	return decode(cursor)
}

func (store *AuthenticationMongoDBStore) filterOne(filter interface{}) (post *domain.UserAuthentication, err error) {
	result := store.users.FindOne(context.TODO(), filter)
	err = result.Decode(&post)
	return
}

func decode(cursor *mongo.Cursor) (posts []*domain.UserAuthentication, err error) {
	for cursor.Next(context.TODO()) {
		var post domain.UserAuthentication
		err = cursor.Decode(&post)
		if err != nil {
			return
		}
		posts = append(posts, &post)
	}
	err = cursor.Err()
	return
}
