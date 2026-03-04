package cmd

import (
	"github.com/gudnuf/mercury/internal/db"
	"github.com/spf13/cobra"
)

func newSubscribeCmd() *cobra.Command {
	var as, channel string

	cmd := &cobra.Command{
		Use:   "subscribe",
		Short: "Subscribe to a channel",
		RunE: func(cmd *cobra.Command, args []string) error {
			d, err := db.Open()
			if err != nil {
				return err
			}
			defer d.Close()
			return d.Subscribe(as, channel)
		},
	}

	cmd.Flags().StringVar(&as, "as", "", "agent name (required)")
	cmd.Flags().StringVar(&channel, "channel", "", "channel to subscribe to (required)")
	cmd.MarkFlagRequired("as")
	cmd.MarkFlagRequired("channel")

	return cmd
}
