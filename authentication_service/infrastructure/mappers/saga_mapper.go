package mappers

import (
	"AuthenticationService/domain"
	pb "github.com/dislinked/common/proto/authentication_service"
	"github.com/dislinked/common/saga/create_order"
	events "github.com/dislinked/common/saga/create_order"
	events2 "github.com/dislinked/common/saga/update_order"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func MapCommandToAuthUser(command *events.RegisterUserCommand) *domain.UserAuthentication {
	id, err := primitive.ObjectIDFromHex(command.User.Id)
	if err != nil {
		println("Error converting id")
		return nil
	}
	println("Id of the new user is: ", id.Hex())
	var userD = &domain.UserAuthentication{
		ID:       id,
		Username: command.User.Username,
		Password: command.User.Password,
		Role:     MapCommandRoleToAuthRole(command.User.Role),
	}
	return userD
}

func MapUpdateCommandToAuthUser(commandUser *events2.UserDetails) *domain.UserAuthentication {
	id, err := primitive.ObjectIDFromHex(commandUser.Id)
	if err != nil {
		println("Error converting id")
		return nil
	}
	userDomain := &domain.UserAuthentication{
		ID:       id,
		Username: commandUser.Username,
		Password: commandUser.Password,
		Role:     mapUpdateCommandRoleToAuthRole(commandUser.Role),
	}
	return userDomain

}

func mapUpdateCommandRoleToAuthRole(role events2.Role) string {
	switch role {
	case events2.Agent:
		return "Agent"
	case events2.Admin:
		return "Admin"
	default:
		return "Regular"
	}

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
