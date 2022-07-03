package main

import "fmt"
import "AuthenticationService/startup/config"
import "AuthenticationService/startup"

func main() {
	config := config.NewConfig()
	server := startup.NewServer(config)
	fmt.Println("Server is starting...")
	server.Start()
}
