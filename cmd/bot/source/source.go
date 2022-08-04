package source

import (
	"github.com/fatih/color"
	"github.com/iyear/searchx/app/bot/source"
	"github.com/spf13/cobra"
)

var (
	src           string
	searchDriver  string
	searchOptions map[string]string
)

var Cmd = &cobra.Command{
	Use:   "source",
	Short: "Import history messages",
	Run: func(cmd *cobra.Command, args []string) {
		if err := source.Start(src, searchDriver, searchOptions); err != nil {
			color.Red("error happens: %v", err)
			return
		}
	},
}

func init() {
	Cmd.PersistentFlags().StringVarP(&src, "file", "f", "result.json", "the path to the JSON file exported by Telegram")
	Cmd.PersistentFlags().StringVarP(&searchDriver, "driver", "d", "", "used search engine driver")
	Cmd.PersistentFlags().StringToStringVarP(&searchOptions, "options", "o", make(map[string]string), "search engine options")
}
