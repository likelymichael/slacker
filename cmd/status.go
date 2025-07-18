package cmd

import (
	"fmt"
	"time"

	"github.com/slack-go/slack"
	"github.com/spf13/cobra"
)

func NewCmdStatus() *cobra.Command {
	statusCmd := &cobra.Command{
		Use:   "status",
		Short: "Show the current Slack status",
		Args:  cobra.NoArgs,
		RunE:  runStatus,
	}
	return statusCmd
}

func runStatus(cmd *cobra.Command, args []string) error {
	api := SlackAPI()

	auth, authErr := api.AuthTest()
	if authErr != nil {
		return fmt.Errorf("auth failed: %w", authErr)
	}

	user, userErr := api.GetUserProfile(&slack.GetUserProfileParameters{
		UserID: auth.UserID,
	})
	if userErr != nil {
		return fmt.Errorf("failed to get user profile: %w", userErr)
	}

	fmt.Printf("Current status: %s %s\n", user.StatusEmoji, user.StatusText)
	if user.StatusExpiration > 0 {
		expiration := time.Unix(int64(user.StatusExpiration), 0)
		fmt.Printf("Expires %s\n", expiration.Local().Format(time.RFC1123))
	} else {
		fmt.Println("Expires: never")
	}

	return nil
}
