package config

import (
	"context"
	"errors"

	"github.com/sethvargo/go-envconfig"
)

type Config struct {
	Port      string `env:"PORT,default=3002"`
	RedisHost string `env:"REDIS_HOST,required"`
	RedisPort int    `env:"REDIS_PORT,required"`
}

var c *Config

func Load(ctx context.Context) (*Config, error) {
	var cfg Config
	if err := envconfig.Process(ctx, &cfg); err != nil {
		return nil, errors.New("failed to load config")
	}

	c = &cfg
	return &cfg, nil
}

func Get() *Config {
	return c
}
