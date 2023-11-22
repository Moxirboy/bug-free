package configs

import (
	"github.com/caarlos0/env/v6"

	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	AppName    string `env:"APP_NAME"`
	AppVersion string `env:"APP_VERSION"`
	Postgres
	InitConfig
}
type Postgres struct {
	Port     string `env:"POSTGRES_PORT"`
	Host     string `env:"POSTGRES_HOST"`
	Password string `env:"POSTGRES_PASSWORD"`
	User     string `env:"POSTGRES_USER"`
	Database string `env:"POSTGRES_DBNAME"`
	SslMode  string `env:"POSTGRES_SSLMODE"`
}
type InitConfig struct {
	RunPort string `env:"RUN_PORT"`
}

var instance Config

func Configuration() *Config {

	if err := env.Parse(&instance); err != nil {
		panic(err)
	}
	return &instance

}
