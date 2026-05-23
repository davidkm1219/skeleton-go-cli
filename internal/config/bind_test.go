package config_test

import (
	"errors"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"

	"github.com/twk/skeleton-go-cli/internal/config"
)

func TestViper_SetFlags(t *testing.T) {
	t.Parallel()

	type args struct {
		binds []config.BindDetail
	}

	type want struct {
		err error
	}

	tests := map[string]struct {
		args args
		want want
	}{
		"Test valid flag": {
			args: args{
				binds: []config.BindDetail{
					{Flag: config.FlagDetail{Name: "boolFlag", Description: "A boolean flag"}, DefaultValue: true},
					{Flag: config.FlagDetail{Name: "stringFlag", Description: "A string flag"}, DefaultValue: "default"},
					{Flag: config.FlagDetail{Name: "intFlag", Description: "An integer flag"}, DefaultValue: 1},
					{Flag: config.FlagDetail{Name: "durationFlag", Description: "A duration flag"}, DefaultValue: 1},
				},
			},
			want: want{err: nil},
		},
		"Test unsupported flag": {
			args: args{
				binds: []config.BindDetail{
					{Flag: config.FlagDetail{Name: "unsupportedFlag", Description: "An unsupported flag"}, DefaultValue: []string{"unsupported"}},
				},
			},
			want: want{err: errors.New("unsupported flag type for flag unsupportedFlag")},
		},
	}
	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			cmd := &cobra.Command{}
			v := config.NewViper()

			err := v.SetFlags(cmd, tt.args.binds)
			if tt.want.err != nil {
				assert.EqualError(t, err, tt.want.err.Error())
			}
		})
	}
}

func TestViper_Binds(t *testing.T) {
	type args struct {
		binds []config.BindDetail
		env   map[string]string
	}

	type want struct {
		expected interface{}
		err      error
	}

	tests := map[string]struct {
		args args
		want want
	}{
		"bind with flag": {
			args: args{
				binds: []config.BindDetail{
					{Flag: config.FlagDetail{Name: "boolFlag", Description: "A boolean flag"}, DefaultValue: true, MapKey: "boolFlag", EnvName: "BOOL_ENV"},
				},
			},
			want: want{expected: true},
		},
		"bind with env": {
			args: args{
				binds: []config.BindDetail{
					{MapKey: "boolFlag", EnvName: "BOOL_ENV"},
				},
				env: map[string]string{
					"BOOL_ENV": "true",
				},
			},
			want: want{expected: "true"},
		},
	}
	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			for k, v := range tt.args.env {
				t.Setenv(k, v)
			}

			cmd := &cobra.Command{}
			v := config.NewViper()

			err := v.SetFlags(cmd, tt.args.binds)
			assert.NoError(t, err)

			err = v.Binds(cmd, tt.args.binds)
			if tt.want.err != nil {
				assert.EqualError(t, err, tt.want.err.Error())
				return
			}

			value := v.Viper.Get(tt.args.binds[0].MapKey)
			assert.Equal(t, tt.want.expected, value)
		})
	}
}
