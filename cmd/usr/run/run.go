package run

import (
	"github.com/iyear/searchx/app/usr/run"
	"github.com/spf13/cobra"
)

var (
	cfg   string
	login bool
)

var Cmd = &cobra.Command{
	Use:   "run",
	Short: "startLogin the (user)bot",
	Run: func(cmd *cobra.Command, args []string) {
		run.Run(cfg, login)
	},
}

func init() {
	Cmd.PersistentFlags().StringVarP(&cfg, "config", "c", "config/usr/config.min.yaml", "the path to the config file")
	Cmd.PersistentFlags().BoolVar(&login, "login", false, "explicitly login to Telegram")
}
