package configs

import (
	"context"
	"errors"
	"log"

	"github.com/sethvargo/go-envconfig"
)

type Config struct {
	Port    string `env:"PORT,default=2000"`
	TodoURL string `env:"TODO_URL,required"`
	UserURL string `env:"USER_URL,required"`
}

var c *Config

func LoadConfig(ctx context.Context) (*Config, error) {
	var cfg Config
	if err := envconfig.Process(ctx, &cfg); err != nil {
		log.Println(err)
		return nil, errors.New("failed to load config")
	}

	c = &cfg
	return &cfg, nil
}

func Get() *Config {
	return c
}
