package services

import (
	"sync"
	"time"

	"github.com/martinrue/parol-api/logger"
)

// Config holds in-memory service config data.
type Config struct {
	AWSKey      string
	AWSSecret   string
	AWSRegion   string
	AWSBucket   string
	MaxRequests int
	Keys        []string
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

	mutex *sync.Mutex
}

func (u *Usage) cleanUsage() {
	validKeys := make([]string, 0)

	for i := 0; i < 5; i++ {
		validKeys = append(validKeys, time.Now().UTC().AddDate(0, 0, i*-1).Format("20060102"))
	}

	expiredKey := func(key string) bool {
		for _, k := range validKeys {
			if k == key[0:8] {
				return false
			}
		}

		return true
	}

	for key := range u.HourlyUsage {
		if expiredKey(key) {
			delete(u.HourlyUsage, key)
		}
	}
}

// TrackRequest adds an entry to the hourly usage map.
func (u *Usage) TrackRequest() {
	u.mutex.Lock()
	defer u.mutex.Unlock()

	hour := time.Now().UTC().Format("2006010215")
	u.HourlyUsage[hour]++
	u.TotalRequests++

	u.cleanUsage()
}

// LimitExceeded returns true if the service has reached its hourly usage limit.
func (u *Usage) LimitExceeded(limit int) bool {
	u.mutex.Lock()
	defer u.mutex.Unlock()

	hour := time.Now().UTC().Format("2006010215")
	return u.HourlyUsage[hour] >= limit
}

// NewUsage constructs a usage object with initialised values.
func NewUsage() *Usage {
	return &Usage{
		Started:       time.Now(),
		TotalRequests: 0,
		HourlyUsage:   make(map[string]int, 0),
		mutex:         &sync.Mutex{},
	}
}

// Services holds shared functionality as independent service functions.
type Services struct {
	Config *Config
	Logger *logger.Logger
	Usage  *Usage
}
