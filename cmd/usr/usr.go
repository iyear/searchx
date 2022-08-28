package usr

import (
	"github.com/iyear/searchx/cmd/usr/login"
	"github.com/iyear/searchx/cmd/usr/run"
	"github.com/iyear/searchx/cmd/usr/source"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "usr",
	Short:   "Userbot for individual users",
	Example: "searchx usr -h",
}

func init() {
	Cmd.AddCommand(run.Cmd, login.Cmd, source.Cmd)

	Cmd.PersistentFlags().StringP("config", "c", "config/usr/config.min.yaml", "the path to the config file")
}
