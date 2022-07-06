package startup

import (
	"NotificationService/application"
	"NotificationService/infrastructure/api"
	"NotificationService/infrastructure/persistence"
	"NotificationService/startup/config"
	"crypto/tls"
	"fmt"
	"log"
	"net"

	notifications "github.com/dislinked/common/proto/notification_service"

	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type Server struct {
	config *config.Config
}

func NewServer(config *config.Config) *Server {
	return &Server{
		config: config,
	}
}

func (server *Server) Start() {

	mongoClient := server.initMongoClient()
	notificationStore := server.initNotificationStore(mongoClient)
	notificationService := server.initNotificationService(notificationStore)
	notificationHandler := server.initNotificationHandler(notificationService)

	fmt.Println("Notification service started.")
	server.startGrpcServer(notificationHandler)
}

func (server *Server) initMongoClient() *mongo.Client {
	client, err := persistence.GetClient(server.config.NotificationDBHost, server.config.NotificationDBPort)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (server *Server) initNotificationStore(client *mongo.Client) persistence.NotificationStore {
	store := persistence.NewNotificationMongoDbStore(client)
	return store
}

func (server *Server) initNotificationService(store persistence.NotificationStore) *application.NotificationService {
	return application.NewNotificationService(store, server.config)
}

func (server *Server) initNotificationHandler(service *application.NotificationService) *api.NotificationHandler {
	return api.NewNotificationHandler(service)
}

func (server *Server) startGrpcServer(notificationHandler *api.NotificationHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	notifications.RegisterNotificationServiceServer(grpcServer, notificationHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
		log.Fatalf("failed to serve: %s", err)
	}
}

func getConnection(address string) (*grpc.ClientConn, error) {
	config := &tls.Config{
		InsecureSkipVerify: true,
	}
	return grpc.Dial(address, grpc.WithTransportCredentials(credentials.NewTLS(config)))
}
