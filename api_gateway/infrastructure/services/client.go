package services

import (
	connGw "github.com/dislinked/common/proto/connection_service"
	postGw "github.com/dislinked/common/proto/post_service"
	userGw "github.com/dislinked/common/proto/user_service"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewPostClient(address string) postGw.PostServiceClient {
	conn, err := getConnection(address)
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to User service: %v", err)
	}
	return postGw.NewPostServiceClient(conn)
}

func NewUserClient(serviceAddress string) userGw.UserServiceClient {
	conn, err := grpc.Dial(serviceAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to User service: %v", err)
	}
	return userGw.NewUserServiceClient(conn)
}

func NewConnectionClient(address string) connGw.ConnectionServiceClient {
	conn, err := getConnection(address)
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to User service: %v", err)
	}
	return connGw.NewConnectionServiceClient(conn)
}

func getConnection(address string) (*grpc.ClientConn, error) {
	return grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
}
