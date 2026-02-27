# qbo — QuickBooks Online CLI

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

## Quick Start

```bash
# Set up OAuth credentials
export QBO_CLIENT_ID=your_client_id
export QBO_CLIENT_SECRET=your_client_secret

# Authenticate
qbo auth login

# List invoices
qbo list invoices

# Get a specific customer
qbo get customer 123

# Query with filters
qbo list invoices --where "Balance > '0'" --json

# Run a report
qbo report profit-and-loss --start-date 2025-01-01 --end-date 2025-12-31
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

## Agent Skill

Install as a Claude Code skill:

```bash
npx skills add -g voska/qbo-cli
```

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
