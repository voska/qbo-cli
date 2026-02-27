# Account

Accounts are what businesses use to track transactions. Accounts can track money coming in (income or revenue) and going out (expenses). They can also track the value of things (assets), like vehicles and equipment. There are five basic account types: asset, liability, income, expense, and equity. Accounts are part of the chart of accounts, the unique list of accounts each business puts together to do their accounting.

Note: If you need to delete an account, set the Active attribute to false in an object update request. This makes it inactive. The account itself is not permanently deleted, but is hidden for display purposes. References to inactive objects remain intact.

## Attributes

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| Id | String | Required for update | Unique identifier for this object. Sort order is ASC by default. |
| Name | String | Required | User recognizable name for the Account. Must not contain double quotes or colon. Max 100 characters. |
| SyncToken | String | Required for update | Version number of the object. Used to lock an object for use by one app at a time. |
| AcctNum | String | Conditionally required | User-defined account number to help identify the account within the chart-of-accounts. Must not contain colon. |
| CurrencyRef | CurrencyRefType | Read-only | Reference to the currency in which this account holds amounts. |
| ParentRef | ReferenceType | Optional | Specifies the Parent AccountId if this represents a SubAccount. |
| Description | String | Optional | User entered description for the account. Max 100 characters. |
| Active | Boolean | Optional | Whether or not active. Inactive accounts may be hidden from most display purposes and may not be posted to. |
| MetaData | ModificationMetaData | Optional | Descriptive information about the object. The MetaData values are set by Data Services and are read only for all applications. |
| SubAccount | Boolean | Read-only | Specifies whether this object represents a parent (false) or subaccount (true). |
| Classification | String | Read-only | The classification of an account. Valid values: Asset, Equity, Expense, Liability, Revenue. |
| FullyQualifiedName | String | Read-only | Fully qualified name of the object. Takes the form of Parent:Account1:SubAccount1:SubAccount2. Limited to 5 levels. |
| TxnLocationType | String | Optional | The account location. Valid values: WithinFrance, FranceOverseas, OutsideFranceWithEU, OutsideEU. For France locales only. |
| AccountType | String | Required | A detailed account classification that specifies the use of this account. The type is based on the Classification. |
| CurrentBalanceWithSubAccounts | Decimal | Read-only | Specifies the cumulative balance amount for the current Account and all its sub-accounts. |
| AccountAlias | String | Optional | A user friendly name for the account. It must be unique across all account categories. For France locales only. |
| TaxCodeRef | ReferenceType | Optional | Reference to the default tax code used by this account. Available with minorversion=3. For global locales only. |
| AccountSubType | String | Optional | The account sub-type classification, based on the AccountType value. |
| CurrentBalance | Decimal | Read-only | Specifies the balance amount for the current Account. Valid for Balance Sheet accounts. |

## Sample Object

```json
{
  "Account": {
    "FullyQualifiedName": "MyJobs",
    "domain": "QBO",
    "Name": "MyJobs",
    "Classification": "Asset",
    "AccountSubType": "AccountsReceivable",
    "CurrencyRef": {
      "name": "United States Dollar",
      "value": "USD"
    },
    "CurrentBalanceWithSubAccounts": 0,
    "sparse": false,
    "MetaData": {
      "CreateTime": "2014-12-31T09:29:05-08:00",
      "LastUpdatedTime": "2014-12-31T09:29:05-08:00"
    },
    "AccountType": "Accounts Receivable",
    "CurrentBalance": 0,
    "Active": true,
    "SyncToken": "0",
    "Id": "94",
    "SubAccount": false
  },
  "time": "2014-12-31T09:29:05.717-08:00"
}
```

## Operations

### Create an account

- **Method**: POST
- **URL**: `/v3/company/{realmID}/account`

• Name must be unique. • The Account.Name attribute must not contain double quotes (") or colon (:).

**Request Body**:
```json
{
  "Name": "MyJobs_test",
  "AccountType": "Accounts Receivable"
}
```

**Response**:
```json
{
  "Account": {
    "FullyQualifiedName": "MyJobs",
    "domain": "QBO",
    "Name": "MyJobs",
    "Classification": "Asset",
    "AccountSubType": "AccountsReceivable",
    "CurrencyRef": {
      "name": "United States Dollar",
      "value": "USD"
    },
    "CurrentBalanceWithSubAccounts": 0,
    "sparse": false,
    "MetaData": {
      "CreateTime": "2014-12-31T09:29:05-08:00",
      "LastUpdatedTime": "2014-12-31T09:29:05-08:00"
    },
    "AccountType": "Accounts Receivable",
    "CurrentBalance": 0,
    "Active": true,
    "SyncToken": "0",
    "Id": "94",
    "SubAccount": false
  },
  "time": "2014-12-31T09:29:05.717-08:00"
}
```

### Query an account

- **Method**: GET
- **URL**: `/v3/company/{realmID}/query?query={selectStatement}`

**Sample Query**: `select * from Account where Metadata.CreateTime > '2014-12-31'`

**Response**:
```json
{
  "QueryResponse": {
    "startPosition": 1,
    "Account": [
      {
        "FullyQualifiedName": "Canadian Accounts Receivable",
        "domain": "QBO",
        "Name": "Canadian Accounts Receivable",
        "Classification": "Asset",
        "AccountSubType": "AccountsReceivable",
        "CurrencyRef": {
          "name": "United States Dollar",
          "value": "USD"
        },
        "CurrentBalanceWithSubAccounts": 0,
        "sparse": false,
        "MetaData": {
          "CreateTime": "2015-06-23T09:38:18-07:00",
          "LastUpdatedTime": "2015-06-23T09:38:18-07:00"
        },
        "AccountType": "Accounts Receivable",
        "CurrentBalance": 0,
        "Active": true,
        "SyncToken": "0",
        "Id": "92",
        "SubAccount": false
      },
      {
        "FullyQualifiedName": "MyClients",
        "domain": "QBO",
        "Name": "MyClients",
        "Classification": "Asset",
        "AccountSubType": "AccountsReceivable",
        "CurrencyRef": {
          "name": "United States Dollar",
          "value": "USD"
        },
        "CurrentBalanceWithSubAccounts": 0,
        "sparse": false,
        "MetaData": {
          "CreateTime": "2015-07-13T12:34:47-07:00",
          "LastUpdatedTime": "2015-07-13T12:34:47-07:00"
        },
        "AccountType": "Accounts Receivable",
        "CurrentBalance": 0,
        "Active": true,
        "SyncToken": "0",
        "Id": "93",
        "SubAccount": false
      },
      {
        "FullyQualifiedName": "MyJobs",
        "domain": "QBO",
        "Name": "MyJobs",
        "Classification": "Asset",
        "AccountSubType": "AccountsReceivable",
        "CurrencyRef": {
          "name": "United States Dollar",
          "value": "USD"
        },
        "CurrentBalanceWithSubAccounts": 0,
        "sparse": false,
        "MetaData": {
          "CreateTime": "2015-01-13T10:29:27-08:00",
          "LastUpdatedTime": "2015-01-13T10:29:27-08:00"
        },
        "AccountType": "Accounts Receivable",
        "CurrentBalance": 0,
        "Active": true,
        "SyncToken": "0",
        "Id": "91",
        "SubAccount": false
      }
    ],
    "maxResults": 3
  },
  "time": "2015-07-13T12:35:57.651-07:00"
}
```

### Read an account

- **Method**: GET
- **URL**: `/v3/company/{realmID}/account/{accountId}`

Retrieves the details of an Account object that has been previously created.

**Response**:
```json
{
  "Account": {
    "FullyQualifiedName": "Accounts Payable (A/P)",
    "domain": "QBO",
    "Name": "Accounts Payable (A/P)",
    "Classification": "Liability",
    "AccountSubType": "AccountsPayable",
    "CurrentBalanceWithSubAccounts": -1091.23,
    "sparse": false,
    "MetaData": {
      "CreateTime": "2014-09-12T10:12:02-07:00",
      "LastUpdatedTime": "2015-06-30T15:09:07-07:00"
    },
    "AccountType": "Accounts Payable",
    "CurrentBalance": -1091.23,
    "Active": true,
    "SyncToken": "0",
    "Id": "33",
    "SubAccount": false
  },
  "time": "2015-07-13T12:50:36.72-07:00"
}
```

### Full update an account

- **Method**: POST
- **URL**: `/v3/company/{realmID}/account`

Use this operation to update any of the writable fields of an existing account object. The request body must include all writable fields of the existing object as returned in a read response. Writable fields omitted from the request body are set to NULL. The ID of the object to update is specified i...

**Request Body**:
```json
{
  "FullyQualifiedName": "Accounts Payable (A/P)",
  "domain": "QBO",
  "SubAccount": false,
  "Description": "Description added during update.",
  "Classification": "Liability",
  "AccountSubType": "AccountsPayable",
  "CurrentBalanceWithSubAccounts": -1091.23,
  "sparse": false,
  "MetaData": {
    "CreateTime": "2014-09-12T10:12:02-07:00",
    "LastUpdatedTime": "2015-06-30T15:09:07-07:00"
  },
  "AccountType": "Accounts Payable",
  "CurrentBalance": -1091.23,
  "Active": true,
  "SyncToken": "0",
  "Id": "33",
  "Name": "Accounts Payable (A/P)"
}
```

**Response**:
```json
{
  "Account": {
    "FullyQualifiedName": "Accounts Payable (A/P)",
    "domain": "QBO",
    "SubAccount": false,
    "Description": "Description added during update.",
    "Classification": "Liability",
    "AccountSubType": "AccountsPayable",
    "CurrentBalanceWithSubAccounts": -1091.23,
    "sparse": false,
    "MetaData": {
      "CreateTime": "2014-09-12T10:12:02-07:00",
      "LastUpdatedTime": "2015-07-13T15:35:13-07:00"
    },
    "AccountType": "Accounts Payable",
    "CurrentBalance": -1091.23,
    "Active": true,
    "SyncToken": "1",
    "Id": "33",
    "Name": "Accounts Payable (A/P)"
  },
  "time": "2015-07-13T15:31:25.618-07:00"
}
```
