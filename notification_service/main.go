package main

import (
	"NotificationService/startup"
	"NotificationService/startup/config"
)

func main() {
	config := config.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
