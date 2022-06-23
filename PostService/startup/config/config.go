package config

import (
	"flag"
	cfg "github.com/dislinked/common/config"
)

type Config struct {
	Port       string
	PostDBHost string
	PostDBPort string
}

func NewConfig() *Config {
	devEnv := flag.Bool("dev", false, "use dev environment variables")
	flag.Parse()

	if *devEnv {
		cfg.LoadEnv()
	}

	return &Config{
		Port:       "8081",
		PostDBHost: "localhost",
		PostDBPort: "27017",
	}
}
