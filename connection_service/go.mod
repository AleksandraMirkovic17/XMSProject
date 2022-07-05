module ConnectionService

go 1.18

replace github.com/dislinked/common => ../common

require (
	github.com/dislinked/common v0.0.0-00010101000000-000000000000
	github.com/joho/godotenv v1.4.0
	github.com/neo4j/neo4j-go-driver/v4 v4.4.3
	google.golang.org/genproto v0.0.0-20220617124728-180714bec0ad
	google.golang.org/grpc v1.47.0
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.10.3 // indirect
	golang.org/x/net v0.0.0-20220617184016-355a448f1bc9 // indirect
	golang.org/x/sys v0.0.0-20220615213510-4f61da869c0c // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
)
