package config

type Config struct {
	Port       string
	AuthDBHost string
	AuthDBPort string
	AuthDBName string
	AuthDBUser string
	AuthDBPass string
}

func NewConfig() *Config {
	return &Config{
		Port:       "4201",
		AuthDBHost: "localhost",
		AuthDBPort: "5432",
		AuthDBName: "authentication_user",
		AuthDBUser: "root",
		AuthDBPass: "root",
	}
}
