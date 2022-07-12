package cmd

import (
	"github.com/fatih/color"
	"github.com/iyear/searchx/cmd/query"
	"github.com/iyear/searchx/cmd/run"
	"github.com/iyear/searchx/cmd/source"
	"github.com/iyear/searchx/global"
	"github.com/spf13/cobra"
)

var version bool

// TODO
var cmd = &cobra.Command{
	Use:   "searchx",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if version {
			color.Blue("%s\n%s", global.Version, global.GetRuntime())
		}
	},
}

func init() {
	cmd.AddCommand(run.Cmd)
	cmd.AddCommand(source.Cmd)
	cmd.AddCommand(query.Cmd)

	cmd.PersistentFlags().BoolVarP(&version, "version", "v", false, "check the version of pure-live")

}

func Execute() {
	cobra.CheckErr(cmd.Execute())
}
