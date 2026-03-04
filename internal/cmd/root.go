package cmd

import (
	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	root := &cobra.Command{
		Use:   "mercury",
		Short: "Inter-agent communication bus",
	}

	root.AddCommand(
		newSendCmd(),
		newReadCmd(),
		newSubscribeCmd(),
		newUnsubscribeCmd(),
		newChannelsCmd(),
		newLogCmd(),
	)

	return root
}
