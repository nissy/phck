package config

import (
	"github.com/BurntSushi/toml"
	"github.com/ngc224/hck/model"
)

type Config struct {
	Listen string       `toml:"listen" json:"listen"`
	Health model.Health `toml:"health" json:"health"`
}

func NewConfig(filePath string) (*Config, error) {
	c := &Config{}

	if _, err := toml.DecodeFile(filePath, &c); err != nil {
		return nil, err
	}

	return c, nil
}
