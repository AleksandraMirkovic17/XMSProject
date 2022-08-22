package main

import (
	"MessageService/startup"
	cfg "MessageService/startup/config"
	"fmt"
)

func main() {
	config := cfg.NewConfig()
	server := startup.NewServer(config)
	fmt.Printf("Working on port:" + cfg.NewConfig().Port)
	server.Start()
}
