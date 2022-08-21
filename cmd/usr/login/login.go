package login

import (
	"context"
	"github.com/fatih/color"
	"github.com/iyear/searchx/app/usr/login"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
)

var Cmd = &cobra.Command{
	Use:   "login",
	Short: "login to Telegram",
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
		defer cancel()

		cfg, err := cmd.Flags().GetString("config")
		if err != nil {
			color.Red("value of config flag not found")
			return
		}
		if err := login.Start(ctx, cfg); err != nil {
			color.Red("login error: %v", err)
		}
	},
}
