package run

import (
	"context"
	"github.com/fatih/color"
	"github.com/iyear/searchx/app/usr/run"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
)

var (
	cfg   string
	login bool
)

var Cmd = &cobra.Command{
	Use:   "run",
	Short: "startLogin the (user)bot",
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
		defer cancel()
		if err := run.Run(ctx, cfg, login); err != nil {
			color.Red("run error: %v", err)
		}
	},
}

func init() {
	Cmd.PersistentFlags().StringVarP(&cfg, "config", "c", "config/usr/config.min.yaml", "the path to the config file")
	Cmd.PersistentFlags().BoolVar(&login, "login", false, "explicitly login to Telegram")
}
