package startup

import (
	"AuthenticationService/application"
	"AuthenticationService/domain"
	"AuthenticationService/infrastructure/api"
	"AuthenticationService/infrastructure/persistence"
	"AuthenticationService/startup/config"
	"fmt"
	pb "github.com/dislinked/common/proto/authentication_service"
	"google.golang.org/grpc"
	"gorm.io/gorm"
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

func (server *Server) Start() {
	postgresClient := server.initPostgresClient()
	productStore := server.initProductStore(postgresClient)
	productService := server.initProductService(productStore)
	productHandler := server.initProductHandler(productService)

	server.startGrpcServer(productHandler)
}

func (server *Server) initPostgresClient() *gorm.DB {
	client, err := persistence.GetClient(
		server.config.AuthDBHost, server.config.AuthDBUser,
		server.config.AuthDBPass, server.config.AuthDBName,
		server.config.AuthDBPort)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (server *Server) initProductStore(client *gorm.DB) domain.UserStore {
	store, err := persistence.NewAuthenticationPostgresStore(client)
	if err != nil {
		log.Fatal(err)
	}
	/*store.DeleteAll()
	for _, Product := range products {
		err := store.Insert(Product)
		if err != nil {
			log.Fatal(err)
		}
	}*/
	return store
}

func (server *Server) initProductService(store domain.UserStore) *application.AuthenticationService {
	return application.NewAuthenticationService(store)
}

func (server *Server) initProductHandler(service *application.AuthenticationService) *api.AuthenticationHandler {
	return api.NewAuthenticationHandler(service)
}

func (server *Server) startGrpcServer(authenticationHandler *api.AuthenticationHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterAuthenticationServiceServer(grpcServer, authenticationHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
