package config

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

// Config represents service configuration for dp-nlp-search-scrubber
type Config struct {
	BindAddr                   string        `envconfig:"BIND_ADDR"`
	GracefulShutdownTimeout    time.Duration `envconfig:"GRACEFUL_SHUTDOWN_TIMEOUT"`
	HealthCheckInterval        time.Duration `envconfig:"HEALTHCHECK_INTERVAL"`
	HealthCheckCriticalTimeout time.Duration `envconfig:"HEALTHCHECK_CRITICAL_TIMEOUT"`
	ScrubberBase               string        `envconfig:"SCRUBBER_BASE"`
	BerlinBase                 string        `envconfig:"BERLIN_BASE"`
	CategoryBase               string        `envconfig:"CATEGORY_BASE"`
}

var cfg *Config

// Get returns the default config with any modifications through environment
// variables
func Get() (*Config, error) {
	cfg := &Config{}

	cfg = &Config{
		BindAddr:                   ":5000",
		GracefulShutdownTimeout:    5 * time.Second,
		HealthCheckInterval:        30 * time.Second,
		HealthCheckCriticalTimeout: 90 * time.Second,
		ScrubberBase:               "http://localhost:3002/scrubber/search",
		BerlinBase:                 "http://localhost:3001/berlin/search",
		CategoryBase:               "http://localhost:80/categories",
	}

	return cfg, envconfig.Process("", cfg)
}
