package cmd

import (
	"github.com/spf13/cobra"
)

func NewCmdClear() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "clear",
		Short: "Clear the current Slack status",
		Args:  cobra.NoArgs,
		RunE:  runClear,
	}
	return cmd
}

func runClear(cmd *cobra.Command, args []string) error {
	api := SlackAPI()
	return api.UnsetUserCustomStatus()
}
