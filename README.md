# slacker

`slacker` is a tool I created so I can update my Slack status within my terminal.

## Prerequisites

- **SLACK_TOKEN** environment variable must be set to a valid Slack API token.

## Usage

```bash
slacker [command] [flags]
```

### Commands

#### status

Manage your Slack status.

```bash
# View your current Slack status
slacker status get

# Clear your Slack status and set yourself as active
slacker status clear
```

#### brb

Set a temporary "brb" status for a specified duration.

```bash
# Set default brb status (:speech_balloon:, 1 hour)
slacker brb

# Customize the brb status
slacker brb --emoji :coffee: --length 2 --message "coffee break"
```

### Help

```bash
slacker --help
```

