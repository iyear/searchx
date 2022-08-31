package source

import (
	"context"
	"fmt"
	"github.com/iyear/searchx/app/bot/source"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
)

var (
	src string
)

var Cmd = &cobra.Command{
	Use:     "source",
	Short:   "Import history messages",
	Example: "searchx bot source -c config/bot/config.min.yaml -f result.json",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
		defer cancel()

		cfg, err := cmd.Flags().GetString("config")
		if err != nil {
			return fmt.Errorf("get config flag failed: %v", err)
		}

		if err := source.Start(ctx, src, cfg); err != nil {
			return fmt.Errorf("source failed: %v", err)
		}
		return nil
	},
}

func init() {
	Cmd.PersistentFlags().StringVarP(&src, "file", "f", "result.json", "the path to the JSON file exported by Telegram")
}
