package source

import (
	"context"
	"github.com/fatih/color"
	"github.com/iyear/searchx/app/usr/source"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
)

var (
	date int
)

var Cmd = &cobra.Command{
	Use:   "source",
	Short: "source history messages",
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
		defer cancel()

		cfg, err := cmd.Flags().GetString("config")
		if err != nil {
			color.Red("value of config flag not found")
			return
		}
		if err := source.Start(ctx, cfg, date); err != nil {
			color.Red("run error: %v", err)
		}
	},
}

func init() {
	Cmd.Flags().IntVarP(&date, "date", "d", 0, "source all messages since the date")
}
