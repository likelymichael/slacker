package cmd

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

func NewCmdDnd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "dnd <minutes> [message]",
		Short: "Set a Do Not Disturb status with duration and optional message",
		Args:  cobra.MinimumNArgs(1),
		RunE:  runDnd,
	}
	return cmd
}

func runDnd(cmd *cobra.Command, args []string) error {
	minutes, err := strconv.Atoi(args[0])
	if err != nil || minutes <= 0 {
		return fmt.Errorf("invalid minutes value: %s", args[0])
	}

	message := "Do Not Disturb"
	if len(args) > 1 {
		message = strings.Join(args[1:], " ")
	}

	expiration := time.Now().Add(time.Duration(minutes) * time.Minute).Unix()

	api := SlackAPI()
	emoji := ":no_bell:"
	return api.SetUserCustomStatus(message, emoji, expiration)
}
