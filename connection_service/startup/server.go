package startup

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"

	"ConnectionService/application"
	"ConnectionService/domain"
	"ConnectionService/infrastructure/api"
	"ConnectionService/infrastructure/persistence"
	"ConnectionService/startup/config"

	connection "github.com/dislinked/common/proto/connection_service"
	saga "github.com/dislinked/common/saga/messaging"
	"github.com/dislinked/common/saga/messaging/nats"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
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

const (
	QueueGroup = "connection_service"
)

func (server *Server) Start() {

	neo4jClient := server.initNeo4J()

	connectionStore := server.initConnectionStore(neo4jClient)

	connectionService := server.initConnectionService(connectionStore)

	connectionHandler := server.initConnectionHandler(connectionService)

	server.startGrpcServer(connectionHandler)
}

func (server *Server) initRegisterUserHandler(connectionService *application.ConnectionService, publisher saga.Publisher, subscriber saga.Subscriber) {
	_, err := api.NewRegisterUserCommandHandler(connectionService, publisher, subscriber)
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

func (server *Server) initNeo4J() *neo4j.Driver {
	fmt.Println(fmt.Sprintf("%s://%s:%s", server.config.Neo4jUri, server.config.Neo4jHost, server.config.Neo4jPort))
	neo4jServer := fmt.Sprintf("%s://%s:%s", server.config.Neo4jUri, server.config.Neo4jHost, server.config.Neo4jPort)

	client, err := persistence.GetClient(neo4jServer, server.config.Neo4jUsername, server.config.Neo4jPassword)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (server *Server) initConnectionStore(client *neo4j.Driver) domain.ConnectionStore {
	store := persistence.NewConnectionDBStore(client)
	store.Init()
	return store
}

func (server *Server) initConnectionService(store domain.ConnectionStore) *application.ConnectionService {
	return application.NewConnectionService(store, server.config)
}

func (server *Server) initConnectionHandler(service *application.ConnectionService) *api.ConnectionHandler {
	return api.NewConnectionHandler(service)
}

func (server *Server) startGrpcServer(connectionHandler *api.ConnectionHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	connection.RegisterConnectionServiceServer(grpcServer, connectionHandler)
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
