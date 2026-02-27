# Payment

A Payment object records a payment in QuickBooks. The payment can be applied for a particular customer against multiple Invoices and Credit Memos. It can also be created without any Invoice or Credit Memo, by just specifying an amount. • A Payment can be updated as a full update or a sparse update. • A Payment can be linked to multiple Invoices and Credit Memos. • A Payment can be created as unapplied to any Invoice or Credit Memo, in which case it is recorded as a credit. • If any element in any line needs to be updated, all the Line elements of the Payment object have to be provided. This is true for full or sparse update. Lines can be updated only ALL or NONE. • To remove all lines, send an empty Line element. • To remove some of the lines, send all the Lines that need to be present MINUS the lines that need to be removed. • To add some lines, send all existing and new Lines that need to be present. • The sequence in which the lines are received is the sequence in which lines are preserved. • If you have a large number of invoice and corresponding payment records that you wish to import to the QuickBooks Online company, sort the invoice and payment records in chronological order and use the batch resource to send invoice and payments batches of 10, one after the other, to ensure any open invoices get credited with their payments.

## Attributes

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| Id | String | Required for update | Unique identifier for this object. Sort order is ASC by default. |
| TotalAmt | Decimal | Required | Indicates the total amount of the transaction. This includes the total of all the charges, allowances, and taxes. |
| CustomerRef | ReferenceType | Required | Reference to a customer or job. Query the Customer name list resource to determine the appropriate Customer object for this reference. Use Customer.Id and Customer. |
| SyncToken | String | Required for update | Version number of the object. It is used to lock an object for use by one app at a time. As soon as an application modifies an object, its SyncToken is incremented. |
| CurrencyRef | CurrencyRefType | Conditionally required | Reference to the currency in which all amounts on the associated transaction are expressed. This must be defined if multicurrency is enabled for the company. |
| ProjectRef | ReferenceType | Conditionally required | Reference to the Project ID associated with this transaction. Available with Minor Version 69 and above |
| PrivateNote | String | Optional | User entered, organization-private note about the transaction. |
| PaymentMethodRef | ReferenceType | Optional | Reference to a PaymentMethod associated with this transaction. Query the PaymentMethod name list resource to determine the appropriate PaymentMethod object for this reference. |
| UnappliedAmt | Decimal | Read-only | Indicates the amount that has not been applied to pay amounts owed for sales transactions. |
| DepositToAccountRef | ReferenceType | Optional | Identifies the account to be used for this payment. Query the Account name list resource to determine the appropriate Account object for this reference, where Account. |
| ExchangeRate | Decimal | Optional | The number of home currency units it takes to equal one unit of currency specified by CurrencyRef. Applicable if multicurrency is enabled for the company |
| Line | Line | Optional | Zero or more transactions accounting for this payment. Values for Line.LinkedTxn.TxnTypecan be one of the following: |
| TxnSource | String | Optional | Used internally to specify originating source of a credit card transaction. |
| TxnDate | Date | Optional | The date entered by the user when this transaction occurred. For posting transactions, this is the posting date that affects the financial statements. |
| CreditCardPayment |  | Optional | CreditCardPayment Information about a payment received by credit card. Inject with data only if the payment was transacted through Intuit Payments API. |
| TransactionLocationType | String | Optional | The account location. Valid values for France: WithinFrance, FranceOverseas, OutsideFranceWithEU, OutsideEU. |
| MetaData | ModificationMetaData | Read-only | Descriptive information about the entity. The MetaData values are set by Data Services and are  for all applications. |
| PaymentRefNum |  | Optional | * Conditionally Required |
| TaxExemptionRef | ReferenceType | Read-only | Reference to the TaxExepmtion ID associated with this object. Available for companies that have |

## Sample Object

```json
{
  "Payment": {
    "SyncToken": "0",
    "domain": "QBO",
    "DepositToAccountRef": {
      "value": "4"
    },
    "UnappliedAmt": 10.0,
    "TxnDate": "2015-01-16",
    "TotalAmt": 65.0,
    "ProjectRef": {
      "value": "39298034"
    },
    "ProcessPayment": false,
    "sparse": false,
    "Line": [
      {
        "Amount": 55.0,
        "LineEx": {
          "any": [
            {
              "name": "{http://schema .intuit.com/finance /v3}NameValue",
              "nil": false,
              "value": {
                "Name": "txnId",
                "Value": "70"
              },
              "declaredType": "com .intuit.schema.finance .v3.NameValue",
              "scope": "javax.xml.bind .JAXBElement$GlobalScop e",
              "globalScope": true,
              "typeSubstituted": false
            },
            {
              "name": "{http://schema .intuit.com/finance /v3}NameValue",
              "nil": false,
              "value": {
                "Name": "txnOpenBalance",
                "Value": "71.00"
              },
              "declaredType": "com .intuit.schema.finance .v3.NameValue",
              "scope": "javax.xml.bind .JAXBElement$GlobalScop e",
              "globalScope": true,
              "typeSubstituted": false
            },
            {
              "name": "{http://schema .intuit.com/finance /v3}NameValue",
              "nil": false,
              "value": {
                "Name": "txnReferenceNumber",
                "Value": "1024"
              },
              "declaredType": "com .intuit.schema.finance .v3.NameValue",
              "scope": "javax.xml.bind .JAXBElement$GlobalScop e",
              "globalScope": true,
              "typeSubstituted": false
            }
          ]
        },
        "LinkedTxn": [
          {
            "TxnId": "70",
            "TxnType": "Invoice"
          }
        ]
      }
    ],
    "CustomerRef": {
      "name": "Red Rock Diner",
      "value": "20"
    },
    "Id": "163",
    "MetaData": {
      "CreateTime": "2015-01-16T15:08:12-08:00",
      "LastUpdatedTime": "2015-01-16T15:08:12-08:00"
    }
  },
  "time": "2015-07-28T15:16:15.435-07:00"
}
```

## Operations

### Create a payment

- **Method**: POST
- **URL**: `/v3/company/{realmID}/payment`

**Request Body**:
```json
{
  "TotalAmt": 25.0,
  "CustomerRef": {
    "value": "20"
  }
}
```

**Response**:
```json
{
  "Payment": {
    "SyncToken": "0",
    "domain": "QBO",
    "DepositToAccountRef": {
      "value": "4"
    },
    "UnappliedAmt": 25.0,
    "TxnDate": "2014-12-30",
    "TotalAmt": 25.0,
    "ProjectRef": {
      "value": "39298034"
    },
    "ProcessPayment": false,
    "sparse": false,
    "Line": [],
    "CustomerRef": {
      "name": "Red Rock Diner",
      "value": "20"
    },
    "Id": "154",
    "MetaData": {
      "CreateTime": "2014-12-30T10:26:03-08:00",
      "LastUpdatedTime": "2014-12-30T10:26:03-08:00"
    }
  },
  "time": "2014-12-30T10:26:03.668-08:00"
}
```

### Delete a payment

- **Method**: POST
- **URL**: `/v3/company/{realmID}/payment?operation=delete`

This operation deletes the Payment object specified in the request body. Include a minimum of Payment.Id and Payment.SyncToken in the request body.

**Request Body**:
```json
{
  "SyncToken": "2",
  "Id": "73"
}
```

**Response**:
```json
{
  "Payment": {
    "status": "Deleted",
    "domain": "QBO",
    "Id": "73"
  },
  "time": "2013-03-14T11:57:42.849-07:00"
}
```

### Void a payment

- **Method**: POST
- **URL**: `/v3/company/{realmID}/payment?operation=update&include=void`

Use a sparse update operation with include=void to void an existing Payment object; include a minimum of Payment.Id and Payment.SyncToken.The transaction remains active but all amounts and quantities are zeroed and the string, Voided, is injected into Payment.PrivateNote, prepended to existing text ...

**Request Body**:
```json
{
  "SyncToken": "1",
  "Id": "33",
  "sparse": true
}
```

**Response**:
```json
{
  "Payment": {
    "SyncToken": "2",
    "domain": "QBO",
    "PaymentMethodRef": {
      "value": "2"
    },
    "DepositToAccountRef": {
      "value": "35"
    },
    "UnappliedAmt": 0,
    "TxnDate": "2014-11-07",
    "TotalAmt": 0,
    "ProjectRef": {
      "value": "39298234"
    },
    "ProcessPayment": false,
    "PrivateNote": "Voided",
    "sparse": false,
    "Line": [],
    "CustomerRef": {
      "name": "Freeman Sporting Goods :55 Twin Lane",
      "value": "9"
    },
    "Id": "33",
    "MetaData": {
      "CreateTime": "2014-11-07T11:07:19-08:00",
      "LastUpdatedTime": "2015-02-23T12:52:07-08:00"
    }
  },
  "time": "2015-02-23T12:52:06.954-08:00"
}
```

### Get a payment as PDF

- **Method**: GET
- **URL**: `/v3/company/{realmID}/payment/{paymentId}/pdf`

### Query a payment

- **Method**: GET
- **URL**: `/v3/company/{realmID}/query?query={selectStatement}`

**Sample Query**: `select * from Payment Where Metadata.LastUpdatedTime>'2015-01-16' Order By Metadata.LastUpdatedTime`

**Response**:
```json
{
  "QueryResponse": {
    "startPosition": 1,
    "Payment": [
      {
        "SyncToken": "0",
        "domain": "QBO",
        "DepositToAccountRef": {
          "value": "4"
        },
        "UnappliedAmt": 55.0,
        "TxnDate": "2015-01-16",
        "TotalAmt": 55.0,
        "ProjectRef": {
          "value": "39298034"
        },
        "ProcessPayment": false,
        "sparse": false,
        "Line": [],
        "CustomerRef": {
          "name": "Red Rock Diner",
          "value": "20"
        },
        "Id": "161",
        "MetaData": {
          "CreateTime": "2015-01-16T14:58:32-08:00",
          "LastUpdatedTime": "2015-01-16T14:58:32-08:00"
        }
      },
      {
        "SyncToken": "0",
        "domain": "QBO",
        "DepositToAccountRef": {
          "value": "4"
        },
        "UnappliedAmt": 65.0,
        "TxnDate": "2015-01-16",
        "TotalAmt": 65.0,
        "ProjectRef": {
          "value": "39298034"
        },
        "ProcessPayment": false,
        "sparse": false,
        "Line": [],
        "CustomerRef": {
          "name": "Red Rock Diner",
          "value": "20"
        },
        "Id": "162",
        "MetaData": {
          "CreateTime": "2015-01-16T14:58:59-08:00",
          "LastUpdatedTime": "2015-01-16T14:58:59-08:00"
        }
      },
      {
        "SyncToken": "0",
        "domain": "QBO",
        "DepositToAccountRef": {
          "value": "4"
        },
        "UnappliedAmt": 10.0,
        "TxnDate": "2015-01-16",
        "TotalAmt": 65.0,
        "ProjectRef": {
          "value": "39298034"
        },
        "ProcessPayment": false,
        "sparse": false,
        "Line": [
          {
            "Amount": 55.0,
            "LineEx": {
              "any": [
                {
                  "name": "{http ://schema.intuit.com/finance /v3}NameValue",
                  "nil": false,
                  "value": {
                    "Name": "txnId",
                    "Value": "70"
                  },
                  "declaredType": "com .intuit.schema .finance.v3 .NameValue",
                  "scope": "javax.xml .bind .JAXBElement$Global Scope",
                  "globalScope": true,
                  "typeSubstituted": false
                },
                {
                  "name": "{http ://schema.intuit.com/finance /v3}NameValue",
                  "nil": false,
                  "value": {
                    "Name": "txnOpenBalance",
                    "Value": "71.00"
                  },
                  "declaredType": "com .intuit.schema .finance.v3 .NameValue",
                  "scope": "javax.xml .bind .JAXBElement$Global Scope",
                  "globalScope": true,
                  "typeSubstituted": false
                },
                {
                  "name": "{http ://schema.intuit.com/finance /v3}NameValue",
                  "nil": false,
                  "value": {
                    "Name": "txnReferenceNumber ",
                    "Value": "1024"
                  },
                  "declaredType": "com .intuit.schema .finance.v3 .NameValue",
                  "scope": "javax.xml .bind .JAXBElement$Global Scope",
                  "globalScope": true,
                  "typeSubstituted": false
                }
              ]
            },
            "LinkedTxn": [
              {
                "TxnId": "70",
                "TxnType": "Invoice"
              }
            ]
          }
        ],
        "CustomerRef": {
          "name": "Red Rock Diner",
          "value": "20"
        },
        "Id": "163",
        "MetaData": {
          "CreateTime": "2015-01-16T15:08:12-08:00",
          "LastUpdatedTime": "2015-01-16T15:08:12-08:00"
        }
      },
      {
        "SyncToken": "0",
        "domain": "QBO",
        "DepositToAccountRef": {
          "value": "4"
        },
        "UnappliedAmt": 245.0,
        "TxnDate": "2015-01-16",
        "TotalAmt": 300.0,
        "ProjectRef": {
          "value": "39298034"
        },
        "ProcessPayment": false,
        "sparse": false,
        "Line": [
          {
            "Amount": 55.0,
            "LineEx": {
              "any": [
                {
                  "name": "{http ://schema.intuit.com/finance /v3}NameValue",
                  "nil": false,
                  "value": {
                    "Name": "txnId",
                    "Value": "70"
                  },
                  "declaredType": "com .intuit.schema .finance.v3 .NameValue",
                  "scope": "javax.xml .bind .JAXBElement$Global Scope",
                  "globalScope": true,
                  "typeSubstituted": false
                },
                {
                  "name": "{http ://schema.intuit.com/finance /v3}NameValue",
                  "nil": false,
                  "value": {
                    "Name": "txnOpenBalance",
                    "Value": "71.00"
                  },
                  "declaredType": "com .intuit.schema .finance.v3 .NameValue",
                  "scope": "javax.xml .bind .JAXBElement$Global Scope",
                  "globalScope": true,
                  "typeSubstituted": false
                },
                {
                  "name": "{http ://schema.intuit.com/finance /v3}NameValue",
                  "nil": false,
                  "value": {
                    "Name": "txnReferenceNumber ",
                    "Value": "1024"
                  },
                  "declaredType": "com .intuit.schema .finance.v3 .NameValue",
                  "scope": "javax.xml .bind .JAXBElement$Global Scope",
                  "globalScope": true,
                  "typeSubstituted": false
                }
              ]
            },
            "LinkedTxn": [
              {
                "TxnId": "70",
                "TxnType": "Invoice"
              }
            ]
          }
        ],
        "CustomerRef": {
          "name": "Red Rock Diner",
          "value": "20"
        },
        "Id": "164",
        "MetaData": {
          "CreateTime": "2015-01-16T15:09:22-08:00",
          "LastUpdatedTime": "2015-01-16T15:09:22-08:00"
        }
      },
      {
        "SyncToken": "0",
        "domain": "QBO",
        "DepositToAccountRef": {
          "value": "4"
        },
        "UnappliedAmt": 0,
        "TxnDate": "2015-02-04",
        "TotalAmt": 15.0,
        "ProjectRef": {
          "value": "39298034"
        },
        "ProcessPayment": false,
        "sparse": false,
        "Line": [
          {
            "Amount": 15.0,
            "LineEx": {
              "any": [
                {
                  "name": "{http ://schema.intuit.com/finance /v3}NameValue",
                  "nil": false,
                  "value": {
                    "Name": "txnId",
                    "Value": "70"
                  },
                  "declaredType": "com .intuit.schema .finance.v3 .NameValue",
                  "scope": "javax.xml .bind .JAXBElement$Global Scope",
                  "globalScope": true,
                  "typeSubstituted": false
                },
                {
                  "name": "{http ://schema.intuit.com/finance /v3}NameValue",
                  "nil": false,
                  "value": {
                    "Name": "txnOpenBalance",
                    "Value": "31.00"
                  },
                  "declaredType": "com .intuit.schema .finance.v3 .NameValue",
                  "scope": "javax.xml .bind .JAXBElement$Global Scope",
                  "globalScope": true,
                  "typeSubstituted": false
                },
                {
                  "name": "{http ://schema.intuit.com/finance /v3}NameValue",
                  "nil": false,
                  "value": {
                    "Name": "txnReferenceNumber ",
                    "Value": "1024"
                  },
                  "declaredType": "com .intuit.schema .finance.v3 .NameValue",
                  "scope": "javax.xml .bind .JAXBElement$Global Scope",
                  "globalScope": true,
                  "typeSubstituted": false
                }
              ]
            },
            "LinkedTxn": [
              {
                "TxnId": "70",
                "TxnType": "Invoice"
              }
            ]
          }
        ],
        "CustomerRef": {
          "name": "Red Rock Diner",
          "value": "20"
        },
        "Id": "170",
        "MetaData": {
          "CreateTime": "2015-02-04T10:42:16-08:00",
          "LastUpdatedTime": "2015-02-04T10:42:16-08:00"
        }
      },
      {
        "SyncToken": "0",
        "domain": "QBO",
        "DepositToAccountRef": {
          "value": "4"
        },
        "UnappliedAmt": 55.0,
        "TxnDate": "2015-02-04",
        "TotalAmt": 55.0,
        "ProjectRef": {
          "value": "39298034"
        },
        "ProcessPayment": false,
        "sparse": false,
        "Line": [],
        "CustomerRef": {
          "name": "Red Rock Diner",
          "value": "20"
        },
        "Id": "171",
        "MetaData": {
          "CreateTime": "2015-02-04T10:42:33-08:00",
          "LastUpdatedTime": "2015-02-04T10:42:33-08:00"
        }
      }
    ],
    "maxResults": 6
  },
  "time": "2015-07-28T15:15:25.802-07:00"
}
```

### Read a payment

- **Method**: GET
- **URL**: `/v3/company/{realmID}/payment/{paymentId}`

Retrieves the details of a Payment object that has been previously created.

**Response**:
```json
{
  "Payment": {
    "SyncToken": "0",
    "domain": "QBO",
    "DepositToAccountRef": {
      "value": "4"
    },
    "UnappliedAmt": 10.0,
    "TxnDate": "2015-01-16",
    "TotalAmt": 65.0,
    "ProjectRef": {
      "value": "39298034"
    },
    "ProcessPayment": false,
    "sparse": false,
    "Line": [
      {
        "Amount": 55.0,
        "LineEx": {
          "any": [
            {
              "name": "{http://schema .intuit.com/finance /v3}NameValue",
              "nil": false,
              "value": {
                "Name": "txnId",
                "Value": "70"
              },
              "declaredType": "com .intuit.schema.finance .v3.NameValue",
              "scope": "javax.xml.bind .JAXBElement$GlobalScop e",
              "globalScope": true,
              "typeSubstituted": false
            },
            {
              "name": "{http://schema .intuit.com/finance /v3}NameValue",
              "nil": false,
              "value": {
                "Name": "txnOpenBalance",
                "Value": "71.00"
              },
              "declaredType": "com .intuit.schema.finance .v3.NameValue",
              "scope": "javax.xml.bind .JAXBElement$GlobalScop e",
              "globalScope": true,
              "typeSubstituted": false
            },
            {
              "name": "{http://schema .intuit.com/finance /v3}NameValue",
              "nil": false,
              "value": {
                "Name": "txnReferenceNumber",
                "Value": "1024"
              },
              "declaredType": "com .intuit.schema.finance .v3.NameValue",
              "scope": "javax.xml.bind .JAXBElement$GlobalScop e",
              "globalScope": true,
              "typeSubstituted": false
            }
          ]
        },
        "LinkedTxn": [
          {
            "TxnId": "70",
            "TxnType": "Invoice"
          }
        ]
      }
    ],
    "CustomerRef": {
      "name": "Red Rock Diner",
      "value": "20"
    },
    "Id": "163",
    "MetaData": {
      "CreateTime": "2015-01-16T15:08:12-08:00",
      "LastUpdatedTime": "2015-01-16T15:08:12-08:00"
    }
  },
  "time": "2015-07-28T15:16:15.435-07:00"
}
```

### Send a payment

- **Method**: POST
- **URL**: `/v3/company/{realmID}/payment/{paymentId}/send?sendTo={emailAddr}`

• The email address should be explicitly mentioned in the POST request

**Response**:
```json
{
  "Payment": {
    "SyncToken": "0",
    "domain": "QBO",
    "DepositToAccountRef": {
      "value": "4"
    },
    "UnappliedAmt": 10.0,
    "TxnDate": "2015-01-16",
    "TotalAmt": 65.0,
    "ProjectRef": {
      "value": "39298034"
    },
    "ProcessPayment": false,
    "sparse": false,
    "Line": [
      {
        "Amount": 55.0,
        "LineEx": {
          "any": [
            {
              "name": "{http://schema .intuit.com/finance /v3}NameValue",
              "nil": false,
              "value": {
                "Name": "txnId",
                "Value": "70"
              },
              "declaredType": "com .intuit.schema.finance .v3.NameValue",
              "scope": "javax.xml.bind .JAXBElement$GlobalScop e",
              "globalScope": true,
              "typeSubstituted": false
            },
            {
              "name": "{http://schema .intuit.com/finance /v3}NameValue",
              "nil": false,
              "value": {
                "Name": "txnOpenBalance",
                "Value": "71.00"
              },
              "declaredType": "com .intuit.schema.finance .v3.NameValue",
              "scope": "javax.xml.bind .JAXBElement$GlobalScop e",
              "globalScope": true,
              "typeSubstituted": false
            },
            {
              "name": "{http://schema .intuit.com/finance /v3}NameValue",
              "nil": false,
              "value": {
                "Name": "txnReferenceNumber",
                "Value": "1024"
              },
              "declaredType": "com .intuit.schema.finance .v3.NameValue",
              "scope": "javax.xml.bind .JAXBElement$GlobalScop e",
              "globalScope": true,
              "typeSubstituted": false
            }
          ]
        },
        "LinkedTxn": [
          {
            "TxnId": "70",
            "TxnType": "Invoice"
          }
        ]
      }
    ],
    "CustomerRef": {
      "name": "Red Rock Diner",
      "value": "20"
    },
    "Id": "163",
    "MetaData": {
      "CreateTime": "2015-01-16T15:08:12-08:00",
      "LastUpdatedTime": "2015-01-16T15:08:12-08:00"
    }
  },
  "time": "2015-07-28T15:16:15.435-07:00"
}
```

### Full update a payment

- **Method**: POST
- **URL**: `/v3/company/{realmID}/payment`

Use this operation to update any of the writable fields of an existing Payment object. The request body must include all writable fields of the existing object as returned in a read response. Writable fields omitted from the request body are set to NULL. The ID of the object to update is specified i...

**Request Body**:
```json
{
  "SyncToken": "0",
  "PaymentMethodRef": {
    "value": "16"
  },
  "ProjectRef": {
    "value": "39298045"
  },
  "PaymentRefNum": "123456",
  "sparse": false,
  "Line": [
    {
      "Amount": 300,
      "LinkedTxn": [
        {
          "TxnId": "67",
          "TxnType": "Invoice"
        }
      ]
    },
    {
      "Amount": 300,
      "LinkedTxn": [
        {
          "TxnId": "68",
          "TxnType": "CreditMemo"
        }
      ]
    }
  ],
  "CustomerRef": {
    "value": "16"
  },
  "Id": "69",
  "MetaData": {
    "CreateTime": "2013-03-13T14:49:21-07:00",
    "LastUpdatedTime": "2013-03-13T14:49:21-07:00"
  }
}
```

**Response**:
```json
{
  "Payment": {
    "SyncToken": "1",
    "domain": "QBO",
    "PaymentMethodRef": {
      "value": "16"
    },
    "UnappliedAmt": 0,
    "TxnDate": "2013-03-13",
    "TotalAmt": 0,
    "ProjectRef": {
      "value": "39298045"
    },
    "PaymentRefNum": "123456",
    "sparse": false,
    "Line": [
      {
        "Amount": 300,
        "LinkedTxn": [
          {
            "TxnId": "67",
            "TxnType": "Invoice"
          }
        ]
      },
      {
        "Amount": 300,
        "LinkedTxn": [
          {
            "TxnId": "68",
            "TxnType": "CreditMemo"
          }
        ]
      }
    ],
    "CustomerRef": {
      "value": "16"
    },
    "Id": "69",
    "MetaData": {
      "CreateTime": "2013-03-13T14:49:21-07:00",
      "LastUpdatedTime": "2013-03-13T14:49:21-07:00"
    }
  },
  "time": "2013-03-13T14:49:41.512-07:00"
}
```
