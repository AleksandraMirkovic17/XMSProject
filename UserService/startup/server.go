package startup

import (
	"UserService/application"
	"UserService/domain"
	"UserService/infrastructure/api"
	"UserService/infrastructure/persistence"
	"UserService/startup/config"
	"fmt"
	"log"
	"net"

	userProto "github.com/dislinked/common/proto/user_service"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type Server struct {
	config *config.Config
	// tracer otgo.Tracer
	// closer io.Closer
}

func NewServer(config *config.Config) *Server {
	// newTracer, closer := tracer.Init(config.JaegerServiceName)
	// otgo.SetGlobalTracer(newTracer)
	return &Server{
		config: config,
		// tracer: newTracer,
		// closer: closer,
	}
}

const (
	QueueGroup = "user_service"
)

func (server *Server) Start() {
	postgresClient := server.initUserClient()
	userStore := server.initUserStore(postgresClient)
	userService := server.initUserService(userStore)
	userHandler := server.initUserHandler(userService)

	//server.initRegisterUserHandler(userService, replyPublisher, commandSubscriber)

	server.startGrpcServer(userHandler)
}

func (server *Server) initUserClient() *gorm.DB {
	client, err := persistence.GetClient(
		server.config.UserDBHost, server.config.UserDBUser,
		server.config.UserDBPass, server.config.UserDBName,
		server.config.UserDBPort)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (server *Server) initUserStore(client *gorm.DB) domain.UserStore {
	store, err := persistence.NewUserPostgresStore(client)
	if err != nil {
		log.Fatal(err)
	}
	// store.DeleteAll()
	// for _, Product := range products {
	// 	err := store.Register(Product)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// }
	return store
}

func (server *Server) initUserService(store domain.UserStore) *application.UserService {
	return application.NewUserService(store)
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
