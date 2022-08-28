package run

import (
	"github.com/fatih/color"
	"github.com/iyear/searchx/app/bot/run"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "run",
	Short: "Run the bot",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := cmd.Flags().GetString("config")
		if err != nil {
			color.Red("value of config flag not found")
			return
		}

		run.Run(cfg)
	},
}
