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
			Port: os.Getenv("PORT"),
		}
	}
	return cfg
}

func loadFromSystem() (Config, bool) {
	port, exists := os.LookupEnv("PORT")
	if !exists {
		return Config{}, false
	}

	return Config{Port: port}, true
}
