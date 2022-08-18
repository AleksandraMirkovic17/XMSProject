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
	Skills      []string
}

type RegisterUserCommandType int8

const (
	CreateUserCredentials RegisterUserCommandType = iota
	UserProfileCreate
	RollebackRegisterUserProfile
	CreateUserNode
	RollebackRegisterConnectionNode
	AuthenticationServiceRegisterUpdate
	RollbackRegisterAuthenticationService
	CreateJobNode
	RollbackRegisterJobNode
	ApproveRegistration
	CancelRegistration
	UnknownCommand
)

type RegisterUserCommand struct {
	User UserDetails
	Type RegisterUserCommandType
}

type RegisterUserReplyType int8

const (
	UserCredentialsCreated RegisterUserReplyType = iota
	UserProfileCreated
	UserProfileNotCreated
	UserProfileRegisterRolledBack
	UserNodeCreated
	UserNodeFailedToCreate
	ConnectionsRolledBack
	AuthenticationServiceUserCreated
	AuthenticationServiceUserNotCreated
	AuthenticationServiceRegisterRolledBack
	JobNodeCreated
	JobNodeFailedToCreate
	JobNodeRegisterRolledBack
	RegistrationCancelled
	RegistrationApproved
	UnknownReply
)

type RegisterUserReply struct {
	User UserDetails
	Type RegisterUserReplyType
}
