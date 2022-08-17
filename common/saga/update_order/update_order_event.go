package update_order

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

type UpdateUserCommandType int8

const (
	CreateUserCredentials UpdateUserCommandType = iota
	UserProfileUpdate
	RollebackUserProfile
	UpdateUserNode
	RollebackConnectionNode
	AuthenticationServiceUpdate
	RollbackAuthenticationService
	UpdateJobNode
	RollbackJobNode
	ApproveUpdate
	CancelUpdate
	UnknownCommand
)

type UpdateUserCommand struct {
	User    UserDetails
	OldUser UserDetails
	Type    UpdateUserCommandType
}

type UpdateUserReplyType int8

const (
	UserCredentialsCreated UpdateUserReplyType = iota
	UserProfileUpdated
	UserProfileNotUpdated
	UserProfileRolledBack
	UserNodeUpdated
	UserNodeFailedToUpdate
	ConnectionsRolledBack
	AuthenticationServiceUpdated
	AuthenticationServiceNotUpdated
	AuthenticationServiceRolledBack
	JobNodeUpdated
	JobNodeFailedToUpdate
	JobNodeRolledBack
	RegistrationCancelled
	RegistrationApproved
	UnknownReply
)

type UpdateUserReply struct {
	User    UserDetails
	UserOld UserDetails
	Type    UpdateUserReplyType
}
