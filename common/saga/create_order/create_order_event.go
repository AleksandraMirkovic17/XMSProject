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

type ConnectionUserDetails struct {
	Id       string
	Username string
	IsPublic bool
}

type JobUserDetails struct {
	UserID   string
	Username string
	Skills   []string
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
	CreateJobNode
	RollbackJobNode
	ApproveRegistration
	CancelRegistration
	UnknownCommand
)

type RegisterUserCommand struct {
	User UserDetails
	Type RegisterUserCommandType
}

type RegisterAuthenticationUserCommand struct {
	User UserDetails
	Type RegisterUserCommandType
}

type RegisterConnectionUserCommand struct {
	User ConnectionUserDetails
	Type RegisterUserCommandType
}

type RegisterJobUserCommand struct {
	User JobUserDetails
	Type RegisterUserCommandType
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
	JobNodeCreated
	JobNodeFailedToCreate
	JobNodeRolledBack
	RegistrationCancelled
	RegistrationApproved
	UnknownReply
)

type RegisterUserReply struct {
	User UserDetails
	Type RegisterUserReplyType
}

type RegisterUserAuthenticationReply struct {
	User UserDetails
	Type RegisterUserReplyType
}

type RegisterUserConnectionReply struct {
	User ConnectionUserDetails
	Type RegisterUserReplyType
}

type RegisterJobUserReply struct {
	User JobUserDetails
	Type RegisterUserReplyType
}
