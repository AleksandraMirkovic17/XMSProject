package main

import (
	"github.com/dislinked/job_service/startup"
	cfg "github.com/dislinked/job_service/startup/config"
)

func main() {
	config := cfg.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
