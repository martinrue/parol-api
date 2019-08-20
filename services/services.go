package services

import (
	"time"

	"github.com/martinrue/parol-api/logger"
)

// Config holds in-memory service config data.
type Config struct {
	AWSKey    string
	AWSSecret string
	AWSRegion string
	AWSBucket string
	Keys      []string
}

// ValidKey determines whether the key is part of the set of valid limit override keys.
func (c *Config) ValidKey(key string) bool {
	for _, k := range c.Keys {
		if k == key {
			return true
		}
	}

	return false
}

// Usage holds in-memory data about the service usage.
type Usage struct {
	Started       time.Time
	TotalRequests int
	HourlyUsage   map[string]int
}

// TrackRequest adds an entry to the hourly usage map.
func (u *Usage) TrackRequest() {
	hour := time.Now().UTC().Format("2006010215")
	u.HourlyUsage[hour]++
	u.TotalRequests++
}

//LimitExceeded returns true if the service has reached its hourly usage limit.
func (u *Usage) LimitExceeded(limit int) bool {
	hour := time.Now().UTC().Format("2006010215")
	return u.HourlyUsage[hour] >= limit
}

// Services holds shared functionality as independent service functions.
type Services struct {
	Config *Config
	Logger *logger.Logger
	Usage  *Usage
}
