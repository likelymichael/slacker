package cmd

import (
	"log"
	"os"

	"github.com/slack-go/slack"
	"github.com/spf13/cobra"
)

var api *slack.Client

var rootCmdLong = `
       .__                 __                 
  _____|  | _____    ____ |  | __ ___________ 
 /  ___/  | \__  \ _/ ___\|  |/ // __ \_  __ \
 \___ \|  |__/ __ \\  \___|    <\  ___/|  | \/
/____  >____(____  /\___  >__|_ \\___  >__|   
     \/          \/     \/     \/    \/       
	`

var rootCmd = &cobra.Command{
	Use:               "slacker",
	Short:             "slacker is a custom CLI tool",
	Long:              rootCmdLong,
	DisableAutoGenTag: true,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		token := os.Getenv("SLACK_TOKEN")
		if token == "" {
			log.Fatal("SLACK_TOKEN env var not set")
		}
		api = slack.New(token)
	},
	Run: func(cmd *cobra.Command, args []string) {
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func SlackAPI() *slack.Client {
	return api
}

// Execute executes the root command.
func Execute() {
	rootCmd.AddCommand(NewCmdStatus())
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
