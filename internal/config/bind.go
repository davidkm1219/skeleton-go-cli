package config

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

// SetFlagAndBind sets flags and binds them to the viper configuration. Use this most of the time.
// Use SetFlags and Binds separately when you share flags between commands.
func (vc *Viper) SetFlagAndBind(cmd *cobra.Command, binds []BindDetail) error {
	if err := vc.SetFlags(cmd, binds); err != nil {
		return fmt.Errorf("failed to set flags: %w", err)
	}

	if err := vc.Binds(cmd, binds); err != nil {
		return fmt.Errorf("failed to bind flags: %w", err)
	}

	return nil
}

// SetFlags sets flags for the command.
func (vc *Viper) SetFlags(cmd *cobra.Command, binds []BindDetail) error {
	for _, b := range binds {
		if b.Flag.Name != "" {
			switch defaultValue := b.Flag.DefaultValue.(type) {
			case bool:
				cmd.PersistentFlags().BoolP(b.Flag.Name, b.Flag.Shorthand, defaultValue, b.Flag.Description)
			case string:
				cmd.PersistentFlags().StringP(b.Flag.Name, b.Flag.Shorthand, defaultValue, b.Flag.Description)
			case int:
				cmd.PersistentFlags().IntP(b.Flag.Name, b.Flag.Shorthand, defaultValue, b.Flag.Description)
			case time.Duration:
				cmd.PersistentFlags().DurationP(b.Flag.Name, b.Flag.Shorthand, defaultValue, b.Flag.Description)
			default:
				return fmt.Errorf("unsupported flag type for flag %s", b.Flag.Name)
			}
		}
	}

	return nil
}

// Binds binds flags based on provided details.
func (vc *Viper) Binds(cmd *cobra.Command, binds []BindDetail) error {
	for _, b := range binds {
		if b.EnvName != "" {
			if err := vc.bindEnvDetails(b.MapKey, b.EnvName); err != nil {
				return fmt.Errorf("failed to bind environment variable: %w", err)
			}
		}

		if b.Flag.Name != "" {
			if err := vc.bindFlag(cmd, b.MapKey, b.Flag); err != nil {
				return fmt.Errorf("failed to bind flag: %w", err)
			}
		}
	}

	return nil
}

// bindFlag binds flags based on provided details.
func (vc *Viper) bindFlag(cmd *cobra.Command, mapKey string, flag FlagDetail) error {
	// Bind the current flag to a configuration key in viper.
	if err := vc.Viper.BindPFlag(mapKey, cmd.PersistentFlags().Lookup(flag.Name)); err != nil {
		return fmt.Errorf("failed to bind flag %s: %w", flag.Name, err)
	}

	return nil
}

func (vc *Viper) bindEnvDetails(mapKey, envName string) error {
	err := vc.Viper.BindEnv(mapKey, envName)
	if err != nil {
		return fmt.Errorf("failed to bind environment variable: %w", err)
	}

	return nil
}
