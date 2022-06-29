package cmd

import (
	"github.com/iyear/searchx/cmd/run"
	"github.com/spf13/cobra"
)

// TODO
var cmd = &cobra.Command{
	Use:   "searchx",
	Short: "Hugo is a very fast static site generator",
	Long: `A Fast and Flexible Static Site Generator built with
                love by spf13 and friends in Go.
                Complete documentation is available at https://gohugo.io/documentation/`,
}

func init() {
	cmd.AddCommand(run.Cmd)
}

func Execute() {
	cobra.CheckErr(cmd.Execute())
}
