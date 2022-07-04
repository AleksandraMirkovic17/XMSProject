package startup

import (
	"AuthenticationService/application"
	"AuthenticationService/domain"
	"AuthenticationService/infrastructure/api"
	"AuthenticationService/infrastructure/persistence"
	"AuthenticationService/startup/config"
	"fmt"
	"log"
	"net"

	pb "github.com/dislinked/common/proto/authentication_service"

	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

type Server struct {
	config *config.Config
}

func NewServer(config *config.Config) *Server {
	return &Server{
		config: config,
	}
}

const (
	QueueGroup = "authentication_service"
)

func (server *Server) Start() {
	mongoClient := server.initMongoClient()
	authenticationStore := server.initAuthenticationStore(mongoClient)

	authenticationService := server.initAuthenticationService(authenticationStore)
	authenticationHandler := server.initAuthenticationHandler(authenticationService)
	server.startGrpcServer(authenticationHandler)
}

func (server *Server) initMongoClient() *mongo.Client {
	client, err := persistence.GetClient(server.config.AuthDBHost, server.config.AuthDBPort)
	if err != nil {
		log.Fatalln(err)
	}
	return client
}

func (server *Server) initAuthenticationStore(client *mongo.Client) domain.UserStore {
	store := persistence.NewAuthenticationMongoDBStore(client)
	return store
}

func (server *Server) initAuthenticationService(store domain.UserStore) *application.AuthenticationService {
	return application.NewAuthenticationService(store)
}

func (server *Server) initAuthenticationHandler(service *application.AuthenticationService) *api.AuthenticationHandler {
	return api.NewAuthenticationHandler(service)
}

func (server *Server) startGrpcServer(authenticationHandler *api.AuthenticationHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterAuthenticationServiceServer(grpcServer, authenticationHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
