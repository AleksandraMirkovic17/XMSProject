package api

import (
	"UserService/domain"
	pb "github.com/dislinked/common/proto/user_service"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func mapUser(user *domain.User) *pb.User {
	var userPb = &pb.User{
		Id:          user.Id.String(),
		Name:        user.Name,
		Surname:     user.Surname,
		Username:    user.Username,
		Email:       user.Email,
		Gender:      mapGenderToPb(user.Gender),
		Role:        mapRoleToPb(user.Role),
		DateOfBirth: timestamppb.New(user.DateOfBirth),
	}

	return userPb
}

func mapRoleToPb(role domain.Role) pb.UserRole {
	switch role {
	case domain.Regular:
		return pb.UserRole_Regular
	case domain.Agent:
		return pb.UserRole_Agent
	case domain.Admin:
		return pb.UserRole_Admin
	}
	return pb.UserRole_Regular
}

func mapGenderToPb(gender domain.Gender) pb.Gender {
	switch gender {
	case domain.FEMALE:
		return pb.Gender_Female
	case domain.MALE:
		return pb.Gender_Male
	case domain.Other:
		return pb.Gender_Other
	}
	return pb.Gender_Other
}

func mapNewUserPbToDomain(userPb *pb.NewUser) *domain.User {
	userD := &domain.User{
		Id:          primitive.ObjectID{},
		Name:        (*userPb).User.Name,
		Surname:     (*userPb).User.Surname,
		Username:    (*userPb).User.Username,
		Email:       (*userPb).User.Email,
		Password:    (*userPb).User.Password,
		Gender:      mapGenderPbToDomainGender(userPb.User.Gender),
		Role:        domain.Regular,
		DateOfBirth: (*((*userPb).User.DateOfBirth)).AsTime(),
	}
	return userD
}

func mapGenderPbToDomainGender(gender pb.Gender) domain.Gender {
	switch gender {
	case pb.Gender_Female:
		return domain.FEMALE
	case pb.Gender_Male:
		return domain.MALE
	case pb.Gender_Other:
		return domain.Other
	}
	return domain.Other

}

/*
func mapNewUser(userPb *pb.NewUser) *domain.User {
	userD := &domain.User{
		Id:          uuid.NewV4(),
		Name:        userPb.Name,
		Surname:     userPb.Surname,
		Username:    &userPb.Username,
		Email:       &userPb.Email,
		Number:      userPb.Number,
		Gender:      mapGenderToDomain(userPb.Gender),
		DateOfBirth: userPb.DateOfBirth,
		Password:    userPb.Password,
		UserRole:    domain.Regular,
		Biography:   userPb.Biography,
		Blocked:     false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Private:     userPb.Private,
	}
	return userD
}*/

/*func mapNewPost(postPb *pb.Post) *domain.Post {
	post := &domain.Post{
		Id:        primitive.NewObjectID(),
		User:      postPb.User,
		PostText:  postPb.Posttext,
		Date:      time.Now(),
		IsDeleted: false,
	}
	for _, image := range postPb.Image {
		post.Images = append(post.Images, image)
	}
	for _, link := range postPb.Links {
		post.Links = append(post.Links, link)
	}
	post.Comments = []domain.Comment{}
	for _, comment := range postPb.Comments {
		id, err := primitive.ObjectIDFromHex(comment.Id)
		if err != nil {
			continue
		}
		post.Comments = append(post.Comments, domain.Comment{
			Id:      id,
			Content: comment.Content,
			Date:    comment.Date.AsTime(),
			User:    comment.Username,
		})
	}

	return post
}
*/
