package cmd

import (
	"github.com/gudnuf/mercury/internal/db"
	"github.com/spf13/cobra"
)

func newUnsubscribeCmd() *cobra.Command {
	var as, channel string

	cmd := &cobra.Command{
		Use:   "unsubscribe",
		Short: "Unsubscribe from a channel",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := requireNonEmpty("--as", as); err != nil {
				return err
			}
			if err := requireNonEmpty("--channel", channel); err != nil {
				return err
			}

			d, err := db.Open()
			if err != nil {
				return err
			}
			defer d.Close()
			return d.Unsubscribe(as, channel)
		},
	}

	cmd.Flags().StringVar(&as, "as", "", "agent name (required)")
	cmd.Flags().StringVar(&channel, "channel", "", "channel to unsubscribe from (required)")
	cmd.MarkFlagRequired("as")
	cmd.MarkFlagRequired("channel")

	return cmd
}
