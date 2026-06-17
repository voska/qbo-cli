# RecurringTransaction

A RecurringTransaction is a *template* QBO uses to create transactions on a
schedule — e.g. a monthly prepaid-expense amortization journal entry, or a
recurring invoice. Native template support lets a tool create one self-posting
schedule instead of pre-creating N future-dated transactions by hand.

## Structure

RecurringTransaction is **not** a flat entity. Each record is a *type wrapper*:
a single key naming the underlying transaction type, whose value is the normal
transaction object for that type plus a nested `RecurringInfo` block.

```json
{
  "JournalEntry": {
    "RecurringInfo": {
      "Name": "Prepaid Insurance Amortization",
      "RecurType": "Automated",
      "Active": true,
      "ScheduleInfo": {
        "IntervalType": "Monthly",
        "NumInterval": 1,
        "DayOfMonth": 1,
        "StartDate": "2026-01-01",
        "MaxOccurrences": 12
      }
    },
    "Line": [ /* normal JournalEntry lines */ ]
  }
}
```

Supported wrapped types: **Bill, CreditMemo, Deposit, Estimate, Invoice,
JournalEntry, Purchase, PurchaseOrder, RefundReceipt, SalesReceipt, Transfer,
VendorCredit**.

### RecurringInfo

| Field | Type | Description |
|-------|------|-------------|
| Name | String | Template name. Required for `Automated`/`Reminded`; should be unique per company. |
| RecurType | String | `Automated` (auto-posts), `Reminded` (reminds only), or `UnScheduled`. |
| Active | Boolean | Whether the schedule is enabled. |
| ScheduleInfo | ScheduleInfo | Recurrence pattern (below). |

### ScheduleInfo

| Field | Type | Description |
|-------|------|-------------|
| IntervalType | String | `Daily`, `Weekly`, `Monthly`, or `Yearly`. |
| NumInterval | Integer | Interval count (e.g. `NumInterval:2` + `Weekly` = every 2 weeks). |
| DayOfMonth | Integer | Day of month for Monthly/Yearly patterns (`0` lets `StartDate` drive the day). |
| DayOfWeek | String | e.g. `Tuesday`, for Weekly / nth-weekday patterns. |
| WeekOfMonth | Integer | `1`–`5` for nth-weekday-of-month patterns. |
| MonthOfYear | String | e.g. `January`, for Yearly patterns. |
| RemindDays | Integer | Days before `StartDate` to remind (for `RecurType: Reminded`). |
| DaysBefore | Integer | Days ahead `Automated` creates the transaction. |
| MaxOccurrences | Integer | Max occurrences (alternative to `EndDate`). |
| StartDate | Date | `YYYY-MM-DD`. Required for `Automated`/`Reminded`. |
| EndDate | Date | Optional end date. |
| NextDate | Date | Read-only; next creation date. |
| PreviousDate | Date | Read-only; last creation date. |

## Operations

### List / query a recurringtransaction

`SELECT * FROM RecurringTransaction` returns `QueryResponse.RecurringTransaction`
as a **heterogeneous array** of type wrappers (`{"Invoice":{…}}`, `{"Bill":{…}}`, …).

```bash
qbo list recurringtransaction --json --results-only
qbo query "SELECT * FROM RecurringTransaction" --json
```

`--results-only` returns the array of templates (an empty array when there are none).

### Read a recurringtransaction

- **Method**: GET
- **URL**: `/v3/company/{realmID}/recurringtransaction/{id}`

```bash
qbo get recurringtransaction <id> --json
```

### Create a recurringtransaction

- **Method**: POST
- **URL**: `/v3/company/{realmID}/recurringtransaction`

The request body is the bare type wrapper (`{"JournalEntry": {…}}`) including the
nested `RecurringInfo`. See [`examples/recurring-je.json`](../../examples/recurring-je.json)
for a worked Automated monthly prepaid-amortization JournalEntry.

```bash
qbo create recurringtransaction -f examples/recurring-je.json --json
```

> Account references in the example are placeholders — replace `AccountRef.value`
> with IDs from your chart of accounts (`qbo list account --json --results-only`).

### Update a recurringtransaction

- **Method**: POST
- **URL**: `/v3/company/{realmID}/recurringtransaction` (full update; the body must
  include the current `SyncToken`)

```bash
qbo update recurringtransaction <id> -f patch.json --json
```

Full update is the verified path. Sparse update (`--sparse`, adds
`?operation=update`) is available but not documented by Intuit for recurring
templates; prefer full updates.

### Delete a recurringtransaction

- **Method**: POST
- **URL**: `/v3/company/{realmID}/recurringtransaction?operation=delete`

Unlike flat entities, the delete body must be the **full type-wrapped object**
(with `Id`, `SyncToken`, and `RecurringInfo`), not just `{Id, SyncToken}`. The CLI
handles this for you: it reads the current template and echoes it back.

```bash
qbo delete recurringtransaction <id>
```
