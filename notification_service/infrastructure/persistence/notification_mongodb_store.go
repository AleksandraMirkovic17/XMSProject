package persistence

import (
	"NotificationService/domain"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE   = "notification"
	COLLECTION = "notification"
)

type NotificationMongoDBStore struct {
	notifications *mongo.Collection
}

func NewNotificationMongoDBStore(client *mongo.Client) domain.NotificationStore {
	notifications := client.Database(DATABASE).Collection(COLLECTION)
	return &NotificationMongoDBStore{notifications: notifications}
}

func (store *NotificationMongoDBStore) Get(id primitive.ObjectID) (*domain.Notification, error) {
	filter := bson.M{"_id": id}
	return store.filterOne(filter)
}

func (store *NotificationMongoDBStore) GetAll() ([]*domain.Notification, error) {
	filter := bson.D{}
	return store.filter(filter)
}

func (store *NotificationMongoDBStore) Insert(notification *domain.Notification) error {
	result, err := store.notifications.InsertOne(context.TODO(), notification)
	if err != nil {
		return err
	}
	notification.Id = result.InsertedID.(primitive.ObjectID)

	return nil
}

func (store *NotificationMongoDBStore) Update(notification *domain.Notification) error {
	_, err := store.notifications.UpdateOne(context.TODO(), bson.M{"_id": notification.Id}, bson.M{"$set": notification})
	if err != nil {
		return err
	}
	return nil
}

func (store *NotificationMongoDBStore) GetAllByUser(uuid string) ([]*domain.Notification, error) {
	filter := bson.M{"user": uuid}
	return store.filter(filter)
}

func (store *NotificationMongoDBStore) filter(filter interface{}) ([]*domain.Notification, error) {
	cursor, err := store.notifications.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}

	return decode(cursor)
}

func (store *NotificationMongoDBStore) filterOne(filter interface{}) (notification *domain.Notification, err error) {
	result := store.notifications.FindOne(context.TODO(), filter)
	err = result.Decode(&notification)
	return
}

func decode(cursor *mongo.Cursor) (notifications []*domain.Notification, err error) {
	for cursor.Next(context.TODO()) {
		var notification domain.Notification
		err = cursor.Decode(&notification)
		if err != nil {
			return
		}
		notifications = append(notifications, &notification)
	}
	err = cursor.Err()
	return
}
