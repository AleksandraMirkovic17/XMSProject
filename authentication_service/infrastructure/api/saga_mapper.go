package api

import (
	"AuthenticationService/domain"
	pb "github.com/dislinked/common/proto/authentication_service"
	"github.com/dislinked/common/saga/create_order"
	events "github.com/dislinked/common/saga/create_order"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func mapCommandToAuthUser(command *events.RegisterUserCommand) *domain.User {
	var userD = &domain.User{
		ID:       primitive.ObjectID{},
		Username: command.Order.Username,
		Password: command.Order.Password,
		Role:     mapCommandRoleToAuthRole(command.Order.Role),
	}
	return userD
}

func mapCommandRoleToAuthRole(role create_order.Role) string {
	switch role {
	case create_order.Agent:
		return "Agent"
	case create_order.Admin:
		return "Admin"
	default:
		return "Regular"

	}
	return "Regular"
}

func mapAuthRoleToCreateOrderRole(role pb.UserRole) events.Role {
	switch role {
	case pb.UserRole_Regular:
		return events.Regular
	case pb.UserRole_Admin:
		return events.Admin
	case pb.UserRole_Agent:
		return events.Regular
	}
	return events.Regular

}
func mapAuthGenderToCreateOrderGender(gender pb.Gender) events.Gender {
	switch gender {
	case pb.Gender_Male:
		return events.MALE
	case pb.Gender_Female:
		return events.FEMALE
	}
	return events.Empty
}
