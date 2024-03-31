package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/viper"
)

// Viper defines the structure for holding the viper configuration.
type Viper struct {
	Viper *viper.Viper
}

// BindDetail defines the structure for holding flag information.
type BindDetail struct {
	Flag    FlagDetail
	EnvName string
	MapKey  string
}

// FlagDetail defines the structure for holding flag information.
type FlagDetail struct {
	Name         string
	Shorthand    string
	Description  string
	DefaultValue interface{}
}

// NewViper creates a new viper configuration.
func NewViper() *Viper {
	v := viper.New()
	v.SetConfigType("yaml")
	// Set the default values for the configuration
	// The order of precedence for the configuration is:
	// 1. overrides
	// 2. flags
	// 3. env. variables
	// 4. config file
	// 5. key/value store
	// 6. defaults

	return &Viper{Viper: v}
}

// BuildConfig will use the Environment variable to decide if it has to use config
// stored in artifactory or local.
func (vc *Viper) BuildConfig() (*Config, error) {
	err := vc.readConfig()
	if err != nil {
		return nil, err // Early return on error
	}

	cfg, err := vc.unmarshall()
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling config: %w", err)
	}

	return cfg, nil
}

func (vc *Viper) readConfig() error {
	configPath := vc.Viper.GetString("config_path")
	vc.Viper.SetConfigFile(configPath)

	err := vc.Viper.ReadInConfig()
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("error reading local config file: %w", err)
	}

	return nil
}

func (vc *Viper) unmarshall() (*Config, error) {
	cfg := Config{}

	if err := vc.Viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("error unmarshalling config: %w", err)
	}

	return &cfg, nil
}
