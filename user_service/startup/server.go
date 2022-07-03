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
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
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
	print("usao u start")
	mongoClient := server.initMongoClient()
	print("init mongo client")
	userStore := server.initUserStore(mongoClient)

	userService := server.initUserService(userStore)
	userHandler := server.initUserHandler(userService)
	server.startGrpcServer(userHandler)
}

func (server *Server) initMongoClient() *mongo.Client {
	client, err := persistence.GetClient(server.config.UserDBHost, server.config.UserDBPort)
	println("Na kom sam portuu")
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
