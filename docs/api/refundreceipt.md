# RefundReceipt

A RefundReceipt object represents a refund to the customer for a product or service that was provided.

## Attributes

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| Id | String | Required for update | Unique identifier for this object. Sort order is ASC by default. |
| DepositToAccountRef | ReferenceType | Required | Account from which payment money is refunded. Query the Account name list resource to determine the appropriate Account object for this reference, where Account. |
| Line | Line | Required | Individual line items of a transaction. Valid Line types include: SalesItemLine, GroupLine, DescriptionOnlyLine, DiscountLine and SubTotalLine(read-only) SalesItemLine GroupLine DescriptionOnlyLine... |
| SyncToken | String | Required for update | Version number of the object. It is used to lock an object for use by one app at a time. As soon as an application modifies an object, its SyncToken is incremented. |
| CurrencyRef | CurrencyRefType | Conditionally required | Reference to the currency in which all amounts on the associated transaction are expressed. This must be defined if multicurrency is enabled for the company. |
| PaymentRefNum |  | Conditionally required | 100 chars String The reference number for the payment received. For example, check # for a check, envelope # for a cash donation. |
| GlobalTaxCalculation | String | Conditionally required | Method in which tax is applied. Allowed values are: TaxExcluded, TaxInclusive, and NotApplicable. |
| ProjectRef | ReferenceType | Conditionally required | Reference to the Project ID associated with this transaction. Available with Minor Version 69 and above |
| BillEmail | EmailAddress | Conditionally required | Identifies the e-mail address where the invoice is sent. If EmailStatus=NeedToSend, BillEmailis a required input. |
| TxnDate | Date | Optional | The date entered by the user when this transaction occurred. For posting transactions, this is the posting date that affects the financial statements. |
| CustomField | CustomField | Optional | One of, up to three custom fields for the transaction. Available for custom fields so configured for the company. Check Preferences.SalesFormsPrefs.CustomField and Preferences. |
| ClassRef | ReferenceType | Optional | Reference to the Class associated with the transaction. Available if Preferences.AccountingInfoPrefs.ClassTrackingPerTxn is set to true. |
| PrintStatus | String | Optional | Printing status of the invoice. Valid values: NotSet, NeedToPrint, PrintComplete . |
| CheckPayment |  | Optional | 21 characters CheckPayment   Information about a check payment for the transaction. Used when PaymentType is Check. |
| TxnSource | String | Optional | The originating source of the credit card transaction. Used in eCommerce apps where credit card transactions are processed by a merchant account. |
| TransactionLocationType | String | Optional | The account location. Valid values for France: WithinFrance, FranceOverseas, OutsideFranceWithEU, OutsideEU. |
| MetaData | ModificationMetaData | Read-only | Descriptive information about the entity. The MetaData values are set by Data Services and are  for all applications. |
| DocNumber | String | Optional | Reference number for the transaction. |
| Recommended |  | Optional | best practice: check the setting of Preferences:CustomTxnNumber before setting DocNumber. |
| PrivateNote | String | Optional | User entered, organization-private note about the transaction. This note does not appear on the refund receipt to the customer. |
| CustomerMemo | MemoRef | Optional | User-entered message to the customer; this message is visible to end user on their transaction. |
| CreditCardPayment |  | Optional | 21 characters CreditCardPayment   Information about a credit card payment for the transaction. Used when PaymentType is CreditCard. |
| CustomerRef | ReferenceType | Optional | Reference to a customer or job. Query the Customer name list resource to determine the appropriate Customer object for this reference. Use Customer.Id and Customer. |
| TxnTaxDetail | TxnTaxDetail | Optional | This data type provides information for taxes charged on the transaction as a whole. |
| PaymentMethodRef | ReferenceType | Optional | Reference to a PaymentMethod associated with this transaction. Query the PaymentMethod name list resource to determine the appropriate PaymentMethod object for this reference. |
| ExchangeRate | Decimal | Optional | The number of home currency units it takes to equal one unit of currency specified by CurrencyRef. Applicable if multicurrency is enabled for the company. |
| ShipAddr | PhysicalAddress | Optional | Identifies the address where the goods must be shipped. |
| DepartmentRef | ReferenceType | Optional | A reference to a Department object specifying the location of the transaction. Available if Preferences.AccountingInfoPrefs.TrackDepartments is set to true. |
| PaymentType |  | Optional | 21 characters PaymentTypeEnum   Valid values are Cash, Check, CreditCard, or Other. |
| BillAddr | PhysicalAddress | Optional | Bill-to address of the Invoice. If BillAddris not specified, and a default Customer:BillingAddr is specified in QuickBooks for this customer, the default bill-to address is us... |
| ApplyTaxAfterDiscount | Boolean | Optional | If false or null, calculate the sales tax first, and then apply the discount. If true, subtract the discount first and then calculate the sales tax. |
| HomeBalance | Decimal | Read-only | Convenience field containing the amount in Balance expressed in terms of the home currency. Calculated by QuickBooks business logic. |
| RecurDataRef | ReferenceType | Read-only | A reference to the Recurring Transaction. It captures what recurring transaction template the RefundReceipt was created from. |
| TotalAmt | BigDecimal | Read-only | Indicates the total amount of the transaction. This includes the total of all the charges, allowances, and taxes. |
| TaxExemptionRef | ReferenceType | Read-only | Reference to the TaxExepmtion ID associated with this object. Available for companies that have |
| Balance | Decimal | Read-only | The balance reflecting any payments made against the transaction. Initially set to the value of TotalAmt. Calculated by QuickBooks business logic; any value you supply is over-written by QuickBooks. |
| HomeTotalAmt | Decimal | Read-only | Total amount of the transaction in the home currency. Includes the total of all the charges, allowances and taxes. Calculated by QuickBooks business logic. |

## Sample Object

```json
{
  "RefundReceipt": {
    "DocNumber": "1020",
    "SyncToken": "0",
    "domain": "QBO",
    "Balance": 0,
    "PaymentMethodRef": {
      "name": "Check",
      "value": "2"
    },
    "BillAddr": {
      "Line4": "South Orange, NJ 07079",
      "Line3": "350 Mountain View Dr.",
      "Line2": "Pye's Cakes",
      "Line1": "Karen Pye",
      "Long": "-74.2609903",
      "Lat": "40.7489277",
      "Id": "73"
    },
    "DepositToAccountRef": {
      "name": "Checking",
      "value": "35"
    },
    "TxnDate": "2014-09-17",
    "TotalAmt": 87.5,
    "CustomerRef": {
      "name": "Pye's Cakes",
      "value": "15"
    },
    "CustomerMemo": {
      "value": "Thank you for your business and have a great day!"
    },
    "ShipAddr": {
      "City": "South Orange",
      "Line1": "350 Mountain View Dr.",
      "PostalCode": "07079",
      "Lat": "40.7633073",
      "Long": "-74.2426072",
      "CountrySubDivisionCode": "NJ",
      "Id": "15"
    },
    "PrintStatus": "NotSet",
    "BillEmail": {
      "Address": "pyescakes@intuit.com"
    },
    "sparse": false,
    "Line": [
      {
        "Description": "Refund- Pest control was ineffective",
        "DetailType": "SalesItemLineDetail",
        "SalesItemLineDetail": {
          "TaxCodeRef": {
            "value": "NON"
          },
          "Qty": 2.5,
          "UnitPrice": 35,
          "ItemRef": {
            "name": "Pest Control",
            "value": "10"
          }
        },
        "LineNum": 1,
        "Amount": 87.5,
        "Id": "1"
      },
      {
        "DetailType": "SubTotalLineDetail",
        "Amount": 87.5,
        "SubTotalLineDetail": {}
      }
    ],
    "ApplyTaxAfterDiscount": false,
    "CustomField": [
      {
        "DefinitionId": "1",
        "Type": "StringType",
        "Name": "Crew #"
      }
    ],
    "Id": "66",
    "TxnTaxDetail": {
      "TotalTax": 0
    },
    "MetaData": {
      "CreateTime": "2014-09-17T15:35:07-07:00",
      "LastUpdatedTime": "2014-09-17T15:35:07-07:00"
    }
  },
  "time": "2015-07-29T08:15:49.421-07:00"
}
```

## Operations

### Create a refund receipt

- **Method**: POST
- **URL**: `/v3/company/{realmID}/refundreceipt`

• A RefundReceipt object must have at least one line that describes an item. • A RefundReceipt object must have a DepositToAccountRef.

**Request Body**:
```json
{
  "Line": [
    {
      "DetailType": "SalesItemLineDetail",
      "Amount": 420.0,
      "SalesItemLineDetail": {
        "ItemRef": {
          "value": "38"
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
  "RefundReceipt": {
    "DocNumber": "1072",
    "SyncToken": "0",
    "domain": "QBO",
    "Balance": 0,
    "DepositToAccountRef": {
      "name": "Checking",
      "value": "35"
    },
    "TxnDate": "2015-07-29",
    "TotalAmt": 420.0,
    "PrintStatus": "NeedToPrint",
    "PaymentRefNum": "To Print",
    "sparse": false,
    "Line": [
      {
        "LineNum": 1,
        "Amount": 420.0,
        "SalesItemLineDetail": {
          "TaxCodeRef": {
            "value": "NON"
          },
          "ItemRef": {
            "name": "Garden Supplies",
            "value": "38"
          }
        },
        "Id": "1",
        "DetailType": "SalesItemLineDetail"
      },
      {
        "DetailType": "SubTotalLineDetail",
        "Amount": 420.0,
        "SubTotalLineDetail": {}
      }
    ],
    "ApplyTaxAfterDiscount": false,
    "CustomField": [
      {
        "DefinitionId": "1",
        "Type": "StringType",
        "Name": "Crew #"
      }
    ],
    "Id": "261",
    "TxnTaxDetail": {
      "TotalTax": 0
    },
    "MetaData": {
      "CreateTime": "2015-07-29T08:07:43-07:00",
      "LastUpdatedTime": "2015-07-29T08:07:43-07:00"
    }
  },
  "time": "2015-07-29T08:07:43.749-07:00"
}
```

### Delete a refund receipt

- **Method**: POST
- **URL**: `/v3/company/{realmID}/refundreceipt?operation=delete`

This operation deletes the RefundReceipt object specified in the request body. Include a minimum of RefundReceipt.Id and RefunReceipt.SyncToken in the request body. You must unlink any linked transactions associated with the RefundReceipt object before deleting it.

**Request Body**:
```json
{
  "SyncToken": "0",
  "Id": "66"
}
```

**Response**:
```json
{
  "RefundReceipt": {
    "status": "Deleted",
    "domain": "QBO",
    "Id": "66"
  },
  "time": "2015-05-27T10:34:32.031-07:00"
}
```

### Get a refund receipt as PDF

- **Method**: GET
- **URL**: `/v3/company/{realmID}/refundreceipt/{refundreceiptId}/pdf`

### Query a refund receipt

- **Method**: GET
- **URL**: `/v3/company/{realmID}/query?query={selectStatement}`

**Sample Query**: `select * from RefundReceipt where Id='66'`

**Response**:
```json
{
  "QueryResponse": {
    "startPosition": 1,
    "totalCount": 1,
    "RefundReceipt": [
      {
        "DocNumber": "1020",
        "SyncToken": "0",
        "domain": "QBO",
        "Balance": 0,
        "PaymentMethodRef": {
          "name": "Check",
          "value": "2"
        },
        "BillAddr": {
          "Line4": "South Orange, NJ 07079",
          "Line3": "350 Mountain View Dr.",
          "Line2": "Pye's Cakes",
          "Line1": "Karen Pye",
          "Long": "-74.2609903",
          "Lat": "40.7489277",
          "Id": "73"
        },
        "DepositToAccountRef": {
          "name": "Checking",
          "value": "35"
        },
        "TxnDate": "2014-09-17",
        "TotalAmt": 87.5,
        "CustomerRef": {
          "name": "Pye's Cakes",
          "value": "15"
        },
        "CustomerMemo": {
          "value": "Thank you for your business and have a great day!"
        },
        "PrintStatus": "NotSet",
        "BillEmail": {
          "Address": "pyescakes@intuit.com"
        },
        "sparse": false,
        "Line": [
          {
            "Description": "Refund- Pest control was ineffective",
            "DetailType": "SalesItemLineDetail",
            "SalesItemLineDetail": {
              "TaxCodeRef": {
                "value": "NON"
              },
              "Qty": 2.5,
              "UnitPrice": 35,
              "ItemRef": {
                "name": "Pest Control",
                "value": "10"
              }
            },
            "LineNum": 1,
            "Amount": 87.5,
            "Id": "1"
          },
          {
            "DetailType": "SubTotalLineDetail",
            "Amount": 87.5,
            "SubTotalLineDetail": {}
          }
        ],
        "ApplyTaxAfterDiscount": false,
        "CustomField": [
          {
            "DefinitionId": "1",
            "Type": "StringType",
            "Name": "Crew #"
          }
        ],
        "Id": "66",
        "TxnTaxDetail": {
          "TotalTax": 0
        },
        "MetaData": {
          "CreateTime": "2014-09-17T15:35:07-07:00",
          "LastUpdatedTime": "2014-09-17T15:35:07-07:00"
        }
      }
    ],
    "maxResults": 1
  },
  "time": "2015-07-29T08:14:41.415-07:00"
}
```

### Read a refund receipt

- **Method**: GET
- **URL**: `/v3/company/{realmID}/refundreceipt/{refundreceiptId}`

Retrieves the details of a RefundReceipt that has been previously created.

**Response**:
```json
{
  "RefundReceipt": {
    "DocNumber": "1020",
    "SyncToken": "0",
    "domain": "QBO",
    "Balance": 0,
    "PaymentMethodRef": {
      "name": "Check",
      "value": "2"
    },
    "BillAddr": {
      "Line4": "South Orange, NJ 07079",
      "Line3": "350 Mountain View Dr.",
      "Line2": "Pye's Cakes",
      "Line1": "Karen Pye",
      "Long": "-74.2609903",
      "Lat": "40.7489277",
      "Id": "73"
    },
    "DepositToAccountRef": {
      "name": "Checking",
      "value": "35"
    },
    "TxnDate": "2014-09-17",
    "TotalAmt": 87.5,
    "CustomerRef": {
      "name": "Pye's Cakes",
      "value": "15"
    },
    "CustomerMemo": {
      "value": "Thank you for your business and have a great day!"
    },
    "ShipAddr": {
      "City": "South Orange",
      "Line1": "350 Mountain View Dr.",
      "PostalCode": "07079",
      "Lat": "40.7633073",
      "Long": "-74.2426072",
      "CountrySubDivisionCode": "NJ",
      "Id": "15"
    },
    "PrintStatus": "NotSet",
    "BillEmail": {
      "Address": "pyescakes@intuit.com"
    },
    "sparse": false,
    "Line": [
      {
        "Description": "Refund- Pest control was ineffective",
        "DetailType": "SalesItemLineDetail",
        "SalesItemLineDetail": {
          "TaxCodeRef": {
            "value": "NON"
          },
          "Qty": 2.5,
          "UnitPrice": 35,
          "ItemRef": {
            "name": "Pest Control",
            "value": "10"
          }
        },
        "LineNum": 1,
        "Amount": 87.5,
        "Id": "1"
      },
      {
        "DetailType": "SubTotalLineDetail",
        "Amount": 87.5,
        "SubTotalLineDetail": {}
      }
    ],
    "ApplyTaxAfterDiscount": false,
    "CustomField": [
      {
        "DefinitionId": "1",
        "Type": "StringType",
        "Name": "Crew #"
      }
    ],
    "Id": "66",
    "TxnTaxDetail": {
      "TotalTax": 0
    },
    "MetaData": {
      "CreateTime": "2014-09-17T15:35:07-07:00",
      "LastUpdatedTime": "2014-09-17T15:35:07-07:00"
    }
  },
  "time": "2015-07-29T08:15:49.421-07:00"
}
```

### Send a refund receipt

- **Method**: POST
- **URL**: `/v3/company/{realmID}/refundreceipt/{refundreceiptId}/send`

• The RefundReceipt.DeliveryInfo element is populated with sending information. • The RefundReceipt.BillEmail.Address parameter is updated to the address specified with the value of the sendTo query parameter, if specified.

**Response**:
```json
{
  "RefundReceipt": {
    "TxnDate": "2014-09-17",
    "domain": "QBO",
    "PrintStatus": "NotSet",
    "DeliveryInfo": {
      "DeliveryType": "Email",
      "DeliveryTime": "2019-09-19T10:43:46-07:00"
    },
    "TotalAmt": 87.5,
    "Line": [
      {
        "Description": "Refund- Pest control was ineffective",
        "DetailType": "SalesItemLineDetail",
        "SalesItemLineDetail": {
          "TaxCodeRef": {
            "value": "NON"
          },
          "Qty": 2.5,
          "UnitPrice": 35,
          "ItemRef": {
            "name": "Pest Control",
            "value": "10"
          }
        },
        "LineNum": 1,
        "Amount": 87.5,
        "Id": "1"
      },
      {
        "DetailType": "SubTotalLineDetail",
        "Amount": 87.5,
        "SubTotalLineDetail": {}
      }
    ],
    "ApplyTaxAfterDiscount": false,
    "DocNumber": "1020",
    "sparse": false,
    "DepositToAccountRef": {
      "name": "Checking",
      "value": "35"
    },
    "CustomerMemo": {
      "value": "Thank you for your business and have a great day!"
    },
    "Balance": 0,
    "CustomerRef": {
      "name": "Pye's Cakes",
      "value": "15"
    },
    "TxnTaxDetail": {
      "TotalTax": 0
    },
    "SyncToken": "0",
    "PaymentMethodRef": {
      "name": "Check",
      "value": "2"
    },
    "CustomField": [
      {
        "DefinitionId": "1",
        "Type": "StringType",
        "Name": "Crew #"
      }
    ],
    "ShipAddr": {
      "City": "South Orange",
      "Line1": "350 Mountain View Dr.",
      "PostalCode": "07079",
      "Lat": "40.7633073",
      "Long": "-74.2426072",
      "CountrySubDivisionCode": "NJ",
      "Id": "15"
    },
    "BillAddr": {
      "Line4": "South Orange, NJ 07079",
      "Line3": "350 Mountain View Dr.",
      "Line2": "Pye's Cakes",
      "Line1": "Karen Pye",
      "Long": "-74.2609903",
      "Lat": "40.7489277",
      "Id": "73"
    },
    "MetaData": {
      "CreateTime": "2014-09-17T15:35:07-07:00",
      "LastUpdatedTime": "2019-09-19T10:43:46-07:00"
    },
    "BillEmail": {
      "Address": "pyescakes@intuit.com"
    },
    "Id": "66"
  },
  "time": "2019-09-19T10:43:46-07:00"
}
```

### Sparse update a refund receipt

- **Method**: POST
- **URL**: `/v3/company/{realmID}/refundreceipt`

Sparse updating provides the ability to update a subset of properties for a given object; only elements specified in the request are updated. Missing elements are left untouched. The ID of the object to update is specified in the request body.

**Request Body**:
```json
{
  "SyncToken": "2",
  "PrivateNote": "This is a new private note added via sparse update.",
  "Id": "66",
  "sparse": true
}
```

**Response**:
```json
{
  "RefundReceipt": {
    "TxnDate": "2014-09-17",
    "domain": "QBO",
    "PrintStatus": "NotSet",
    "TotalAmt": 87.5,
    "Line": [
      {
        "Description": "Refund- Pest control was ineffective",
        "DetailType": "SalesItemLineDetail",
        "SalesItemLineDetail": {
          "TaxCodeRef": {
            "value": "NON"
          },
          "Qty": 2.5,
          "UnitPrice": 35,
          "ItemRef": {
            "name": "Pest Control",
            "value": "10"
          }
        },
        "LineNum": 1,
        "Amount": 87.5,
        "Id": "1"
      },
      {
        "DetailType": "SubTotalLineDetail",
        "Amount": 87.5,
        "SubTotalLineDetail": {}
      }
    ],
    "ApplyTaxAfterDiscount": false,
    "DocNumber": "1020",
    "PrivateNote": "This is a new private note added via sparse update.",
    "sparse": false,
    "DepositToAccountRef": {
      "name": "Checking",
      "value": "35"
    },
    "CustomerMemo": {
      "value": "Thank you for your business and have a great day!"
    },
    "ProjectRef": {
      "value": "39298034"
    },
    "Balance": 0,
    "CustomerRef": {
      "name": "Pye's Cakes",
      "value": "15"
    },
    "TxnTaxDetail": {
      "TotalTax": 0
    },
    "SyncToken": "3",
    "PaymentMethodRef": {
      "name": "Check",
      "value": "2"
    },
    "CustomField": [
      {
        "DefinitionId": "1",
        "Type": "StringType",
        "Name": "Crew #"
      }
    ],
    "ShipAddr": {
      "City": "South Orange",
      "Line1": "350 Mountain View Dr.",
      "PostalCode": "07079",
      "Lat": "40.7633073",
      "Long": "-74.2426072",
      "CountrySubDivisionCode": "NJ",
      "Id": "15"
    },
    "BillAddr": {
      "Line4": "South Orange, NJ 07079",
      "Line3": "350 Mountain View Dr.",
      "Id": "73",
      "Line1": "Karen Pye",
      "Line2": "Pye's Cakes"
    },
    "MetaData": {
      "CreateTime": "2014-09-17T15:35:07-07:00",
      "LastUpdatedTime": "2015-07-29T08:59:30-07:00"
    },
    "BillEmail": {
      "Address": "pyescakes@intuit.com"
    },
    "Id": "66"
  },
  "time": "2015-07-29T08:59:32.061-07:00"
}
```

### Full update a refund receipt

- **Method**: POST
- **URL**: `/v3/company/{realmID}/refundreceipt`

Use this operation to update any of the writable fields of an existing RefundReceipt object. The request body must include all writable fields of the existing object as returned in a read response. Writable fields omitted from the request body are set to NULL. The ID of the object to update is speci...

**Request Body**:
```json
{
  "TxnDate": "2014-09-17",
  "domain": "QBO",
  "PrintStatus": "NotSet",
  "TotalAmt": 87.5,
  "Line": [
    {
      "Description": "Refund- Pest control was ineffective",
      "DetailType": "SalesItemLineDetail",
      "SalesItemLineDetail": {
        "TaxCodeRef": {
          "value": "NON"
        },
        "Qty": 2.5,
        "UnitPrice": 35,
        "ItemRef": {
          "name": "Pest Control",
          "value": "10"
        }
      },
      "LineNum": 1,
      "Amount": 87.5,
      "Id": "1"
    },
    {
      "DetailType": "SubTotalLineDetail",
      "Amount": 87.5,
      "SubTotalLineDetail": {}
    }
  ],
  "ApplyTaxAfterDiscount": false,
  "DocNumber": "1020",
  "sparse": false,
  "DepositToAccountRef": {
    "name": "Checking",
    "value": "35"
  },
  "CustomerMemo": {
    "value": "Updated customer memo"
  },
  "ProjectRef": {
    "value": "39298034"
  },
  "Balance": 0,
  "CustomerRef": {
    "name": "Pye's Cakes",
    "value": "15"
  },
  "TxnTaxDetail": {
    "TotalTax": 0
  },
  "SyncToken": "0",
  "PaymentMethodRef": {
    "name": "Check",
    "value": "2"
  },
  "CustomField": [
    {
      "DefinitionId": "1",
      "Type": "StringType",
      "Name": "Crew #"
    }
  ],
  "ShipAddr": {
    "City": "South Orange",
    "Line1": "350 Mountain View Dr.",
    "PostalCode": "07079",
    "Lat": "40.7633073",
    "Long": "-74.2426072",
    "CountrySubDivisionCode": "NJ",
    "Id": "15"
  },
  "BillAddr": {
    "Line4": "South Orange, NJ 07079",
    "Line3": "350 Mountain View Dr.",
    "Line2": "Pye's Cakes",
    "Line1": "Karen Pye",
    "Long": "-74.2609903",
    "Lat": "40.7489277",
    "Id": "73"
  },
  "MetaData": {
    "CreateTime": "2014-09-17T15:35:07-07:00",
    "LastUpdatedTime": "2014-09-17T15:35:07-07:00"
  },
  "BillEmail": {
    "Address": "pyescakes@intuit.com"
  },
  "Id": "66"
}
```

**Response**:
```json
{
  "RefundReceipt": {
    "TxnDate": "2014-09-17",
    "domain": "QBO",
    "PrintStatus": "NotSet",
    "TotalAmt": 87.5,
    "Line": [
      {
        "Description": "Refund- Pest control was ineffective",
        "DetailType": "SalesItemLineDetail",
        "SalesItemLineDetail": {
          "TaxCodeRef": {
            "value": "NON"
          },
          "Qty": 2.5,
          "UnitPrice": 35,
          "ItemRef": {
            "name": "Pest Control",
            "value": "10"
          }
        },
        "LineNum": 1,
        "Amount": 87.5,
        "Id": "1"
      },
      {
        "DetailType": "SubTotalLineDetail",
        "Amount": 87.5,
        "SubTotalLineDetail": {}
      }
    ],
    "ApplyTaxAfterDiscount": false,
    "DocNumber": "1020",
    "sparse": false,
    "DepositToAccountRef": {
      "name": "Checking",
      "value": "35"
    },
    "CustomerMemo": {
      "value": "Updated customer memo"
    },
    "ProjectRef": {
      "value": "39298034"
    },
    "Balance": 0,
    "CustomerRef": {
      "name": "Pye's Cakes",
      "value": "15"
    },
    "TxnTaxDetail": {
      "TotalTax": 0
    },
    "SyncToken": "1",
    "PaymentMethodRef": {
      "name": "Check",
      "value": "2"
    },
    "CustomField": [
      {
        "DefinitionId": "1",
        "Type": "StringType",
        "Name": "Crew #"
      }
    ],
    "ShipAddr": {
      "City": "South Orange",
      "Line1": "350 Mountain View Dr.",
      "PostalCode": "07079",
      "Lat": "40.7633073",
      "Long": "-74.2426072",
      "CountrySubDivisionCode": "NJ",
      "Id": "15"
    },
    "BillAddr": {
      "Line4": "South Orange, NJ 07079",
      "Line3": "350 Mountain View Dr.",
      "Id": "73",
      "Line1": "Karen Pye",
      "Line2": "Pye's Cakes"
    },
    "MetaData": {
      "CreateTime": "2014-09-17T15:35:07-07:00",
      "LastUpdatedTime": "2015-07-29T08:18:49-07:00"
    },
    "BillEmail": {
      "Address": "pyescakes@intuit.com"
    },
    "Id": "66"
  },
  "time": "2015-07-29T08:18:48.873-07:00"
}
```
