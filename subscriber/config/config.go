package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	App struct {
		Port string `yaml:"port" env:"PORT"`
		Host string `yaml:"host"`
	} `yaml:"app"`

	Db struct {
		DSN string `env:"POSTGRES_DSN" env-required:"true"`
	}

	Nats struct {
		Port      string `yaml:"port"`
		Host      string `yaml:"host"`
		ClusterID string `yaml:"cluster_id"`
		ClientID  string `yaml:"client_id"`
		Channel   string `yaml:"channel"`
	} `yaml:"nats-server"`
}

func LoadConfig(configPath string) (*Config, error) {
	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		return nil, fmt.Errorf("parse env: %w", err)
	}

	return &cfg, nil
}
