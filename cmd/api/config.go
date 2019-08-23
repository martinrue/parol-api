package main

import (
	"errors"

	"github.com/BurntSushi/toml"
)

type config struct {
	Bind        string   `toml:"bind"`
	AWSKey      string   `toml:"aws-key"`
	AWSSecret   string   `toml:"aws-secret"`
	AWSRegion   string   `toml:"aws-region"`
	AWSBucket   string   `toml:"aws-bucket"`
	MaxRequests int      `toml:"max-hourly-requests"`
	Keys        []string `toml:"full-access-keys"`
	Development bool     `toml:"development"`
}

func readConfig(path string) (*config, error) {
	if path == "" {
		return nil, errors.New("missing config file path")
	}

	conf := &config{
		MaxRequests: -1,
	}

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

	if conf.AWSRegion == "" {
		return nil, errors.New("config file is missing aws-region key")
	}

	if conf.AWSBucket == "" {
		return nil, errors.New("config file is missing aws-bucket key")
	}

	if conf.MaxRequests == -1 {
		return nil, errors.New("config file is missing max-hourly-requests")
	}

	return conf, nil
}
