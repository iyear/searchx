package bot

import (
	"github.com/iyear/searchx/cmd/bot/query"
	"github.com/iyear/searchx/cmd/bot/run"
	"github.com/iyear/searchx/cmd/bot/source"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "bot",
	Short:   "Official Telegram Bot for group/channel owner",
	Example: "searchx bot -h",
}

func init() {
	Cmd.AddCommand(run.Cmd, query.Cmd, source.Cmd)

	Cmd.PersistentFlags().StringP("config", "c", "config/bot/config.min.yaml", "the path to the config file")
}
