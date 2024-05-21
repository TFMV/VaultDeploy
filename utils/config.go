package utils

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Postgres  PostgresConfig  `yaml:"postgres" form:"postgres"`
	DataVault DataVaultConfig `yaml:"data_vault" form:"data_vault"`
}

type PostgresConfig struct {
	Host     string `yaml:"host" form:"host"`
	Port     int    `yaml:"port" form:"port"`
	User     string `yaml:"user" form:"user"`
	Password string `yaml:"password" form:"password"`
	DBName   string `yaml:"dbname" form:"dbname"`
	SSLMode  string `yaml:"sslmode" form:"sslmode"`
}

type DataVaultConfig struct {
	Schema     string            `yaml:"schema" form:"schema"`
	Hubs       []HubConfig       `yaml:"hubs" form:"hubs"`
	Links      []LinkConfig      `yaml:"links" form:"links"`
	Satellites []SatelliteConfig `yaml:"satellites" form:"satellites"`
}

type HubConfig struct {
	Name    string   `yaml:"name" form:"name"`
	Columns []string `yaml:"columns" form:"columns"`
}

type LinkConfig struct {
	Name    string   `yaml:"name" form:"name"`
	Columns []string `yaml:"columns" form:"columns"`
}

type SatelliteConfig struct {
	Name    string   `yaml:"name" form:"name"`
	Columns []string `yaml:"columns" form:"columns"`
}

func LoadConfig(configPath string) (*Config, error) {
	config := &Config{}
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(data, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
