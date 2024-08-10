package config

import "github.com/BurntSushi/toml"

type Config struct {
	Store struct {
		DBurl string `toml:"database_url"`
	} `toml:"store"`
}

func NewConfig(configPath string) (*Config, error) {
	var cfg Config
	_, err := toml.DecodeFile(configPath, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}