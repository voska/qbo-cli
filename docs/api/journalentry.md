# JournalEntry

JournalEntry is a transaction in which: • There are at least one pair of lines, a debit and a credit, called distribution lines. • Each distribution line has an account from the Chart of Accounts. Query the Account resource for a listing of the Chart of Accounts. • The total of the debit column equals the total of the credit column. When you record a transaction with a JournalEntry object, the QuickBooks Online UI labels the transaction as JRNL in the register and as General Journal on reports that list transactions.

## Business Rules

- "• Accounts Receivable (A/R) account: needs to have a Customer in the Name Field. The A/R account is visible only after there are A/R transactions such as receive payments from invoices."
- "• Accounts Payable (A/P) account: needs to have a Vendor in the Name Field. The A/P account is visible only after there are A/P transactions such Bill objects."
- There are both Sales Tax and Purchase Tax.
- "• On the transaction line , if TaxCodeRef is specified, TaxApplicableOn and TaxAmount are required. Each TaxCodeRef can result in one or more tax lines. For AU locale : On the transaction line, if GlobalTaxCalculation is TaxInclusive andTaxCodeRef is specified, TaxInclusiveAmt is required."
- Any TxnTaxDetail lines specified are not overridden. That is, if a user provides incorrect values such that the total amount on debit is not equal to total amount on credit, an error is returned.
- Not SKU specific.

## Attributes

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| Id | String | Required for update | Unique identifier for this object. Sort order is ASC by default. |
| Line | Line | Required | Individual line items of a transaction. There must be at least one pair of Journal Entry Line elements, representing a debit and a credit, called distribution lines. |
| SyncToken | String | Required for update | Version number of the object. It is used to lock an object for use by one app at a time. As soon as an application modifies an object, its SyncToken is incremented. |
| CurrencyRef | CurrencyRefType | Conditionally required | Reference to the currency in which all amounts on the associated transaction are expressed. This must be defined if multicurrency is enabled for the company. |
| GlobalTaxCalculation | String | Conditionally required | Method in which tax is applied. Allowed values are: TaxExcluded, TaxInclusive, and NotApplicable. |
| DocNumber | String | Optional | Reference number for the transaction. Throws an error when duplicate DocNumber is sent in the request and if Preferences:OtherPrefs:NameValue. |
| PrivateNote | String | Optional | User entered, organization-private note about the transaction. |
| TxnDate | Date | Optional | The date entered by the user when this transaction occurred. For posting transactions, this is the posting date that affects the financial statements. |
| ExchangeRate | Decimal | Optional | The number of home currency units it takes to equal one unit of currency specified by CurrencyRef. Applicable if multicurrency is enabled for the company. |
| TaxRateRef | ReferenceType | Optional | Reference to the Tax Adjustment Rate Ids for this item. Query the TaxRate list resource to determine the appropriate TaxRate object for this reference. Use TaxRate. |
| TransactionLocationType | String | Optional | The account location. Valid values for France: WithinFrance, FranceOverseas, OutsideFranceWithEU, OutsideEU. |
| TxnTaxDetail | TxnTaxDetail | Optional | This data type provides information for taxes charged on the transaction as a whole. |
| Adjustment |  | Optional |  |
| MetaData | ModificationMetaData | Read-only | Descriptive information about the object. The MetaData values are set by Data Services and are  for all applications. |
| RecurDataRef | ReferenceType | Read-only | A reference to the Recurring Transaction. It captures what recurring transaction template the JournalEntry was created from. |
| TotalAmt | BigDecimal | Read-only | The value of this field will always be set to zero. Calculated by QuickBooks business logic; any value you supply is over-written by QuickBooks. |
| HomeTotalAmt | Decimal | Read-only | The value of this field will always be set to zero. Applicable if multicurrency is enabled for the company. |

## Sample Object

```json
{
  "time": "2015-06-29T12:43:42.132-07:00",
  "JournalEntry": {
    "SyncToken": "0",
    "domain": "QBO",
    "TxnDate": "2015-06-29",
    "sparse": false,
    "Line": [
      {
        "Description": "Four sprinkler heads damaged",
        "JournalEntryLineDetail": {
          "PostingType": "Debit",
          "AccountRef": {
            "name": "Job Expenses:Job Materials:Fountain and Garden Lighting",
            "value": "65"
          },
          "Entity": {
            "Type": "Customer",
            "EntityRef": {
              "name": "Amy's Bird Sanctuary",
              "value": "1"
            }
          }
        },
        "DetailType": "JournalEntryLineDetail",
        "ProjectRef": {
          "value": "39298034"
        },
        "Amount": 25.54,
        "Id": "0"
      },
      {
        "JournalEntryLineDetail": {
          "PostingType": "Credit",
          "AccountRef": {
            "name": "Notes Payable",
            "value": "44"
          },
          "Entity": {
            "Type": "Vendor",
            "EntityRef": {
              "name": "IDX Vendor",
              "value": "2"
            }
          }
        },
        "DetailType": "JournalEntryLineDetail",
        "Amount": 25.54,
        "Id": "1",
        "Description": "Sprinkler Hds- Sprinkler Hds Inventory Adjustment"
      }
    ],
    "Adjustment": false,
    "Id": "227",
    "TxnTaxDetail": {},
    "MetaData": {
      "CreateTime": "2015-06-29T12:33:57-07:00",
      "LastUpdatedTime": "2015-06-29T12:33:57-07:00"
    }
  }
}
```

## Operations

### Create a journalentry

- **Method**: POST
- **URL**: `/v3/company/{realmID}/journalentry`

**Request Body**:
```json
{
  "Line": [
    {
      "JournalEntryLineDetail": {
        "PostingType": "Debit",
        "AccountRef": {
          "name": "Opening Bal Equity",
          "value": "39"
        }
      },
      "DetailType": "JournalEntryLineDetail",
      "Amount": 100.0,
      "Id": "0",
      "Description": "nov portion of rider insurance"
    },
    {
      "JournalEntryLineDetail": {
        "PostingType": "Credit",
        "AccountRef": {
          "name": "Notes Payable",
          "value": "44"
        }
      },
      "DetailType": "JournalEntryLineDetail",
      "Amount": 100.0,
      "Description": "nov portion of rider insurance"
    }
  ]
}
```

**Response**:
```json
{
  "time": "2015-06-29T12:45:32.183-07:00",
  "JournalEntry": {
    "SyncToken": "0",
    "domain": "QBO",
    "TxnDate": "2015-06-29",
    "sparse": false,
    "Line": [
      {
        "JournalEntryLineDetail": {
          "PostingType": "Debit",
          "AccountRef": {
            "name": "Truck :Depreciation",
            "value": "39"
          }
        },
        "DetailType": "JournalEntryLineDetail",
        "Amount": 100.0,
        "Id": "0",
        "Description": "nov portion of rider insurance"
      },
      {
        "JournalEntryLineDetail": {
          "PostingType": "Credit",
          "AccountRef": {
            "name": "Notes Payable",
            "value": "44"
          }
        },
        "DetailType": "JournalEntryLineDetail",
        "Amount": 100.0,
        "Id": "1",
        "Description": "nov portion of rider insurance"
      }
    ],
    "Adjustment": false,
    "Id": "228",
    "TxnTaxDetail": {},
    "MetaData": {
      "CreateTime": "2015-06-29T12:45:32-07:00",
      "LastUpdatedTime": "2015-06-29T12:45:32-07:00"
    }
  }
}
```

### Delete a journalentry

- **Method**: POST
- **URL**: `/v3/company/{realmID}/journalentry?operation=delete`

This operation deletes the JournalEntry object specified in the request body. Include a minimum of JournalEntry.Id and JournalEntry.SyncToken in the request body.

**Request Body**:
```json
{
  "SyncToken": "0",
  "Id": "228"
}
```

**Response**:
```json
{
  "time": "2015-05-26T14:03:31.321-07:00",
  "JournalEntry": {
    "status": "Deleted",
    "domain": "QBO",
    "Id": "228"
  }
}
```

### Query a journalentry

- **Method**: GET
- **URL**: `/v3/company/{realmID}/query?query={selectStatement}`

**Sample Query**: `select * from JournalEntry Where Metadata.LastUpdatedTime>'2014-09-15T00: 00:00-07:00' Order By Metadata.LastUpdatedTime`

**Response**:
```json
{
  "QueryResponse": {
    "startPosition": 1,
    "totalCount": 3,
    "maxResults": 3,
    "JournalEntry": [
      {
        "SyncToken": "0",
        "domain": "QBO",
        "TxnDate": "2014-09-16",
        "PrivateNote": "Opening Balance",
        "sparse": false,
        "Line": [
          {
            "JournalEntryLineDetail": {
              "PostingType": "Credit",
              "AccountRef": {
                "name": "Notes Payable",
                "value": "44"
              }
            },
            "DetailType": "JournalEntryLineDetail",
            "Amount": 25000.0,
            "Id": "0",
            "Description": "Opening Balance"
          },
          {
            "JournalEntryLineDetail": {
              "PostingType": "Debit",
              "AccountRef": {
                "name": "Opening Balance Equity",
                "value": "34"
              }
            },
            "DetailType": "JournalEntryLineDetail",
            "Amount": 25000.0,
            "Id": "1",
            "Description": "Opening Balance"
          }
        ],
        "Adjustment": false,
        "Id": "8",
        "TxnTaxDetail": {},
        "MetaData": {
          "CreateTime": "2014-09-16T10:04:24-07:00",
          "LastUpdatedTime": "2014-09-16T10:04:24-07:00"
        }
      },
      {
        "SyncToken": "0",
        "domain": "QBO",
        "TxnDate": "2014-09-16",
        "PrivateNote": "Opening Balance",
        "sparse": false,
        "Line": [
          {
            "JournalEntryLineDetail": {
              "PostingType": "Credit",
              "AccountRef": {
                "name": "Loan Payable",
                "value": "43"
              }
            },
            "DetailType": "JournalEntryLineDetail",
            "Amount": 4000.0,
            "Id": "0",
            "Description": "Opening Balance"
          },
          {
            "JournalEntryLineDetail": {
              "PostingType": "Debit",
              "AccountRef": {
                "name": "Opening Balance Equity",
                "value": "34"
              }
            },
            "DetailType": "JournalEntryLineDetail",
            "Amount": 4000.0,
            "Id": "1",
            "Description": "Opening Balance"
          }
        ],
        "Adjustment": false,
        "Id": "7",
        "TxnTaxDetail": {},
        "MetaData": {
          "CreateTime": "2014-09-16T10:03:25-07:00",
          "LastUpdatedTime": "2014-09-16T10:03:25-07:00"
        }
      },
      {
        "SyncToken": "0",
        "domain": "QBO",
        "TxnDate": "2014-09-03",
        "PrivateNote": "Opening Balance",
        "sparse": false,
        "Line": [
          {
            "JournalEntryLineDetail": {
              "PostingType": "Debit",
              "AccountRef": {
                "name": "Truck :Original Cost",
                "value": "38"
              }
            },
            "DetailType": "JournalEntryLineDetail",
            "Amount": 13495.0,
            "Id": "0",
            "Description": "Opening Balance"
          },
          {
            "JournalEntryLineDetail": {
              "PostingType": "Credit",
              "AccountRef": {
                "name": "Opening Balance Equity",
                "value": "34"
              }
            },
            "DetailType": "JournalEntryLineDetail",
            "Amount": 13495.0,
            "Id": "1",
            "Description": "Opening Balance"
          }
        ],
        "Adjustment": false,
        "Id": "6",
        "TxnTaxDetail": {},
        "MetaData": {
          "CreateTime": "2014-09-15T12:11:06-07:00",
          "LastUpdatedTime": "2014-09-15T12:11:06-07:00"
        }
      }
    ]
  },
  "time": "2015-01-16T09:05:53.455-08:00"
}
```

### Read a journalentry

- **Method**: GET
- **URL**: `/v3/company/{realmID}/journalentry/{journalentryId}`

Retrieves the details of an JournalEntry that has been previously created.

**Response**:
```json
{
  "time": "2015-06-29T12:43:42.132-07:00",
  "JournalEntry": {
    "SyncToken": "0",
    "domain": "QBO",
    "TxnDate": "2015-06-29",
    "sparse": false,
    "Line": [
      {
        "Description": "Four sprinkler heads damaged",
        "JournalEntryLineDetail": {
          "PostingType": "Debit",
          "AccountRef": {
            "name": "Job Expenses:Job Materials:Fountain and Garden Lighting",
            "value": "65"
          },
          "Entity": {
            "Type": "Customer",
            "EntityRef": {
              "name": "Amy's Bird Sanctuary",
              "value": "1"
            }
          }
        },
        "DetailType": "JournalEntryLineDetail",
        "ProjectRef": {
          "value": "39298034"
        },
        "Amount": 25.54,
        "Id": "0"
      },
      {
        "JournalEntryLineDetail": {
          "PostingType": "Credit",
          "AccountRef": {
            "name": "Notes Payable",
            "value": "44"
          },
          "Entity": {
            "Type": "Vendor",
            "EntityRef": {
              "name": "IDX Vendor",
              "value": "2"
            }
          }
        },
        "DetailType": "JournalEntryLineDetail",
        "Amount": 25.54,
        "Id": "1",
        "Description": "Sprinkler Hds- Sprinkler Hds Inventory Adjustment"
      }
    ],
    "Adjustment": false,
    "Id": "227",
    "TxnTaxDetail": {},
    "MetaData": {
      "CreateTime": "2015-06-29T12:33:57-07:00",
      "LastUpdatedTime": "2015-06-29T12:33:57-07:00"
    }
  }
}
```

### Full update a journalentry

- **Method**: POST
- **URL**: `/v3/company/{realmID}/journalentry`

Use this operation to update any of the writable fields of an existing JournalEntry object. The request body must include all writable fields of the existing object, including all lines, as returned in a read response. Writable fields omitted from the request body are set to NULL. The ID of the obje...

**Request Body**:
```json
{
  "SyncToken": "1",
  "domain": "QBO",
  "TxnDate": "2015-06-29",
  "sparse": false,
  "Line": [
    {
      "JournalEntryLineDetail": {
        "PostingType": "Debit",
        "AccountRef": {
          "name": "Job Expenses:Job Materials:Fountain and Garden Lighting",
          "value": "65"
        }
      },
      "DetailType": "JournalEntryLineDetail",
      "Amount": 25.54,
      "Id": "0",
      "Description": "Updated description"
    },
    {
      "JournalEntryLineDetail": {
        "PostingType": "Credit",
        "AccountRef": {
          "name": "Notes Payable",
          "value": "44"
        }
      },
      "DetailType": "JournalEntryLineDetail",
      "Amount": 25.54,
      "Id": "1",
      "Description": "Sprinkler Hds- Sprinkler Hds Inventory Adjustment"
    }
  ],
  "Adjustment": false,
  "Id": "227",
  "TxnTaxDetail": {},
  "MetaData": {
    "CreateTime": "2015-06-29T12:33:57-07:00",
    "LastUpdatedTime": "2015-06-29T12:33:57-07:00"
  }
}
```

**Response**:
```json
{
  "time": "2015-06-29T12:57:14.02-07:00",
  "JournalEntry": {
    "SyncToken": "2",
    "domain": "QBO",
    "TxnDate": "2015-06-29",
    "sparse": false,
    "Line": [
      {
        "JournalEntryLineDetail": {
          "PostingType": "Debit",
          "AccountRef": {
            "name": "Job Expenses:Job Materials:Fountain and Garden Lighting",
            "value": "65"
          }
        },
        "DetailType": "JournalEntryLineDetail",
        "Amount": 25.54,
        "Id": "0",
        "Description": "Updated description"
      },
      {
        "JournalEntryLineDetail": {
          "PostingType": "Credit",
          "AccountRef": {
            "name": "Notes Payable",
            "value": "44"
          }
        },
        "DetailType": "JournalEntryLineDetail",
        "Amount": 25.54,
        "Id": "1",
        "Description": "Sprinkler Hds- Sprinkler Hds Inventory Adjustment"
      }
    ],
    "Adjustment": false,
    "Id": "227",
    "TxnTaxDetail": {},
    "MetaData": {
      "CreateTime": "2015-06-29T12:33:57-07:00",
      "LastUpdatedTime": "2015-06-29T12:57:15-07:00"
    }
  }
}
```

### Sparse update a journalentry

- **Method**: POST
- **URL**: `/v3/company/{realmID}/journalentry`

Sparse updating provides the ability to update a subset of properties for a given object; only elements specified in the request are updated. Missing elements are left untouched. The ID of the object to update is specified in the request body.

**Request Body**:
```json
{
  "SyncToken": "1",
  "domain": "QBO",
  "TxnDate": "2015-11-30",
  "PrivateNote": "Revised private note via sparse update",
  "sparse": true,
  "Adjustment": false,
  "Id": "227"
}
```

**Response**:
```json
{
  "time": "2015-06-29T12:54:38.135-07:00",
  "JournalEntry": {
    "DocNumber": "1112",
    "SyncToken": "1",
    "domain": "QBO",
    "TxnDate": "2015-11-30",
    "PrivateNote": "Revised private note via sparse update",
    "sparse": false,
    "Line": [
      {
        "JournalEntryLineDetail": {
          "PostingType": "Debit",
          "AccountRef": {
            "name": "Job Expenses:Job Materials:Fountain and Garden Lighting",
            "value": "65"
          }
        },
        "DetailType": "JournalEntryLineDetail",
        "Amount": 25.54,
        "Id": "0",
        "Description": "Four sprinkler heads damaged"
      },
      {
        "JournalEntryLineDetail": {
          "PostingType": "Credit",
          "AccountRef": {
            "name": "Notes Payable",
            "value": "44"
          }
        },
        "DetailType": "JournalEntryLineDetail",
        "Amount": 25.54,
        "Id": "1",
        "Description": "Sprinkler Hds- Sprinkler Hds Inventory Adjustment"
      }
    ],
    "Adjustment": false,
    "Id": "227",
    "TxnTaxDetail": {},
    "MetaData": {
      "CreateTime": "2015-06-29T12:33:57-07:00",
      "LastUpdatedTime": "2015-06-29T12:54:38-07:00"
    }
  }
}
```
