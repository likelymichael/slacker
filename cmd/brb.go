package cmd

import (
	"time"

	"github.com/spf13/cobra"
)

func BrbStatus() *cobra.Command {
	var emoji string
	var length int
	var msg string

	brbCmd := &cobra.Command{
		Use:   "brb",
		Short: "Set status to 'brb'",
		Args:  cobra.RangeArgs(0, 2),
		RunE:  runBrbCmd,
	}

	brbCmd.Flags().StringVarP(&emoji, "emoji", "e", ":speech_balloon:", "(Optional) Emoji to display. Default: :speech_balloon:")
	brbCmd.Flags().IntVarP(&length, "length", "l", 1, "Length of status in 1 hour increments. Default: 1")
	brbCmd.Flags().StringVarP(&msg, "message", "m", "brb", "(Optional) Custom message to display. Default: brb")

	return brbCmd
}

func runBrbCmd(cmd *cobra.Command, args []string) error {
	length, lenErr := cmd.Flags().GetInt("length")
	if lenErr != nil {
		return lenErr
	}

	emoji, emojiErr := cmd.Flags().GetString("emoji")
	if emojiErr != nil {
		return emojiErr
	}

	msg, msgErr := cmd.Flags().GetString("message")
	if msgErr != nil {
		return msgErr
	}

	var expiration int64
	if length > 0 {
		expiration = time.Now().Add(time.Duration(length) * time.Hour).Unix()
	} else {
		expiration = 0
	}

	api := SlackAPI()
	return api.SetUserCustomStatus(msg, emoji, expiration)
}
