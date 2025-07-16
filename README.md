# slacker

`slacker` is a tool I created so I can update my Slack status within my terminal.

## Prerequisites

- **SLACK_TOKEN** environment variable must be set to a valid Slack API token.

## Setup

1. Obtain a Slack API token with the necessary permissions to update your status.
2. Set the environment variable `SLACK_TOKEN` in your shell:

```bash
export SLACK_TOKEN="your-slack-token-here"
```

## Usage

```bash
slacker [command] [flags]
```

### Commands

#### status

Manage your Slack status.

```bash
# View your current Slack status
slacker status

# Clear your Slack status and set yourself as active
slacker clear
```

#### brb

Set a temporary "brb" status for a specified duration with an optional message and emoji.

```bash
# Set a brb status for 5 minutes with default message and emoji
slacker brb 5

# Set a brb status for 10 minutes with a custom message
slacker brb 10 "coffee break"

# Set a brb status for 15 minutes with a custom message and emoji
slacker brb 15 "coffee break" :coffee:
```

#### dnd

Set a temporary "dnd" specific status with a predefined emoji and an optional message.

```bash
slacker dnd 10
```

```bash
slacker dnd 60 "focus time"
```

#### ooo

Set a quick out-of-office status that expires at the end of the day.

```bash
slacker ooo
```

```bash
slacker ooo "not feeling great"
```

### Help

```bash
slacker --help
```

