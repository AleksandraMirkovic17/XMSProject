package api

import (
	"UserService/domain"

	dislinked "github.com/dislinked/common/proto/user_service"
	pb "github.com/dislinked/common/proto/user_service"
	uuid "github.com/satori/go.uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func mapUser(user *domain.User) *pb.User {
	var userPb = &pb.User{
		Id:          user.Id.String(),
		Name:        user.Name,
		Surname:     user.Surname,
		Username:    *user.Username,
		Email:       *user.Email,
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
		Id:          uuid.NewV4(),
		Name:        (*userPb).User.Name,
		Surname:     (*userPb).User.Surname,
		Username:    &(*userPb).User.Username,
		Email:       &(*userPb).User.Email,
		Password:    (*userPb).User.Password,
		Gender:      mapGenderPbToDomainGender((*userPb).User.Gender),
		Role:        domain.Regular,
		DateOfBirth: (*((*userPb).User.DateOfBirth)).AsTime(),
		Public:      (*userPb).User.Public,
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
