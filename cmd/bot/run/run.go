package run

import (
	"github.com/iyear/searchx/app/bot/run"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "run",
	Short:   "Run the bot",
	Example: "searchx bot run -c config/bot/config.min.yaml",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := cmd.Flags().GetString("config")
		if err != nil {
			return err
		}

		return run.Run(cfg)
	},
}
