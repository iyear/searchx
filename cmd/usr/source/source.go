package source

import (
	"context"
	"github.com/fatih/color"
	"github.com/iyear/searchx/app/usr/source"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"time"
)

var (
	from int
	to   int
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

		if from > to {
			color.Red("`from` must be less than `to`")
			return
		}

		if err := source.Start(ctx, cfg, from, to); err != nil {
			color.Red("source error: %v", err)
		}
	},
}

func init() {
	Cmd.Flags().IntVar(&from, "from", 0, "source from this timestamp")
	Cmd.Flags().IntVar(&to, "to", int(time.Now().Unix()), "source to this timestamp, default value is NOW")
}
