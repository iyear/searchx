package run

import (
	"context"
	"fmt"
	"github.com/iyear/searchx/app/usr/run"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
)

var Cmd = &cobra.Command{
	Use:     "run",
	Short:   "Run the user and search bot",
	Example: "searchx usr run -c config/usr/config.min.yaml",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
		defer cancel()

		cfg, err := cmd.Flags().GetString("config")
		if err != nil {
			return fmt.Errorf("get config flag failed: %v", err)
		}

		if err := run.Run(ctx, cfg); err != nil {
			return fmt.Errorf("run failed: %v", err)
		}
		return nil
	},
}
