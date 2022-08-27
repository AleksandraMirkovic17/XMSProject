package startup

import (
	"PostService/application"
	"PostService/domain"
	"PostService/infrastructure/api"
	"PostService/infrastructure/handlers"
	"PostService/infrastructure/orchestrators"
	"PostService/infrastructure/persistence"
	"PostService/startup/config"
	"fmt"
	"log"
	"net"

	postProto "github.com/dislinked/common/proto/post_service"
	saga "github.com/dislinked/common/saga/messaging"
	"github.com/dislinked/common/saga/messaging/nats"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

const (
	QueueGroupFriendPosted = "friend_posted_notification"
)

type Server struct {
	config *config.Config
}

func NewServer(config *config.Config) *Server {
	return &Server{config: config}
}

func (server *Server) Start() {
	print("usao u start")
	mongoClient := server.initMongoClient()
	print("init mongo client")
	postStore := server.initPostStore(mongoClient)

	//friend posted orchestrator
	commandPublisher := server.initPublisher(server.config.FriendPostedCommandSubject)
	replySubsciber := server.initSubscriber(server.config.FriendPostedReplySubject, QueueGroupFriendPosted)
	orchestrator := server.InitOrchestrator(commandPublisher, replySubsciber)

	postService := server.initPostService(postStore, orchestrator)
	postHandler := server.initPostHandler(postService)

	//friend posted orchestrator
	commandSuscriberFriendPosted := server.initSubscriber(server.config.FriendPostedCommandSubject, QueueGroupFriendPosted)
	replyPublisherFriendPosted := server.initPublisher(server.config.FriendPostedReplySubject)
	server.initFriendPostedNotificationHandler(postService, replyPublisherFriendPosted, commandSuscriberFriendPosted)

	server.startGrpcServer(postHandler)
}

func (server *Server) initMongoClient() *mongo.Client {
	client, err := persistence.GetClient(server.config.PostDBHost, server.config.PostDBPort)
	println("Na kom sam portuu")
	fmt.Println(server.config.PostDBHost)
	if err != nil {
		log.Fatalln(err)
	}
	return client
}

func (server *Server) initPostStore(client *mongo.Client) domain.PostStore {
	store := persistence.NewPostMongoDBStore(client)
	return store
}

func (server *Server) initFriendPostedNotificationHandler(service *application.PostService, publisher saga.Publisher, subscriber saga.Subscriber) {
	_, err := handlers.NewFriendPostedNotificationHandler(service, publisher, &subscriber)
	if err != nil {
		log.Fatal(err)
	}
}

func (server *Server) initPostService(store domain.PostStore, orchestrator *orchestrators.FriendPostedNotificationOrchestrator) *application.PostService {
	return application.NewPostService(store, orchestrator)
}

func (server *Server) initPostHandler(service *application.PostService) *api.PostHandler {
	return api.NewPostHandler(service)
}

func (server *Server) startGrpcServer(postHandler *api.PostHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	postProto.RegisterPostServiceServer(grpcServer, postHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
		log.Fatalf("failed to serve: %s", err)
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

func (server *Server) InitOrchestrator(publisher saga.Publisher, subscriber saga.Subscriber) *orchestrators.FriendPostedNotificationOrchestrator {
	orchestrator, err := orchestrators.NewFriendPostedNotificationOrchestrator(publisher, subscriber)
	if err != nil {
		log.Fatal(err)
	}
	return orchestrator
}
