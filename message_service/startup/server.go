package startup

import (
	"MessageService/application"
	"MessageService/domain"
	"MessageService/infrastructure/api"
	"MessageService/infrastructure/handler"
	"MessageService/infrastructure/orchestrator"
	"MessageService/infrastructure/persistence"
	"MessageService/startup/config"
	"fmt"
	saga "github.com/dislinked/common/saga/messaging"
	"github.com/dislinked/common/saga/messaging/nats"
	"log"
	"net"

	messageProto "github.com/dislinked/common/proto/message_service"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

type Server struct {
	config *config.Config
}

func NewServer(config *config.Config) *Server {
	return &Server{config: config}
}

const (
	QueueGroup = "message_service"
)

func (server *Server) Start() {
	mongoClient := server.initMongoClient()
	messageStore := server.initMessageStore(mongoClient)

	//Message notification orchestrator
	commandPublisher := server.initPublisher(server.config.MessageNotificationCommandSubject)
	replySubscriber := server.initSubscriber(server.config.MessageNotificationReplySubject, QueueGroup)
	orchestrator := server.InitMessageNotificationOrchestrator(commandPublisher, replySubscriber)

	messageService := server.initMessageService(messageStore, orchestrator)
	messageHandler := server.initMessageHandler(messageService)

	//Message notification handler
	commandSubscriber := server.initSubscriber(server.config.MessageNotificationCommandSubject, QueueGroup)
	replyPublisher := server.initPublisher(server.config.MessageNotificationReplySubject)
	server.initMessageNotificationHandler(messageService, replyPublisher, commandSubscriber)

	server.startGrpcServer(messageHandler)
}

func (server *Server) initMongoClient() *mongo.Client {
	client, err := persistence.GetClient(server.config.MessageDBHost, server.config.MessageDBPort)
	if err != nil {
		log.Fatalln(err)
	}
	return client
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

func (server *Server) InitMessageNotificationOrchestrator(publisher saga.Publisher, subscriber saga.Subscriber) *orchestrator.MessageNotificationOrchestrator {
	orchestrator, err := orchestrator.NewMessageNotificationOrchestrator(publisher, subscriber)
	if err != nil {
		log.Fatal(err)
	}
	return orchestrator
}

func (server *Server) initMessageNotificationHandler(service *application.MessageService, publisher saga.Publisher, subscriber saga.Subscriber) {
	_, err := handler.NewMessageNotificatioHandler(service, publisher, subscriber)
	if err != nil {
		log.Fatal(err)
	}
}

func (server *Server) initMessageStore(client *mongo.Client) domain.MessageStore {
	store := persistence.NewMessageMongoDBStore(client)
	return store
}

func (server *Server) initMessageService(store domain.MessageStore, orchestrator *orchestrator.MessageNotificationOrchestrator) *application.MessageService {
	return application.NewMessageService(store, orchestrator)
}

func (server *Server) initMessageHandler(service *application.MessageService) *api.MessageHandler {
	return api.NewMessageHandler(service)
}

func (server *Server) startGrpcServer(messageHandler *api.MessageHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	messageProto.RegisterMessageServiceServer(grpcServer, messageHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
