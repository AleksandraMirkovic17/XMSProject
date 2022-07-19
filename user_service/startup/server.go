package startup

import (
	"UserService/application"
	"UserService/domain"
	"UserService/infrastructure/api"
	"UserService/infrastructure/handlers"
	"UserService/infrastructure/orchestrators"
	"UserService/infrastructure/persistence"
	"UserService/startup/config"
	"fmt"
	userProto "github.com/dislinked/common/proto/user_service"
	saga "github.com/dislinked/common/saga/messaging"
	"github.com/dislinked/common/saga/messaging/nats"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
	config *config.Config
	// tracer otgo.Tracer
	// closer io.Closer
}

func NewServer(config *config.Config) *Server {
	return &Server{
		config: config,
	}
}

const (
	QueueGroup = "user_service"
)

func (server *Server) Start() {
	mongoClient := server.initMongoClient()
	userStore := server.initUserStore(mongoClient)

	//orchestrator
	commandPublisher := server.initPublisher(server.config.RegisterUserCommandSubject)
	replySubscriber := server.initSubscriber(server.config.RegisterUserReplySubject, QueueGroup)
	orchestrator := server.InitOrchestrator(commandPublisher, replySubscriber)

	userService := server.initUserService(userStore, orchestrator)
	userHandler := server.initUserHandler(userService)

	//handler
	commandSubscriber := server.initSubscriber(server.config.RegisterUserCommandSubject, QueueGroup)
	replyPublisher := server.initPublisher(server.config.RegisterUserReplySubject)
	server.initRegisterUserHandler(userService, replyPublisher, commandSubscriber)

	server.startGrpcServer(userHandler)
}

func (server *Server) InitOrchestrator(publisher saga.Publisher, subscriber saga.Subscriber) *orchestrators.UserOrchestrator {
	orchestrator, err := orchestrators.NewUserOrchestrator(publisher, subscriber)
	if err != nil {
		log.Fatal(err)
	}
	return orchestrator
}

func (server *Server) initRegisterUserHandler(userService *application.UserService, publisher saga.Publisher, subscriber saga.Subscriber) {
	_, err := handlers.NewRegisterUserCommandHandler(userService, publisher, subscriber)
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

func (server *Server) initMongoClient() *mongo.Client {
	client, err := persistence.GetClient(server.config.UserDBHost, server.config.UserDBPort)
	fmt.Println(server.config.UserDBHost)
	if err != nil {
		log.Fatalln(err)
	}
	return client
}

func (server *Server) initUserStore(client *mongo.Client) domain.UserStore {
	store := persistence.NewUserMongoDBStore(client)
	return store
}

func (server *Server) initUserService(store domain.UserStore, orchestrator *orchestrators.UserOrchestrator) *application.UserService {
	return application.NewUserService(store, orchestrator)
}

func (server *Server) initUserHandler(service *application.UserService) *api.UserHandler {
	return api.NewUserHandler(service)
}

func (server *Server) startGrpcServer(userHandler *api.UserHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	userProto.RegisterUserServiceServer(grpcServer, userHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
		log.Fatalf("failed to serve: %s", err)
	}
}
