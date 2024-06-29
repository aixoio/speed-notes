package env

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadConfig() Config {
	cfg, done := loadFromSystem()
	if !done {
		godotenv.Load()
		return Config{
			Port:           os.Getenv("PORT"),
			Postgresql_url: os.Getenv("POSTGRESQL_URL"),
		}
	}
	return cfg
}

func loadFromSystem() (Config, bool) {
	port, exists := os.LookupEnv("PORT")
	if !exists {
		return Config{}, false
	}

	postgresql_url, exists := os.LookupEnv("POSTGRESQL_URL")
	if !exists {
		return Config{}, false
	}

	return Config{
		Port:           port,
		Postgresql_url: postgresql_url,
	}, true
}
