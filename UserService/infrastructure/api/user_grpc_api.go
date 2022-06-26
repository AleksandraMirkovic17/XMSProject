package api

import (
	"UserService/application"
	"UserService/domain"
	"context"
	"fmt"

	pb "github.com/dislinked/common/proto/user_service"
)

type UserHandler struct {
	pb.UnimplementedUserServiceServer
	service *application.UserService
}

func NewUserHandler(service *application.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (handler *UserHandler) GetAll(ctx context.Context, request *pb.EmptyUser) (*pb.GetAllUserResponse, error) {
	// span := tracer.StartSpanFromContextMetadata(ctx, "GetAllAPI")
	// defer span.Finish()

	// ctx = tracer.ContextWithSpan(context.Background(), span)
	// users, err := handler.service.GetAll(ctx)
	var users *[]domain.User
	var err error
	users, err = handler.service.GetAll()
	if err != nil || *users == nil {
		return nil, err
	}
	response := &pb.GetAllUserResponse{
		Users: []*pb.User{},
	}
	for _, user := range *users {
		current := mapUser(&user)
		response.Users = append(response.Users, current)
	}
	return response, nil
}

func (handler *UserHandler) Insert(ctx context.Context, request *pb.RegisterUserRequest) (*pb.User, error) {
	user := mapNewUserPbToDomain(request.User)
	fmt.Println("mapper zavrsio")

	err := handler.service.Insert(user)
	fmt.Println("Register zavrsio")
	if err != nil {
		return nil, err
	}

	return mapUser(user), nil
}

/*import (
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
}

func NewServer(config *config.Config) *Server {
	return &Server{config: config}
}

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
	store := persistence.NewUserPostgresStore(client)
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
}*/
