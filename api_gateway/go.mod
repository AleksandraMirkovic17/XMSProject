module github.com/dislinked/api_gateway

go 1.18

replace github.com/dislinked/common => ../common

require (
	github.com/dislinked/common v0.0.0-00010101000000-000000000000
	github.com/gorilla/handlers v1.5.1
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.10.0
	github.com/joho/godotenv v1.4.0
	go.mongodb.org/mongo-driver v1.10.1
	google.golang.org/grpc v1.47.0
	google.golang.org/protobuf v1.28.0
)

require (
	github.com/felixge/httpsnoop v1.0.1 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	golang.org/x/net v0.0.0-20220127200216-cd36cc0744dd // indirect
	golang.org/x/sys v0.0.0-20211216021012-1d35b9e2eb4e // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20220317150908-0efb43f6373e // indirect
)
