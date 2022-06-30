package source

import (
	"github.com/fatih/color"
	"github.com/iyear/searchx/app/source"
	"github.com/spf13/cobra"
)

var (
	src           string
	searchDriver  string
	searchOptions map[string]string
)

var Cmd = &cobra.Command{
	Use:   "source",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if err := source.Start(src, searchDriver, searchOptions); err != nil {
			color.Red("error happens: %v", err)
			return
		}
	},
}

func init() {
	Cmd.PersistentFlags().StringVarP(&src, "file", "f", "result.json", "The path to the JSON file exported by Telegram")
	Cmd.PersistentFlags().StringVarP(&searchDriver, "driver", "d", "", "Used search driver")
	Cmd.PersistentFlags().StringToStringVarP(&searchOptions, "options", "o", make(map[string]string), "")
}
