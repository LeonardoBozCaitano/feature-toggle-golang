package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURI  string
	DatabaseName string
}

func NewConfig() *Config {
	godotenv.Load(".env")
	return &Config{
		DatabaseURI:  os.Getenv("DATABASE_URI"),
		DatabaseName: "feature_toggle",
	}
}
