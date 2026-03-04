package cmd

import (
	"fmt"

	"github.com/gudnuf/mercury/internal/db"
	"github.com/spf13/cobra"
)

func newChannelsCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "channels",
		Short: "List channels with messages",
		RunE: func(cmd *cobra.Command, args []string) error {
			d, err := db.Open()
			if err != nil {
				return err
			}
			defer d.Close()
			channels, err := d.Channels()
			if err != nil {
				return err
			}
			for _, ch := range channels {
				fmt.Println(ch)
			}
			return nil
		},
	}
}
