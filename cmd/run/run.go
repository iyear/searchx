package run

import (
	"github.com/iyear/searchx/app/bot"
	"github.com/spf13/cobra"
)

var (
	cfg string
)

var Cmd = &cobra.Command{
	Use:   "run",
	Short: "Run the bot",
	Run: func(cmd *cobra.Command, args []string) {
		bot.Run(cfg)
	},
}

func init() {
	Cmd.PersistentFlags().StringVarP(&cfg, "config", "c", "config/config.yaml", "the path to the config file")
}
