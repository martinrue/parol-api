package services

import (
	"github.com/martinrue/parol-api/logger"
)

// Config holds in-memory service config data.
type Config struct {
	AWSKey    string
	AWSSecret string
}

// Services holds shared functionality as independent service functions.
type Services struct {
	Config   *Config
	Logger   *logger.Logger
}
