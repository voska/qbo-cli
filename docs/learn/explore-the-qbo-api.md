# Explore the QuickBooks Online API

The QuickBooks Online Accounting API provides a comprehensive set of entities for managing business financial data.

## API Base URLs

- **Production**: `https://quickbooks.api.intuit.com`
- **Sandbox**: `https://sandbox-quickbooks.api.intuit.com`

## API Endpoint Pattern

```
https://quickbooks.api.intuit.com/v3/company/<realmId>/<entity>
```

- `realmId` - The QuickBooks Online company ID
- `entity` - The name of the API entity (e.g., `invoice`, `customer`, `payment`)

## Available Entities

### Most Commonly Used
- Account
- Bill
- BillPayment
- CompanyInfo
- CreditMemo
- Customer
- Deposit
- Employee
- Estimate
- Invoice
- Item
- JournalEntry
- Payment
- Purchase
- PurchaseOrder
- SalesReceipt
- Vendor
- VendorCredit

### All Entities

Account, Attachable, Bill, BillPayment, Budget, ChangeDataCapture, Class, CompanyCurrency, CompanyInfo, CreditMemo, CreditCardPayment, Customer, CustomerType, Department, Deposit, Employee, Entitlements, Estimate, Exchangerate, InventoryAdjustment, Invoice, Item, JournalCode, JournalEntry, Payment, PaymentMethod, Preferences, Purchase, PurchaseOrder, RecurringTransaction, RefundReceipt, SalesReceipt, TaxAgency, TaxClassification, TaxCode, TaxRate, TaxService, Term, TimeActivity, Transfer, Vendor, VendorCredit

### Report Entities

- BalanceSheet
- CashFlow
- CustomerBalance
- CustomerBalanceDetail
- CustomerIncome
- GeneralLedger
- InventoryValuationDetail
- InventoryValuationSummary
- ProfitAndLoss
- ProfitAndLossDetail
- SalesByClassSummary
- SalesByCustomer
- SalesByDepartment
- SalesByProduct
- TaxSummary
- TransactionList
- TransactionListByVendor
- TransactionListByCustomer
- TransactionListWithSplits
- TrialBalance
- VendorBalance
- VendorBalanceDetail
- VendorExpenses

## Common Operations

Most entities support the following operations:

- **Create** - `POST /v3/company/<realmId>/<entity>`
- **Read** - `GET /v3/company/<realmId>/<entity>/<entityId>`
- **Update** (Full) - `POST /v3/company/<realmId>/<entity>` (include all fields)
- **Update** (Sparse) - `POST /v3/company/<realmId>/<entity>` (include `sparse: true`)
- **Delete** - `POST /v3/company/<realmId>/<entity>?operation=delete`
- **Query** - `GET /v3/company/<realmId>/query?query=<selectStatement>`

## Content Types

The API supports both JSON and XML formats:
- **JSON**: `Content-Type: application/json`
- **XML**: `Content-Type: application/xml`

## Key Concepts

### Entity Types

- **List entities**: Customer, Vendor, Item, Account, etc. These have an `Active` attribute and can be made inactive.
- **Transaction entities**: Invoice, Payment, Bill, etc. These represent financial transactions.
- **Report entities**: BalanceSheet, ProfitAndLoss, etc. These are read-only reports.

### Important Fields

- **Id**: Unique identifier for each entity (read-only, system-generated)
- **SyncToken**: Version number used for optimistic locking (required for updates)
- **MetaData**: Contains `CreateTime` and `LastUpdatedTime` (read-only)
- **realmId**: The QuickBooks Online company identifier

## Related Topics

- [Data Queries](./data-queries.md) - Query syntax and operations
- [Batch Operations](./batch-operations.md) - Send multiple requests in a single call
- [Minor Versions](./minor-versions.md) - API versioning
- [API Explorer](https://developer.intuit.com/app/developer/qbo/docs/api/accounting/all-entities/account) - Interactive API reference
