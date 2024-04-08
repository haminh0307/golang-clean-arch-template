package config

import (
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"go.uber.org/fx"
)

type (
	Config struct {
		HTTP  `yaml:"http"`
		Mongo `yaml:"mongo"`
		JWT   `yaml:"jwt"`
	}

	HTTP struct {
		fx.Out

		Host string `yaml:"host" env:"HTTP_HOST" name:"HTTP_HOST"`
		Port string `yaml:"port" env:"HTTP_PORT" name:"HTTP_PORT"`
	}

	Mongo struct {
		fx.Out

		DBName  string `yaml:"db_name" env:"MONGO_DB_NAME" name:"MONGO_DB_NAME"`
		ConnURI string `yaml:"conn_uri" env:"MONGO_CONN_URI" name:"MONGO_CONN_URI"`
	}

	JWT struct {
		fx.Out

		Expiry time.Duration `yaml:"expiry" env:"JWT_EXPIRY" name:"JWT_EXPIRY"`
		Alg    string        `yaml:"alg" env:"JWT_ALG" name:"JWT_ALG"`
		Key    []byte        `yaml:"key" env:"JWT_KEY" name:"JWT_KEY"`
	}
)

func NewConfig() (Config, error) {
	var cfg Config
	err := godotenv.Load()
	if err != nil {
		return cfg, err
	}

	err = cleanenv.ReadEnv(&cfg)

	return cfg, err
}
