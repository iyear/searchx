package cmd

import (
	"github.com/iyear/searchx/cmd/query"
	"github.com/iyear/searchx/cmd/run"
	"github.com/iyear/searchx/cmd/source"
	"github.com/iyear/searchx/cmd/version"
	"github.com/spf13/cobra"
)

// TODO
var cmd = &cobra.Command{
	Use:   "searchx",
	Short: "Enhance Telegram Group/Channel Search In 5 Minutes",
}

func init() {
	cmd.AddCommand(run.Cmd)
	cmd.AddCommand(source.Cmd)
	cmd.AddCommand(query.Cmd)
	cmd.AddCommand(version.Cmd)
}

func Execute() {
	cobra.CheckErr(cmd.Execute())
}
