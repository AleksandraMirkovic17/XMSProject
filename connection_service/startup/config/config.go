package config

type Config struct {
	Port          string
	Host          string
	Neo4jUri      string
	Neo4jHost     string
	Neo4jPort     string
	Neo4jUsername string
	Neo4jPassword string

	UserHost string
	UserPort string
}

func NewConfig() *Config {
	return &Config{
		Host: "localhost",
		Port: "8001",

		Neo4jUri:      "bolt",
		Neo4jHost:     "localhost",
		Neo4jPort:     "7687",
		Neo4jUsername: "neo4j",
		Neo4jPassword: "connection",

		UserHost: "localhost",
		UserPort: "8089",
	}
}
