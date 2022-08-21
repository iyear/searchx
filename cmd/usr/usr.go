package usr

import (
	"github.com/iyear/searchx/cmd/usr/login"
	"github.com/iyear/searchx/cmd/usr/run"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "usr",
	Short: "'UserBot' for individual users",
}

func init() {
	Cmd.AddCommand(run.Cmd, login.Cmd)

	Cmd.PersistentFlags().StringP("config", "c", "config/usr/config.min.yaml", "the path to the config file")
}
