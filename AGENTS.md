# Agent Guidelines for slacker Repository

## Build, Lint, and Test Commands

- Build the project: `go build` (builds the main binary)
- Run the CLI tool: `./slacker [command] [flags]`
- Run all tests: `go test ./...`
- Run a single test: `go test -run ^TestName$ ./...`
- Format code: `gofmt -w .`
- Lint code (if golangci-lint installed): `golangci-lint run`

## Code Style Guidelines

- Use `gofmt` for formatting; always keep code formatted.
- Use idiomatic Go naming conventions: MixedCaps for functions, variables, and types.
- Use explicit error handling with wrapped errors using `fmt.Errorf("...: %w", err)`.
- Keep imports grouped and clean; use `goimports` if possible.
- Use `cobra` for CLI commands; define commands with `RunE` for error returns.
- Use clear and concise comments where necessary.
- Use `log.Fatal` only in `main` or CLI root command for unrecoverable errors.
- Avoid global variables except for CLI state or config.

## Environment

- Requires `SLACK_TOKEN` environment variable set for Slack API access.

## Cursor and Copilot Rules

- No Cursor rules or Copilot instructions found in this repository.

---

This file is intended to guide agentic coding agents working in this repository.
