package main

import (
	"ConnectionService/startup"
	cfg "ConnectionService/startup/config"
)

func main() {
	config := cfg.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
