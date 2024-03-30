// Package config provides the configuration for the application. It contains the configuration schema and the default values for the configuration.
package config

import "time"

// Config represents the configuration for the application.
type Config struct {
	ConfigPath string `mapstructure:"config_path"`
	LogLevel   string `mapstructure:"log_level"`
	Stacktrace bool   `mapstructure:"stacktrace"`
	Get        Get    `mapstructure:"get"`
}

// Get represents the configuration for the get command.
type Get struct {
	Timeout time.Duration `mapstructure:"timeout"`
}
