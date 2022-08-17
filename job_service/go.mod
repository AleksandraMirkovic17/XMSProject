module github.com/dislinked/job_service

go 1.18

replace github.com/dislinked/common => ../common

require (
	github.com/dislinked/common v0.0.0-00010101000000-000000000000
	github.com/joho/godotenv v1.4.0
	github.com/neo4j/neo4j-go-driver/v4 v4.4.4
	go.mongodb.org/mongo-driver v1.10.1
	golang.org/x/net v0.0.0-20220127200216-cd36cc0744dd
	google.golang.org/grpc v1.47.0
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.10.0 // indirect
	github.com/nats-io/nats-server/v2 v2.8.4 // indirect
	github.com/nats-io/nats.go v1.16.0 // indirect
	github.com/nats-io/nkeys v0.3.0 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	golang.org/x/crypto v0.0.0-20220622213112-05595931fe9d // indirect
	golang.org/x/sys v0.0.0-20220111092808-5a964db01320 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20220317150908-0efb43f6373e // indirect
	google.golang.org/protobuf v1.28.0 // indirect
)
