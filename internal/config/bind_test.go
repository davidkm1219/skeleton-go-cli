package config_test

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/twk/skeleton-go-cli/internal/config"
)

func TestViper_SetFlags(t *testing.T) {
	type args struct {
		cmd   *cobra.Command
		binds []config.BindDetail
	}
	type want struct {
		err error
	}
	tests := map[string]struct {
		name    string
		v       *Viper
		args    args
		wantErr bool
	}{
		//{
		//	name: "Test bool flag",
		//	v:    &Viper{},
		//	args: args{
		//		cmd: &cobra.Command{},
		//		binds: []BindDetail{
		//			{
		//				Flag: FlagDetail{
		//					Name:         "boolFlag",
		//					DefaultValue: true,
		//					Description:  "A boolean flag",
		//				},
		//			},
		//		},
		//	},
		//	wantErr: false,
		//},
		//{
		//	name: "Test string flag",
		//	v:    &Viper{},
		//	args: args{
		//		cmd: &cobra.Command{},
		//		binds: []BindDetail{
		//			{
		//				Flag: FlagDetail{
		//					Name:         "stringFlag",
		//					DefaultValue: "default",
		//					Description:  "A string flag",
		//				},
		//			},
		//		},
		//	},
		//	wantErr: false,
		//},
		//{
		//	name: "Test int flag",
		//	v:    &Viper{},
		//	args: args{
		//		cmd: &cobra.Command{},
		//		binds: []BindDetail{
		//			{
		//				Flag: FlagDetail{
		//					Name:         "intFlag",
		//					DefaultValue: 1,
		//					Description:  "An integer flag",
		//				},
		//			},
		//		},
		//	},
		//	wantErr: false,
		//},
		//{
		//	name: "Test duration flag",
		//	v:    &Viper{},
		//	args: args{
		//		cmd: &cobra.Command{},
		//		binds: []BindDetail{
		//			{
		//				Flag: FlagDetail{
		//					Name:         "durationFlag",
		//					DefaultValue: time.Second,
		//					Description:  "A duration flag",
		//				},
		//			},
		//		},
		//	},
		//	wantErr: false,
		//},
		//{
		//	name: "Test unsupported flag",
		//	v:    &Viper{},
		//	args: args{
		//		cmd: &cobra.Command{},
		//		binds: []BindDetail{
		//			{
		//				Flag: FlagDetail{
		//					Name:         "unsupportedFlag",
		//					DefaultValue: []string{"unsupported"},
		//					Description:  "An unsupported flag",
		//				},
		//			},
		//		},
		//	},
		//	wantErr: true,
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.v.SetFlags(tt.args.cmd, tt.args.binds); (err != nil) != tt.wantErr {
				t.Errorf("Viper.SetFlags() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
