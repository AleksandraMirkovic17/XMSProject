package mappers

import (
	"UserService/domain"
	pb "github.com/dislinked/common/proto/user_service"
	events "github.com/dislinked/common/saga/update_order"
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
	domainUser := &domain.User{
		Name:     eventUser.Name,
		Surname:  eventUser.Surname,
		Username: eventUser.Username,
		Email:    eventUser.Email,
		Password: eventUser.Password,
		Phone:    eventUser.PhoneNumber,
		Gender:   mapEventGenderToDomainGender(eventUser.Gender),
		Role:     mapEventRoleToDomainRole(eventUser.Role),
		Public:   eventUser.IsPublic,
	}

	for _, skill := range eventUser.Skills {
		domainUser.Skills = append(domainUser.Skills, domain.Skill{
			Name: skill,
		})

	}
	return domainUser
}