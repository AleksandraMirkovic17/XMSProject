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

func mapAuthRoleToCreateOrderRole(role pb.UserRoleAuth) events.Role {
	switch role {
	case pb.UserRoleAuth_RegularAuth:
		return events.Regular
	case pb.UserRoleAuth_AdminAuth:
		return events.Admin
	case pb.UserRoleAuth_AgentAuth:
		return events.Agent
	}
	return events.Regular

}
func mapAuthGenderToCreateOrderGender(gender pb.GenderAuth) events.Gender {
	switch gender {
	case pb.GenderAuth_MaleAuth:
		return events.MALE
	case pb.GenderAuth_FemaleAuth:
		return events.FEMALE
	}
	return events.Empty
}
