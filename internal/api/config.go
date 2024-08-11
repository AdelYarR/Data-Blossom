package api


import "github.com/BurntSushi/toml"


type Config struct {
	BindAddr string `toml:"bind_addr"`
}


func NewConfig(configPath string) (*Config, error) {
	var cfg Config
	if _, err := toml.DecodeFile(configPath, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
