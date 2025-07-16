package cmd

import (
	"strings"
	"time"

	"github.com/spf13/cobra"
)

func NewCmdOoo() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ooo [message]",
		Short: "Set an out of office status with optional message",
		Args:  cobra.MaximumNArgs(1),
		RunE:  runOoo,
	}
	return cmd
}

func runOoo(cmd *cobra.Command, args []string) error {
	message := "Out of office"
	if len(args) == 1 {
		message = strings.Join(args, " ")
	}

	now := time.Now()
	endOfDay := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, now.Location())
	expiration := endOfDay.Unix()

	api := SlackAPI()
	emoji := ":palm_tree:"
	return api.SetUserCustomStatus(message, emoji, expiration)
}
