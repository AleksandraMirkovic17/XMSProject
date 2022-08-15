package mappers

import (
	"AuthenticationService/domain"
	pb "github.com/dislinked/common/proto/authentication_service"
	"github.com/dislinked/common/saga/create_order"
	events "github.com/dislinked/common/saga/create_order"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func MapCommandToAuthUser(command *events.RegisterUserCommand) *domain.UserAuthentication {
	var userD = &domain.UserAuthentication{
		ID:       primitive.ObjectID{},
		Username: command.User.Username,
		Password: command.User.Password,
		Role:     MapCommandRoleToAuthRole(command.User.Role),
	}
	return userD
}

func MapCommandRoleToAuthRole(role create_order.Role) string {
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

func MapAuthRoleToCreateOrderRole(role pb.UserRoleAuth) events.Role {
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
func MapAuthGenderToCreateOrderGender(gender pb.GenderAuth) events.Gender {
	switch gender {
	case pb.GenderAuth_MaleAuth:
		return events.MALE
	case pb.GenderAuth_FemaleAuth:
		return events.FEMALE
	}
	return events.Empty
}
