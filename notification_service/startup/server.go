package startup

import (
	"NotificationService/application"
	"NotificationService/domain"
	"NotificationService/infrastructure/api"
	"NotificationService/infrastructure/handlers"
	orchestrators "NotificationService/infrastructure/orchestrator"
	"NotificationService/infrastructure/persistence"
	"NotificationService/startup/config"
	"fmt"
	"log"
	"net"

	pb "github.com/dislinked/common/proto/notification_service"
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
	QueueGroup = "notification_service"
)

func (server *Server) Start() {
	mongoClient := server.initMongoClient()
	notificationStore := server.initNotificationStore(mongoClient)

	//orchestratorCreate
	commandPublisher := server.initPublisher(server.config.CreateNotificationCommandSubject)
	replySubscriber := server.initSubscriber(server.config.CreateNotificationReplySubject, QueueGroup)
	orchestrator := server.InitOrchestrator(commandPublisher, replySubscriber)

	notificationService := server.initNotificationService(notificationStore, orchestrator)
	notificationHandler := server.initNotificationHandler(notificationService)

	//handler
	commandSubscriber := server.initSubscriber(server.config.CreateNotificationCommandSubject, QueueGroup)
	replyPublisher := server.initPublisher(server.config.CreateNotificationReplySubject)
	server.initCreateNotificationHandler(notificationService, replyPublisher, commandSubscriber)

	server.startGrpcServer(notificationHandler)

}

func (server *Server) initMongoClient() *mongo.Client {
	client, err := persistence.GetClient(server.config.AuthDBHost, server.config.AuthDBPort)
	if err != nil {
		log.Fatalln(err)
	}
	return client
}

func (server *Server) initNotificationStore(client *mongo.Client) domain.NotificationStore {
	store := persistence.NewNotificationMongoDBStore(client)
	return store
}

func (server *Server) initNotificationService(store domain.NotificationStore, orchestrator *orchestrators.CreateNotificationOrchestrator) *application.NotificationService {
	return application.NewNotificationService(store, orchestrator)
}

func (server *Server) initNotificationHandler(service *application.NotificationService) *api.NotificationHandler {
	return api.NewNotificationHandler(service)
}

func (server *Server) initCreateNotificationHandler(notificationService *application.NotificationService, publisher saga.Publisher, subscriber saga.Subscriber) {
	_, err := handlers.NewCreateNotificationCommandHandler(notificationService, publisher, subscriber)
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

func (server *Server) initMessageNotificationOrchestrator(publisher saga.Publisher, subscriber saga.Subscriber) *orchestrators.CreateNotificationOrchestrator {
	orchestrator, err := orchestrators.NewCreateNotificationOrchestrator(publisher, subscriber)
	if err != nil {
		log.Fatal(err)
	}
	return orchestrator
}

func (server *Server) InitOrchestrator(publisher saga.Publisher, subscriber saga.Subscriber) *orchestrators.CreateNotificationOrchestrator {
	orchestrator, err := orchestrators.NewCreateNotificationOrchestrator(publisher, subscriber)
	if err != nil {
		log.Fatal(err)
	}
	return orchestrator
}

func (server *Server) startGrpcServer(notificationHandler *api.NotificationHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterNotificationServiceServer(grpcServer, notificationHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
