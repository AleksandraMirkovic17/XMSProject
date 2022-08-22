package mappers

import (
	"UserService/domain"
	pb "github.com/dislinked/common/proto/user_service"
	events "github.com/dislinked/common/saga/update_order"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func MapEventGenderToDomainGender(gender events.Gender) domain.Gender {
	switch gender {
	case events.MALE:
		return domain.MALE
	case events.FEMALE:
		return domain.FEMALE
	default:
		return domain.Other
	}
}

func mapEventRoleToDomainRole(role events.Role) domain.Role {
	switch role {
	case events.Agent:
		return domain.Agent
	case events.Admin:
		return domain.Admin
	case events.Regular:
		return domain.Regular
	default:
		return domain.Regular
	}
}

func mapPbRoleToEventRole(role pb.UserRole) events.Role {
	switch role {
	case pb.UserRole_Agent:
		return events.Agent
	case pb.UserRole_Admin:
		return events.Admin
	default:
		return events.Regular
	}
}

func mapPbGenderToEventGender(gender pb.Gender) events.Gender {
	switch gender {
	case pb.Gender_Female:
		return events.FEMALE
	case pb.Gender_Male:
		return events.MALE
	default:
		return events.Empty
	}
}

func mapEventGenderToDomainGender(gender events.Gender) domain.Gender {
	switch gender {
	case events.FEMALE:
		return domain.FEMALE
	case events.MALE:
		return domain.MALE
	default:
		return domain.Other

	}
}

func MapPbUserToEventUpdateUser(pbUser pb.User) *events.UserDetails {
	println("pb user id: ", pbUser.Id)
	println("name is: ", pbUser.Name)
	println("surname is: ", pbUser.Surname)
	println("username is: " + pbUser.Username)
	println("contact phone is " + pbUser.ContactPhone)
	println("password is " + pbUser.Password)

	eventUser := &events.UserDetails{
		Id:          pbUser.Id,
		Name:        pbUser.Name,
		Surname:     pbUser.Surname,
		Username:    pbUser.Username,
		Password:    pbUser.Password,
		Email:       pbUser.Email,
		Gender:      mapPbGenderToEventGender(pbUser.Gender),
		Role:        mapPbRoleToEventRole(pbUser.Role),
		PhoneNumber: pbUser.ContactPhone,
		IsPublic:    pbUser.Public,
	}

	for _, skill := range pbUser.Skills {
		eventUser.Skills = append(eventUser.Skills, skill.Name)
	}

	return eventUser
}

func MapEventUserToDomainUser(eventUser events.UserDetails) *domain.User {
	id, err := primitive.ObjectIDFromHex(eventUser.Id)
	if err != nil {
		println("Error convering in command_mapper!")
		return nil
	}
	println("Phone number: " + eventUser.PhoneNumber)
	domainUser := &domain.User{
		Id:       id,
		Name:     eventUser.Name,
		Surname:  eventUser.Surname,
		Username: eventUser.Username,
		Email:    eventUser.Email,
		Password: eventUser.Password,
		Phone:    eventUser.PhoneNumber,
		Gender:   mapEventGenderToDomainGender(eventUser.Gender),
		Role:     mapEventRoleToDomainRole(eventUser.Role),
		Public:   eventUser.IsPublic,
		Skills:   []domain.Skill{},
	}

	for _, skill := range eventUser.Skills {
		domainUser.Skills = append(domainUser.Skills, domain.Skill{
			Name: skill,
		})

	}
	return domainUser
}
