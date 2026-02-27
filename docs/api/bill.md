# Bill

A Bill object is an AP transaction representing a request-for-payment from a third party for goods/services rendered, received, or both.

## Attributes

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| Id | String | Required for update | Unique identifier for this object. Sort order is ASC by default. |
| VendorRef | ReferenceType | Required | Reference to the vendor for this transaction. Query the Vendor name list resource to determine the appropriate Vendor object for this reference. Use Vendor.Id and Vendor. |
| Line | Line | Required | Individual line items of a transaction. Valid Line types include: ItemBasedExpenseLine and AccountBasedExpenseLine ItemBasedExpenseLine AccountBasedExpenseLine |
| SyncToken | String | Required for update | Version number of the object. It is used to lock an object for use by one app at a time. As soon as an application modifies an object, its SyncToken is incremented. |
| CurrencyRef | CurrencyRefType | Conditionally required | Reference to the currency in which all amounts on the associated transaction are expressed. This must be defined if multicurrency is enabled for the company. |
| GlobalTaxCalculation | String | Conditionally required | Method in which tax is applied. Allowed values are: TaxExcluded, TaxInclusive, and NotApplicable. |
| TxnDate | Date | Optional | The date entered by the user when this transaction occurred. For posting transactions, this is the posting date that affects the financial statements. |
| APAccountRef | ReferenceType | Optional | Specifies to which AP account the bill is credited. Query the Account name list resource to determine the appropriate Account object for this reference. Use Account. |
| SalesTermRef | ReferenceType | Optional | Reference to the Term associated with the transaction. Query the Term name list resource to determine the appropriate Term object for this reference. Use Term.Id and Term. |
| LinkedTxn | LinkedTxn | Optional | Zero or more transactions linked to this Bill object. The LinkedTxn.TxnType can be set to PurchaseOrder, BillPaymentCheck or if using Minor Version 55 and above ReimburseCharge. |
| TotalAmt | BigDecimal | Read-only | Indicates the total amount of the transaction. This includes the total of all the charges, allowances, and taxes. |
| TransactionLocationType | String | Optional | The account location. Valid values for France: WithinFrance, FranceOverseas, OutsideFranceWithEU, OutsideEU. |
| DueDate | Date | Optional | Date when the payment of the transaction is due. If date is not provided, the number of days specified in SalesTermRef added the transaction date will be used. |
| MetaData | ModificationMetaData | Read-only | Descriptive information about the object. The MetaData values are set by Data Services and are  for all applications. |
| DocNumber | String | Optional | Reference number for the transaction. If not explicitly provided at create time, a custom value can be provided. |
| PrivateNote | String | Optional | User entered, organization-private note about the transaction. This note does not appear on the invoice to the customer. |
| TxnTaxDetail |  | Optional |  |
| ExchangeRate | Decimal | Optional | The number of home currency units it takes to equal one unit of currency specified by CurrencyRef. Applicable if multicurrency is enabled for the company. |
| DepartmentRef | ReferenceType | Optional | A reference to a Department object specifying the location of the transaction, as defined using location tracking in QuickBooks Online. |
| IncludeInAnnualTPAR |  | Optional |  |
| HomeBalance | Decimal | Read-only | Convenience field containing the amount in Balance expressed in terms of the home currency. Calculated by QuickBooks business logic. |
| RecurDataRef | ReferenceType | Read-only | A reference to the Recurring Transaction. It captures what recurring transaction template the Bill was created from. |
| Balance | Decimal | Read-only | The balance reflecting any payments made against the transaction. Initially set to the value of TotalAmt. A Balance of 0 indicates the bill is fully paid. |

## Sample Object

```json
{
  "Bill": {
    "SyncToken": "2",
    "domain": "QBO",
    "APAccountRef": {
      "name": "Accounts Payable (A/P)",
      "value": "33"
    },
    "VendorRef": {
      "name": "Norton Lumber and Building Materials",
      "value": "46"
    },
    "TxnDate": "2014-11-06",
    "TotalAmt": 103.55,
    "CurrencyRef": {
      "name": "United States Dollar",
      "value": "USD"
    },
    "LinkedTxn": [
      {
        "TxnId": "118",
        "TxnType": "BillPaymentCheck"
      }
    ],
    "SalesTermRef": {
      "value": "3"
    },
    "DueDate": "2014-12-06",
    "sparse": false,
    "Line": [
      {
        "Description": "Lumber",
        "DetailType": "AccountBasedExpenseLineDetai l",
        "ProjectRef": {
          "value": "39298034"
        },
        "Amount": 103.55,
        "Id": "1",
        "AccountBasedExpenseLineDetail": {
          "TaxCodeRef": {
            "value": "TAX"
          },
          "AccountRef": {
            "name": "Job Expenses:Job Materials:Decks and Patios",
            "value": "64"
          },
          "BillableStatus": "Billable",
          "CustomerRef": {
            "name": "Travis Waldron",
            "value": "26"
          }
        }
      }
    ],
    "Balance": 0,
    "Id": "25",
    "MetaData": {
      "CreateTime": "2014-11-06T15:37:25-08:00",
      "LastUpdatedTime": "2015-02-09T10:11:11-08:00"
    }
  },
  "time": "2015-02-09T10:17:20.251-08:00"
}
```

## Operations

### Create a bill

- **Method**: POST
- **URL**: `/v3/company/{realmID}/bill`

**Request Body**:
```json
{
  "Line": [
    {
      "DetailType": "AccountBasedExpenseLineDetail",
      "Amount": 200.0,
      "Id": "1",
      "AccountBasedExpenseLineDetail": {
        "AccountRef": {
          "value": "7"
        }
      }
    }
  ],
  "VendorRef": {
    "value": "56"
  }
}
```

**Response**:
```json
{
  "Bill": {
    "SyncToken": "0",
    "domain": "QBO",
    "VendorRef": {
      "name": "Bob's Burger Joint",
      "value": "56"
    },
    "TxnDate": "2014-12-31",
    "TotalAmt": 200.0,
    "APAccountRef": {
      "name": "Accounts Payable (A/P)",
      "value": "33"
    },
    "Id": "151",
    "sparse": false,
    "Line": [
      {
        "DetailType": "AccountBasedExpenseLineDetai l",
        "Amount": 200.0,
        "Id": "1",
        "AccountBasedExpenseLineDetail": {
          "TaxCodeRef": {
            "value": "NON"
          },
          "AccountRef": {
            "name": "Advertising",
            "value": "7"
          },
          "BillableStatus": "NotBillable"
        }
      }
    ],
    "Balance": 200.0,
    "DueDate": "2014-12-31",
    "MetaData": {
      "CreateTime": "2014-12-31T09:59:18-08:00",
      "LastUpdatedTime": "2014-12-31T09:59:18-08:00"
    }
  },
  "time": "2014-12-31T09:59:17.449-08:00"
}
```

### Delete a bill

- **Method**: POST
- **URL**: `/v3/company/{realmID}/bill?operation=delete`

This operation deletes the bill object specified in the request body. Include a minimum of Bill.Id and Bill.SyncToken in the request body. You must unlink any linked transactions associated with the bill object before deleting it.

**Request Body**:
```json
{
  "SyncToken": "0",
  "Id": "108"
}
```

**Response**:
```json
{
  "Bill": {
    "status": "Deleted",
    "domain": "QBO",
    "Id": "108"
  },
  "time": "2015-05-26T13:14:34.775-07:00"
}
```

### Query a bill

- **Method**: GET
- **URL**: `/v3/company/{realmID}/query?query={selectStatement}`

**Sample Query**: `select * from bill maxresults 2`

**Response**:
```json
{
  "QueryResponse": {
    "startPosition": 1,
    "totalCount": 2,
    "Bill": [
      {
        "SyncToken": "2",
        "domain": "QBO",
        "VendorRef": {
          "name": "Norton Lumber and Building Materials",
          "value": "46"
        },
        "TxnDate": "2014-10-07",
        "TotalAmt": 225.0,
        "APAccountRef": {
          "name": "Accounts Payable (A /P)",
          "value": "33"
        },
        "Id": "150",
        "sparse": false,
        "Line": [
          {
            "DetailType": "ItemBasedExpenseLineDet ail",
            "Amount": 100.0,
            "Id": "1",
            "ItemBasedExpenseLineDet ail": {
              "TaxCodeRef": {
                "value": "NON"
              },
              "Qty": 8,
              "BillableStatus": "NotBillable",
              "UnitPrice": 10,
              "ItemRef": {
                "name": "Pump",
                "value": "11"
              }
            },
            "Description": "Fountain Pump"
          },
          {
            "DetailType": "ItemBasedExpenseLineDet ail",
            "Amount": 125.0,
            "Id": "2",
            "ItemBasedExpenseLineDet ail": {
              "TaxCodeRef": {
                "value": "NON"
              },
              "Qty": 1,
              "BillableStatus": "NotBillable",
              "UnitPrice": 125,
              "ItemRef": {
                "name": "Rock Fountain",
                "value": "5"
              }
            },
            "Description": "Rock Fountain"
          }
        ],
        "Balance": 225.0,
        "DueDate": "2014-10-07",
        "MetaData": {
          "CreateTime": "2014-10-15T13:55:31-07:00",
          "LastUpdatedTime": "2014-10-15T14:24:54-07:00"
        }
      },
      {
        "SyncToken": "0",
        "domain": "QBO",
        "VendorRef": {
          "name": "Bob's Burger Joint",
          "value": "56"
        },
        "TxnDate": "2014-10-15",
        "TotalAmt": 200.0,
        "APAccountRef": {
          "name": "Accounts Payable (A /P)",
          "value": "33"
        },
        "Id": "149",
        "sparse": false,
        "Line": [
          {
            "DetailType": "AccountBasedExpenseLine Detail",
            "Amount": 200.0,
            "Id": "1",
            "AccountBasedExpenseLine Detail": {
              "TaxCodeRef": {
                "value": "NON"
              },
              "AccountRef": {
                "name": "Advertising",
                "value": "7"
              },
              "BillableStatus": "NotBillable"
            }
          }
        ],
        "Balance": 200.0,
        "DueDate": "2014-10-15",
        "MetaData": {
          "CreateTime": "2014-10-15T13:48:00-07:00",
          "LastUpdatedTime": "2014-10-15T13:48:00-07:00"
        }
      }
    ],
    "maxResults": 2
  },
  "time": "2014-10-15T14:41:39.98-07:00"
}
```

### Read a bill

- **Method**: GET
- **URL**: `/v3/company/{realmID}/bill/{billId}`

Retrieves the details of a bill that has been previously created.

**Response**:
```json
{
  "Bill": {
    "SyncToken": "2",
    "domain": "QBO",
    "APAccountRef": {
      "name": "Accounts Payable (A/P)",
      "value": "33"
    },
    "VendorRef": {
      "name": "Norton Lumber and Building Materials",
      "value": "46"
    },
    "TxnDate": "2014-11-06",
    "TotalAmt": 103.55,
    "CurrencyRef": {
      "name": "United States Dollar",
      "value": "USD"
    },
    "LinkedTxn": [
      {
        "TxnId": "118",
        "TxnType": "BillPaymentCheck"
      }
    ],
    "SalesTermRef": {
      "value": "3"
    },
    "DueDate": "2014-12-06",
    "sparse": false,
    "Line": [
      {
        "Description": "Lumber",
        "DetailType": "AccountBasedExpenseLineDetai l",
        "ProjectRef": {
          "value": "39298034"
        },
        "Amount": 103.55,
        "Id": "1",
        "AccountBasedExpenseLineDetail": {
          "TaxCodeRef": {
            "value": "TAX"
          },
          "AccountRef": {
            "name": "Job Expenses:Job Materials:Decks and Patios",
            "value": "64"
          },
          "BillableStatus": "Billable",
          "CustomerRef": {
            "name": "Travis Waldron",
            "value": "26"
          }
        }
      }
    ],
    "Balance": 0,
    "Id": "25",
    "MetaData": {
      "CreateTime": "2014-11-06T15:37:25-08:00",
      "LastUpdatedTime": "2015-02-09T10:11:11-08:00"
    }
  },
  "time": "2015-02-09T10:17:20.251-08:00"
}
```

### Full update a bill

- **Method**: POST
- **URL**: `/v3/company/{realmID}/bill`

Use this operation to update any of the writable fields of an existing bill object. The request body must include all writable fields of the existing object as returned in a read response. Writable fields omitted from the request body are set to NULL. The ID of the object to update is specified in t...

**Request Body**:
```json
{
  "DocNumber": "56789",
  "SyncToken": "1",
  "domain": "QBO",
  "APAccountRef": {
    "name": "Accounts Payable",
    "value": "49"
  },
  "VendorRef": {
    "name": "Bayshore CalOil Service",
    "value": "81"
  },
  "TxnDate": "2014-04-04",
  "TotalAmt": 200.0,
  "CurrencyRef": {
    "name": "United States Dollar",
    "value": "USD"
  },
  "PrivateNote": "This is a updated memo.",
  "SalesTermRef": {
    "value": "12"
  },
  "DepartmentRef": {
    "name": "Garden Services",
    "value": "1"
  },
  "DueDate": "2013-06-09",
  "sparse": false,
  "Line": [
    {
      "Description": "Gasoline",
      "DetailType": "AccountBasedExpenseLineDetail",
      "ProjectRef": {
        "value": "39298034"
      },
      "Amount": 200.0,
      "Id": "1",
      "AccountBasedExpenseLineDetail": {
        "TaxCodeRef": {
          "value": "TAX"
        },
        "AccountRef": {
          "name": "Automobile",
          "value": "75"
        },
        "BillableStatus": "Billable",
        "CustomerRef": {
          "name": "Blackwell, Edward",
          "value": "20"
        },
        "MarkupInfo": {
          "Percent": 10
        }
      }
    }
  ],
  "Balance": 200.0,
  "Id": "890",
  "MetaData": {
    "CreateTime": "2014-04-04T12:38:01-07:00",
    "LastUpdatedTime": "2014-04-04T12:48:56-07:00"
  }
}
```

**Response**:
```json
{
  "Bill": {
    "DocNumber": "56789",
    "SyncToken": "2",
    "domain": "QBO",
    "APAccountRef": {
      "name": "Accounts Payable",
      "value": "49"
    },
    "VendorRef": {
      "name": "Bayshore CalOil Service",
      "value": "81"
    },
    "TxnDate": "2014-04-04",
    "TotalAmt": 200.0,
    "CurrencyRef": {
      "name": "United States Dollar",
      "value": "USD"
    },
    "PrivateNote": "This is a updated memo.",
    "SalesTermRef": {
      "value": "12"
    },
    "DepartmentRef": {
      "name": "Garden Services",
      "value": "1"
    },
    "DueDate": "2013-06-09",
    "sparse": false,
    "Line": [
      {
        "Description": "Gasoline",
        "DetailType": "AccountBasedExpenseLineDetai l",
        "ProjectRef": {
          "value": "39298034"
        },
        "Amount": 200.0,
        "Id": "1",
        "AccountBasedExpenseLineDetail": {
          "TaxCodeRef": {
            "value": "TAX"
          },
          "AccountRef": {
            "name": "Automobile",
            "value": "75"
          },
          "BillableStatus": "Billable",
          "CustomerRef": {
            "name": "Blackwell, Edward",
            "value": "20"
          },
          "MarkupInfo": {
            "Percent": 10
          }
        }
      }
    ],
    "Balance": 200.0,
    "Id": "890",
    "MetaData": {
      "CreateTime": "2014-04-04T12:38:01-07:00",
      "LastUpdatedTime": "2014-04-04T12:58:16-07:00"
    }
  },
  "time": "2014-04-04T12:58:16.491-07:00"
}
```
