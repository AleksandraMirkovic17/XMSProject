package main

import (
	"UserService/startup"
	cfg "UserService/startup/config"
)

func main() {
	config := cfg.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
