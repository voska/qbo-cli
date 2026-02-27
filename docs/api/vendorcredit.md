# VendorCredit

The VendorCredit object is an accounts payable transaction that represents a refund or credit of payment for goods or services. It is a credit that a vendor owes you for various reasons such as overpaid bill, returned merchandise, or other reasons.

## Attributes

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| Id | String | Required for update | Unique identifier for this object. Sort order is ASC by default. |
| VendorRef | ReferenceType | Required | Reference to the vendor for this transaction. Query the Vendor name list resource to determine the appropriate Vendor object for this reference. Use Vendor.Id and Vendor. |
| Line | Line | Required | Individual line items of a transaction. Valid Line types include: ItemBasedExpenseLine and AccountBasedExpenseLine ItemBasedExpenseLine AccountBasedExpenseLine |
| SyncToken | String | Required for update | Version number of the object. It is used to lock an object for use by one app at a time. As soon as an application modifies an object, its SyncToken is incremented. |
| GlobalTaxCalculation | String | Conditionally required | Method in which tax is applied. Allowed values are: TaxExcluded, TaxInclusive, and NotApplicable. |
| CurrencyRef | CurrencyRefType | Conditionally required | Reference to the currency in which all amounts on the associated transaction are expressed. This must be defined if multicurrency is enabled for the company. |
| DocNumber | String | Optional | Reference number for the transaction. |
| Throws |  | Optional | an error when duplicate DocNumber is sent in the request. Recommended best practice: check the setting of Preferences:OtherPrefs:NameValue.Name = VendorAndPurchasesPrefs. |
| PrivateNote | String | Optional | User entered, organization-private note about the transaction. This note does not appear on the transaction to the vendor. |
| LinkedTxn | LinkedTxn | Optional | Zero or more transactions linked to this object. The LinkedTxn.TxnType can be set to ReimburseCharge. The LinkedTxn.TxnId can be set as the ID of the transaction. |
| ExchangeRate |  | Optional | default: 1 Decimal The number of home currency units it takes to equal one unit of currency specified by CurrencyRef. Applicable if multicurrency is enabled for the company. |
| APAccountRef | ReferenceType | Optional | Specifies to which AP account the bill is credited. Query the Account name list resource to determine the appropriate Account object for this reference. Use Account. |
| DepartmentRef | ReferenceType | Optional | A reference to a Department object specifying the location of the transaction. Available if Preferences.AccountingInfoPrefs.TrackDepartments is set to true. |
| TxnDate | Date | Optional | The date entered by the user when this transaction occurred. For posting transactions, this is the posting date that affects the financial statements. |
| IncludeInAnnualTPAR |  | Optional |  |
| TransactionLocationType | String | Optional | The account location. Valid values for France: WithinFrance, FranceOverseas, OutsideFranceWithEU, OutsideEU. |
| Balance | Decimal | Read-only | The current amount of the vendor credit reflecting any adjustments to the original credit amount. Initially set to the value of TotalAmt. |
| MetaData | ModificationMetaData | Read-only | Descriptive information about the object. The MetaData values are set by Data Services and are  for all applications. |
| RecurDataRef | ReferenceType | Read-only | A reference to the Recurring Transaction. It captures what recurring transaction template the VendorCredit was created from. |
| TotalAmt | BigDecimal | Read-only | Indicates the total credit amount, determined by taking the total of all all lines of the transaction. This includes all charges, allowances, discounts, and taxes. |

## Sample Object

```json
{
  "VendorCredit": {
    "SyncToken": "0",
    "domain": "QBO",
    "VendorRef": {
      "name": "Books by Bessie",
      "value": "30"
    },
    "TxnDate": "2014-12-23",
    "TotalAmt": 90.0,
    "APAccountRef": {
      "name": "Accounts Payable (A/P)",
      "value": "33"
    },
    "sparse": false,
    "Line": [
      {
        "DetailType": "AccountBasedExpenseLineDetai l",
        "Amount": 90.0,
        "ProjectRef": {
          "value": "39298045"
        },
        "Id": "1",
        "AccountBasedExpenseLineDetail": {
          "TaxCodeRef": {
            "value": "TAX"
          },
          "AccountRef": {
            "name": "Bank Charges",
            "value": "8"
          },
          "BillableStatus": "Billable",
          "CustomerRef": {
            "name": "Amy's Bird Sanctuary",
            "value": "1"
          }
        }
      }
    ],
    "Id": "255",
    "MetaData": {
      "CreateTime": "2015-07-28T14:13:30-07:00",
      "LastUpdatedTime": "2015-07-28T14:13:30-07:00"
    }
  },
  "time": "2015-07-28T14:16:42.709-07:00"
}
```

## Operations

### Create a vendorcredit

- **Method**: POST
- **URL**: `/v3/company/{realmID}/vendorcredit`

The VendorRef attribute must be specified.At lease one Line with Line.Amount must be specified.

**Request Body**:
```json
{
  "TotalAmt": 90.0,
  "TxnDate": "2014-12-23",
  "Line": [
    {
      "DetailType": "AccountBasedExpenseLineDetail",
      "Amount": 90.0,
      "ProjectRef": {
        "value": "39298045"
      },
      "Id": "1",
      "AccountBasedExpenseLineDetail": {
        "TaxCodeRef": {
          "value": "TAX"
        },
        "AccountRef": {
          "name": "Bank Charges",
          "value": "8"
        },
        "BillableStatus": "Billable",
        "CustomerRef": {
          "name": "Amy's Bird Sanctuary",
          "value": "1"
        }
      }
    }
  ],
  "APAccountRef": {
    "name": "Accounts Payable (A/P)",
    "value": "33"
  },
  "VendorRef": {
    "name": "Books by Bessie",
    "value": "30"
  }
}
```

**Response**:
```json
{
  "VendorCredit": {
    "SyncToken": "0",
    "domain": "QBO",
    "VendorRef": {
      "name": "Books by Bessie",
      "value": "30"
    },
    "TxnDate": "2014-12-23",
    "TotalAmt": 90.0,
    "APAccountRef": {
      "name": "Accounts Payable (A/P)",
      "value": "33"
    },
    "sparse": false,
    "Line": [
      {
        "DetailType": "AccountBasedExpenseLineDetai l",
        "Amount": 90.0,
        "ProjectRef": {
          "value": "39298045"
        },
        "Id": "1",
        "AccountBasedExpenseLineDetail": {
          "TaxCodeRef": {
            "value": "TAX"
          },
          "AccountRef": {
            "name": "Bank Charges",
            "value": "8"
          },
          "BillableStatus": "Billable",
          "CustomerRef": {
            "name": "Amy's Bird Sanctuary",
            "value": "1"
          }
        }
      }
    ],
    "Id": "157",
    "MetaData": {
      "CreateTime": "2014-12-23T11:14:15-08:00",
      "LastUpdatedTime": "2014-12-23T11:14:15-08:00"
    }
  },
  "time": "2014-12-23T11:14:15.462-08:00"
}
```

### Delete a vendorcredit

- **Method**: POST
- **URL**: `/v3/company/{realmID}/vendorcredit?operation=delete`

This operation deletes the vendorcredit object specified in the request body. Include a minimum of VendorCredit.Id and VendorCredit.SyncToken in the request body.

**Request Body**:
```json
{
  "SyncToken": "0",
  "Id": "13"
}
```

**Response**:
```json
{
  "VendorCredit": {
    "status": "Deleted",
    "domain": "QBO",
    "Id": "13"
  },
  "time": "2015-05-27T10:42:58.468-07:00"
}
```

### Query a vendorcredit

- **Method**: GET
- **URL**: `/v3/company/{realmID}/query?query={selectStatement}`

**Sample Query**: `select * from vendorcredit`

**Response**:
```json
{
  "QueryResponse": {
    "startPosition": 1,
    "totalCount": 2,
    "VendorCredit": [
      {
        "SyncToken": "0",
        "domain": "QBO",
        "VendorRef": {
          "name": "Books by Bessie",
          "value": "30"
        },
        "TxnDate": "2014-12-23",
        "TotalAmt": 90.0,
        "APAccountRef": {
          "name": "Accounts Payable (A /P)",
          "value": "33"
        },
        "sparse": false,
        "Line": [
          {
            "DetailType": "AccountBasedExpenseLineD etail",
            "Amount": 90.0,
            "ProjectRef": {
              "value": "39298045"
            },
            "Id": "1",
            "AccountBasedExpenseLineD etail": {
              "TaxCodeRef": {
                "value": "TAX"
              },
              "AccountRef": {
                "name": "Bank Charges",
                "value": "8"
              },
              "BillableStatus": "Billable",
              "CustomerRef": {
                "name": "Amy's Bird Sanctuary",
                "value": "1"
              }
            }
          }
        ],
        "Id": "255",
        "MetaData": {
          "CreateTime": "2015-07-28T14:13:30-07:00",
          "LastUpdatedTime": "2015-07-28T14:13:30-07:00"
        }
      },
      {
        "SyncToken": "0",
        "domain": "QBO",
        "VendorRef": {
          "name": "Books by Bessie",
          "value": "30"
        },
        "TxnDate": "2014-12-23",
        "TotalAmt": 90.0,
        "APAccountRef": {
          "name": "Accounts Payable (A /P)",
          "value": "33"
        },
        "sparse": false,
        "Line": [
          {
            "DetailType": "AccountBasedExpenseLineD etail",
            "Amount": 90.0,
            "ProjectRef": {
              "value": "39298045"
            },
            "Id": "1",
            "AccountBasedExpenseLineD etail": {
              "TaxCodeRef": {
                "value": "TAX"
              },
              "AccountRef": {
                "name": "Bank Charges",
                "value": "8"
              },
              "BillableStatus": "Billable",
              "CustomerRef": {
                "name": "Amy's Bird Sanctuary",
                "value": "1"
              }
            }
          }
        ],
        "Id": "253",
        "MetaData": {
          "CreateTime": "2015-07-28T14:13:08-07:00",
          "LastUpdatedTime": "2015-07-28T14:13:08-07:00"
        }
      }
    ],
    "maxResults": 2
  },
  "time": "2015-07-28T14:14:36.327-07:00"
}
```

### Read a vendorcredit

- **Method**: GET
- **URL**: `/v3/company/{realmID}/vendorcredit/{vendorcreditId}`

Retrieves the details of a VendorcCredit object that has been previously created.

**Response**:
```json
{
  "VendorCredit": {
    "SyncToken": "0",
    "domain": "QBO",
    "VendorRef": {
      "name": "Books by Bessie",
      "value": "30"
    },
    "TxnDate": "2014-12-23",
    "TotalAmt": 90.0,
    "APAccountRef": {
      "name": "Accounts Payable (A/P)",
      "value": "33"
    },
    "sparse": false,
    "Line": [
      {
        "DetailType": "AccountBasedExpenseLineDetai l",
        "Amount": 90.0,
        "ProjectRef": {
          "value": "39298045"
        },
        "Id": "1",
        "AccountBasedExpenseLineDetail": {
          "TaxCodeRef": {
            "value": "TAX"
          },
          "AccountRef": {
            "name": "Bank Charges",
            "value": "8"
          },
          "BillableStatus": "Billable",
          "CustomerRef": {
            "name": "Amy's Bird Sanctuary",
            "value": "1"
          }
        }
      }
    ],
    "Id": "255",
    "MetaData": {
      "CreateTime": "2015-07-28T14:13:30-07:00",
      "LastUpdatedTime": "2015-07-28T14:13:30-07:00"
    }
  },
  "time": "2015-07-28T14:16:42.709-07:00"
}
```

### Full update a vendorcredit

- **Method**: POST
- **URL**: `/v3/company/{realmID}/vendorcredit`

Use this operation to update any of the writable fields of an existing vendorcredit object. The request body must include all writable fields of the existing object as returned in a read response. Writable fields omitted from the request body are set to NULL. The ID of the object to update is specif...

**Request Body**:
```json
{
  "SyncToken": "1",
  "domain": "QBO",
  "VendorRef": {
    "name": "Books by Bessie",
    "value": "30"
  },
  "TxnDate": "2014-12-23",
  "TotalAmt": 140.0,
  "APAccountRef": {
    "name": "Accounts Payable (A/P)",
    "value": "33"
  },
  "sparse": false,
  "Line": [
    {
      "DetailType": "AccountBasedExpenseLineDetail",
      "Amount": 140.0,
      "ProjectRef": {
        "value": "39298045"
      },
      "Id": "1",
      "AccountBasedExpenseLineDetail": {
        "TaxCodeRef": {
          "value": "TAX"
        },
        "AccountRef": {
          "name": "Bank Charges",
          "value": "8"
        },
        "BillableStatus": "Billable",
        "CustomerRef": {
          "name": "Amy's Bird Sanctuary",
          "value": "1"
        }
      }
    }
  ],
  "Id": "255",
  "MetaData": {
    "CreateTime": "2015-07-28T14:13:30-07:00",
    "LastUpdatedTime": "2015-07-28T14:22:05-07:00"
  }
}
```

**Response**:
```json
{
  "VendorCredit": {
    "SyncToken": "2",
    "domain": "QBO",
    "VendorRef": {
      "name": "Books by Bessie",
      "value": "30"
    },
    "TxnDate": "2014-12-23",
    "TotalAmt": 140.0,
    "APAccountRef": {
      "name": "Accounts Payable (A/P)",
      "value": "33"
    },
    "sparse": false,
    "Line": [
      {
        "DetailType": "AccountBasedExpenseLineDetai l",
        "Amount": 140.0,
        "ProjectRef": {
          "value": "39298045"
        },
        "Id": "1",
        "AccountBasedExpenseLineDetail": {
          "TaxCodeRef": {
            "value": "TAX"
          },
          "AccountRef": {
            "name": "Bank Charges",
            "value": "8"
          },
          "BillableStatus": "Billable",
          "CustomerRef": {
            "name": "Amy's Bird Sanctuary",
            "value": "1"
          }
        }
      }
    ],
    "Id": "255",
    "MetaData": {
      "CreateTime": "2015-07-28T14:13:30-07:00",
      "LastUpdatedTime": "2015-07-28T14:23:50-07:00"
    }
  },
  "time": "2015-07-28T14:23:52.196-07:00"
}
```
