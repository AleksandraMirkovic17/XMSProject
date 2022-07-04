package persistence

import (
	"UserService/domain"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE   = "user"
	COLLECTION = "user"
)

type UserMongoDBStore struct {
	users *mongo.Collection
}

func NewUserMongoDBStore(client *mongo.Client) domain.UserStore {
	users := client.Database(DATABASE).Collection(COLLECTION)
	return &UserMongoDBStore{users: users}
}

func (store *UserMongoDBStore) GetAll() ([]*domain.User, error) {
	filter := bson.D{}
	return store.filter(filter)
}

func (store *UserMongoDBStore) FindByID(id primitive.ObjectID) (*domain.User, error) {
	filter := bson.M{"_id": id}
	return store.filterOne(filter)
}

func (store *UserMongoDBStore) FindByUsername(username string) (user *domain.User, err error) {
	filter := bson.M{"username": username}
	return store.filterOne(filter)
}

func (store *UserMongoDBStore) FindByUsernameAndNameAndSurname(username string, name string, surname string) (users []*domain.User, err error) {
	filter := bson.M{
		"$or": bson.A{
			bson.M{"username": bson.M{"$regex": username, "$options": "i"}, "public": true},
			bson.M{"name": bson.M{"$regex": name, "$options": "i"}, "public": true},
			bson.M{"surname": bson.M{"$regex": surname, "$options": "i"}, "public": true},
		},
	}
	return store.filter(filter)
}

func (store *UserMongoDBStore) Insert(user *domain.User) error {
	result, err := store.users.InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}
	user.Id = result.InsertedID.(primitive.ObjectID)

	return nil
}

func (store *UserMongoDBStore) Update(user *domain.User) error {
	filter := bson.M{"username": user.Username}
	oldUser, err := store.filterOne(filter)
	if err != nil {
		return err
	}
	_, err = store.users.UpdateOne(
		context.TODO(),
		bson.M{"_id": oldUser.Id},
		bson.D{
			{"$set", bson.D{{"name", user.Name}}},
			{"$set", bson.D{{"surname", user.Surname}}},
			{"$set", bson.D{{"phone", user.Phone}}},
			{"$set", bson.D{{"date_of_birth", user.DateOfBirth}}},
			{"$set", bson.D{{"skills", user.Skills}}},
			{"$set", bson.D{{"interests", user.Interests}}},
			{"$set", bson.D{{"education_experience", user.EducationExperiences}}},
			{"$set", bson.D{{"work_experience", user.WorkExperiences}}},
		},
	)
	if err != nil {
		return err
	}

	return nil
}

func (store *UserMongoDBStore) Search(searchText string) (*[]domain.User, error) {
	/*var users []domain.User
	args := strings.TrimSpace(searchText)
	splitArgs := strings.Split(args, " ")
	allUsers, _ := store.GetAll()
	for _, cuser := range *allUsers {
		nameSearch := strings.ToLower(cuser.Name)
		println(nameSearch)
		surnameSearch := strings.ToLower(cuser.Surname)
		println(surnameSearch)
		usernameSearch := strings.ToLower(*cuser.Username)
		println(usernameSearch)
		for _, searchParam := range splitArgs {
			searchParamLower := strings.ToLower(searchParam)
			println(searchParamLower)
			if !(strings.Contains(nameSearch, searchParamLower) || strings.Contains(surnameSearch, searchParamLower) || strings.Contains(usernameSearch, searchParamLower)) {
				break
			}
			if cuser.Public {
				users = append(users, cuser)
			}
		}
	}

	return &users, nil*/
	panic("implement me")
}

func (store *UserMongoDBStore) Delete(user *domain.User) error {
	/*result := store.db.Delete(user)
	if result.Error != nil {
		return result.Error
	}
	return nil*/
	panic("implement me")
}

func (store *UserMongoDBStore) filter(filter interface{}) ([]*domain.User, error) {
	cursor, err := store.users.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}

	return decode(cursor)
}

func (store *UserMongoDBStore) filterOne(filter interface{}) (post *domain.User, err error) {
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
