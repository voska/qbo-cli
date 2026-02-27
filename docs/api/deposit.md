# Deposit

A deposit object is a transaction that records one or more deposits of the following types: • A customer payment, originally held in the Undeposited Funds account, into the Asset Account specified by the Deposit.DepositToAccountRef attribute. The Deposit.line.LinkedTxn element is used in this case to hold deposit information. • A new, direct deposit specified by Deposit.Line.DepositLineDetail line detail.

## Business Rules

- There must be at least one line item included in a create request.
- Any transaction that funds the Undeposited Funds account can be linked to a Deposit object with a Deposit.Line.LinkedTxn element.

## Attributes

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| Id | String | Required for update | Unique identifier for this object. Sort order is ASC by default. |
| DepositToAccountRef | ReferenceType | Required | Identifies the account to be used for this deposit. Query the Account name list resource to determine the appropriate Account object for this reference, where Account. |
| Line | Line | Required | Individual line items comprising the deposit. Specify a Line.LinkedTxn element along with DepositLine detail type if this line is to record a deposit for an existing transaction. |
| SyncToken | String | Required for update | Version number of the object. It is used to lock an object for use by one app at a time. As soon as an application modifies an object, its SyncToken is incremented. |
| GlobalTaxCalculation | String | Conditionally required | Method in which tax is applied. Allowed values are: TaxExcluded, TaxInclusive, and NotApplicable. |
| CurrencyRef | CurrencyRefType | Conditionally required | Reference to the currency in which all amounts on the associated transaction are expressed. This must be defined if multicurrency is enabled for the company. |
| PrivateNote | String | Optional | User entered, organization-private note about the transaction. This note does not appear on the invoice to the customer. |
| ExchangeRate | Decimal | Optional | The number of home currency units it takes to equal one unit of currency specified by CurrencyRef. Applicable if multicurrency is enabled for the company. |
| DepartmentRef | ReferenceType | Optional | A reference to a Department object specifying the location of the transaction, as defined using location tracking in QuickBooks Online. Available if Preferences. |
| TxnSource | String | Optional | Used internally to specify originating source of a credit card transaction. |
| TxnDate | Date | Optional | The date entered by the user when this transaction occurred. For posting transactions, this is the posting date that affects the financial statements. |
| CashBack |  | Optional | CashBackInfo |
| TransactionLocationType | String | Optional | The account location. Valid values for France: WithinFrance, FranceOverseas, OutsideFranceWithEU, OutsideEU. |
| TxnTaxDetail |  | Optional |  |
| MetaData | ModificationMetaData | Read-only | Descriptive information about the object. The MetaData values are set by Data Services and are  for all applications. |
| RecurDataRef | ReferenceType | Read-only | A reference to the Recurring Transaction. It captures what recurring transaction template the Deposit was created from. |
| TotalAmt | BigDecimal | Read-only | Indicates the total amount of the transaction. This includes the total of all the charges, allowances, and taxes. |
| HomeTotalAmt | Decimal | Read-only | Total amount of the transaction in the home currency. Includes the total of all the charges, allowances and taxes. Calculated by QuickBooks business logic. |

## Sample Object

```json
{
  "Deposit": {
    "SyncToken": "0",
    "domain": "QBO",
    "DepositToAccountRef": {
      "name": "Checking",
      "value": "35"
    },
    "TxnDate": "2014-12-22",
    "TotalAmt": 1675.52,
    "sparse": false,
    "Line": [
      {
        "Amount": 1675,
        "LinkedTxn": [
          {
            "TxnLineId": "0",
            "TxnId": "120",
            "TxnType": "Payment"
          }
        ]
      }
    ],
    "Id": "148",
    "MetaData": {
      "CreateTime": "2014-12-22T12:46:52-08:00",
      "LastUpdatedTime": "2014-12-22T12:46:52-08:00"
    }
  },
  "time": "2014-12-22T13:39:35.449-08:00"
}
```

## Operations

### Create a deposit

- **Method**: POST
- **URL**: `/v3/company/{realmID}/deposit`

A create request includes at least one line representing a deposit--either a direct deposit or linked deposit. More than one deposit can be included in the request; types can be mixed. A direct deposit must have at least: -One line that specifies Deposit.Line.DepositLineDetail.AccountRef. -The Depos...

**Request Body**:
```json
{
  "Line": [
    {
      "DetailType": "DepositLineDetail",
      "Amount": 20.0,
      "ProjectRef": {
        "value": "42991284"
      },
      "DepositLineDetail": {
        "AccountRef": {
          "name": "Unapplied Cash Payment Income",
          "value": "87"
        }
      }
    }
  ],
  "DepositToAccountRef": {
    "name": "Checking",
    "value": "35"
  }
}
```

**Response**:
```json
{
  "Deposit": {
    "SyncToken": "0",
    "domain": "QBO",
    "DepositToAccountRef": {
      "name": "Checking",
      "value": "35"
    },
    "TxnDate": "2014-12-22",
    "TotalAmt": 20.0,
    "sparse": false,
    "Line": [
      {
        "DetailType": "DepositLineDetail",
        "ProjectRef": {
          "value": "42991284"
        },
        "LineNum": 1,
        "Amount": 20.0,
        "Id": "1",
        "DepositLineDetail": {
          "AccountRef": {
            "name": "Unapplied Cash Payment Income",
            "value": "87"
          }
        }
      }
    ],
    "Id": "149",
    "MetaData": {
      "CreateTime": "2014-12-22T14:46:36-08:00",
      "LastUpdatedTime": "2014-12-22T14:46:36-08:00"
    }
  },
  "time": "2014-12-22T14:46:36.084-08:00"
}
```

### Delete a deposit

- **Method**: POST
- **URL**: `/v3/company/{realmID}/deposit?operation=delete`

This operation deletes the Deposit object specified in the request body. Include a minimum of Id and SyncToken in the request body.

**Request Body**:
```json
{
  "SyncToken": "3",
  "Id": "148"
}
```

**Response**:
```json
{
  "Deposit": {
    "status": "Deleted",
    "domain": "QBO",
    "Id": "148"
  },
  "time": "2014-12-22T14:07:19.053-08:00"
}
```

### Query a deposit

- **Method**: GET
- **URL**: `/v3/company/{realmID}/query?query={selectStatement}`

**Sample Query**: `select * from Deposit`

**Response**:
```json
{
  "QueryResponse": {
    "startPosition": 1,
    "totalCount": 8,
    "Deposit": [
      {
        "SyncToken": "0",
        "domain": "QBO",
        "DepositToAccountRef": {
          "name": "Checking",
          "value": "35"
        },
        "TxnDate": "2014-12-22",
        "TotalAmt": 1675.52,
        "sparse": false,
        "Line": [
          {
            "Amount": 1675,
            "ProjectRef": {
              "value": "39298045"
            },
            "LinkedTxn": [
              {
                "TxnLineId": "0",
                "TxnId": "120",
                "TxnType": "Payment"
              }
            ]
          }
        ],
        "Id": "148",
        "MetaData": {
          "CreateTime": "2014-12-22T12:46:52-08:00",
          "LastUpdatedTime": "2014-12-22T12:46:52-08:00"
        }
      },
      {
        "SyncToken": "0",
        "domain": "QBO",
        "DepositToAccountRef": {
          "name": "Checking",
          "value": "35"
        },
        "TxnDate": "2014-12-10",
        "TotalAmt": 20.0,
        "sparse": false,
        "Line": [
          {
            "Amount": 1675,
            "LinkedTxn": [
              {
                "TxnLineId": "0",
                "TxnId": "120",
                "TxnType": "Payment"
              }
            ]
          },
          {
            "LineNum": 1,
            "Amount": 20.0,
            "Id": "1",
            "DepositLineDetail": {
              "AccountRef": {
                "name": "Unapplied Cash Payment Income",
                "value": "87"
              }
            },
            "DetailType": "DepositLineDetail"
          }
        ],
        "Id": "145",
        "MetaData": {
          "CreateTime": "2014-12-10T15:21:44-08:00",
          "LastUpdatedTime": "2014-12-10T15:21:44-08:00"
        }
      },
      {
        "SyncToken": "0",
        "domain": "QBO",
        "DepositToAccountRef": {
          "name": "Savings",
          "value": "36"
        },
        "TxnDate": "2014-11-05",
        "TotalAmt": 600.0,
        "PrivateNote": "Opening Balance",
        "sparse": false,
        "Line": [
          {
            "Amount": 140,
            "LinkedTxn": [
              {
                "TxnLineId": "0",
                "TxnId": "47",
                "TxnType": "SalesReceipt"
              }
            ]
          },
          {
            "Amount": 78,
            "LinkedTxn": [
              {
                "TxnLineId": "0",
                "TxnId": "38",
                "TxnType": "SalesReceipt"
              }
            ]
          },
          {
            "LineNum": 1,
            "Amount": 600.0,
            "Id": "1",
            "DepositLineDetail": {
              "AccountRef": {
                "name": "Opening Balance Equity",
                "value": "34"
              }
            },
            "DetailType": "DepositLineDetail"
          }
        ],
        "Id": "5",
        "MetaData": {
          "CreateTime": "2014-11-05T12:09:00-08:00",
          "LastUpdatedTime": "2014-11-05T12:09:00-08:00"
        }
      }
    ],
    "maxResults": 8
  },
  "time": "2014-12-22T13:11:37.977-08:00"
}
```

### Read a deposit

- **Method**: GET
- **URL**: `/v3/company/{realmID}/deposit/{depositId}`

Retrieves the details of a Deposit object that has been previously created.

**Response**:
```json
{
  "Deposit": {
    "SyncToken": "0",
    "domain": "QBO",
    "DepositToAccountRef": {
      "name": "Checking",
      "value": "35"
    },
    "TxnDate": "2014-12-22",
    "TotalAmt": 1675.52,
    "sparse": false,
    "Line": [
      {
        "Amount": 1675,
        "LinkedTxn": [
          {
            "TxnLineId": "0",
            "TxnId": "120",
            "TxnType": "Payment"
          }
        ]
      }
    ],
    "Id": "148",
    "MetaData": {
      "CreateTime": "2014-12-22T12:46:52-08:00",
      "LastUpdatedTime": "2014-12-22T12:46:52-08:00"
    }
  },
  "time": "2014-12-22T13:39:35.449-08:00"
}
```

### Full update a deposit

- **Method**: POST
- **URL**: `/v3/company/{realmID}/deposit`

Use this operation to update any of the writable fields of an existing deposit object. The request body must include all writable fields of the existing object as returned in a read response. Writable fields omitted from the request body are set to NULL. The ID of the object to update is specified i...

**Request Body**:
```json
{
  "SyncToken": "1",
  "domain": "QBO",
  "DepositToAccountRef": {
    "name": "Checking",
    "value": "35"
  },
  "TxnDate": "2014-12-15",
  "TotalAmt": 1675.52,
  "sparse": false,
  "Line": [
    {
      "Amount": 1675,
      "LinkedTxn": [
        {
          "TxnLineId": "0",
          "TxnId": "120",
          "TxnType": "Payment"
        }
      ]
    }
  ],
  "Id": "148",
  "MetaData": {
    "CreateTime": "2014-12-22T12:46:52-08:00",
    "LastUpdatedTime": "2014-12-22T12:46:52-08:00"
  }
}
```

**Response**:
```json
{
  "Deposit": {
    "SyncToken": "2",
    "domain": "QBO",
    "DepositToAccountRef": {
      "name": "Checking",
      "value": "35"
    },
    "TxnDate": "2014-12-07",
    "TotalAmt": 1675.52,
    "sparse": false,
    "Line": [
      {
        "Amount": 1675,
        "LinkedTxn": [
          {
            "TxnLineId": "0",
            "TxnId": "120",
            "TxnType": "Payment"
          }
        ]
      }
    ],
    "Id": "148",
    "MetaData": {
      "CreateTime": "2014-12-22T12:46:52-08:00",
      "LastUpdatedTime": "2014-12-22T14:04:10-08:00"
    }
  },
  "time": "2014-12-22T14:04:10.815-08:00"
}
```

### Sparse update a deposit

- **Method**: POST
- **URL**: `/v3/company/{realmID}/deposit`

Sparse updating provides the ability to update a subset of properties for a given object; only elements specified in the request are updated. Missing elements are left untouched. The ID of the object to update is specified in the request body.

**Request Body**:
```json
{
  "SyncToken": "0",
  "domain": "QBO",
  "DepositToAccountRef": {
    "name": "Checking",
    "value": "35"
  },
  "TxnDate": "2014-12-02",
  "sparse": true,
  "Id": "146",
  "MetaData": {
    "CreateTime": "2014-12-22T12:46:52-08:00",
    "LastUpdatedTime": "2014-12-22T12:46:52-08:00"
  }
}
```

**Response**:
```json
{
  "Deposit": {
    "SyncToken": "1",
    "domain": "QBO",
    "DepositToAccountRef": {
      "name": "Checking",
      "value": "35"
    },
    "TxnDate": "2014-12-02",
    "TotalAmt": 20.0,
    "sparse": false,
    "Line": [
      {
        "LineNum": 1,
        "Amount": 20.0,
        "Id": "1",
        "DepositLineDetail": {
          "AccountRef": {
            "name": "Unapplied Cash Payment Income",
            "value": "87"
          }
        },
        "DetailType": "DepositLineDetail"
      }
    ],
    "Id": "146",
    "MetaData": {
      "CreateTime": "2014-12-10T15:24:40-08:00",
      "LastUpdatedTime": "2014-12-22T15:13:18-08:00"
    }
  },
  "time": "2014-12-22T15:13:17.913-08:00"
}
```
