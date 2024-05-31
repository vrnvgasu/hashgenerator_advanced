package config

import (
	"fmt"
	"net/url"
)

var Cfg *Config

type Config struct {
	User     string
	Password string
	Host     string
	Port     string
	DBName   string
}

func NewConfig(user string, pass string, host string, port string, dbName string) *Config {
	return &Config{
		User:     user,
		Password: pass,
		Host:     host,
		Port:     port,
		DBName:   dbName,
	}
}

func (c *Config) PostgresCon() string {
	return fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable&connect_timeout=%d",
		"postgres",
		url.QueryEscape(c.User),
		c.Password,
		c.Host,
		c.Port,
		c.DBName,
		5,
	)
}
