package main

import (
	"github.com/dislinked/api_gateway/startup"
	cfg "github.com/dislinked/api_gateway/startup/config"
)

func main() {
	config := cfg.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
