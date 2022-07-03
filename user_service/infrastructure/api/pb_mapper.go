package api

import (
	"UserService/domain"

	dislinked "github.com/dislinked/common/proto/user_service"
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
		Public:      user.Public,
		Skills:      []*dislinked.Skill{},
	}

	for _, skill := range user.Skills {
		userPb.Skills = append(userPb.Skills, &dislinked.Skill{
			Id:   skill.Id.String(),
			Name: skill.Name,
		})
	}

	return userPb
}

func mapUserPbToDomain(userPb *pb.NewUser) *domain.User {
	userD := &domain.User{
		Id:          primitive.NewObjectID(),
		Name:        (*userPb).User.Name,
		Surname:     (*userPb).User.Surname,
		Username:    (*userPb).User.Username,
		Email:       (*userPb).User.Email,
		Password:    (*userPb).User.Password,
		Gender:      mapGenderPbToDomainGender((*userPb).User.Gender),
		Role:        domain.Regular,
		DateOfBirth: (*((*userPb).User.DateOfBirth)).AsTime(),
		Public:      (*userPb).User.Public,
	}

	for _, skill := range userPb.User.Skills {
		userD.Skills = append(userD.Skills, domain.Skill{
			Id:   primitive.NewObjectID(),
			Name: skill.Name,
		})
	}
	return userD
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
