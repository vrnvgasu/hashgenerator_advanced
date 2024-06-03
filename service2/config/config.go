package config

import (
	"context"
	"fmt"
	"net/url"
	"service2/internal/lg"

	"github.com/ilyakaznacheev/cleanenv"
)

var Cfg *Config

type DebugConfig struct {
	Level int `env:"APP_CONFIG_DEBUG_LEVEL" env-default:"5"`
}

type HTTPConfig struct {
	Port int `env:"APP_CONFIG_PORT" env-default:"5"`
}

type DBConfig struct {
	User     string `env:"APP_CONFIG_POSTGRES_USER" env-required:"true"`
	Password string `env:"APP_CONFIG_POSTGRES_PASSWORD" env-required:"true"`
	Host     string `env:"APP_CONFIG_POSTGRES_HOST" env-required:"true"`
	Port     string `env:"APP_CONFIG_POSTGRES_PORT" env-required:"true"`
	DBName   string `env:"APP_CONFIG_POSTGRES_DB" env-required:"true"`
}

type GRPCConfig struct {
	Host string `env:"APP_CONFIG_GRPC_HOST" env-required:"true"`
	Port string `env:"APP_CONFIG_GRPC_PORT" env-required:"true"`
}

type Config struct {
	DBConfig
	GRPCConfig
	DebugConfig
	HTTPConfig
}

func (c *Config) PostgresCon() string {
	return fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable&connect_timeout=%d",
		"postgres",
		url.QueryEscape(c.User),
		c.DBConfig.Password,
		c.DBConfig.Host,
		c.DBConfig.Port,
		c.DBConfig.DBName,
		5,
	)
}

func MustLoad() *Config {
	cfg := &Config{}
	err := cleanenv.ReadEnv(cfg)
	if err != nil {
		lg.Fatal(context.Background(), "Load", "config", err)
	}

	return cfg
}
