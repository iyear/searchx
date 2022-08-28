package run

import (
	"github.com/iyear/searchx/app/bot/run"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "run",
	Short: "Run the bot",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := cmd.Flags().GetString("config")
		if err != nil {
			return err
		}

		return run.Run(cfg)
	},
}
