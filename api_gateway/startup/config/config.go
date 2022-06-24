package config

import (
	"flag"
	cfg "github.com/dislinked/common/config"
)

type Config struct {
	Port     string
	PostHost string
	PostPort string
	UserHost string
	UserPort string
}

func NewConfig() *Config {
	devEnv := flag.Bool("dev", false, "use dev environment variables")
	flag.Parse()

	if *devEnv {
		cfg.LoadEnv()
	}
	return &Config{
		Port:     "4200",
		PostHost: "localhost",
		PostPort: "8081",
		UserHost: "localhost",
		UserPort: "8082",
	}

}
