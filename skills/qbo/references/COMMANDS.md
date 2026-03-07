# qbo Command Reference

## Global Flags

| Flag | Short | Description |
|------|-------|-------------|
| `--json` | `-j` | Output JSON to stdout |
| `--plain` | `-p` | Output TSV, no color |
| `--select` | | Project output to fields (top-level keys only) |
| `--results-only` | | Strip QueryResponse wrapper, return entity array |
| `--company-id` | | Target company realm ID |
| `--sandbox` | | Use sandbox environment (sandbox-quickbooks.api.intuit.com) |
| `--dry-run` | `-n` | Show what would happen |
| `--no-input` | | Never prompt |
| `--force` | | Skip confirmations |
| `--verbose` | `-v` | Verbose output to stderr |
| `--minor-version` | | QBO API minor version (default: 75) |

### Aliases

- `--fields` = `--select`
- `--machine` = `--json`
- `--tsv` = `--plain`
- `--yes` / `--assume-yes` = `--force`
- `--noop` / `--preview` = `--dry-run`

## Response Shapes

All QBO responses are wrapped. Use `--json` for parsing.

| Command | Shape | Tip |
|---------|-------|-----|
| `list`, `query` | `{"QueryResponse": {"Entity": [...], "startPosition": N, ...}}` | Use `--results-only` to get just the array |
| `get` | `{"Entity": {...}}` | Pipe through `jq '.Entity'` |
| `create`, `update` | `{"Entity": {...}}` | Same as `get` |
| `delete` | `{"Entity": {"Id": "...", "status": "Deleted"}}` | |
| `report` | `{"Header": {...}, "Rows": {...}}` | Nested row structure |

## Commands

### auth

```bash
qbo auth login [--manual] [--redirect-uri URI]  # OAuth flow
qbo auth logout               # Remove stored credentials
qbo auth status               # Show token status and company
qbo auth refresh              # Force token refresh
```

For sandbox, the default `http://localhost:8844/callback` redirect works automatically. For production, pass `--redirect-uri` (or set `QBO_REDIRECT_URI`) to a registered public URI. Non-localhost URIs use a manual flow: paste the callback URL after authorizing in the browser.

### list

```bash
qbo list <entity> [--where CLAUSE] [--order-by EXPR] [--limit N] [--offset N]
```

Sugar for `query`. Entity names are case-insensitive and accept plurals (`invoices` → `Invoice`).

```bash
qbo list invoices --where "Balance > '0'" --sandbox --json --results-only
qbo list customers --where "Active = true" --limit 10 --sandbox --json --results-only
```

### get

```bash
qbo get <entity> <id>
```

Response is wrapped: `{"Invoice": {...}}`. Use jq to unwrap:

```bash
qbo get invoice 145 --sandbox --json | jq '.Invoice'
```

### create

```bash
qbo create <entity> --file <path>   # From file
qbo create <entity> -f -            # From stdin
```

Pipe JSON directly:

```bash
echo '{"DisplayName":"New Customer"}' | qbo create customer -f - --sandbox --json
```

### update

```bash
qbo update <entity> <id> --file <path> [--sparse]
```

Use `--sparse` for partial updates. Body must include `Id` and `SyncToken`:

```bash
echo '{"Id":"58","SyncToken":"0","DisplayName":"Updated Name"}' \
  | qbo update customer 58 --sparse -f - --sandbox --json
```

### delete

```bash
qbo delete <entity> <id> [--force]
```

Prompts for confirmation unless `--force` is set.

### query

```bash
qbo query "SELECT * FROM Invoice WHERE Balance > '0' ORDERBY DueDate"
```

Raw QBO SQL-like query. Combine with `--results-only` for clean arrays.

### batch

```bash
qbo batch --file batch.json
```

### cdc

```bash
qbo cdc --entities Customer,Invoice --since 2026-02-20T00:00:00Z
```

### report

```bash
qbo report <type> [--start-date YYYY-MM-DD] [--end-date YYYY-MM-DD]
```

Types: `ProfitAndLoss` / `profit-and-loss`, `BalanceSheet` / `balance-sheet`, `CashFlow` / `cash-flow`, `TrialBalance` / `trial-balance`, `GeneralLedger` / `general-ledger`, `AccountList` / `account-list`.

### company

```bash
qbo company info              # Show company info from API
qbo company list              # List configured companies
qbo company switch <realm-id> # Set default company
```

### attach

```bash
qbo attach [entity-type entity-id] <file> [--note TEXT] [--include-on-send]
```

### download

```bash
qbo download <id> [-o path] [--url]
```

`--url` prints the temporary download URL instead of saving the file.

### schema

```bash
qbo schema --json             # Full CLI tree with all entities
qbo schema get --json         # Schema for a specific command
```

### exit-codes

```bash
qbo exit-codes --json
```

## Supported Entities

Account, Attachable, Bill, BillPayment, Budget, Class, CompanyInfo, CreditMemo, Customer, Department, Deposit, Employee, Estimate, Invoice, Item, JournalEntry, Payment, PaymentMethod, Preferences, Purchase, PurchaseOrder, RefundReceipt, SalesReceipt, TaxCode, TaxRate, Term, TimeActivity, Transfer, Vendor, VendorCredit.

## Environment Variables

| Variable | Description |
|----------|-------------|
| `QBO_CLIENT_ID` | OAuth client ID (required) |
| `QBO_CLIENT_SECRET` | OAuth client secret (required) |
| `QBO_COMPANY_ID` | Default company realm ID |
| `QBO_AUTO_JSON` | Set to `1` for auto-JSON when stdout is piped |
| `QBO_CONFIG_DIR` | Override config directory (default: `~/.config/qbo/`) |
