package api


import "github.com/BurntSushi/toml"


type Config struct {
	BindAddr string `toml:"bind_addr"`
}


func NewConfig(configPath string) (*Config, error) {
	var cfg Config
	_, err := toml.DecodeFile(configPath, &cfg)

	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
