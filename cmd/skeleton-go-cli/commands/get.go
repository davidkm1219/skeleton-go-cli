package commands

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/twk/skeleton-go-cli/internal/client"
	"github.com/twk/skeleton-go-cli/internal/photos"
	"go.uber.org/zap"

	"github.com/twk/skeleton-go-cli/internal/config"
)

// NewGetCmd creates a new cobra command for the get command
func NewGetCmd(v *config.Viper, l *zap.Logger) *cobra.Command {
	b := []config.BindDetail{
		{Flag: config.FlagDetail{Name: "timeout", Shorthand: "t", Description: "Sets the maximum duration for the request to complete before it is forcefully terminated.", DefaultValue: "5s"}, MapKey: "get.timeout"},
	}

	cmd := &cobra.Command{
		Use:   "get <concurrency>",
		Short: "make a get request to the provided url",
		Long:  `The 'get' command makes a get request to the provided url.`,
		Args:  cobra.MinimumNArgs(1),
		RunE: func(_ *cobra.Command, args []string) error {
			concurrency, err := strconv.Atoi(args[0])
			if err != nil {
				return fmt.Errorf("error converting argument to integer: %w", err)
			}
			return get(v, l, concurrency)
		},
	}

	if err := v.SetFlagAndBind(cmd, b); err != nil {
		return nil
	}

	return cmd
}

func get(v *config.Viper, l *zap.Logger, concurrency int) error {
	cfg, err := v.BuildConfig()
	if err != nil {
		return fmt.Errorf("error building config: %w", err)
	}

	l.Info("making get request", zap.Int("concurrency", concurrency), zap.Any("config", cfg))

	httpClient := &http.Client{
		Timeout: cfg.Get.Timeout,
	}
	hc := client.NewClient(httpClient)
	ps := photos.NewService(hc, l)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	result := ps.GetPhotosConcurrently(ctx, concurrency)

	l.Info("get request completed", zap.Any("result", result))

	return nil
}
