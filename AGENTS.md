# qbo — QuickBooks Online CLI

CLI for humans and AI agents. Data goes to stdout (parseable), hints/progress to stderr.

## Build & Test

```bash
make build    # Build to bin/qbo (uses CGO_ENABLED=0 for pure Go)
make test     # Run tests with race detector
make lint     # golangci-lint
make vet      # go vet
make fmt      # gofmt
```

## Project Structure

- `cmd/qbo/main.go` — Thin entrypoint, wires Kong CLI
- `internal/cmd/` — CLI commands (Kong structs with `Run(g *Globals)` methods)
- `internal/api/` — QBO HTTP client, entity registry, query builder
- `internal/auth/` — OAuth 2.0 flow, keyring token storage
- `internal/output/` — JSON/plain/human output modes, field projection
- `internal/config/` — ~/.config/qbo/ company profiles
- `internal/errfmt/` — Exit codes, structured error types
- `skills/qbo/` — Agent skill (SKILL.md + references/)
- `docs/` — Scraped QBO API documentation

## Output Modes

- Human (default): colored tables, summaries on stderr.
- `--json` / `-j`: structured JSON to stdout. Auto-enabled when stdout is not a TTY and `QBO_AUTO_JSON=1`.
- `--plain` / `-p`: tab-separated values, no color.
- `--results-only`: strip pagination metadata, return data array only.
- `--select field1,field2`: project output to specified fields (dot paths supported).
- Desire-path aliases: `--fields` → `--select`, `--machine` → `--json`, `--tsv` → `--plain`, `--yes` → `--force`.

## Exit Codes

0=success, 1=error, 2=usage, 3=empty results, 4=auth required, 5=not found, 6=permission denied, 7=rate limited, 8=retryable, 10=config error.

## CLI Conventions

- Flags: long-form kebab-case (`--company-id`, `--start-date`).
- `--dry-run` / `-n` on all mutating commands: show what would happen, no API calls.
- `--no-input`: never prompt; fail if confirmation needed.
- `--company-id` / `QBO_COMPANY_ID`: target company. Multiple companies supported.
- `--sandbox`: target QBO sandbox environment.
- `qbo schema`: dump CLI command tree as JSON for agent introspection.

## Commits

Conventional Commits (`feat:`, `fix:`, `chore:`) with imperative summary. Keep changesets focused.
