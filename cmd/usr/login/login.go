package login

import (
	"context"
	"fmt"
	"github.com/iyear/searchx/app/usr/login"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
)

var Cmd = &cobra.Command{
	Use:     "login",
	Short:   "Login to Telegram",
	Example: "searchx usr login -c config/usr/config.min.yaml",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
		defer cancel()

		cfg, err := cmd.Flags().GetString("config")
		if err != nil {
			return fmt.Errorf("get config flag failed: %v", err)
		}

		if err := login.Start(ctx, cfg); err != nil {
			return fmt.Errorf("login failed: %v", err)
		}
		return nil
	},
}
