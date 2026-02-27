# qbo — QuickBooks Online CLI

[Website](https://voska.github.io/qbo-cli/) · [Releases](https://github.com/voska/qbo-cli/releases)

CLI for humans and AI agents. Data goes to stdout (parseable), hints/progress to stderr.

## Install

**Homebrew** (macOS / Linux):

```bash
brew install voska/tap/qbo
```

**Scoop** (Windows):

```powershell
scoop bucket add voska https://github.com/voska/scoop-bucket
scoop install qbo
```

**Go**:

```bash
go install github.com/voska/qbo-cli/cmd/qbo@latest
```

**Binary**: download from [Releases](https://github.com/voska/qbo-cli/releases).

**Linux (deb)**:

```bash
curl -LO https://github.com/voska/qbo-cli/releases/latest/download/qbo_linux_amd64.deb
sudo dpkg -i qbo_linux_amd64.deb
```

## Agent Skill

Install as a [Claude Code skill](https://docs.anthropic.com/en/docs/agents-and-tools/claude-code/skills) for AI-assisted QuickBooks workflows:

```bash
npx skills add -g voska/qbo-cli
```

The skill includes setup guidance, usage patterns, troubleshooting, and a full command reference.

## Getting Credentials

You need an Intuit Developer account to get OAuth credentials.

1. Sign up at [developer.intuit.com](https://developer.intuit.com) and create an app.
2. Select **QuickBooks Online and Payments** as the platform.
3. Under **Keys & credentials**, grab your **Client ID** and **Client Secret**.
4. Add `http://localhost:8844/callback` as a **Redirect URI**.

See [Intuit's OAuth 2.0 guide](https://developer.intuit.com/app/developer/qbo/docs/develop/authentication-and-authorization/oauth-2.0) for the full walkthrough.

> **Sandbox vs Production:** Development keys only work with sandbox companies. For production, you must complete Intuit's app assessment and use a public redirect URI (not localhost).

## Quick Start

```bash
export QBO_CLIENT_ID=your_client_id
export QBO_CLIENT_SECRET=your_client_secret

# Authenticate (sandbox)
qbo auth login --sandbox

# Verify
qbo auth status
qbo company info --sandbox --json

# List invoices
qbo list invoices --sandbox

# Get a specific customer
qbo get customer 123 --sandbox

# Query with filters
qbo list invoices --where "Balance > '0'" --sandbox --json

# Run a report
qbo report profit-and-loss --start-date 2025-01-01 --end-date 2025-12-31 --sandbox
```

## Output Modes

| Flag | Description |
|------|-------------|
| (default) | Colored tables, summaries on stderr |
| `--json` / `-j` | Structured JSON to stdout |
| `--plain` / `-p` | Tab-separated values, no color |
| `--results-only` | Strip pagination metadata |
| `--select f1,f2` | Project output to specific fields |

Auto-JSON: when stdout is not a TTY and `QBO_AUTO_JSON=1`, defaults to JSON output.

## Commands

| Command | Description |
|---------|-------------|
| `auth login\|logout\|status\|refresh` | OAuth 2.0 authentication |
| `list <entity>` | List entities with optional filters |
| `get <entity> <id>` | Read a single entity |
| `create <entity>` | Create from JSON file or stdin |
| `update <entity> <id>` | Update (full or `--sparse`) |
| `delete <entity> <id>` | Delete with confirmation |
| `query "<sql>"` | Raw QBO query |
| `batch --file` | Batch operations |
| `cdc --entities --since` | Change data capture |
| `report <type>` | Financial reports |
| `company info\|list\|switch` | Company management |
| `schema` | CLI schema as JSON (agent introspection) |
| `exit-codes` | Exit code reference |

## Exit Codes

| Code | Name | Description |
|------|------|-------------|
| 0 | success | Operation completed successfully |
| 1 | error | General error |
| 2 | usage | Invalid usage or arguments |
| 3 | empty | No results found |
| 4 | auth_required | Authentication required |
| 5 | not_found | Resource not found |
| 6 | forbidden | Permission denied |
| 7 | rate_limited | API rate limit exceeded |
| 8 | retryable | Transient error, safe to retry |
| 10 | config_error | Configuration error |

## Development

```bash
make build    # Build to bin/qbo
make test     # Run tests
make lint     # Run linter
make vet      # Run go vet
```

## License

MIT
