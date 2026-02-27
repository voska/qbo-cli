# Purchase

A Purchase object represents an expense, such as a purchase made from a vendor. Of note, • You must specify an AccountRef for all purchases. • The TotalAmtattribute must add up to sum of Line.Amount attributes. There are three types of purchases: Cash, Check, and Credit Card. • Cash Purchase contains information regarding a payment made in cash. • Check Purchase contains information regarding a payment made by check. • Credit Card Purchase contains information regarding a payment made by credit card or refunded/credited back to a credit card. For example, to create a transaction that sends a check to a vendor, create a Purchase object with PaymentType set to Check.

## Attributes

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| Id | String | Required for update | Unique identifier for this object. Sort order is ASC by default. |
| Line | Line | Required | Individual line items of a transaction. Valid Line types include ItemBasedExpenseLine (Available if Preferences.ProductAndServicesPrefs. |
| PaymentType | String | Required | Type can be Cash, Check, or CreditCard. |
| AccountRef | ReferenceType | Required | Specifies the account reference to which this purchase is applied based on the PaymentType. A type of Check should have bank account, CreditCard should specify credit card account, etc. |
| SyncToken | String | Required for update | Version number of the object. It is used to lock an object for use by one app at a time. As soon as an application modifies an object, its SyncToken is incremented. |
| CurrencyRef | CurrencyRefType | Conditionally required | Reference to the currency in which all amounts on the associated transaction are expressed. This must be defined if multicurrency is enabled for the company. |
| TxnDate | Date | Optional | The date entered by the user when this transaction occurred. For posting transactions, this is the posting date that affects the financial statements. |
| PrintStatus | String | Optional | PrintStatus is applicable only for Check. Ignored for CreditCardcharge or refund. Valid values: NotSet, NeedToPrint, PrintComplete. |
| RemitToAddr | PhysicalAddress | Read-only | Address to which the payment should be sent. This attribute is applicable only for Check. Ignored for CreditCard charge or refund. |
| TxnSource | String | Optional | Used internally to specify originating source of a credit card transaction. |
| LinkedTxn | LinkedTxn | Optional | Zero or more transactions linked to this object. The LinkedTxn.TxnType can be set to ReimburseCharge. The LinkedTxn.TxnId can be set as the ID of the transaction. |
| GlobalTaxCalculation | String | Optional | Method in which tax is applied. Allowed values are: TaxExcluded, TaxInclusive, and NotApplicable. |
| TransactionLocationType | String | Optional | The account location. Valid values for France: WithinFrance, FranceOverseas, OutsideFranceWithEU, OutsideEU. |
| MetaData | ModificationMetaData | Read-only | Descriptive information about the object. The MetaData values are set by Data Services and are  for all applications. |
| DocNumber | String | Optional | Reference number for the transaction. |
| PrivateNote | String | Optional | User-entered, organization-private note about the transaction. This field maps to the Memo field on the Expense form in the QuickBooks UI. |
| Credit | Boolean | Optional | False—it represents a charge. True—it represents a refund. Valid only for CreditCardpayment type. Validation Rules: Valid only for CreditCardtransactions. |
| TxnTaxDetail | TxnTaxDetail | Optional | This data type provides information for taxes charged on the transaction as a whole. |
| PaymentMethodRef | ReferenceType | Optional | Reference to a PaymentMethod associated with this transaction. Query the PaymentMethod name list resource to determine the appropriate PaymentMethod object for this reference. |
| PurchaseEx | Int | Optional | ernal use For internal use. |
| ExchangeRate | Decimal | Optional | The number of home currency units it takes to equal one unit of currency specified by CurrencyRef. Applicable if multicurrency is enabled for the company |
| DepartmentRef | ReferenceType | Optional | A reference to a Department object specifying the location of the transaction. Available if Preferences.AccountingInfoPrefs.TrackDepartments is set to true. |
| EntityRef | ReferenceType | Optional | Specifies the party with whom an expense is associated. Can be Customer, Vendor, or Employee. |
| IncludeInAnnualTPAR |  | Optional |  |
| TotalAmt | BigDecimal | Read-only | Indicates the total amount of the transaction. This includes the total of all the charges, allowances, and taxes. |
| RecurDataRef | ReferenceType | Read-only | A reference to the Recurring Transaction. It captures what recurring transaction template the Purchase was created from. |

## Sample Object

```json
{
  "Purchase": {
    "SyncToken": "0",
    "domain": "QBO",
    "PurchaseEx": {
      "any": [
        {
          "name": "{http://schema .intuit.com/finance /v3}NameValue",
          "nil": false,
          "value": {
            "Name": "TxnType",
            "Value": "54"
          },
          "declaredType": "com.intuit .schema.finance.v3 .NameValue",
          "scope": "javax.xml.bind .JAXBElement$GlobalScope",
          "globalScope": true,
          "typeSubstituted": false
        }
      ]
    },
    "TxnDate": "2015-07-27",
    "TotalAmt": 10.0,
    "PaymentType": "Cash",
    "sparse": false,
    "Line": [
      {
        "DetailType": "AccountBasedExpenseLineDetai l",
        "Amount": 10.0,
        "ProjectRef": {
          "value": "39298034"
        },
        "Id": "1",
        "AccountBasedExpenseLineDetail": {
          "TaxCodeRef": {
            "value": "NON"
          },
          "AccountRef": {
            "name": "Meals and Entertainment",
            "value": "13"
          },
          "BillableStatus": "NotBillable"
        }
      }
    ],
    "AccountRef": {
      "name": "Checking",
      "value": "35"
    },
    "CustomField": [],
    "Id": "252",
    "MetaData": {
      "CreateTime": "2015-07-27T10:37:26-07:00",
      "LastUpdatedTime": "2015-07-27T10:37:26-07:00"
    }
  },
  "time": "2015-07-27T10:39:33.171-07:00"
}
```

## Operations

### Create a purchase

- **Method**: POST
- **URL**: `/v3/company/{realmID}/purchase`

You must specify an AccountRef for all purchases.

**Request Body**:
```json
{
  "PaymentType": "CreditCard",
  "AccountRef": {
    "name": "Visa",
    "value": "42"
  },
  "Line": [
    {
      "DetailType": "AccountBasedExpenseLineDetail",
      "Amount": 10.0,
      "AccountBasedExpenseLineDetail": {
        "AccountRef": {
          "name": "Meals and Entertainment",
          "value": "13"
        },
        "ProjectRef": {
          "value": "42991284"
        }
      }
    }
  ]
}
```

**Response**:
```json
{
  "Purchase": {
    "SyncToken": "0",
    "domain": "QBO",
    "PurchaseEx": {
      "any": [
        {
          "name": "{http://schema .intuit.com/finance /v3}NameValue",
          "nil": false,
          "value": {
            "Name": "TxnType",
            "Value": "54"
          },
          "declaredType": "com.intuit .schema.finance.v3 .NameValue",
          "scope": "javax.xml.bind .JAXBElement$GlobalScope",
          "globalScope": true,
          "typeSubstituted": false
        }
      ]
    },
    "Credit": false,
    "TotalAmt": 10.0,
    "PaymentType": "CreditCard",
    "TxnDate": "2015-07-27",
    "sparse": false,
    "Line": [
      {
        "DetailType": "AccountBasedExpenseLineDetai l",
        "Amount": 10.0,
        "ProjectRef": {
          "value": "42991284"
        },
        "Id": "1",
        "AccountBasedExpenseLineDetail": {
          "TaxCodeRef": {
            "value": "NON"
          },
          "AccountRef": {
            "name": "Meals and Entertainment",
            "value": "13"
          },
          "BillableStatus": "NotBillable"
        }
      }
    ],
    "AccountRef": {
      "name": "Visa",
      "value": "42"
    },
    "CustomField": [],
    "Id": "247",
    "MetaData": {
      "CreateTime": "2015-07-27T10:27:01-07:00",
      "LastUpdatedTime": "2015-07-27T10:27:01-07:00"
    }
  },
  "time": "2015-07-27T10:27:01.593-07:00"
}
```

### Delete a purchase

- **Method**: POST
- **URL**: `/v3/company/{realmID}/purchase?operation=delete`

This operation deletes the Purchase object specified in the request body. Include a minimum of Purchase.Id and Purchase.SyncToken in the request body.

**Request Body**:
```json
{
  "SyncToken": "2",
  "Id": "595"
}
```

**Response**:
```json
{
  "Purchase": {
    "status": "Deleted",
    "domain": "QBO",
    "Id": "595"
  },
  "time": "2014-04-22T12:00:52.298-07:00"
}
```

### Query a purchase

- **Method**: GET
- **URL**: `/v3/company/{realmID}/query?query={selectStatement}`

**Sample Query**: `select * from Purchase where TotalAmt < '100.00'`

**Response**:
```json
{
  "QueryResponse": {
    "startPosition": 1,
    "Purchase": [
      {
        "SyncToken": "0",
        "domain": "QBO",
        "PurchaseEx": {
          "any": [
            {
              "name": "{http://schema .intuit.com/finance /v3}NameValue",
              "nil": false,
              "value": {
                "Name": "TxnType",
                "Value": "11"
              },
              "declaredType": "com .intuit.schema.finance .v3.NameValue",
              "scope": "javax.xml.bind .JAXBElement$GlobalSco pe",
              "globalScope": true,
              "typeSubstituted": false
            }
          ]
        },
        "Credit": true,
        "TotalAmt": 900.0,
        "PrivateNote": "Monthly Payment",
        "PaymentType": "CreditCard",
        "TxnDate": "2014-10-03",
        "sparse": false,
        "Line": [
          {
            "DetailType": "AccountBasedExpenseLine Detail",
            "Amount": 900.0,
            "ProjectRef": {
              "value": "39298034"
            },
            "Id": "1",
            "AccountBasedExpenseLine Detail": {
              "TaxCodeRef": {
                "value": "NON"
              },
              "AccountRef": {
                "name": "Checking",
                "value": "35"
              },
              "BillableStatus": "NotBillable"
            }
          }
        ],
        "AccountRef": {
          "name": "Mastercard",
          "value": "41"
        },
        "Id": "139",
        "MetaData": {
          "CreateTime": "2014-10-03T14:35:37-07:00",
          "LastUpdatedTime": "2014-10-03T14:35:37-07:00"
        }
      },
      {
        "DocNumber": "70",
        "SyncToken": "0",
        "domain": "QBO",
        "PurchaseEx": {
          "any": [
            {
              "name": "{http://schema .intuit.com/finance /v3}NameValue",
              "nil": false,
              "value": {
                "Name": "TxnType",
                "Value": "3"
              },
              "declaredType": "com .intuit.schema.finance .v3.NameValue",
              "scope": "javax.xml.bind .JAXBElement$GlobalSco pe",
              "globalScope": true,
              "typeSubstituted": false
            }
          ]
        },
        "TxnDate": "2014-09-11",
        "TotalAmt": 185.0,
        "PrintStatus": "NotSet",
        "PaymentType": "Check",
        "EntityRef": {
          "type": "Vendor",
          "name": "Chin's Gas and Oil",
          "value": "33"
        },
        "sparse": false,
        "Line": [
          {
            "DetailType": "AccountBasedExpenseLine Detail",
            "Amount": 185.0,
            "Id": "1",
            "AccountBasedExpenseLine Detail": {
              "TaxCodeRef": {
                "value": "NON"
              },
              "AccountRef": {
                "name": "Maintenance and Repair",
                "value": "72"
              },
              "BillableStatus": "NotBillable"
            }
          }
        ],
        "AccountRef": {
          "name": "Checking",
          "value": "35"
        },
        "Id": "133",
        "MetaData": {
          "CreateTime": "2014-10-03T14:17:55-07:00",
          "LastUpdatedTime": "2014-10-03T14:17:55-07:00"
        }
      },
      {
        "DocNumber": "75",
        "SyncToken": "0",
        "domain": "QBO",
        "PurchaseEx": {
          "any": [
            {
              "name": "{http://schema .intuit.com/finance /v3}NameValue",
              "nil": false,
              "value": {
                "Name": "TxnType",
                "Value": "3"
              },
              "declaredType": "com .intuit.schema.finance .v3.NameValue",
              "scope": "javax.xml.bind .JAXBElement$GlobalSco pe",
              "globalScope": true,
              "typeSubstituted": false
            }
          ]
        },
        "TxnDate": "2014-09-19",
        "TotalAmt": 228.75,
        "Id": "115",
        "PrintStatus": "NotSet",
        "PaymentType": "Check",
        "EntityRef": {
          "type": "Vendor",
          "name": "Hicks Hardware",
          "value": "41"
        },
        "sparse": false,
        "Line": [
          {
            "DetailType": "ItemBasedExpenseLineDet ail",
            "Amount": 125.0,
            "Id": "1",
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
          },
          {
            "DetailType": "ItemBasedExpenseLineDet ail",
            "Amount": 11.25,
            "Id": "2",
            "ItemBasedExpenseLineDet ail": {
              "TaxCodeRef": {
                "value": "NON"
              },
              "Qty": 15,
              "BillableStatus": "NotBillable",
              "UnitPrice": 0.75,
              "ItemRef": {
                "name": "Sprinkler Heads",
                "value": "16"
              }
            },
            "Description": "Sprinkler Heads"
          },
          {
            "DetailType": "ItemBasedExpenseLineDet ail",
            "Amount": 62.5,
            "Id": "3",
            "ItemBasedExpenseLineDet ail": {
              "TaxCodeRef": {
                "value": "NON"
              },
              "Qty": 25,
              "BillableStatus": "NotBillable",
              "UnitPrice": 2.5,
              "ItemRef": {
                "name": "Sprinkler Pipes",
                "value": "17"
              }
            },
            "Description": "Sprinkler Pipes"
          },
          {
            "DetailType": "ItemBasedExpenseLineDet ail",
            "Amount": 30.0,
            "Id": "4",
            "ItemBasedExpenseLineDet ail": {
              "TaxCodeRef": {
                "value": "NON"
              },
              "Qty": 3,
              "BillableStatus": "NotBillable",
              "UnitPrice": 10,
              "ItemRef": {
                "name": "Pump",
                "value": "11"
              }
            },
            "Description": "Fountain Pump"
          }
        ],
        "AccountRef": {
          "name": "Checking",
          "value": "35"
        },
        "RemitToAddr": {
          "City": "Middlefield",
          "Line1": "42 Main St.",
          "PostalCode": "94303",
          "Lat": "37.445013",
          "Long": "-122.1391443",
          "CountrySubDivisionCode": "CA",
          "Id": "37"
        },
        "MetaData": {
          "CreateTime": "2014-09-19T12:51:46-07:00",
          "LastUpdatedTime": "2014-09-19T12:51:46-07:00"
        }
      },
      {
        "DocNumber": "12",
        "SyncToken": "0",
        "domain": "QBO",
        "PurchaseEx": {
          "any": [
            {
              "name": "{http://schema .intuit.com/finance /v3}NameValue",
              "nil": false,
              "value": {
                "Name": "TxnType",
                "Value": "54"
              },
              "declaredType": "com .intuit.schema.finance .v3.NameValue",
              "scope": "javax.xml.bind .JAXBElement$GlobalSco pe",
              "globalScope": true,
              "typeSubstituted": false
            }
          ]
        },
        "TxnDate": "2014-07-09",
        "TotalAmt": 250.0,
        "PaymentType": "Cash",
        "EntityRef": {
          "type": "Vendor",
          "name": "Robertson & Associates",
          "value": "49"
        },
        "sparse": false,
        "Line": [
          {
            "DetailType": "AccountBasedExpenseLine Detail",
            "Amount": 250.0,
            "Id": "1",
            "AccountBasedExpenseLine Detail": {
              "TaxCodeRef": {
                "value": "NON"
              },
              "AccountRef": {
                "name": "Legal & Professional Fees :Accounting",
                "value": "69"
              },
              "BillableStatus": "NotBillable"
            }
          }
        ],
        "AccountRef": {
          "name": "Checking",
          "value": "35"
        },
        "Id": "107",
        "MetaData": {
          "CreateTime": "2014-09-19T12:36:23-07:00",
          "LastUpdatedTime": "2014-09-19T12:36:23-07:00"
        }
      },
      {
        "DocNumber": "15",
        "SyncToken": "1",
        "domain": "QBO",
        "PurchaseEx": {
          "any": [
            {
              "name": "{http://schema .intuit.com/finance /v3}NameValue",
              "nil": false,
              "value": {
                "Name": "TxnType",
                "Value": "54"
              },
              "declaredType": "com .intuit.schema.finance .v3.NameValue",
              "scope": "javax.xml.bind .JAXBElement$GlobalSco pe",
              "globalScope": true,
              "typeSubstituted": false
            }
          ]
        },
        "TxnDate": "2014-08-16",
        "TotalAmt": 108.09,
        "PaymentType": "Cash",
        "EntityRef": {
          "type": "Vendor",
          "name": "Tania's Nursery",
          "value": "50"
        },
        "sparse": false,
        "Line": [
          {
            "DetailType": "AccountBasedExpenseLine Detail",
            "Amount": 108.09,
            "Id": "1",
            "AccountBasedExpenseLine Detail": {
              "TaxCodeRef": {
                "value": "NON"
              },
              "AccountRef": {
                "name": "Job Expenses",
                "value": "58"
              },
              "BillableStatus": "NotBillable"
            }
          }
        ],
        "AccountRef": {
          "name": "Checking",
          "value": "35"
        },
        "Id": "87",
        "MetaData": {
          "CreateTime": "2014-09-18T13:14:42-07:00",
          "LastUpdatedTime": "2014-09-18T13:17:06-07:00"
        }
      },
      {
        "DocNumber": "3",
        "SyncToken": "0",
        "domain": "QBO",
        "PurchaseEx": {
          "any": [
            {
              "name": "{http://schema .intuit.com/finance /v3}NameValue",
              "nil": false,
              "value": {
                "Name": "TxnType",
                "Value": "54"
              },
              "declaredType": "com .intuit.schema.finance .v3.NameValue",
              "scope": "javax.xml.bind .JAXBElement$GlobalSco pe",
              "globalScope": true,
              "typeSubstituted": false
            }
          ]
        },
        "Credit": false,
        "TotalAmt": 158.08,
        "PaymentType": "CreditCard",
        "TxnDate": "2014-07-16",
        "EntityRef": {
          "type": "Vendor",
          "name": "Tania's Nursery",
          "value": "50"
        },
        "sparse": false,
        "Line": [
          {
            "DetailType": "AccountBasedExpenseLine Detail",
            "Amount": 158.08,
            "Id": "1",
            "AccountBasedExpenseLine Detail": {
              "TaxCodeRef": {
                "value": "NON"
              },
              "AccountRef": {
                "name": "Job Expenses :Job Materials :Plants and Soil",
                "value": "66"
              },
              "BillableStatus": "NotBillable"
            }
          }
        ],
        "AccountRef": {
          "name": "Mastercard",
          "value": "41"
        },
        "Id": "85",
        "MetaData": {
          "CreateTime": "2014-09-18T13:12:01-07:00",
          "LastUpdatedTime": "2014-09-18T13:12:01-07:00"
        }
      },
      {
        "DocNumber": "13",
        "SyncToken": "0",
        "domain": "QBO",
        "PurchaseEx": {
          "any": [
            {
              "name": "{http://schema .intuit.com/finance /v3}NameValue",
              "nil": false,
              "value": {
                "Name": "TxnType",
                "Value": "54"
              },
              "declaredType": "com .intuit.schema.finance .v3.NameValue",
              "scope": "javax.xml.bind .JAXBElement$GlobalSco pe",
              "globalScope": true,
              "typeSubstituted": false
            }
          ]
        },
        "TxnDate": "2014-09-13",
        "TotalAmt": 215.66,
        "PaymentType": "Cash",
        "EntityRef": {
          "type": "Vendor",
          "name": "Hicks Hardware",
          "value": "41"
        },
        "sparse": false,
        "Line": [
          {
            "DetailType": "AccountBasedExpenseLine Detail",
            "Amount": 215.66,
            "Id": "1",
            "AccountBasedExpenseLine Detail": {
              "TaxCodeRef": {
                "value": "NON"
              },
              "AccountRef": {
                "name": "Job Expenses :Job Materials :Sprinklers and Drip Systems",
                "value": "67"
              },
              "BillableStatus": "NotBillable"
            }
          }
        ],
        "AccountRef": {
          "name": "Checking",
          "value": "35"
        },
        "Id": "83",
        "MetaData": {
          "CreateTime": "2014-09-18T13:08:20-07:00",
          "LastUpdatedTime": "2014-09-18T13:08:20-07:00"
        }
      },
      {
        "DocNumber": "1",
        "SyncToken": "0",
        "domain": "QBO",
        "PurchaseEx": {
          "any": [
            {
              "name": "{http://schema .intuit.com/finance /v3}NameValue",
              "nil": false,
              "value": {
                "Name": "TxnType",
                "Value": "54"
              },
              "declaredType": "com .intuit.schema.finance .v3.NameValue",
              "scope": "javax.xml.bind .JAXBElement$GlobalSco pe",
              "globalScope": true,
              "typeSubstituted": false
            }
          ]
        },
        "Credit": false,
        "TotalAmt": 112.0,
        "PaymentType": "CreditCard",
        "TxnDate": "2014-09-17",
        "EntityRef": {
          "type": "Vendor",
          "name": "Ellis Equipment Rental",
          "value": "38"
        },
        "sparse": false,
        "Line": [
          {
            "DetailType": "AccountBasedExpenseLine Detail",
            "Amount": 112.0,
            "Id": "1",
            "AccountBasedExpenseLine Detail": {
              "TaxCodeRef": {
                "value": "NON"
              },
              "AccountRef": {
                "name": "Equipment Rental",
                "value": "29"
              },
              "BillableStatus": "NotBillable"
            },
            "Description": "Equipment rental for 5 days"
          }
        ],
        "AccountRef": {
          "name": "Mastercard",
          "value": "41"
        },
        "Id": "51",
        "MetaData": {
          "CreateTime": "2014-09-17T11:45:45-07:00",
          "LastUpdatedTime": "2014-09-17T11:45:45-07:00"
        }
      }
    ],
    "maxResults": 8
  },
  "time": "2015-07-27T09:09:11.269-07:00"
}
```

### Read a purchase

- **Method**: GET
- **URL**: `/v3/company/{realmID}/purchase/{purchaseId}`

Retrieves the details of a Purchase object that has been previously created.

**Response**:
```json
{
  "Purchase": {
    "SyncToken": "0",
    "domain": "QBO",
    "PurchaseEx": {
      "any": [
        {
          "name": "{http://schema .intuit.com/finance /v3}NameValue",
          "nil": false,
          "value": {
            "Name": "TxnType",
            "Value": "54"
          },
          "declaredType": "com.intuit .schema.finance.v3 .NameValue",
          "scope": "javax.xml.bind .JAXBElement$GlobalScope",
          "globalScope": true,
          "typeSubstituted": false
        }
      ]
    },
    "TxnDate": "2015-07-27",
    "TotalAmt": 10.0,
    "PaymentType": "Cash",
    "sparse": false,
    "Line": [
      {
        "DetailType": "AccountBasedExpenseLineDetai l",
        "Amount": 10.0,
        "ProjectRef": {
          "value": "39298034"
        },
        "Id": "1",
        "AccountBasedExpenseLineDetail": {
          "TaxCodeRef": {
            "value": "NON"
          },
          "AccountRef": {
            "name": "Meals and Entertainment",
            "value": "13"
          },
          "BillableStatus": "NotBillable"
        }
      }
    ],
    "AccountRef": {
      "name": "Checking",
      "value": "35"
    },
    "CustomField": [],
    "Id": "252",
    "MetaData": {
      "CreateTime": "2015-07-27T10:37:26-07:00",
      "LastUpdatedTime": "2015-07-27T10:37:26-07:00"
    }
  },
  "time": "2015-07-27T10:39:33.171-07:00"
}
```

### Full update a purchase

- **Method**: POST
- **URL**: `/v3/company/{realmID}/purchase`

Use this operation to update any of the writable fields of an existing purchase object. The request body must include all writable fields of the existing object as returned in a read response. Writable fields omitted from the request body are set to NULL. The ID of the object to update is specified ...

**Request Body**:
```json
{
  "SyncToken": "1",
  "domain": "QBO",
  "PurchaseEx": {
    "any": [
      {
        "name": "{http://schema.intuit.com/finance/v3}NameValue",
        "nil": false,
        "value": {
          "Name": "TxnType",
          "Value": "54"
        },
        "declaredType": "com.intuit .schema.finance.v3.NameValue",
        "scope": "javax.xml.bind .JAXBElement$GlobalScope",
        "globalScope": true,
        "typeSubstituted": false
      }
    ]
  },
  "TxnDate": "2015-07-27",
  "TotalAmt": 10.0,
  "PrivateNote": "Added an updated private note via update.",
  "PaymentType": "Cash",
  "sparse": false,
  "Line": [
    {
      "DetailType": "AccountBasedExpenseLineDetail",
      "Amount": 10.0,
      "ProjectRef": {
        "value": "42991284"
      },
      "Id": "1",
      "AccountBasedExpenseLineDetail": {
        "TaxCodeRef": {
          "value": "NON"
        },
        "AccountRef": {
          "name": "Meals and Entertainment",
          "value": "13"
        },
        "BillableStatus": "NotBillable"
      }
    }
  ],
  "AccountRef": {
    "name": "Checking",
    "value": "35"
  },
  "CustomField": [],
  "Id": "252",
  "MetaData": {
    "CreateTime": "2015-07-27T10:37:26-07:00",
    "LastUpdatedTime": "2015-07-27T10:42:11-07:00"
  }
}
```

**Response**:
```json
{
  "Purchase": {
    "SyncToken": "2",
    "domain": "QBO",
    "PurchaseEx": {
      "any": [
        {
          "name": "{http://schema .intuit.com/finance /v3}NameValue",
          "nil": false,
          "value": {
            "Name": "TxnType",
            "Value": "54"
          },
          "declaredType": "com.intuit .schema.finance.v3 .NameValue",
          "scope": "javax.xml.bind .JAXBElement$GlobalScope",
          "globalScope": true,
          "typeSubstituted": false
        }
      ]
    },
    "TxnDate": "2015-07-27",
    "TotalAmt": 10.0,
    "PrivateNote": "Added an updated private note via update.",
    "PaymentType": "Cash",
    "sparse": false,
    "Line": [
      {
        "DetailType": "AccountBasedExpenseLineDetai l",
        "Amount": 10.0,
        "ProjectRef": {
          "value": "42991284"
        },
        "Id": "1",
        "AccountBasedExpenseLineDetail": {
          "TaxCodeRef": {
            "value": "NON"
          },
          "AccountRef": {
            "name": "Meals and Entertainment",
            "value": "13"
          },
          "BillableStatus": "NotBillable"
        }
      }
    ],
    "AccountRef": {
      "name": "Checking",
      "value": "35"
    },
    "CustomField": [],
    "Id": "252",
    "MetaData": {
      "CreateTime": "2015-07-27T10:37:26-07:00",
      "LastUpdatedTime": "2015-07-27T10:45:20-07:00"
    }
  },
  "time": "2015-07-27T10:45:20.806-07:00"
}
```
