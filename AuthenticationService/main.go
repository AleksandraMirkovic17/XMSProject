package main

import "fmt"

func main() {
	config := config.NewConfig()
	server := startup.NewServer(config)
	fmt.Println("Server is starting...")
	server.Start()
}
