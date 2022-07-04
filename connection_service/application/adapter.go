package application

import (
	"crypto/tls"
	"fmt"
	"log"

	userService "github.com/dislinked/common/proto/user_service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func getConnection(address string) (*grpc.ClientConn, error) {
	config := &tls.Config{
		InsecureSkipVerify: true,
	}
	return grpc.Dial(address, grpc.WithTransportCredentials(credentials.NewTLS(config)))
}

func NewUserClient(address string) userService.UserServiceClient {
	conn, err := getConnection(address)
	if err != nil {
		fmt.Println("Gateway failed to start", "Failed to start")
		log.Fatalf("Failed to start gRPC connection to Catalogue service: %v", err)
	}
	return userService.NewUserServiceClient(conn)
}
