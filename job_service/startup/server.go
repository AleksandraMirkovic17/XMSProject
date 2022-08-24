package startup

import (
	"crypto/tls"
	"fmt"
	job "github.com/dislinked/common/proto/job_service"
	saga "github.com/dislinked/common/saga/messaging"
	"github.com/dislinked/common/saga/messaging/nats"
	"github.com/dislinked/job_service/application"
	"github.com/dislinked/job_service/domain"
	"github.com/dislinked/job_service/infrastructure/api"
	"github.com/dislinked/job_service/infrastructure/handlers"
	"github.com/dislinked/job_service/persistence"
	"github.com/dislinked/job_service/startup/config"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
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
	QueueGroupConnection = "job_service_connection"
)

func (server *Server) Start() {
	fmt.Println("starting job service server")
	neo4jClient := server.initNeo4J()

	connectionStore := server.initJobStore(neo4jClient)

	connectionService := server.initJobService(connectionStore)

	connectionHandler := server.initConnectionHandler(connectionService)

	//register handler
	commandSubscriber := server.initSubscriber(server.config.RegisterUserCommandSubject, QueueGroupConnection)
	replyPublisher := server.initPublisher(server.config.RegisterUserReplySubject)
	server.initRegisterUserHandler(connectionService, replyPublisher, commandSubscriber)

	//update handler
	commandSubscriberUpdate := server.initSubscriber(server.config.UpdateUserCommandSubject, QueueGroupConnection)
	replyPublisherUpdate := server.initPublisher(server.config.UpdateUserReplySubject)
	server.initUpdateUserHandler(connectionService, replyPublisherUpdate, commandSubscriberUpdate)

	server.startGrpcServer(connectionHandler)
}

func (server *Server) initRegisterUserHandler(service *application.JobService, publisher saga.Publisher, subscriber saga.Subscriber) {
	_, err := handlers.NewJobHandler(service, publisher, subscriber)
	if err != nil {
		log.Fatal(err)
	}
}

func (server *Server) initUpdateUserHandler(service *application.JobService, publisher saga.Publisher, subscriber saga.Subscriber) {
	_, err := handlers.NewUpdateUserHandler(service, publisher, subscriber)
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

func (server *Server) initJobStore(client *neo4j.Driver) domain.JobStore {
	store := persistence.NewJobDBStore(client)
	store.Init()
	return store
}

func (server *Server) initJobService(store domain.JobStore) *application.JobService {
	return application.NewJobService(store, server.config)
}

func (server *Server) initConnectionHandler(service *application.JobService) *api.JobHandler {
	return api.NewConnectionHandler(service)
}

func (server *Server) startGrpcServer(jobHandler *api.JobHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	job.RegisterJobServiceServer(grpcServer, jobHandler)
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
