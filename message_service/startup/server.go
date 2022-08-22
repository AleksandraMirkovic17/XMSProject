package startup

import (
	"MessageService/application"
	"MessageService/domain"
	"MessageService/infrastructure/api"
	"MessageService/infrastructure/persistence"
	"MessageService/startup/config"
	"fmt"
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

func (server *Server) Start() {
	mongoClient := server.initMongoClient()
	messageStore := server.initMessageStore(mongoClient)

	messageService := server.initMessageService(messageStore)
	messageHandler := server.initMessageHandler(messageService)
	server.startGrpcServer(messageHandler)
}

func (server *Server) initMongoClient() *mongo.Client {
	client, err := persistence.GetClient(server.config.MessageDBHost, server.config.MessageDBPort)
	if err != nil {
		log.Fatalln(err)
	}
	return client
}

func (server *Server) initMessageStore(client *mongo.Client) domain.MessageStore {
	store := persistence.NewMessageMongoDBStore(client)
	return store
}

func (server *Server) initMessageService(store domain.MessageStore) *application.MessageService {
	return application.NewMessageService(store)
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
