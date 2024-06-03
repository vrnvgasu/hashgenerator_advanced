package config

import (
	"service1/internal/lg"

	"github.com/ilyakaznacheev/cleanenv"
)

var Cfg *Config

type Config struct {
	DebugLevel int    `env:"APP_CONFIG_DEBUG_LEVEL" env-default:"5"`
	Port       string `env:"APP_CONFIG_PORT" env-required:"true"`
}

func MustLoad() *Config {
	cfg := &Config{}
	err := cleanenv.ReadEnv(cfg)
	if err != nil {
		lg.Fatal("Load", "config", err)
	}

	return cfg
}
