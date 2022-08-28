package startup

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"

	handlers1 "github.com/dislinked/api_gateway/infrastructure/handlers"
	cfg "github.com/dislinked/api_gateway/startup/config"
	authGw "github.com/dislinked/common/proto/authentication_service"
	connGw "github.com/dislinked/common/proto/connection_service"
	jobGw "github.com/dislinked/common/proto/job_service"
	messageGw "github.com/dislinked/common/proto/message_service"
	postGw "github.com/dislinked/common/proto/post_service"
	userGw "github.com/dislinked/common/proto/user_service"
	"github.com/gorilla/handlers"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Server struct {
	config *cfg.Config
	mux    *runtime.ServeMux
	closer io.Closer
}

func NewServer(config *cfg.Config) *Server {
	server := &Server{
		config: config,
		mux:    runtime.NewServeMux(),
	}
	server.initHandlers()
	server.initCustomHandlers()
	return server
}

//endpointi iz api_gateway-a
func (server *Server) initCustomHandlers() {
	/*postEndpoint := fmt.Sprintf("%s:%s", server.config.PostHost, server.config.PostPort)
	connectionEndpoint := fmt.Sprintf("%s:%s", server.config.ConnHost, server.config.ConnPort)*/

	feedHandler := handlers1.NewFeedHandler(server.config)
	feedHandler.Init(server.mux)

	postHandler := handlers1.NewPostHandler(server.config)
	postHandler.Init(server.mux)

	searchHandler := handlers1.NewSearchHandlerc(server.config)
	searchHandler.Init(server.mux)

}

func (server *Server) initHandlers() {
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	userEndpoint := fmt.Sprintf("%s:%s", server.config.UserHost, server.config.UserPort)
	err := userGw.RegisterUserServiceHandlerFromEndpoint(context.TODO(), server.mux, userEndpoint, opts)
	if err != nil {
		panic(err)
	}

	authEndpoint := fmt.Sprintf("%s:%s", server.config.AuthHost, server.config.AuthPort)
	err = authGw.RegisterAuthenticationServiceHandlerFromEndpoint(context.TODO(), server.mux, authEndpoint, opts)
	if err != nil {
		panic(err)
	}
	postEndpoint := fmt.Sprintf("%s:%s", server.config.PostHost, server.config.PostPort)
	err = postGw.RegisterPostServiceHandlerFromEndpoint(context.TODO(), server.mux, postEndpoint, opts)
	if err != nil {
		panic(err)
	}
	connEndpoint := fmt.Sprintf("%s:%s", server.config.ConnHost, server.config.ConnPort)
	err = connGw.RegisterConnectionServiceHandlerFromEndpoint(context.TODO(), server.mux, connEndpoint, opts)
	if err != nil {
		panic(err)
	}
	jobEndpoint := fmt.Sprintf("%s:%s", server.config.JobHost, server.config.JobPort)
	err = jobGw.RegisterJobServiceHandlerFromEndpoint(context.TODO(), server.mux, jobEndpoint, opts)
	if err != nil {
		panic(err)
	}
	messageEndpoint := fmt.Sprintf("%s:%s", server.config.MessageHost, server.config.MessagePort)
	err = messageGw.RegisterMessageServiceHandlerFromEndpoint(context.TODO(), server.mux, messageEndpoint, opts)
	if err != nil {
		panic(err)
	}
}

func (server *Server) Start() {
	print("Server started apig!")
	cors := handlers.CORS(
		handlers.AllowedOrigins([]string{
			"http://localhost:4200",
			"http://localhost:4200/**",
			"http://localhost:8080/**",
			"http://localhost:8080",
			"http://localhost:8081/**",
			"http://localhost:8081",
			"http://localhost:8082/**",
			"http://localhost:8082",
		}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"}),
		handlers.AllowedHeaders([]string{"Accept", "Accept-Language", "Content-Type", "Content-Language", "Origin", "Authorization", "Access-Control-Allow-Origin", "*"}),
		handlers.AllowCredentials(),
	)
	print("Cors is set up...")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", server.config.Port), cors(muxMiddleware(server))))

}

func muxMiddleware(server *Server) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		server.mux.ServeHTTP(w, r)
	})
}
