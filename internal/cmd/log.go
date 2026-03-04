package cmd

import (
	"fmt"

	"github.com/gudnuf/mercury/internal/db"
	"github.com/spf13/cobra"
)

func newLogCmd() *cobra.Command {
	var channel string
	var limit int

	cmd := &cobra.Command{
		Use:   "log",
		Short: "Show message history",
		RunE: func(cmd *cobra.Command, args []string) error {
			d, err := db.Open()
			if err != nil {
				return err
			}
			defer d.Close()
			msgs, err := d.Log(channel, limit)
			if err != nil {
				return err
			}
			for _, m := range msgs {
				fmt.Printf("%s [%s] %s: %s\n", m.CreatedAt, m.Channel, m.Sender, m.Body)
			}
			return nil
		},
	}

	cmd.Flags().StringVar(&channel, "channel", "", "filter by channel")
	cmd.Flags().IntVar(&limit, "limit", 50, "max messages to show")

	return cmd
}
