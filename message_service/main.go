package main

import (
	"MessageService/startup"
	cfg "MessageService/startup/config"
)

func main() {
	config := cfg.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
