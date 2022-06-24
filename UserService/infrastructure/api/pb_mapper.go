package api

import (
	"UserService/domain"
	pb "github.com/dislinked/common/proto/user_service"
)

/*func mapUser(userD *domain.User) *pb.User {
	userPb := &pb.User{
		Id:          userD.Id.String(),
		Name:        userD.Name,
		Surname:     userD.Surname,
		Username:    *userD.Username,
		Email:       *userD.Email,
		Number:      userD.Number,
		Gender:      mapGenderToPb(userD.Gender),
		DateOfBirth: userD.DateOfBirth,
		Password:    userD.Password,
		UserRole:    mapUserRole(userD.UserRole),
		Biography:   userD.Biography,
		Blocked:     userD.Blocked,
		CreatedAt:   timestamppb.New(userD.CreatedAt),
		UpdatedAt:   timestamppb.New(userD.UpdatedAt),
		Private:     userD.Private,
	}
	return userPb
}*/
func mapUser(user *domain.User) *pb.User {
	var userPb = &pb.User{
		Id:          user.Id.String(),
		Username:    user.Username,
		Email:       user.Email,
		Gender:      mapGenderToPb(user.Gender),
		DateOfBirth: user.DateOfBirth.String(),
	}

	return userPb
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
