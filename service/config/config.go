package config

import (
	"github.com/caarlos0/env/v6"
	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	Postgres
	JWT
	InitConfig
}
type Postgres struct {
	Port     string `env:"PORT"`
	Host     string `env:"HOST"`
	Password string `env:"PASSWORD"`
	User     string `env:"USER"`
	Database string `env:"DBNAME"`
	SslMode  string `env:"SSLMODE"`
}
type JWT struct {
	SigningKey string `env:"SIGNING"`
	Salt       string `env:"SALT"`
	TokenTTL   string `env:"TOKEN_TTL"`
}
type InitConfig struct {
	RunPort string `env:"RUN_PORT"`
}

var instance Config

func Configs() *Config {
	if err := env.Parse(&instance); err != nil {
		panic(err)
	}
	return &instance
}
