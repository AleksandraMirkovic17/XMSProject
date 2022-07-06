package adapters

import (
	"crypto/tls"

	connectionService "github.com/dislinked/common/proto/connection_service"
	userService "github.com/dislinked/common/proto/user_service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func NewUserClient(address string) userService.UserServiceClient {
	conn, _ := getConnection(address)
	return userService.NewUserServiceClient(conn)
}

func getConnection(address string) (*grpc.ClientConn, error) {
	config := &tls.Config{
		InsecureSkipVerify: true,
	}
	return grpc.Dial(address, grpc.WithTransportCredentials(credentials.NewTLS(config)))
}

func NewConnectionClient(address string) connectionService.ConnectionServiceClient {
	conn, _ := getConnection(address)
	return connectionService.NewConnectionServiceClient(conn)
}
