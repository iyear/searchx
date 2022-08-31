package cmd

import (
	"github.com/fatih/color"
	"github.com/iyear/searchx/cmd/bot"
	"github.com/iyear/searchx/cmd/usr"
	"github.com/iyear/searchx/cmd/version"
	"github.com/iyear/searchx/global"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
	"path/filepath"
)

var cmd = &cobra.Command{
	Use:               "searchx",
	Short:             "Enhance Telegram Search In 5 Minutes",
	Example:           "searchx -h",
	DisableAutoGenTag: true,
	SilenceErrors:     true,
	SilenceUsage:      true,
}

func init() {
	cmd.AddCommand(bot.Cmd, usr.Cmd, version.Cmd)

	if err := doc.GenMarkdownTree(cmd, filepath.Join(global.DocsPath, "command")); err != nil {
		color.Red("generate cmd docs failed: %v", err)
		return
	}
}

func Execute() {
	if err := cmd.Execute(); err != nil {
		color.Red("%v", err)
	}
}
