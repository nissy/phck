package config

import (
	"github.com/BurntSushi/toml"
	"github.com/ngc224/phck"
)

type Config struct {
	Listen  string      `toml:"listen"  json:"listen"`
	LOGFile string      `toml:"logfile" json:"-"`
	PIDFile string      `toml:"pidfile" json:"-"`
	Health  phck.Health `toml:"health"  json:"health"`
}

func NewConfig(filePath string) (*Config, error) {
	c := &Config{}

	if _, err := toml.DecodeFile(filePath, &c); err != nil {
		return nil, err
	}

	return c, nil
}
