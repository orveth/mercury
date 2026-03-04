package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/gudnuf/mercury/internal/db"
	"github.com/spf13/cobra"
)

func newSendCmd() *cobra.Command {
	var as, to string

	cmd := &cobra.Command{
		Use:   "send [body...]",
		Short: "Send a message to a channel",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := requireNonEmpty("--as", as); err != nil {
				return err
			}
			if err := requireNonEmpty("--to", to); err != nil {
				return err
			}

			var body string
			if len(args) > 0 {
				body = strings.Join(args, " ")
			} else {
				// Read from stdin if piped
				stat, _ := os.Stdin.Stat()
				if (stat.Mode() & os.ModeCharDevice) == 0 {
					scanner := bufio.NewScanner(os.Stdin)
					var lines []string
					for scanner.Scan() {
						lines = append(lines, scanner.Text())
					}
					if err := scanner.Err(); err != nil {
						return fmt.Errorf("reading stdin: %w", err)
					}
					body = strings.Join(lines, "\n")
				}
			}
			if body == "" {
				return fmt.Errorf("message body required")
			}

			d, err := db.Open()
			if err != nil {
				return err
			}
			defer d.Close()
			return d.Send(to, as, body)
		},
	}

	cmd.Flags().StringVar(&as, "as", "", "sender name (required)")
	cmd.Flags().StringVar(&to, "to", "", "channel name (required)")
	cmd.MarkFlagRequired("as")
	cmd.MarkFlagRequired("to")

	return cmd
}
