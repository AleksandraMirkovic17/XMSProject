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
	saga "github.com/dislinked/common/saga/messaging"
	"github.com/dislinked/common/saga/messaging/nats"

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

func (server *Server) initRegisterUserHandler(authenticationService *application.AuthenticationService, publisher saga.Publisher, subscriber saga.Subscriber) {
	_, err := api.NewRegisterUserCommandHandler(authenticationService, publisher, subscriber)
	if err != nil {
		log.Fatal(err)
	}
}

func (server *Server) initPublisher(subject string) saga.Publisher {
	publisher, err := nats.NewNATSPublisher(
		server.config.NatsHost, server.config.NatsPort,
		server.config.NatsUser, server.config.NatsPass, subject)
	if err != nil {
		log.Fatal(err)
	}
	return publisher
}

func (server *Server) initSubscriber(subject, queueGroup string) saga.Subscriber {
	subscriber, err := nats.NewNATSSubscriber(
		server.config.NatsHost, server.config.NatsPort,
		server.config.NatsUser, server.config.NatsPass, subject, queueGroup)
	if err != nil {
		log.Fatal(err)
	}
	return subscriber
}

func (server *Server) initCreateOrderOrchestrator(publisher saga.Publisher, subscriber saga.Subscriber) *application.RegisterUserOrchestrator {
	orchestrator, err := application.NewRegisterUserOrchestrator(publisher, subscriber)
	if err != nil {
		log.Fatal(err)
	}
	return orchestrator
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
