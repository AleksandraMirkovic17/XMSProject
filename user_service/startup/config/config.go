package config

import (
	"flag"

	cfg "github.com/dislinked/common/config"
)

type Config struct {
	Port       string
	UserDBHost string
	UserDBPort string
	UserDBName string
	UserDBUser string
	UserDBPass string
}

func NewConfig() *Config {
	devEnv := flag.Bool("dev", false, "use dev environment variables")
	flag.Parse()

	if *devEnv {
		cfg.LoadEnv()
	}

	return &Config{
		Port:       "8089",
		UserDBHost: "localhost",
		UserDBPort: "5432",
		UserDBName: "xws_user",
		UserDBUser: "admin",
		UserDBPass: "ftn",
	}
}
