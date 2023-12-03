package config

import (
	"context"
	"errors"

	"github.com/sethvargo/go-envconfig"
)

type Config struct {
	Port string `env:"PORT,default=3002"`
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
