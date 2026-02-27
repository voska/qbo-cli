---
name: qbo
description: Queries, creates, updates, and manages QuickBooks Online data via the qbo CLI. Use when working with QBO entities (invoices, customers, bills, payments, vendors, accounts, items, estimates), running reports, or setting up install, auth, sandbox, and company switching.
allowed-tools: Bash(qbo *)
---

# qbo — QuickBooks Online CLI

## Install

```bash
# macOS / Linux
brew install voska/tap/qbo

# Scoop (Windows)
scoop bucket add voska https://github.com/voska/scoop-bucket && scoop install qbo

# Go (any platform)
go install github.com/voska/qbo-cli/cmd/qbo@latest

# Binary: https://github.com/voska/qbo-cli/releases
```

## Getting Credentials

Requires an Intuit Developer account and a QuickBooks app for OAuth credentials.

1. Sign up at https://developer.intuit.com and create an app from the dashboard.
2. Select **QuickBooks Online and Payments** as the platform.
3. Under **Keys & credentials**, find your **Client ID** and **Client Secret**.
4. Add `http://localhost:8844/callback` as a **Redirect URI**.

See [Intuit's OAuth 2.0 guide](https://developer.intuit.com/app/developer/qbo/docs/develop/authentication-and-authorization/oauth-2.0) for the full walkthrough.

**Sandbox vs Production:** Development keys only work with sandbox companies. For production access, complete Intuit's app assessment and use a public redirect URI (not localhost). Use a domain routed through a tunnel, or a non-resolving domain — when the redirect errors in the browser, copy the full callback URL and paste it into the terminal.

## Setup

```bash
export QBO_CLIENT_ID=your_client_id
export QBO_CLIENT_SECRET=your_client_secret

# Authenticate (opens browser, listens on localhost:8844)
qbo auth login

# Or for sandbox
qbo auth login --sandbox

# Or print the URL without opening a browser (useful for agents/remote)
qbo auth login --manual
```

Tokens are stored at `~/.config/qbo/tokens/`.

### Verify

After auth, confirm everything works:

```bash
qbo auth status
qbo company info --sandbox --json
```

If you get "no company ID", set one: `export QBO_COMPANY_ID=<realm_id>` or `qbo company switch <realm_id>`.

## Troubleshooting

| Problem | Fix |
|---------|-----|
| `no company ID` | `export QBO_COMPANY_ID=<realm>` or `qbo company switch <realm>` |
| `ApplicationAuthorizationFailed` on production endpoint | Use `--sandbox` or re-authorize the app for a production company |
| OAuth state mismatch | Restart `qbo auth login` and use the newly generated URL |
| Token expired | `qbo auth refresh` or re-run `qbo auth login` |

## Response Structure

QBO wraps all responses. Know the shapes:

- **`list`/`query`** returns `{"QueryResponse": {"Invoice": [...], "startPosition": 1, ...}}`.
  Use `--results-only` to unwrap to just the array.
- **`get`** returns `{"Invoice": {...}}`. Use jq to drill in: `qbo get invoice 123 --json | jq '.Invoice'`
- **`create`/`update`** returns the same wrapper as `get`.
- **`report`** returns `{"Header": {...}, "Rows": {...}}`.

Always use `--json` when parsing output programmatically.

## Common Patterns

```bash
# List with filtering — --results-only gives you the array directly
qbo list customers --sandbox --json --results-only
qbo list invoices --where "Balance > '0'" --sandbox --json --results-only

# Get by ID — drill into entity key with jq
qbo get invoice 145 --sandbox --json | jq '.Invoice'
qbo get customer 58 --sandbox --json | jq '.Customer | {Id, DisplayName, Balance}'

# Create from stdin
echo '{"DisplayName":"Acme Corp"}' | qbo create customer -f - --sandbox --json

# Create from file
qbo create invoice -f invoice.json --sandbox --json

# Record a payment against an invoice
echo '{"CustomerRef":{"value":"58"},"TotalAmt":500,"Line":[{"Amount":500,"LinkedTxn":[{"TxnId":"145","TxnType":"Invoice"}]}]}' | qbo create payment -f - --sandbox --json

# Sparse update (partial)
echo '{"Id":"58","SyncToken":"0","DisplayName":"New Name"}' | qbo update customer 58 --sparse -f - --sandbox --json

# Reports
qbo report profit-and-loss --start-date 2026-01-01 --end-date 2026-12-31 --sandbox --json
qbo report balance-sheet --sandbox --json

# Raw query
qbo query "SELECT Id, DisplayName, Balance FROM Customer WHERE Active = true" --sandbox --json --results-only
```

## Workflow: Create -> Invoice -> Payment

```bash
# 1. Create customer
echo '{"DisplayName":"New Client","PrimaryEmailAddr":{"Address":"client@example.com"}}' \
  | qbo create customer -f - --sandbox --json | jq '.Customer.Id'

# 2. Create item (service)
echo '{"Name":"Consulting","Type":"Service","IncomeAccountRef":{"value":"1"},"UnitPrice":150}' \
  | qbo create item -f - --sandbox --json | jq '.Item.Id'

# 3. Create invoice (use IDs from above)
echo '{"CustomerRef":{"value":"CUSTOMER_ID"},"Line":[{"Amount":1500,"DetailType":"SalesItemLineDetail","SalesItemLineDetail":{"ItemRef":{"value":"ITEM_ID"},"Qty":10,"UnitPrice":150}}]}' \
  | qbo create invoice -f - --sandbox --json | jq '.Invoice | {Id, DocNumber, TotalAmt}'

# 4. Record payment
echo '{"CustomerRef":{"value":"CUSTOMER_ID"},"TotalAmt":1500,"Line":[{"Amount":1500,"LinkedTxn":[{"TxnId":"INVOICE_ID","TxnType":"Invoice"}]}]}' \
  | qbo create payment -f - --sandbox --json
```

## Agent Introspection

```bash
qbo schema --json             # Full CLI tree, all entities, all flags
qbo schema get --json         # Schema for a specific command
qbo exit-codes --json         # Exit codes as JSON
```

## Exit Codes

0=success, 1=error, 2=usage, 3=empty, 4=auth required, 5=not found, 6=forbidden, 7=rate limited, 8=retryable, 10=config error.

## Reference

See [references/COMMANDS.md](references/COMMANDS.md) for full command reference.
