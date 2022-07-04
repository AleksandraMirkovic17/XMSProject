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
		Id:                   user.Id.String(),
		Name:                 user.Name,
		Surname:              user.Surname,
		Username:             user.Username,
		Email:                user.Email,
		Gender:               mapGenderToPb(user.Gender),
		Role:                 mapRoleToPb(user.Role),
		DateOfBirth:          timestamppb.New(user.DateOfBirth),
		Public:               user.Public,
		Skills:               []*dislinked.Skill{},
		Interests:            []*dislinked.Interest{},
		EducationExperiences: []*dislinked.EducationExperience{},
		WorkExperiences:      []*dislinked.WorkExperience{},
	}

	for _, skill := range user.Skills {
		userPb.Skills = append(userPb.Skills, &dislinked.Skill{
			Id:   skill.Id.String(),
			Name: skill.Name,
		})
	}

	for _, interest := range user.Interests {
		userPb.Interests = append(userPb.Interests, &dislinked.Interest{
			Id:   interest.Id.String(),
			Name: interest.Name,
		})
	}

	for _, educationExperience := range user.EducationExperiences {
		userPb.EducationExperiences = append(userPb.EducationExperiences, &dislinked.EducationExperience{
			Id:              educationExperience.Id.String(),
			InstitutionName: educationExperience.InstitutionName,
		})
	}

	for _, workExperience := range user.WorkExperiences {
		userPb.WorkExperiences = append(userPb.WorkExperiences, &dislinked.WorkExperience{
			Id:               workExperience.Id.String(),
			OrganizationName: workExperience.OrganizationName,
			PositionName:     workExperience.PositionName,
		})
	}

	return userPb
}

func mapUserPbToDomain(userPb *pb.NewUser) *domain.User {
	userD := &domain.User{
		Id:                   primitive.NewObjectID(),
		Name:                 (*userPb).User.Name,
		Surname:              (*userPb).User.Surname,
		Username:             (*userPb).User.Username,
		Email:                (*userPb).User.Email,
		Password:             (*userPb).User.Password,
		Gender:               mapGenderPbToDomainGender((*userPb).User.Gender),
		Role:                 domain.Regular,
		DateOfBirth:          (*((*userPb).User.DateOfBirth)).AsTime(),
		Public:               (*userPb).User.Public,
		Skills:               []domain.Skill{},
		Interests:            []domain.Interest{},
		EducationExperiences: []domain.EducationExperience{},
		WorkExperiences:      []domain.WorkExperience{},
	}

	for _, skill := range userPb.User.Skills {
		userD.Skills = append(userD.Skills, domain.Skill{
			Id:   primitive.NewObjectID(),
			Name: skill.Name,
		})
	}

	for _, interest := range userPb.User.Interests {
		userD.Interests = append(userD.Interests, domain.Interest{
			Id:   primitive.NewObjectID(),
			Name: interest.Name,
		})
	}

	for _, educationExperience := range userPb.User.EducationExperiences {
		userD.EducationExperiences = append(userD.EducationExperiences, domain.EducationExperience{
			Id:              primitive.NewObjectID(),
			InstitutionName: educationExperience.InstitutionName,
		})
	}

	for _, workExperience := range userPb.User.WorkExperiences {
		userD.WorkExperiences = append(userD.WorkExperiences, domain.WorkExperience{
			Id:               primitive.NewObjectID(),
			OrganizationName: workExperience.OrganizationName,
			PositionName:     workExperience.PositionName,
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
