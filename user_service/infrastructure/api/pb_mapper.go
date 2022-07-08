package api

import (
	"UserService/domain"
	events "github.com/dislinked/common/saga/create_order"
	"time"

	dislinked "github.com/dislinked/common/proto/user_service"
	pb "github.com/dislinked/common/proto/user_service"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
		DateOfBirth:          user.DateOfBirth.Format("2006-01-02T15:04"),
		Public:               user.Public,
		Skills:               []*dislinked.Skill{},
		Interests:            []*dislinked.Interest{},
		EducationExperiences: []*dislinked.EducationExperience{},
		WorkExperiences:      []*dislinked.WorkExperience{},
		Biography:            user.Biography,
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
			Type:            pb.EducationType(educationExperience.Type),
			StartDate:       educationExperience.StartDate.Format("2006-01-02T15:04"),
			EndDate:         educationExperience.EndDate.Format("2006-01-02T15:04"),
		})
	}

	for _, workExperience := range user.WorkExperiences {
		userPb.WorkExperiences = append(userPb.WorkExperiences, &dislinked.WorkExperience{
			Id:               workExperience.Id.String(),
			OrganizationName: workExperience.OrganizationName,
			PositionName:     workExperience.PositionName,
			StartDate:        workExperience.StartDate.Format("2006-01-02T15:04"),
			EndDate:          workExperience.EndDate.Format("2006-01-02T15:04"),
		})
	}

	return userPb
}

func mapUserPbToDomain(userPb *pb.NewUser) *domain.User {
	dateOfBirth, _ := time.Parse("2006-01-02T15:04", (*userPb).User.DateOfBirth)
	userD := &domain.User{
		Id:                   primitive.NewObjectID(),
		Name:                 (*userPb).User.Name,
		Surname:              (*userPb).User.Surname,
		Username:             (*userPb).User.Username,
		Email:                (*userPb).User.Email,
		Password:             (*userPb).User.Password,
		Gender:               mapGenderPbToDomainGender((*userPb).User.Gender),
		Role:                 domain.Regular,
		DateOfBirth:          dateOfBirth,
		Public:               (*userPb).User.Public,
		Skills:               []domain.Skill{},
		Interests:            []domain.Interest{},
		EducationExperiences: []domain.EducationExperience{},
		WorkExperiences:      []domain.WorkExperience{},
		Biography:            (*userPb).User.Biography,
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
		startDate, _ := time.Parse("2006-01-02T15:04", educationExperience.StartDate)
		endDate, _ := time.Parse("2006-01-02T15:04", educationExperience.EndDate)
		userD.EducationExperiences = append(userD.EducationExperiences, domain.EducationExperience{
			Id:              primitive.NewObjectID(),
			InstitutionName: educationExperience.InstitutionName,
			Type:            domain.EducationType(educationExperience.Type),
			StartDate:       startDate,
			EndDate:         endDate,
		})
	}

	for _, workExperience := range userPb.User.WorkExperiences {
		startDate, _ := time.Parse("2006-01-02T15:04", workExperience.StartDate)
		endDate, _ := time.Parse("2006-01-02T15:04", workExperience.EndDate)
		userD.WorkExperiences = append(userD.WorkExperiences, domain.WorkExperience{
			Id:               primitive.NewObjectID(),
			OrganizationName: workExperience.OrganizationName,
			PositionName:     workExperience.PositionName,
			StartDate:        startDate,
			EndDate:          endDate,
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

func mapCommandRoleToDomainRole(role events.Role) domain.Role {
	switch role {
	case events.Agent:
		return domain.Agent
	case events.Admin:
		return domain.Admin
	case events.Regular:
		return domain.Regular

	}
	return domain.Regular
}

func mapCommandGenderToDomainGender(gender events.Gender) domain.Gender {
	switch gender {
	case events.MALE:
		return domain.MALE
	case events.FEMALE:
		return domain.FEMALE
	case events.Empty:
		return domain.Other
	}
	return domain.Other

}

func mapCommandToUser(command *events.RegisterUserCommand) *domain.User {
	id, err := primitive.ObjectIDFromHex(command.Order.Id)
	if err != nil {
		return nil
	}
	user := &domain.User{
		Id:          id,
		Name:        command.Order.Name,
		Surname:     command.Order.Surname,
		Username:    command.Order.Username,
		Email:       command.Order.Email,
		Password:    command.Order.Password,
		Phone:       command.Order.PhoneNumber,
		Gender:      mapCommandGenderToDomainGender(command.Order.Gender),
		Role:        mapCommandRoleToDomainRole(command.Order.Role),
		DateOfBirth: command.Order.Birthday,
		Public:      command.Order.IsPublic,
	}
	return user

}

/*func mapEducationTypeToPb(educationType domain.EducationType) pb.EducationType {
	switch educationType {
	case domain.PRIMARY_EDUCATION:
		return pb.EducationType_PRIMARY_EDUCATION
	case domain.SECONDARY_EDUCATION:
		return pb.EducationType_SECONDARY_EDUCATION
	case domain.COLLEGE_EDUCATION:
		return pb.EducationType_COLLEGE_EDUCATION
	}
	return pb.EducationType_PRIMARY_EDUCATION
}

func mapEducationTypePbToDomain(educationType pb.EducationType) domain.EducationType {
	switch educationType {
	case pb.EducationType_PRIMARY_EDUCATION:
		return domain.PRIMARY_EDUCATION
	case pb.EducationType_SECONDARY_EDUCATION:
		return domain.SECONDARY_EDUCATION
	case pb.EducationType_COLLEGE_EDUCATION:
		return domain.COLLEGE_EDUCATION
	}
	return domain.PRIMARY_EDUCATION
}*/
