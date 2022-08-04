package cmd

import (
	"github.com/iyear/searchx/cmd/bot"
	"github.com/iyear/searchx/cmd/version"
	"github.com/spf13/cobra"
)

var cmd = &cobra.Command{
	Use:   "searchx",
	Short: "Enhance Telegram Group/Channel Search In 5 Minutes",
}

func init() {
	cmd.AddCommand(bot.Cmd, version.Cmd)
}

func Execute() {
	cobra.CheckErr(cmd.Execute())
}
