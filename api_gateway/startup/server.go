package startup

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"

	cfg "github.com/dislinked/api_gateway/startup/config"
	authGw "github.com/dislinked/common/proto/authentication_service"
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
	return server
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
}

func (server *Server) Start() {
	print("Server started apig!")
	cors := handlers.CORS(
		handlers.AllowedOrigins([]string{
			"http://localhost:4200",
			"http://localhost:4200/**",
			"http://localhost:8080/**",
			"http://localhost:8080",
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
