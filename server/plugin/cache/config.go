package cache

import (
	"fmt"
	"strings"

	"github.com/liftbridge-io/liftbridge/server/conf"
)

const (
	defaultValue = 10
)

type Config struct {
	Value int
}

// NewDefaultConfig creates a new Config with default settings.
func NewDefaultConfig() *Config {
	config := &Config{
		Value: defaultValue,
	}
	return config
}

// NewConfig creates a new Config with default settings and applies any
// settings from the given configuration file.
func NewConfig(configFile string) (*Config, error) {
	config := NewDefaultConfig()

	if configFile == "" {
		return config, nil
	}
	c, err := conf.ParseFile(configFile)
	if err != nil {
		return nil, err
	}

	for k, v := range c {
		switch strings.ToLower(k) {
		case "value":
			config.Value = int(v.(int64))
		default:
			return nil, fmt.Errorf("Unknown configuration setting %q", k)
		}
	}

	return config, nil
}
