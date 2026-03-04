# qbo — QuickBooks Online CLI

[![CI](https://github.com/voska/qbo-cli/actions/workflows/ci.yml/badge.svg)](https://github.com/voska/qbo-cli/actions/workflows/ci.yml)
[![Release](https://img.shields.io/github/v/release/voska/qbo-cli)](https://github.com/voska/qbo-cli/releases)
[![Go](https://img.shields.io/github/go-mod/go-version/voska/qbo-cli)](https://go.dev/)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

CLI for humans and AI agents. Data goes to stdout (parseable), hints/progress to stderr.

```bash
$ qbo list invoices --where "Balance > '0'" --sandbox --json
[
  {"Id": "130", "DocNumber": "1038", "CustomerRef": {"name": "Amy's Bird Sanctuary"}, "Balance": 629.10},
  {"Id": "131", "DocNumber": "1039", "CustomerRef": {"name": "Bill's Windsurf Shop"}, "Balance": 239.00}
]

$ qbo report profit-and-loss --start-date 2025-01-01 --sandbox
Profit and Loss (2025-01-01 — 2025-12-31)
──────────────────────────────────────────
  Income                          $12,450.00
  Cost of Goods Sold               $4,200.00
  ─────────────────────────────────
  Gross Profit                     $8,250.00
  Expenses                         $3,100.00
  ─────────────────────────────────
  Net Income                       $5,150.00
```

Run `qbo --help` for the full command tree, or `qbo schema` for machine-readable introspection.

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

**Linux (deb)** — also available for arm64:

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
4. Add a **Redirect URI** — `http://localhost:8844/callback` for sandbox, or a public URI for production.

See [Intuit's OAuth 2.0 guide](https://developer.intuit.com/app/developer/qbo/docs/develop/authentication-and-authorization/oauth-2.0) for the full walkthrough.

> **Sandbox vs Production:** Development keys only work with sandbox companies. For production, complete Intuit's app assessment and use `--redirect-uri` (or `QBO_REDIRECT_URI`) with a public URI. You can use a tunnel, or any registered domain — after authorizing, paste the callback URL back into the CLI.

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
| `list\|get\|create\|update\|delete <entity>` | CRUD operations |
| `query "<sql>"` | Raw QBO query |
| `report <type>` | Financial reports |
| `batch --file` | Batch operations |
| `cdc --entities --since` | Change data capture |
| `company info\|list\|switch` | Company management |
| `schema` | CLI command tree as JSON |

All commands support `--dry-run`, `--no-input`, and `--sandbox`. Run `qbo exit-codes` for the full exit code reference.

## Development

```bash
make build    # Build to bin/qbo
make test     # Run tests
make lint     # Run linter
make vet      # Run go vet
```

## License

MIT
