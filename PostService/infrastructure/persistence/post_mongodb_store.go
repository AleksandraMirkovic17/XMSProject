package persistence

import (
	"PostService/domain"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE   = "post"
	COLLECTION = "post"
)

type PostMongoDBStore struct {
	posts *mongo.Collection
}

func (store *PostMongoDBStore) GetAllByUser(uuid string) ([]*domain.Post, error) {
	filter := bson.M{"user_id": uuid}
	return store.filter(filter)
}

func (store *PostMongoDBStore) GetAllByConnections(uuids []string) ([]*domain.Post, error) {
	//TODO implement me
	panic("implement me")
}

func (store *PostMongoDBStore) ReactToPost(post *domain.Post, username string, reaction *domain.Reaction) error {
	//TODO implement me
	panic("implement me")
}

func (store *PostMongoDBStore) DeleteAll() {
	//TODO implement me
	panic("implement me")
}

func NewPostMongoDBStore(client *mongo.Client) domain.PostStore {
	posts := client.Database(DATABASE).Collection(COLLECTION)
	return &PostMongoDBStore{posts: posts}
}

func (store *PostMongoDBStore) Get(id primitive.ObjectID) (*domain.Post, error) {
	filter := bson.M{"_id": id}
	return store.filterOne(filter)
}

func (store *PostMongoDBStore) GetAll() ([]*domain.Post, error) {
	filter := bson.D{}
	return store.filter(filter)
}

func (store *PostMongoDBStore) GetAllByConnectionIds(uuids []string) ([]*domain.Post, error) {
	var posts []*domain.Post

	for _, uuid := range uuids {
		postsByUser, err := store.GetAllByUser(uuid)
		posts = append(posts, postsByUser...)

		if err != nil {
			return nil, err
		}
	}

	return posts, nil
}

func (store *PostMongoDBStore) Insert(post *domain.Post) error {
	result, err := store.posts.InsertOne(context.TODO(), post)
	if err != nil {
		return err
	}
	post.Id = result.InsertedID.(primitive.ObjectID)

	return nil
}

func (store *PostMongoDBStore) CreateComment(post *domain.Post, comment *domain.Comment) error {
	comments := append(post.Comments, *comment)

	_, err := store.posts.UpdateOne(context.TODO(), bson.M{"_id": post.Id}, bson.D{
		{"$set", bson.D{{"comments", comments}}},
	},
	)
	if err != nil {
		return err
	}

	return nil
}

func (store *PostMongoDBStore) LikePost(post *domain.Post, username string) error {

	var reactions []domain.Reaction

	reactionExists := false
	for _, reaction := range post.Reactions {
		if reaction.User != username {
			reactions = append(reactions, reaction)
		} else {
			if reaction.Reaction != domain.LIKE {
				reaction.Reaction = domain.LIKE
				reactions = append(reactions, reaction)
			}
			reactionExists = true
		}

	}
	if !reactionExists {
		reaction := domain.Reaction{
			User:     username,
			Reaction: domain.LIKE,
		}
		reactions = append(reactions, reaction)
	}

	_, err := store.posts.UpdateOne(context.TODO(), bson.M{"_id": post.Id}, bson.D{
		{"$set", bson.D{{"reactions", reactions}}},
	},
	)
	if err != nil {
		return err
	}

	return nil
}

func (store *PostMongoDBStore) DislikePost(post *domain.Post, username string) error {
	var reactions []domain.Reaction

	reactionExists := false
	for _, reaction := range post.Reactions {
		if reaction.User != username {
			reactions = append(reactions, reaction)
		} else {
			if reaction.Reaction != domain.DISLIKED {
				reaction.Reaction = domain.DISLIKED
				reactions = append(reactions, reaction)
			}
			reactionExists = true
		}

	}
	if !reactionExists {
		reaction := domain.Reaction{
			User:     username,
			Reaction: domain.DISLIKED,
		}
		reactions = append(reactions, reaction)
	}

	_, err := store.posts.UpdateOne(context.TODO(), bson.M{"_id": post.Id}, bson.D{
		{"$set", bson.D{{"reactions", reactions}}},
	},
	)
	if err != nil {
		return err
	}

	return nil
}

func (store *PostMongoDBStore) filter(filter interface{}) ([]*domain.Post, error) {
	cursor, err := store.posts.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}

	return decode(cursor)
}

func (store *PostMongoDBStore) filterOne(filter interface{}) (post *domain.Post, err error) {
	result := store.posts.FindOne(context.TODO(), filter)
	err = result.Decode(&post)
	return
}

func decode(cursor *mongo.Cursor) (posts []*domain.Post, err error) {
	for cursor.Next(context.TODO()) {
		var post domain.Post
		err = cursor.Decode(&post)
		if err != nil {
			return
		}
		posts = append(posts, &post)
	}
	err = cursor.Err()
	return
}
