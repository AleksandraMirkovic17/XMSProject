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

	commandPublisherUpdateUser := server.initPublisher(server.config.UpdateUserCommandSubject)
	replySubscriberUpdateUser := server.initSubscriber(server.config.UpdateUserReplySubject, QueueGroup)
	orchestrator2 := server.InitUpdateOrchestrator(commandPublisherUpdateUser, replySubscriberUpdateUser)

	userService := server.initUserService(userStore, orchestrator, orchestrator2)
	userHandler := server.initUserHandler(userService, orchestrator2)

	commandSubscriber := server.initSubscriber(server.config.RegisterUserCommandSubject, QueueGroup)
	replyPublisher := server.initPublisher(server.config.RegisterUserReplySubject)
	server.initRegisterUserHandler(userService, replyPublisher, commandSubscriber)

	//Update handler
	commandSubscriberUpdateUser := server.initSubscriber(server.config.UpdateUserCommandSubject, QueueGroup)
	replyPublisherUpdateUser := server.initPublisher(server.config.UpdateUserReplySubject)
	server.initUpdateUserHandler(userService, replyPublisherUpdateUser, commandSubscriberUpdateUser)

	//friend posted notification handler
	commandSubscriberFriendPosted := server.initSubscriber(server.config.FriendPostedCommandSubject, QueueGroup)
	replyPublisherFriendPosted := server.initPublisher(server.config.FriendPostedReplySubject)
	server.initFriendPostedNotificationHandler(userService, replyPublisherFriendPosted, commandSubscriberFriendPosted)

	//connection notification handler
	commandSubscriberConNot := server.initSubscriber(server.config.ConnectionNotificationCommandSubject, QueueGroup)
	replyPublisherConNot := server.initPublisher(server.config.ConnectionNotificationReplySubject)
	server.initConnectionNotificationHandler(userService, replyPublisherConNot, commandSubscriberConNot)

	//message notificatio handler
	commandSubscriberMessNot := server.initSubscriber(server.config.MessageNotificationCommandSubject, QueueGroup)
	replyPublisherMessNot := server.initPublisher(server.config.MessageNotificationReplySubject)
	server.initMessageNotificationHandler(userService, replyPublisherMessNot, commandSubscriberMessNot)

	server.startGrpcServer(userHandler)
}

func (server *Server) InitOrchestrator(publisher saga.Publisher, subscriber saga.Subscriber) *orchestrators.UserOrchestrator {
	orchestrator, err := orchestrators.NewUserOrchestrator(publisher, subscriber)
	if err != nil {
		log.Fatal(err)
	}
	return orchestrator
}

func (server *Server) InitUpdateOrchestrator(publisher saga.Publisher, subscriber saga.Subscriber) *orchestrators.UpdateUserOrchestrator {
	orchestrator, err := orchestrators.NewUpdateUserOrchestrator(publisher, subscriber)
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

func (server *Server) initUpdateUserHandler(userService *application.UserService, publisher saga.Publisher, subscriber saga.Subscriber) {
	_, err := handlers.NewUpdateUserCommandHandler(userService, publisher, subscriber)
	if err != nil {
		log.Fatal(err)
	}
}

func (server *Server) initFriendPostedNotificationHandler(userService *application.UserService, publisher saga.Publisher, subscriber saga.Subscriber) {
	_, err := handlers.NewFriendPostedNotificationHandler(userService, publisher, subscriber)
	if err != nil {
		log.Fatal(err)
	}
}

func (server *Server) initConnectionNotificationHandler(userService *application.UserService, publisher saga.Publisher, subscriber saga.Subscriber) {
	_, err := handlers.NewConnectionNotificatioHandler(userService, publisher, subscriber)
	if err != nil {
		log.Fatal(err)
	}
}

func (server *Server) initMessageNotificationHandler(service *application.UserService, publisher saga.Publisher, subscriber saga.Subscriber) {
	_, err := handlers.NewMessageNotificatioHandler(service, publisher, subscriber)
	if err != nil {
		log.Fatal(err)
	}
}

func (server *Server) initPublisher(subject string) saga.Publisher {
	publisher, err := nats.NewNATSPublisher(
		server.config.NatsHost, server.config.NatsPort,
		server.config.NatsUser, server.config.NatsPass, subject)
	if err != nil {
		println("Da li ovdee!")
		log.Fatal(err)
		println("Posle log fatal")

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

func (server *Server) initUserService(store domain.UserStore, orchestrator *orchestrators.UserOrchestrator, orchestratorUpdate *orchestrators.UpdateUserOrchestrator) *application.UserService {
	return application.NewUserService(store, orchestrator, orchestratorUpdate)
}

func (server *Server) initUserHandler(service *application.UserService, updateOrchestrator *orchestrators.UpdateUserOrchestrator) *api.UserHandler {
	return api.NewUserHandler(service, updateOrchestrator)
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


