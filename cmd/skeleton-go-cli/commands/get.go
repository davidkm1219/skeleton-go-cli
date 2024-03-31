package commands

import (
	"fmt"

	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"github.com/twk/skeleton-go-cli/internal/config"
)

// NewGetCmd creates a new cobra command for the get command
func NewGetCmd(v *config.Viper, l *zap.Logger) *cobra.Command {
	b := []config.BindDetail{
		{Flag: config.FlagDetail{Name: "timeout", Shorthand: "t", Description: "Sets the maximum duration for the request to complete before it is forcefully terminated.", DefaultValue: "5s"}, MapKey: "get.timeout"},
	}

	cmd := &cobra.Command{
		Use:   "get <url> [flags]",
		Short: "make a get request to the provided url",
		Long:  `The 'get' command makes a get request to the provided url.`,
		Args:  cobra.MinimumNArgs(1),
		RunE: func(_ *cobra.Command, args []string) error {
			return get(v, l, args[0])
		},
	}

	if err := v.SetFlagAndBind(cmd, b); err != nil {
		return nil
	}

	return cmd
}

func get(v *config.Viper, l *zap.Logger, url string) error {
	cfg, err := v.BuildConfig()
	if err != nil {
		return fmt.Errorf("error building config: %w", err)
	}

	l.Info("making get request", zap.String("url", url), zap.Any("config", cfg))

	return nil
}
