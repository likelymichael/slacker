package cmd

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

func NewCmdBrb() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "brb <minutes> [message] [emoji]",
		Short: "Set a BRB status with duration, optional message, and optional emoji",
		Args:  cobra.MinimumNArgs(1),
		RunE:  runBrb,
	}
	return cmd
}

func runBrb(cmd *cobra.Command, args []string) error {
	minutes, err := strconv.Atoi(args[0])
	if err != nil || minutes <= 0 {
		return fmt.Errorf("invalid minutes value: %s", args[0])
	}

	message := "BRB"
	emoji := ":speech_balloon:"

	if len(args) > 1 {
		lastArg := args[len(args)-1]
		if strings.HasPrefix(lastArg, ":") && strings.HasSuffix(lastArg, ":") {
			emoji = lastArg
			if len(args) > 2 {
				message = strings.Join(args[1:len(args)-1], " ")
			}
		} else {
			message = strings.Join(args[1:], " ")
		}
	}

	expiration := time.Now().Add(time.Duration(minutes) * time.Minute).Unix()

	api := SlackAPI()
	return api.SetUserCustomStatus(message, emoji, expiration)
}
