package config

import (
	"flag"

	cfg "github.com/dislinked/common/config"
)

type Config struct {
	Port       string
	AuthDBHost string
	AuthDBPort string
}

func NewConfig() *Config {
	devEnv := flag.Bool("dev", false, "use dev environment variables")
	flag.Parse()

	if *devEnv {
		cfg.LoadEnv()
	}

	return &Config{
		Port:       "4201",
		AuthDBHost: "localhost",
		AuthDBPort: "27017",
	}
}
