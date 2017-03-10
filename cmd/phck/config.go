package main

import (
	"errors"

	"github.com/nissy/phck"
	"gopkg.in/BurntSushi/toml.v0"
)

type config struct {
	Listen  string      `toml:"listen"`
	PIDFile string      `toml:"pidfile"`
	Health  phck.Health `toml:"health"`
}

func newConfig() *config {
	return &config{
		Listen: ":0",
	}
}

func (cfg *config) read(filename string) error {
	if _, err := toml.DecodeFile(filename, &cfg); err != nil {
		return err
	}

	if err := cfg.validate(); err != nil {
		return err
	}

	return nil
}

func (cfg *config) validate() error {
	if len(cfg.Health.Process) == 0 {
		return errors.New("There is no process")
	}

	return nil
}
