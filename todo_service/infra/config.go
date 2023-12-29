package infra

import (
	"context"
	"fmt"

	"github.com/sethvargo/go-envconfig"
)

type Config struct {
	Port       string `env:"PORT,default=8080"`
	DBUser     string `env:"DB_USER,required"`
	DBPassword string `env:"DB_PASSWORD,required"`
	DBHost     string `env:"DB_HOST,required"`
	DBPort     string `env:"DB_PORT,required"`
	DBName     string `env:"DB_NAME,required"`
	DBSSLMode  string `env:"DB_SSM_MODE,default=disable"`
	DBTimeZone string `env:"DB_TIME_ZONE,default=Asia/Tokyo"`
	DBMaxConn  int    `env:"DB_MAX_CONN,default=0"`
	DBMaxIdle  int    `env:"DB_MAX_IDLE,default=0"`
}

var c *Config

func LoadConfig(ctx context.Context) (*Config, error) {
	var cfg Config
	if err := envconfig.Process(ctx, &cfg); err != nil {
		return nil, err
	}

	c = &cfg
	return &cfg, nil
}

func GetConfig() *Config {
	return c
}

func (r *Config) DNS() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		r.DBHost, r.DBUser, r.DBPassword, r.DBName, r.DBPort, r.DBSSLMode, r.DBTimeZone,
	)
}
