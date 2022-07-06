package create_order

import "time"

type Gender int

const (
	Empty Gender = iota
	MALE
	FEMALE
)

type Role int

const (
	Regular Role = iota
	Admin
	Agent
)

type UserDetails struct {
	Id          string
	Name        string
	Surname     string
	Username    string
	Password    string
	Email       string
	Birthday    time.Time
	Gender      Gender
	Role        Role
	PhoneNumber string
	IsPublic    bool
}

type RegisterUserCommandType int8

const (
	CreateUserCredentials RegisterUserCommandType = iota
	UserProfileCreate
	RollebackUserProfile
	CreateUserNode
	RollebackConnectionNode
	AuthenticationServiceUpdate
	RollbackAuthenticationService
	ApproveRegistration
	CancelRegistration
	UnknownCommand
)

type RegisterUserCommand struct {
	Order UserDetails
	Type  RegisterUserCommandType
}

type RegisterUserReplyType int8

const (
	UserCredentialsCreated RegisterUserReplyType = iota
	UserProfileCreated
	UserProfileNotCreated
	UserProfileRolledBack

	UserNodeCreated
	UserNodeFailedToCreate
	ConnectionsRolledBack

	AuthenticationServiceUpdated
	AuthenticationServiceNotUpdated
	AuthenticationServiceRolledBack
	RegistrationCancelled
	RegistrationApproved

	UnknownReply
)

type RegisterUserReply struct {
	Order UserDetails
	Type  RegisterUserReplyType
}
