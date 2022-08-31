package source

import (
	"context"
	"fmt"
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
	Use:     "source",
	Short:   "Source history messages",
	Example: "searchx usr source -c config/usr/config.min.yaml --from 1661681097 --to 1661691097",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
		defer cancel()

		cfg, err := cmd.Flags().GetString("config")
		if err != nil {
			return fmt.Errorf("get config flag failed: %v", err)
		}

		if from > to {
			return fmt.Errorf("`from` must be less than `to`")
		}

		// set default `to` value
		if to == 0 {
			to = int(time.Now().Unix())
		}

		if err := source.Start(ctx, cfg, from, to); err != nil {
			return fmt.Errorf("source failed: %v", err)
		}
		return nil
	},
}

func init() {
	Cmd.Flags().IntVar(&from, "from", 0, "source from this timestamp")
	Cmd.Flags().IntVar(&to, "to", 0, "source to this timestamp, default value is NOW")
}
