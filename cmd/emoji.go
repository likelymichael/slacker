package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func NewCmdEmoji() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "emoji",
		Short: "List all available Slack emojis with inline search",
		RunE:  runEmoji,
	}
	return cmd
}

func runEmoji(cmd *cobra.Command, args []string) error {
	api := SlackAPI()
	emojis, err := api.GetEmoji()
	if err != nil {
		return fmt.Errorf("failed to get emoji list: %w", err)
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Search emoji (or press enter to list all, type 'exit' to quit): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if input == "exit" {
			break
		}

		for name := range emojis {
			if input == "" || strings.Contains(name, input) {
				fmt.Println(name)
			}
		}
	}
	return nil
}
