package cmd

import (
	"context"
	"errors"
	"github.com/fatih/color"
	"github.com/iyear/searchx/cmd/bot"
	"github.com/iyear/searchx/cmd/usr"
	"github.com/iyear/searchx/cmd/version"
	"github.com/spf13/cobra"
)

var cmd = &cobra.Command{
	Use:           "searchx",
	Short:         "Enhance Telegram Group/Channel Search In 5 Minutes",
	SilenceErrors: true,
	SilenceUsage:  true,
}

func init() {
	cmd.AddCommand(bot.Cmd, usr.Cmd, version.Cmd)
}

func Execute() {
	if err := cmd.Execute(); err != nil && !errors.As(err, &context.Canceled) {
		color.Red("Error happens: %v", err)
	}
}
