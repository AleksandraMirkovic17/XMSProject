package main

import (
	"AuthenticationService/startup"
	"AuthenticationService/startup/config"
	"fmt"
)

func main() {
	config := config.NewConfig()
	server := startup.NewServer(config)
	fmt.Println("Server is starting...")
	server.Start()
}
