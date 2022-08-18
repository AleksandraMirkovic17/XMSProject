package mappers

import (
	"UserService/domain"
	dislinked "github.com/dislinked/common/proto/user_service"
	pb "github.com/dislinked/common/proto/user_service"
	events "github.com/dislinked/common/saga/create_order"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func MapUser(user *domain.User) *pb.User {
	var userPb = &pb.User{
		Id:                   user.Id.Hex(),
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

func MapUserPbToDomain(userPb *pb.User) *domain.User {
	println("Mapping pb user to domain")
	dateOfBirth, _ := time.Parse("2006-01-02T15:04", (*userPb).DateOfBirth)
	println("name" + (*userPb).Name)
	println("surname" + (*userPb).Surname)
	println("username" + (*userPb).Username)
	userD := &domain.User{
		Id:                   primitive.NewObjectID(),
		Name:                 (*userPb).Name,
		Surname:              (*userPb).Surname,
		Username:             (*userPb).Username,
		Email:                (*userPb).Email,
		Password:             (*userPb).Password,
		Gender:               mapGenderPbToDomainGender((*userPb).Gender),
		Role:                 domain.Regular,
		DateOfBirth:          dateOfBirth,
		Public:               (*userPb).Public,
		Skills:               []domain.Skill{},
		Interests:            []domain.Interest{},
		EducationExperiences: []domain.EducationExperience{},
		WorkExperiences:      []domain.WorkExperience{},
		Biography:            (*userPb).Biography,
	}

	for _, skill := range userPb.Skills {
		userD.Skills = append(userD.Skills, domain.Skill{
			Id:   primitive.NewObjectID(),
			Name: skill.Name,
		})
	}

	for _, interest := range userPb.Interests {
		userD.Interests = append(userD.Interests, domain.Interest{
			Id:   primitive.NewObjectID(),
			Name: interest.Name,
		})
	}

	for _, educationExperience := range userPb.EducationExperiences {
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

	for _, workExperience := range userPb.WorkExperiences {
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

func MapCommandToUser(command *events.RegisterUserCommand) *domain.User {
	id, err := primitive.ObjectIDFromHex(command.User.Id)
	if err != nil {
		return nil
	}
	user := &domain.User{
		Id:          id,
		Name:        command.User.Name,
		Surname:     command.User.Surname,
		Username:    command.User.Username,
		Email:       command.User.Email,
		Password:    command.User.Password,
		Phone:       command.User.PhoneNumber,
		Gender:      mapCommandGenderToDomainGender(command.User.Gender),
		Role:        mapCommandRoleToDomainRole(command.User.Role),
		DateOfBirth: command.User.Birthday,
		Public:      command.User.IsPublic,
	}
	return user

}

func MapDomainUserToCommandUser(user *domain.User) *events.UserDetails {
	id := user.Id.Hex()

	command := events.UserDetails{
		Id:          id,
		Name:        user.Name,
		Surname:     user.Surname,
		Username:    user.Username,
		Password:    user.Password,
		Email:       user.Email,
		Birthday:    time.Time{},
		Gender:      0,
		Role:        0,
		PhoneNumber: user.Phone,
		IsPublic:    user.Public,
	}

	for _, skill := range user.Skills {
		command.Skills = append(command.Skills, skill.Name)
	}

	return &command
}

func MapDomainUserToConnectionCommandUser(user *domain.User) *events.UserDetails {
	id := user.Id.Hex()

	command := events.UserDetails{
		Id:       id,
		IsPublic: user.Public,
		Username: user.Username,
	}

	return &command
}

func MapDomainUserToJobCommandUser(user *domain.User) *events.UserDetails {
	id := user.Id.Hex()

	command := events.UserDetails{
		Id:       id,
		Username: user.Username,
	}
	for _, skill := range user.Skills {
		command.Skills = append(command.Skills, skill.Name)
	}

	return &command
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
