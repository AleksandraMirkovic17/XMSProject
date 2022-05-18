package config

import (
	"flag"
	"os"

	cfg "github.com/dislinked/common/config"
)

type Config struct {
	Port     string
	PostHost string
	PostPort string
}

func NewConfig() *Config {
	devEnv := flag.Bool("dev", false, "use dev environment variables")
	flag.Parse()

	if *devEnv {
		cfg.LoadEnv()
	}
	return &Config{
		Port:     os.Getenv("GATEWAY_PORT"),
		PostHost: os.Getenv("POST_SERVICE_HOST"),
		PostPort: os.Getenv("POST_SERVICE_PORT"),
	}

}
