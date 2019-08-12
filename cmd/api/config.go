package main

import (
	"errors"

	"github.com/BurntSushi/toml"
)

type config struct {
	Bind        string `toml:"bind"`
	AWSKey      string `toml:"aws-key"`
	AWSSecret   string `toml:"aws-secret"`
	Development bool   `toml:"development"`
}

func readConfig(path string) (*config, error) {
	if path == "" {
		return nil, errors.New("missing config file path")
	}

	conf := &config{}

	if _, err := toml.DecodeFile(path, &conf); err != nil {
		return nil, err
	}

	if conf.Bind == "" {
		return nil, errors.New("config file is missing bind key")
	}

	if conf.AWSKey == "" {
		return nil, errors.New("config file is missing aws-key key")
	}

	if conf.AWSSecret == "" {
		return nil, errors.New("config file is missing aws-secret key")
	}

	return conf, nil
}
