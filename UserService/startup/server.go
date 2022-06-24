package startup

import (
	"UserService/application"
	"UserService/domain"
	"UserService/infrastructure/persistence"
	"UserService/startup/config"
	"gorm.io/gorm"
	"log"
)

var grpcGatewayTag = otgo.Tag{Key: string(ext.Component), Value: "grpc-gateway"}

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

	server.initRegisterUserHandler(userService, replyPublisher, commandSubscriber)

	userHandler := server.initUserHandler(userService)

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
