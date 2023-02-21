package config

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetDefaultConfig(t *testing.T) {
	// Call the Get function to get the default configuration
	config, err := Get()
	assert.Nil(t, err)

	// Assert that the configuration has the default values
	assert.Equal(t, ":5000", config.BindAddr)
	assert.Equal(t, 5*time.Second, config.GracefulShutdownTimeout)
	assert.Equal(t, 30*time.Second, config.HealthCheckInterval)
	assert.Equal(t, 90*time.Second, config.HealthCheckCriticalTimeout)
	assert.Equal(t, "http://localhost:3002/scrubber/search", config.ScrubberBase)
	assert.Equal(t, "http://localhost:3001/berlin/search", config.BerlinBase)
	assert.Equal(t, "http://localhost:80/categories", config.CategoryBase)
}

func TestGetConfigFromEnv(t *testing.T) {
	// Set environment variables to modify the default configuration
	os.Setenv("BIND_ADDR", ":8080")
	os.Setenv("GRACEFUL_SHUTDOWN_TIMEOUT", "10s")
	os.Setenv("HEALTHCHECK_INTERVAL", "60s")
	os.Setenv("HEALTHCHECK_CRITICAL_TIMEOUT", "180s")
	os.Setenv("SCRUBBER_BASE", "scrubber")
	os.Setenv("BERLIN_BASE", "berlin")
	os.Setenv("CATEGORY_BASE", "category")

	// Call the Get function to get the modified configuration
	config, err := Get()
	assert.Nil(t, err)

	// Assert that the configuration has the modified values
	assert.Equal(t, ":8080", config.BindAddr)
	assert.Equal(t, 10*time.Second, config.GracefulShutdownTimeout)
	assert.Equal(t, 60*time.Second, config.HealthCheckInterval)
	assert.Equal(t, 180*time.Second, config.HealthCheckCriticalTimeout)
	assert.Equal(t, "scrubber", config.ScrubberBase)
	assert.Equal(t, "berlin", config.BerlinBase)
	assert.Equal(t, "category", config.CategoryBase)

	// Unset the environment variables
	os.Unsetenv("BIND_ADDR")
	os.Unsetenv("GRACEFUL_SHUTDOWN_TIMEOUT")
	os.Unsetenv("HEALTHCHECK_INTERVAL")
	os.Unsetenv("HEALTHCHECK_CRITICAL_TIMEOUT")
	os.Unsetenv("SCRUBBER_BASE")
	os.Unsetenv("BERLIN_BASE")
	os.Unsetenv("CATEGORY_BASE")
}
