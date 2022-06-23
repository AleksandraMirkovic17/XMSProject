package main

import (
	"PostService/startup"
	cfg "PostService/startup/config"
)

func main() {
	config := cfg.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
