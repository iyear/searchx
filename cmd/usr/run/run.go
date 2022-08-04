package run

import "github.com/spf13/cobra"

var Cmd = &cobra.Command{
	Use:   "run",
	Short: "Run the (user)bot",
	Run: func(cmd *cobra.Command, args []string) {

	},
}
