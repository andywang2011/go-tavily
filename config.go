package tavily

import (
	"fmt"
	"time"
)

const (
	// DefaultHost is the Tavily API base URL.
	DefaultHost = "https://api.tavily.com"
	defaultTimeout = 10 * time.Second
)

// Config holds Tavily API client settings.
// Fill this struct in the application layer (not from viper inside this library).
type Config struct {
	// Host is the API base URL; defaults to DefaultHost when empty.
	Host string
	// Secret is the Tavily API key (Bearer token).
	Secret string
	// Timeout is the HTTP client timeout.
	Timeout time.Duration
}

// DefaultConfig returns defaults for tests and local development.
func DefaultConfig() Config {
	return Config{
		Host:    DefaultHost,
		Timeout: defaultTimeout,
	}
}

// ConfigDefault is an alias for DefaultConfig (naming consistency across go-* libs).
func ConfigDefault() Config { return DefaultConfig() }

// Validate checks required fields and applies defaults.
func (c *Config) Validate() error {
	if c.Host == "" {
		c.Host = DefaultHost
	}
	if c.Secret == "" {
		return fmt.Errorf("tavily config: secret is required")
	}
	if c.Timeout <= 0 {
		c.Timeout = defaultTimeout
	}
	return nil
}
