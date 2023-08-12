package config

import (
	"github.com/caarlos0/env/v6"
	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	Port     string `env:"PORT"`
	Host     string `env:"HOST"`
	Password string `env:"PASSWORD"`
	User     string `env:"USER"`
	Database string `env:"DATABASE"`
}

var instance Config

func Configs() *Config {
	if err := env.Parse(&instance); err != nil {
		panic(err)
	}
	return &instance
}
