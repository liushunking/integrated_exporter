package cmd

import (
	"fmt"
	"integrated-exporter/config"
	"integrated-exporter/internal/server"
	"time"

	"github.com/spf13/cobra"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: `integrated-exporter server`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		interval, err := time.ParseDuration(config.C.Server.Interval)
		if err != nil {
			return err
		}
		for _, service := range config.C.Server.HttpServices {
			duration, err := time.ParseDuration(service.Timeout)
			if err != nil {
				return err
			}
			if duration >= interval {
				return fmt.Errorf("%s: HttpService.Timeout should be smaller than Server.Interval", service.Name)
			}
		}
		for _, service := range config.C.Server.RpcServices {
			duration, err := time.ParseDuration(service.Timeout)
			if err != nil {
				return err
			}
			if duration >= interval {
				return fmt.Errorf("%s: RpcService.Timeout should be smaller than Server.Interval", service.Name)
			}
		}
		for _, service := range config.C.Server.GethServices {
			duration, err := time.ParseDuration(service.Timeout)
			if err != nil {
				return err
			}
			if duration >= interval {
				return fmt.Errorf("%s: GethService.Timeout should be smaller than Server.Interval", service.Name)
			}
		}
		for _, service := range config.C.Server.ApiServices {
			duration, err := time.ParseDuration(service.Timeout)
			if err != nil {
				return err
			}
			if duration >= interval {
				return fmt.Errorf("%s: ApiService.Timeout should be smaller than Server.Interval", service.Name)
			}
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return server.Run(config.C.Server)
	},
}

func init() {
	{
		rootCmd.AddCommand(serverCmd)

		serverCmd.Flags().StringP("port", "p", "6070", "exporter server port")
		serverCmd.Flags().StringP("interval", "i", "5s", "exporter server interval for probing")
		serverCmd.Flags().StringP("route", "r", "/metrics", "exporter server metrics route")
	}
}
