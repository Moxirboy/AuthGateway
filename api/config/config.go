package config

import (
	"github.com/caarlos0/env/v6"
	_ "github.com/joho/godotenv/autoload"
)

// Config ...
type Config struct {
	Environment string `env:"ENVIRONMENT"`

	UserServiceUrl string `env:"PORT"`

	UserServiceHost string `env:"USER_SERVICE_HOST"`
	UserServicePort int    `env:"USER_SERVICE_PORT"`

	LogLevel string `env:"LOG_LEVEL"`

	HttpPort string `env:"HTTP_PORT"`
}

var instance Config

// Load loads environment vars and inflates Config
func Load() Config {
	err := env.Parse(&instance)
	if err != nil {
		panic(err)
	}
	return instance
}
