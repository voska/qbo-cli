# BillPayment

A BillPayment object represents the payment transaction for a bill that the business owner receives from a vendor for goods or services purchased from the vendor. QuickBooks Online supports bill payments through a credit card or a checking account. BillPayment.TotalAmt is the total amount associated with this payment. This includes the total of all the payments from the payment line details. If TotalAmt is greater than the total on the lines being paid, the overpayment is treated as a credit and exposed as such on the QuickBooks UI. The total amount cannot be negative.

## Attributes

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| Id | String | Required for update | Unique Identifier for an Intuit entity (object). Sort order is ASC by default. |
| VendorRef | ReferenceType | Required | Reference to the vendor for this transaction. Query the Vendor name list resource to determine the appropriate Vendor object for this reference. Use Vendor.Id and Vendor. |
| Line | Line | Required | Individual line items representing zero or more Bill, VendorCredit, and JournalEntry objects linked to this BillPayment object. Use Line.LinkedTxn. |
| TotalAmt | BigDecimal | Required | Indicates the total amount associated with this payment. This includes the total of all the payments from the payment line details. |
| PayType |  | Required | BillPaymentTypeEnum The payment type. Valid values include: Check, CreditCard |
| SyncToken | String | Required for update | Version number of the object. It is used to lock an object for use by one app at a time. As soon as an application modifies an object, its SyncToken is incremented. |
| CurrencyRef | CurrencyRefType | Conditionally required | Reference to the currency in which all amounts on the associated transaction are expressed. This must be defined if multicurrency is enabled for the company. |
| DocNumber | String | Optional | Reference number for the transaction. If not explicitly provided at create time, a custom value can be provided. |
| PrivateNote | String | Optional | User entered, organization-private note about the transaction. This note does not appear on the invoice to the customer. |
| TxnDate | Date | Optional | The date entered by the user when this transaction occurred. For posting transactions, this is the posting date that affects the financial statements. |
| ExchangeRate | Decimal | Optional | The number of home currency units it takes to equal one unit of currency specified by CurrencyRef. Applicable if multicurrency is enabled for the company. |
| APAccountRef | ReferenceType | Optional | Specifies to which AP account the bill is credited. Query the Account name list resource to determine the appropriate Account object for this reference. Use Account. |
| DepartmentRef | ReferenceType | Optional | A reference to a Department object specifying the location of the transaction, as defined using location tracking in QuickBooks Online. |
| TransactionLocationType | String | Optional | The account location. Valid values for France: WithinFrance, FranceOverseas, OutsideFranceWithEU, OutsideEU. |
| ProcessBillPayment | Boolean | Optional | Indicates that the payment should be processed by merchant account service. Valid for QuickBooks companies with credit card processing. |
| MetaData | ModificationMetaData | Read-only | Descriptive information about the object. The MetaData values are set by Data Services and are  for all applications. |
| CheckPayment |  | Optional | BillPaymentCheck   Information about a check payment for the transaction. Not applicable to Estimate and SalesOrder. Used when PayType is Check. |
| CreditCardPayment |  | Optional | BillPaymentCreditCard   Information about a credit card payment for the transaction. Not applicable to Estimate and SalesOrder. Used when PayType is CreditCard. |

## Sample Object

```json
{
  "BillPayment": {
    "SyncToken": "0",
    "domain": "QBO",
    "VendorRef": {
      "name": "Bob's Burger Joint",
      "value": "56"
    },
    "TxnDate": "2015-07-14",
    "TotalAmt": 200.0,
    "PayType": "Check",
    "PrivateNote": "Acct. 1JK90",
    "sparse": false,
    "Line": [
      {
        "Amount": 200.0,
        "LinkedTxn": [
          {
            "TxnId": "234",
            "TxnType": "Bill"
          }
        ]
      }
    ],
    "Id": "236",
    "CheckPayment": {
      "PrintStatus": "NeedToPrint",
      "BankAccountRef": {
        "name": "Checking",
        "value": "35"
      }
    },
    "MetaData": {
      "CreateTime": "2015-07-14T12:34:04-07:00",
      "LastUpdatedTime": "2015-07-14T12:34:04-07:00"
    }
  },
  "time": "2015-07-14T12:34:03.964-07:00"
}
```

## Operations

### Create a billpayment

- **Method**: POST
- **URL**: `/v3/company/{realmID}/billpayment`

**Request Body**:
```json
{
  "PrivateNote": "Acct. 1JK90",
  "VendorRef": {
    "name": "Bob's Burger Joint",
    "value": "56"
  },
  "TotalAmt": 200.0,
  "PayType": "Check",
  "Line": [
    {
      "Amount": 200.0,
      "LinkedTxn": [
        {
          "TxnId": "234",
          "TxnType": "Bill"
        }
      ]
    }
  ],
  "CheckPayment": {
    "BankAccountRef": {
      "name": "Checking",
      "value": "35"
    }
  }
}
```

**Response**:
```json
{
  "BillPayment": {
    "SyncToken": "0",
    "domain": "QBO",
    "VendorRef": {
      "name": "Bob's Burger Joint",
      "value": "56"
    },
    "TxnDate": "2015-07-14",
    "TotalAmt": 200.0,
    "PayType": "Check",
    "PrivateNote": "Acct. 1JK90",
    "sparse": false,
    "Line": [
      {
        "Amount": 200.0,
        "LinkedTxn": [
          {
            "TxnId": "234",
            "TxnType": "Bill"
          }
        ]
      }
    ],
    "Id": "236",
    "CheckPayment": {
      "PrintStatus": "NeedToPrint",
      "BankAccountRef": {
        "name": "Checking",
        "value": "35"
      }
    },
    "MetaData": {
      "CreateTime": "2015-07-14T12:34:04-07:00",
      "LastUpdatedTime": "2015-07-14T12:34:04-07:00"
    }
  },
  "time": "2015-07-14T12:34:03.964-07:00"
}
```

### Void a billpayment

- **Method**: POST
- **URL**: `/v3/company/{realmID}/billpayment?operation=update&include=void`

Use a sparse update operation with include=void to void an existing BillPayment object; include a minimum of BillPayment.Id and BillPayment.SyncToken. The transaction remains active but all amounts and quantities are zeroed, all lines are cleared, and the string, Voided, is injected into BillPayment...

**Request Body**:
```json
{
  "SyncToken": "0",
  "Id": "104",
  "sparse": true
}
```

**Response**:
```json
{
  "BillPayment": {
    "DocNumber": "11",
    "SyncToken": "1",
    "domain": "QBO",
    "VendorRef": {
      "name": "Hall Properties",
      "value": "40"
    },
    "TxnDate": "2016-08-18",
    "TotalAmt": 0,
    "CurrencyRef": {
      "name": "United States Dollar",
      "value": "USD"
    },
    "PayType": "Check",
    "PrivateNote": "Voided",
    "sparse": false,
    "Line": [],
    "Id": "104",
    "CheckPayment": {
      "PrintStatus": "NotSet",
      "BankAccountRef": {
        "name": "Cash on hand",
        "value": "9"
      }
    },
    "MetaData": {
      "CreateTime": "2016-08-18T13:11:14-07:00",
      "LastUpdatedTime": "2016-08-18T13:27:13-07:00"
    }
  },
  "time": "2016-08-18T13:27:13.323-07:00"
}
```

### Delete a billpayment

- **Method**: POST
- **URL**: `/v3/company/{realmID}/billpayment?operation=delete`

This operation deletes the billpayment object specified in the request body. Include a minimum of BillPayment.Id and BillPayment.SyncToken in the request body.

**Request Body**:
```json
{
  "SyncToken": "0",
  "Id": "117"
}
```

**Response**:
```json
{
  "BillPayment": {
    "status": "Deleted",
    "domain": "QBO",
    "Id": "117"
  },
  "time": "2015-05-26T13:17:25.316-07:00"
}
```

### Query a billpayment

- **Method**: GET
- **URL**: `/v3/company/{realmID}/query?query={selectStatement}`

**Sample Query**: `select * from billpayment Where Metadata.LastUpdatedTime>'2014-12-12T14: 50:22-08:00' Order By Metadata.LastUpdatedTime`

**Response**:
```json
{
  "QueryResponse": {
    "BillPayment": [
      {
        "DocNumber": "1",
        "SyncToken": "0",
        "domain": "QBO",
        "VendorRef": {
          "name": "PG&E",
          "value": "48"
        },
        "TxnDate": "2015-01-16",
        "TotalAmt": 86.44,
        "CurrencyRef": {
          "name": "United States Dollar",
          "value": "USD"
        },
        "PayType": "CreditCard",
        "PrivateNote": "00649587213",
        "sparse": false,
        "CreditCardPayment": {
          "CCAccountRef": {
            "name": "Mastercard",
            "value": "41"
          }
        },
        "Line": [
          {
            "Amount": 86.44,
            "LinkedTxn": [
              {
                "TxnId": "78",
                "TxnType": "Bill"
              }
            ]
          }
        ],
        "Id": "165",
        "MetaData": {
          "CreateTime": "2015-01-16T15:36:20-08:00",
          "LastUpdatedTime": "2015-01-16T15:36:20-08:00"
        }
      },
      {
        "SyncToken": "0",
        "domain": "QBO",
        "VendorRef": {
          "name": "Bob's Burger Joint",
          "value": "56"
        },
        "TxnDate": "2015-01-16",
        "TotalAmt": 200.0,
        "PayType": "CreditCard",
        "sparse": false,
        "CreditCardPayment": {
          "CCAccountRef": {
            "name": "Mastercard",
            "value": "41"
          }
        },
        "Line": [
          {
            "Amount": 200.0,
            "LinkedTxn": [
              {
                "TxnId": "157",
                "TxnType": "Bill"
              }
            ]
          }
        ],
        "Id": "166",
        "MetaData": {
          "CreateTime": "2015-01-16T15:40:26-08:00",
          "LastUpdatedTime": "2015-01-16T15:40:26-08:00"
        }
      },
      {
        "SyncToken": "0",
        "domain": "QBO",
        "VendorRef": {
          "name": "Norton Lumber and Building Materials",
          "value": "46"
        },
        "TxnDate": "2015-01-16",
        "TotalAmt": 205.0,
        "PayType": "CreditCard",
        "sparse": false,
        "CreditCardPayment": {
          "CCAccountRef": {
            "name": "Mastercard",
            "value": "41"
          }
        },
        "Line": [
          {
            "Amount": 205.0,
            "LinkedTxn": [
              {
                "TxnId": "126",
                "TxnType": "Bill"
              }
            ]
          }
        ],
        "Id": "169",
        "MetaData": {
          "CreateTime": "2015-01-16T16:00:29-08:00",
          "LastUpdatedTime": "2015-01-16T16:00:29-08:00"
        }
      },
      {
        "SyncToken": "0",
        "domain": "QBO",
        "VendorRef": {
          "name": "Robertson & Associates",
          "value": "49"
        },
        "TxnDate": "2015-06-30",
        "TotalAmt": 110.0,
        "PayType": "Check",
        "PrivateNote": "Acct. 1JK90",
        "sparse": false,
        "Line": [
          {
            "Amount": 110.0,
            "LinkedTxn": [
              {
                "TxnId": "108",
                "TxnType": "Bill"
              }
            ]
          }
        ],
        "Id": "231",
        "CheckPayment": {
          "PrintStatus": "NeedToPrint",
          "BankAccountRef": {
            "name": "Checking",
            "value": "35"
          }
        },
        "MetaData": {
          "CreateTime": "2015-06-30T15:05:30-07:00",
          "LastUpdatedTime": "2015-06-30T15:05:30-07:00"
        }
      },
      {
        "SyncToken": "0",
        "domain": "QBO",
        "VendorRef": {
          "name": "Robertson & Associates",
          "value": "49"
        },
        "TxnDate": "2015-06-30",
        "TotalAmt": 110.0,
        "PayType": "Check",
        "PrivateNote": "Acct. 1JK90",
        "sparse": false,
        "Line": [
          {
            "Amount": 110.0,
            "LinkedTxn": [
              {
                "TxnId": "108",
                "TxnType": "Bill"
              }
            ]
          }
        ],
        "Id": "232",
        "CheckPayment": {
          "PrintStatus": "NeedToPrint",
          "BankAccountRef": {
            "name": "Checking",
            "value": "35"
          }
        },
        "MetaData": {
          "CreateTime": "2015-06-30T15:09:06-07:00",
          "LastUpdatedTime": "2015-06-30T15:09:06-07:00"
        }
      },
      {
        "DocNumber": "1",
        "SyncToken": "2",
        "domain": "QBO",
        "VendorRef": {
          "name": "Brosnahan Insurance Agency",
          "value": "31"
        },
        "TxnDate": "2014-09-16",
        "TotalAmt": 2000.0,
        "PayType": "Check",
        "PrivateNote": "Add private note",
        "sparse": false,
        "Line": [
          {
            "Amount": 2000.0,
            "LinkedTxn": [
              {
                "TxnId": "1",
                "TxnType": "Bill"
              }
            ]
          }
        ],
        "Id": "22",
        "CheckPayment": {
          "PrintStatus": "NotSet",
          "BankAccountRef": {
            "name": "Checking",
            "value": "35"
          }
        },
        "MetaData": {
          "CreateTime": "2014-09-16T15:28:48-07:00",
          "LastUpdatedTime": "2015-06-30T15:24:40-07:00"
        }
      },
      {
        "SyncToken": "0",
        "domain": "QBO",
        "VendorRef": {
          "name": "Bob's Burger Joint",
          "value": "56"
        },
        "TxnDate": "2015-07-14",
        "TotalAmt": 200.0,
        "PayType": "Check",
        "PrivateNote": "Acct. 1JK90",
        "sparse": false,
        "Line": [
          {
            "Amount": 200.0,
            "LinkedTxn": [
              {
                "TxnId": "234",
                "TxnType": "Bill"
              }
            ]
          }
        ],
        "Id": "236",
        "CheckPayment": {
          "PrintStatus": "NeedToPrint",
          "BankAccountRef": {
            "name": "Checking",
            "value": "35"
          }
        },
        "MetaData": {
          "CreateTime": "2015-07-14T12:34:04-07:00",
          "LastUpdatedTime": "2015-07-14T12:34:04-07:00"
        }
      },
      {
        "SyncToken": "0",
        "domain": "QBO",
        "VendorRef": {
          "name": "Bob's Burger Joint",
          "value": "56"
        },
        "TxnDate": "2015-07-14",
        "TotalAmt": 110.0,
        "PayType": "Check",
        "sparse": false,
        "Line": [],
        "Id": "237",
        "CheckPayment": {
          "PrintStatus": "NotSet",
          "BankAccountRef": {
            "name": "Checking",
            "value": "35"
          }
        },
        "MetaData": {
          "CreateTime": "2015-07-14T12:37:57-07:00",
          "LastUpdatedTime": "2015-07-14T12:37:57-07:00"
        }
      }
    ],
    "startPosition": 1,
    "maxResults": 8,
    "totalCount": 8
  },
  "time": "2015-07-14T12:48:36.854-07:00"
}
```

### Read a billpayment

- **Method**: GET
- **URL**: `/v3/company/{realmID}/billpayment/{billpaymentId}`

Retrieves the details of a billpayment that has been previously created.

**Response**:
```json
{
  "BillPayment": {
    "SyncToken": "0",
    "domain": "QBO",
    "VendorRef": {
      "name": "Bob's Burger Joint",
      "value": "56"
    },
    "TxnDate": "2015-07-14",
    "TotalAmt": 200.0,
    "PayType": "Check",
    "PrivateNote": "Acct. 1JK90",
    "sparse": false,
    "Line": [
      {
        "Amount": 200.0,
        "LinkedTxn": [
          {
            "TxnId": "234",
            "TxnType": "Bill"
          }
        ]
      }
    ],
    "Id": "236",
    "CheckPayment": {
      "PrintStatus": "NeedToPrint",
      "BankAccountRef": {
        "name": "Checking",
        "value": "35"
      }
    },
    "MetaData": {
      "CreateTime": "2015-07-14T12:34:04-07:00",
      "LastUpdatedTime": "2015-07-14T12:34:04-07:00"
    }
  },
  "time": "2015-07-14T12:39:40.606-07:00"
}
```

### Full update a billpayment

- **Method**: POST
- **URL**: `/v3/company/{realmID}/billpayment`

Use this operation to update any of the writable fields of an existing billpayment object. The request body must include all writable fields of the existing object as returned in a read response. Writable fields omitted from the request body are set to NULL. The ID of the object to update is specifi...

**Request Body**:
```json
{
  "SyncToken": "2",
  "domain": "QBO",
  "VendorRef": {
    "name": "Bob's Burger Joint",
    "value": "56"
  },
  "TxnDate": "2015-07-14",
  "TotalAmt": 200.0,
  "PayType": "Check",
  "PrivateNote": "A new private note",
  "sparse": false,
  "Line": [
    {
      "Amount": 200.0,
      "LinkedTxn": [
        {
          "TxnId": "234",
          "TxnType": "Bill"
        }
      ]
    }
  ],
  "Id": "236",
  "CheckPayment": {
    "PrintStatus": "NeedToPrint",
    "BankAccountRef": {
      "name": "Checking",
      "value": "35"
    }
  },
  "MetaData": {
    "CreateTime": "2015-07-14T12:34:04-07:00",
    "LastUpdatedTime": "2015-07-14T13:17:22-07:00"
  }
}
```

**Response**:
```json
{
  "BillPayment": {
    "SyncToken": "3",
    "domain": "QBO",
    "VendorRef": {
      "name": "Bob's Burger Joint",
      "value": "56"
    },
    "TxnDate": "2015-07-14",
    "TotalAmt": 200.0,
    "PayType": "Check",
    "PrivateNote": "A new private note",
    "sparse": false,
    "Line": [
      {
        "Amount": 200.0,
        "LinkedTxn": [
          {
            "TxnId": "234",
            "TxnType": "Bill"
          }
        ]
      }
    ],
    "Id": "236",
    "CheckPayment": {
      "PrintStatus": "NeedToPrint",
      "BankAccountRef": {
        "name": "Checking",
        "value": "35"
      }
    },
    "MetaData": {
      "CreateTime": "2015-07-14T12:34:04-07:00",
      "LastUpdatedTime": "2015-07-30T09:55:19-07:00"
    }
  },
  "time": "2015-07-30T09:55:20.597-07:00"
}
```
