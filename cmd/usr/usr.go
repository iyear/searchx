package usr

import (
	"github.com/iyear/searchx/cmd/usr/run"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "usr",
	Short: "'UserBot' for individual users",
}

func init() {
	Cmd.AddCommand(run.Cmd)
}
