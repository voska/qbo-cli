# SalesReceipt

A SalesReceipt object represents the sales receipt that is given to a customer. A sales receipt is similar to an invoice. However, for a sales receipt, payment is received as part of the sale of goods and services. The sales receipt specifies a deposit account where the customer's payment is deposited. If the deposit account is not specified, the Undeposited Account is used.

## Attributes

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| Id | String | Required for update | Unique identifier for this object. Sort order is ASC by default. |
| Line | Line | Required | Individual line items of a transaction. Valid Line types include: SalesItemLine, GroupLine, DescriptionOnlyLine, DiscountLine and SubTotalLine (read-only). |
| CustomerRef | ReferenceType | Required | Reference to a customer or job. Query the Customer name list resource to determine the appropriate Customer object for this reference. Use Customer.Id and Customer. |
| SyncToken | String | Required for update | Version number of the object. It is used to lock an object for use by one app at a time. As soon as an application modifies an object, its SyncToken is incremented. |
| ShipFromAddr | PhysicalAddress | Conditionally required | Identifies the address where the goods are shipped from. For transactions without shipping, it represents the address where the sale took place. If automated sales tax is enabled (Preferences. |
| CurrencyRef | CurrencyRefType | Conditionally required | Reference to the currency in which all amounts on the associated transaction are expressed. This must be defined if multicurrency is enabled for the company. |
| GlobalTaxCalculation | String | Conditionally required | Method in which tax is applied. Allowed values are: TaxExcluded, TaxInclusive, and NotApplicable. |
| ProjectRef | ReferenceType | Conditionally required | Reference to the Project ID associated with this transaction. Available with Minor Version 69 and above |
| BillEmail | EmailAddress | Conditionally required | Identifies the e-mail address where the invoice is sent. Required if EmailStatus=NeedToSend |
| TxnDate | Date | Optional | The date entered by the user when this transaction occurred. For posting transactions, this is the posting date that affects the financial statements. |
| CustomField | CustomField | Optional | One of, up to three custom fields for the transaction. Available for custom fields so configured for the company. Check Preferences.SalesFormsPrefs.CustomField and Preferences. |
| ShipDate | Date | Optional | Location of the transaction, as defined using location tracking in QuickBooks Online. |
| TrackingNum | String | Optional | Shipping provider's tracking number for the delivery of the goods associated with the transaction. |
| ClassRef | ReferenceType | Optional | Reference to the Class associated with the transaction. Available if Preferences.AccountingInfoPrefs.ClassTrackingPerTxn is set to true. |
| PrintStatus | String | Optional | Printing status of the invoice. Valid values: NotSet, NeedToPrint, PrintComplete . |
| PaymentRefNum |  | Optional | 21 characters String The reference number for the payment received. For example, Â Check # for a check, envelope # for a cash donation. |
| TxnSource | String | Optional | Used internally to specify originating source of a credit card transaction. |
| LinkedTxn | LinkedTxn | Optional | Zero or more related transactions to this sales receipt object. The following linked relationships are supported: |
| TransactionLocationType | String | Optional | The account location. Valid values for France: WithinFrance, FranceOverseas, OutsideFranceWithEU, OutsideEU. |
| ApplyTaxAfterDiscount | Boolean | Optional | If false or null, calculate the sales tax first, and then apply the discount. If true, subtract the discount first and then calculate the sales tax. |
| DocNumber | String | Optional | Reference number for the transaction. |
| PrivateNote | String | Optional | User entered, organization-private note about the transaction. This note does not appear on the transaction form to the customer. |
| DepositToAccountRef | ReferenceType | Optional | Account to which payment money is deposited. Query the Account name list resource to determine the appropriate Account object for this reference, where Account. |
| CustomerMemo | MemoRef | Optional | User-entered message to the customer; this message is visible to end user on their transactions. |
| EmailStatus | String | Optional | Email status of the receipt. Valid values: NotSet, NeedToSend, EmailSent . |
| CreditCardPayment |  | Optional | CreditCardPayment Information about a credit card payment for the transaction. Used when PaymentType is CreditCard. |
| TxnTaxDetail | TxnTaxDetail | Optional | This element provides information for taxes charged on the transaction as a whole. |
| PaymentMethodRef | ReferenceType | Optional | Reference to a PaymentMethod associated with this transaction. Query the PaymentMethod name list resource to determine the appropriate PaymentMethod object for this reference. |
| ExchangeRate | Decimal | Optional | The number of home currency units it takes to equal one unit of currency specified by CurrencyRef. Applicable if multicurrency is enabled for the company. |
| ShipAddr | PhysicalAddress | Optional | Identifies the address where the goods must be shipped. |
| DepartmentRef | ReferenceType | Optional | A reference to a Department object specifying the location of the transaction. Available if Preferences.AccountingInfoPrefs.TrackDepartments is set to true. |
| ShipMethodRef | ReferenceType | Optional | Reference to the ShipMethod associated with the transaction. There is no shipping method list. Reference resolves to a string. |
| BillAddr | PhysicalAddress | Optional | Bill-to address of the Invoice. If BillAddris not specified, and a default Customer:BillingAddr is specified in QuickBooks for this customer, the default bill-to address is us... |
| MetaData | ModificationMetaData | Read-only | Descriptive information about the object. The MetaData values are set by Data Services and are  for all applications. |
| HomeBalance | Decimal | Read-only | Convenience field containing the amount in Balance expressed in terms of the home currency. Calculated by QuickBooks business logic. |
| DeliveryInfo | DeliveryInfo | Read-only | Email delivery information. Returned when a request has been made to deliver email with the send operation. |
| RecurDataRef | ReferenceType | Read-only | A reference to the Recurring Transaction. It captures what recurring transaction template the SalesReceipt was created from. |
| TotalAmt | BigDecimal | Read-only | Indicates the total amount of the transaction. This includes the total of all the charges, allowances, and taxes. |
| Balance | Decimal | Read-only | The balance reflecting any payments made against the transaction. Initially set to the value of TotalAmt. A Balance of 0 indicates the invoice is fully paid. |
| HomeTotalAmt | Decimal | Read-only | Total amount of the transaction in the home currency. Includes the total of all the charges, allowances and taxes. Calculated by QuickBooks business logic. |
| FreeFormAddress | Boolean | Optional | Denotes how ShipAddr is stored: formatted or unformatted. The value of this flag is  based on format of shipping address at object create time. |

## Sample Object

```json
{
  "SalesReceipt": {
    "TxnDate": "2014-09-14",
    "domain": "QBO",
    "PrintStatus": "NotSet",
    "PaymentRefNum": "10264",
    "TotalAmt": 337.5,
    "Line": [
      {
        "Description": "Custom Design",
        "DetailType": "SalesItemLineDetail",
        "SalesItemLineDetail": {
          "TaxCodeRef": {
            "value": "NON"
          },
          "Qty": 4.5,
          "UnitPrice": 75,
          "ItemRef": {
            "name": "Design",
            "value": "4"
          }
        },
        "LineNum": 1,
        "Amount": 337.5,
        "Id": "1"
      },
      {
        "DetailType": "SubTotalLineDetail",
        "Amount": 337.5,
        "SubTotalLineDetail": {}
      }
    ],
    "ApplyTaxAfterDiscount": false,
    "DocNumber": "1003",
    "sparse": false,
    "DepositToAccountRef": {
      "name": "Checking",
      "value": "35"
    },
    "CustomerMemo": {
      "value": "Thank you for your business and have a great day!"
    },
    "ProjectRef": {
      "value": "39298243"
    },
    "Balance": 0,
    "CustomerRef": {
      "name": "Dylan Sollfrank",
      "value": "6"
    },
    "TxnTaxDetail": {
      "TotalTax": 0
    },
    "SyncToken": "0",
    "PaymentMethodRef": {
      "name": "Check",
      "value": "2"
    },
    "EmailStatus": "NotSet",
    "BillAddr": {
      "Lat": "INVALID",
      "Long": "INVALID",
      "Id": "49",
      "Line1": "Dylan Sollfrank"
    },
    "MetaData": {
      "CreateTime": "2014-09-16T14:59:48-07:00",
      "LastUpdatedTime": "2014-09-16T14:59:48-07:00"
    },
    "CustomField": [
      {
        "DefinitionId": "1",
        "Type": "StringType",
        "Name": "Crew #"
      }
    ],
    "Id": "11"
  },
  "time": "2015-07-29T09:29:56.229-07:00"
}
```

## Operations

### Create a salesreceipt

- **Method**: POST
- **URL**: `/v3/company/{realmID}/salesreceipt`

• A SalesReceipt object must have at least one line that describes an item and an amount.

**Request Body**:
```json
{
  "Line": [
    {
      "Description": "Pest Control Services",
      "DetailType": "SalesItemLineDetail",
      "SalesItemLineDetail": {
        "TaxCodeRef": {
          "value": "NON"
        },
        "Qty": 1,
        "UnitPrice": 35,
        "ItemRef": {
          "name": "Pest Control",
          "value": "10"
        }
      },
      "LineNum": 1,
      "Amount": 35.0,
      "Id": "1"
    }
  ]
}
```

**Response**:
```json
{
  "SalesReceipt": {
    "DocNumber": "1074",
    "SyncToken": "0",
    "domain": "QBO",
    "Balance": 0,
    "DepositToAccountRef": {
      "name": "Undeposited Funds",
      "value": "4"
    },
    "TxnDate": "2015-07-29",
    "TotalAmt": 35.0,
    "PrintStatus": "NeedToPrint",
    "EmailStatus": "NotSet",
    "sparse": false,
    "Line": [
      {
        "Description": "Pest Control Services",
        "DetailType": "SalesItemLineDetail",
        "SalesItemLineDetail": {
          "TaxCodeRef": {
            "value": "NON"
          },
          "Qty": 1,
          "UnitPrice": 35,
          "ItemRef": {
            "name": "Pest Control",
            "value": "10"
          }
        },
        "LineNum": 1,
        "Amount": 35.0,
        "Id": "1"
      },
      {
        "DetailType": "SubTotalLineDetail",
        "Amount": 35.0,
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
    "Id": "263",
    "TxnTaxDetail": {
      "TotalTax": 0
    },
    "MetaData": {
      "CreateTime": "2015-07-29T09:25:02-07:00",
      "LastUpdatedTime": "2015-07-29T09:25:02-07:00"
    }
  },
  "time": "2015-07-29T09:25:04.214-07:00"
}
```

### Delete a salesreceipt

- **Method**: POST
- **URL**: `/v3/company/{realmID}/salesreceipt?operation=delete`

This operation deletes the SalesReceipt object specified in the request body. Include a minimum of SalesReceipt.Id and SalesReceipt.SyncToken in the request body.

**Request Body**:
```json
{
  "SyncToken": "1",
  "Id": "98"
}
```

**Response**:
```json
{
  "SalesReceipt": {
    "status": "Deleted",
    "domain": "QBO",
    "Id": "98"
  },
  "time": "2013-03-13T13:39:58.505-07:00"
}
```

### Void a salesreceipt

- **Method**: POST
- **URL**: `/v3/company/{realmID}/salesreceipt?operation=update&include=void`

Use a sparse update operation with include=void to void an existing SalesReceipt object; include a minimum of SalesReceipt.Id and SalesReceipt.SyncToken. The transaction remains active but all amounts and quantities are zeroed and the string, Voided, is injected into SalesReceipt.PrivateNote, prepen...

**Request Body**:
```json
{
  "SyncToken": "0",
  "Id": "161",
  "sparse": true
}
```

**Response**:
```json
{
  "SalesReceipt": {
    "TxnDate": "2014-12-31",
    "domain": "QBO",
    "CurrencyRef": {
      "name": "United States Dollar",
      "value": "USD"
    },
    "PrintStatus": "NeedToPrint",
    "TotalAmt": 0,
    "Line": [
      {
        "LineNum": 1,
        "Amount": 0,
        "SalesItemLineDetail": {
          "TaxCodeRef": {
            "value": "NON"
          },
          "Qty": 0,
          "ItemRef": {
            "name": "Services",
            "value": "1"
          }
        },
        "Id": "1",
        "DetailType": "SalesItemLineDetail"
      },
      {
        "DetailType": "SubTotalLineDetail",
        "Amount": 0,
        "SubTotalLineDetail": {}
      }
    ],
    "ApplyTaxAfterDiscount": false,
    "DocNumber": "1038",
    "PrivateNote": "Voided",
    "sparse": false,
    "DepositToAccountRef": {
      "name": "Undeposited Funds",
      "value": "4"
    },
    "CustomerMemo": {
      "value": "Thank you for your business and have a great day!"
    },
    "ProjectRef": {
      "value": "39298243"
    },
    "Balance": 0,
    "CustomerRef": {
      "name": "Amy's Bird Sanctuary",
      "value": "1"
    },
    "TxnTaxDetail": {
      "TotalTax": 0
    },
    "SyncToken": "1",
    "CustomField": [
      {
        "DefinitionId": "1",
        "Type": "StringType",
        "Name": "Crew #"
      }
    ],
    "ShipAddr": {
      "CountrySubDivisionCode": "CA",
      "City": "Bayshore",
      "PostalCode": "94326",
      "Id": "98",
      "Line1": "4581 Finch St."
    },
    "EmailStatus": "NotSet",
    "BillAddr": {
      "CountrySubDivisionCode": "CA",
      "City": "Bayshore",
      "PostalCode": "94326",
      "Id": "97",
      "Line1": "4581 Finch St."
    },
    "MetaData": {
      "CreateTime": "2014-12-31T16:17:23-08:00",
      "LastUpdatedTime": "2015-02-09T12:29:53-08:00"
    },
    "BillEmail": {
      "Address": "virti_vora@Intuit.com"
    },
    "Id": "161"
  },
  "time": "2015-02-09T12:29:52.970-08:00"
}
```

### Get a salesreceipt as PDF

- **Method**: GET
- **URL**: `/v3/company/{realmID}/salesreceipt/{salesreceiptId}/pdf`

### Query a salesreceipt

- **Method**: GET
- **URL**: `/v3/company/{realmID}/query?query={selectStatement}`

**Sample Query**: `select * from SalesReceipt where id='11'`

**Response**:
```json
{
  "QueryResponse": {
    "SalesReceipt": [
      {
        "TxnDate": "2014-09-14",
        "domain": "QBO",
        "PrintStatus": "NotSet",
        "PaymentRefNum": "10264",
        "TotalAmt": 337.5,
        "Line": [
          {
            "Description": "Custom Design",
            "DetailType": "SalesItemLineDetail",
            "SalesItemLineDetail": {
              "TaxCodeRef": {
                "value": "NON"
              },
              "Qty": 4.5,
              "UnitPrice": 75,
              "ItemRef": {
                "name": "Design",
                "value": "4"
              }
            },
            "LineNum": 1,
            "Amount": 337.5,
            "Id": "1"
          },
          {
            "DetailType": "SubTotalLineDetail",
            "Amount": 337.5,
            "SubTotalLineDetail": {}
          }
        ],
        "ApplyTaxAfterDiscount": false,
        "DocNumber": "1003",
        "sparse": false,
        "DepositToAccountRef": {
          "name": "Checking",
          "value": "35"
        },
        "CustomerMemo": {
          "value": "Updated customer memo via sparse update operation."
        },
        "ProjectRef": {
          "value": "39298243"
        },
        "Balance": 0,
        "CustomerRef": {
          "name": "Dylan Sollfrank",
          "value": "6"
        },
        "TxnTaxDetail": {
          "TotalTax": 0
        },
        "SyncToken": "3",
        "PaymentMethodRef": {
          "name": "Check",
          "value": "2"
        },
        "EmailStatus": "NotSet",
        "BillAddr": {
          "Id": "122",
          "Line1": "Dylan Sollfrank"
        },
        "MetaData": {
          "CreateTime": "2014-09-16T14:59:48-07:00",
          "LastUpdatedTime": "2015-07-29T09:48:53-07:00"
        },
        "CustomField": [
          {
            "DefinitionId": "1",
            "Type": "StringType",
            "Name": "Crew #"
          }
        ],
        "Id": "11"
      }
    ],
    "startPosition": 1,
    "maxResults": 1
  },
  "time": "2015-07-29T09:50:39.882-07:00"
}
```

### Read a salesreceipt

- **Method**: GET
- **URL**: `/v3/company/{realmID}/salesreceipt/{salesreceiptId}`

Retrieves the details of a SalesReceipt object that has been previously created.

**Response**:
```json
{
  "SalesReceipt": {
    "TxnDate": "2014-09-14",
    "domain": "QBO",
    "PrintStatus": "NotSet",
    "PaymentRefNum": "10264",
    "TotalAmt": 337.5,
    "Line": [
      {
        "Description": "Custom Design",
        "DetailType": "SalesItemLineDetail",
        "SalesItemLineDetail": {
          "TaxCodeRef": {
            "value": "NON"
          },
          "Qty": 4.5,
          "UnitPrice": 75,
          "ItemRef": {
            "name": "Design",
            "value": "4"
          }
        },
        "LineNum": 1,
        "Amount": 337.5,
        "Id": "1"
      },
      {
        "DetailType": "SubTotalLineDetail",
        "Amount": 337.5,
        "SubTotalLineDetail": {}
      }
    ],
    "ApplyTaxAfterDiscount": false,
    "DocNumber": "1003",
    "sparse": false,
    "DepositToAccountRef": {
      "name": "Checking",
      "value": "35"
    },
    "CustomerMemo": {
      "value": "Thank you for your business and have a great day!"
    },
    "ProjectRef": {
      "value": "39298243"
    },
    "Balance": 0,
    "CustomerRef": {
      "name": "Dylan Sollfrank",
      "value": "6"
    },
    "TxnTaxDetail": {
      "TotalTax": 0
    },
    "SyncToken": "0",
    "PaymentMethodRef": {
      "name": "Check",
      "value": "2"
    },
    "EmailStatus": "NotSet",
    "BillAddr": {
      "Lat": "INVALID",
      "Long": "INVALID",
      "Id": "49",
      "Line1": "Dylan Sollfrank"
    },
    "MetaData": {
      "CreateTime": "2014-09-16T14:59:48-07:00",
      "LastUpdatedTime": "2014-09-16T14:59:48-07:00"
    },
    "CustomField": [
      {
        "DefinitionId": "1",
        "Type": "StringType",
        "Name": "Crew #"
      }
    ],
    "Id": "11"
  },
  "time": "2015-07-29T09:29:56.229-07:00"
}
```

### Send a salesreceipt

- **Method**: POST
- **URL**: `/v3/company/{realmID}/salesreceipt/{salesreceiptId}/send`

• The SalesReceipt.EmailStatus parameter is set to EmailSent. • The SalesReceipt.DeliveryInfo element is populated with sending information.

**Response**:
```json
{
  "SalesReceipt": {
    "DocNumber": "1047",
    "SyncToken": "0",
    "domain": "QBO",
    "Balance": 0,
    "DepositToAccountRef": {
      "name": "Undeposited Funds",
      "value": "4"
    },
    "TxnDate": "2013-03-13",
    "TotalAmt": 5,
    "CurrencyRef": {
      "name": "United States Dollar",
      "value": "USD"
    },
    "ProjectRef": {
      "value": "39298243"
    },
    "PrivateNote": "Memo for SalesReceipt",
    "PrintStatus": "NeedToPrint",
    "DepartmentRef": {
      "name": "Department1",
      "value": "1"
    },
    "DeliveryInfo": {
      "DeliveryType": "Email",
      "DeliveryTime": "2014-12-17T11:50:52-08:00"
    },
    "EmailStatus": "EmailSent",
    "sparse": false,
    "Line": [
      {
        "Description": "123189403765",
        "DetailType": "SalesItemLineDetail",
        "SalesItemLineDetail": {
          "TaxCodeRef": {
            "value": "NON"
          },
          "Qty": 10,
          "UnitPrice": 0.5,
          "ItemRef": {
            "name": "Sales",
            "value": "1"
          }
        },
        "LineNum": 1,
        "Amount": 5,
        "Id": "1"
      },
      {
        "DetailType": "SubTotalLineDetail",
        "Amount": 5,
        "SubTotalLineDetail": {}
      }
    ],
    "ApplyTaxAfterDiscount": false,
    "CustomField": [
      {
        "Type": "StringType",
        "Name": "Custom 1"
      },
      {
        "Type": "StringType",
        "Name": "Custom 2"
      },
      {
        "Type": "StringType",
        "Name": "Custom 3"
      }
    ],
    "Id": "97",
    "TxnTaxDetail": {
      "TotalTax": 0
    },
    "MetaData": {
      "CreateTime": "2013-03-13T13:31:43-07:00",
      "LastUpdatedTime": "2014-12-17T11:50:54-08:00"
    }
  },
  "time": "2013-03-13T13:31:42.956-07:00"
}
```

### Full update a salesreceipt

- **Method**: POST
- **URL**: `/v3/company/{realmID}/salesreceipt`

Use this operation to update any of the writable fields of an existing SalesReceipt object. The request body must include all writable fields of the existing object as returned in a read response. Writable fields omitted from the request body are set to NULL. The ID of the object to update is specif...

**Request Body**:
```json
{
  "TxnDate": "2014-09-14",
  "domain": "QBO",
  "PrintStatus": "NotSet",
  "PaymentRefNum": "10264",
  "TotalAmt": 337.5,
  "Line": [
    {
      "Description": "Custom Design",
      "DetailType": "SalesItemLineDetail",
      "SalesItemLineDetail": {
        "TaxCodeRef": {
          "value": "NON"
        },
        "Qty": 4.5,
        "UnitPrice": 75,
        "ItemRef": {
          "name": "Design",
          "value": "4"
        }
      },
      "LineNum": 1,
      "Amount": 337.5,
      "Id": "1"
    },
    {
      "DetailType": "SubTotalLineDetail",
      "Amount": 337.5,
      "SubTotalLineDetail": {}
    }
  ],
  "ApplyTaxAfterDiscount": false,
  "DocNumber": "1003",
  "sparse": false,
  "DepositToAccountRef": {
    "name": "Checking",
    "value": "35"
  },
  "CustomerMemo": {
    "value": "An updated customer memo."
  },
  "ProjectRef": {
    "value": "39298243"
  },
  "Balance": 0,
  "CustomerRef": {
    "name": "Dylan Sollfrank",
    "value": "6"
  },
  "TxnTaxDetail": {
    "TotalTax": 0
  },
  "SyncToken": "0",
  "PaymentMethodRef": {
    "name": "Check",
    "value": "2"
  },
  "EmailStatus": "NotSet",
  "BillAddr": {
    "Lat": "INVALID",
    "Long": "INVALID",
    "Id": "49",
    "Line1": "Dylan Sollfrank"
  },
  "MetaData": {
    "CreateTime": "2014-09-16T14:59:48-07:00",
    "LastUpdatedTime": "2014-09-16T14:59:48-07:00"
  },
  "CustomField": [
    {
      "DefinitionId": "1",
      "Type": "StringType",
      "Name": "Crew #"
    }
  ],
  "Id": "11"
}
```

**Response**:
```json
{
  "SalesReceipt": {
    "TxnDate": "2014-09-14",
    "domain": "QBO",
    "PrintStatus": "NotSet",
    "PaymentRefNum": "10264",
    "TotalAmt": 337.5,
    "Line": [
      {
        "Description": "Custom Design",
        "DetailType": "SalesItemLineDetail",
        "SalesItemLineDetail": {
          "TaxCodeRef": {
            "value": "NON"
          },
          "Qty": 4.5,
          "UnitPrice": 75,
          "ItemRef": {
            "name": "Design",
            "value": "4"
          }
        },
        "LineNum": 1,
        "Amount": 337.5,
        "Id": "1"
      },
      {
        "DetailType": "SubTotalLineDetail",
        "Amount": 337.5,
        "SubTotalLineDetail": {}
      }
    ],
    "ApplyTaxAfterDiscount": false,
    "DocNumber": "1003",
    "sparse": false,
    "DepositToAccountRef": {
      "name": "Checking",
      "value": "35"
    },
    "CustomerMemo": {
      "value": "An updated customer memo."
    },
    "ProjectRef": {
      "value": "39298243"
    },
    "Balance": 0,
    "CustomerRef": {
      "name": "Dylan Sollfrank",
      "value": "6"
    },
    "TxnTaxDetail": {
      "TotalTax": 0
    },
    "SyncToken": "1",
    "PaymentMethodRef": {
      "name": "Check",
      "value": "2"
    },
    "EmailStatus": "NotSet",
    "BillAddr": {
      "Lat": "INVALID",
      "Long": "INVALID",
      "Id": "49",
      "Line1": "Dylan Sollfrank"
    },
    "MetaData": {
      "CreateTime": "2014-09-16T14:59:48-07:00",
      "LastUpdatedTime": "2015-07-29T09:43:18-07:00"
    },
    "CustomField": [
      {
        "DefinitionId": "1",
        "Type": "StringType",
        "Name": "Crew #"
      }
    ],
    "Id": "11"
  },
  "time": "2015-07-29T09:43:01.436-07:00"
}
```

### Sparse update a salesreceipt

- **Method**: POST
- **URL**: `/v3/company/{realmID}/salesreceipt`

Sparse updating provides the ability to update a subset of properties for a given object; only elements specified in the request are updated. Missing elements are left untouched. The ID of the object to update is specified in the request body.

**Request Body**:
```json
{
  "SyncToken": "0",
  "Line": [
    {
      "DetailType": "SalesItemLineDetail",
      "Amount": 5,
      "Description": "UpdatedDescription",
      "SalesItemLineDetail": {
        "Qty": 10,
        "UnitPrice": 0.5,
        "ItemRef": {
          "value": "1"
        }
      }
    }
  ],
  "Id": "97",
  "sparse": true,
  "MetaData": {
    "CreateTime": "2013-03-13T13:39:57-07:00",
    "LastUpdatedTime": "2013-03-13T13:39:57-07:00"
  }
}
```

**Response**:
```json
{
  "SalesReceipt": {
    "DocNumber": "1003",
    "SyncToken": "2",
    "domain": "QBO",
    "Balance": 0,
    "DepositToAccountRef": {
      "name": "Undeposited Funds",
      "value": "4"
    },
    "TxnDate": "2015-07-29",
    "TotalAmt": 337.5,
    "CustomerMemo": {
      "value": "A sparsely updated customer memo."
    },
    "PrintStatus": "NeedToPrint",
    "EmailStatus": "NotSet",
    "sparse": false,
    "Line": [
      {
        "Description": "Custom Design",
        "DetailType": "SalesItemLineDetail",
        "SalesItemLineDetail": {
          "TaxCodeRef": {
            "value": "NON"
          },
          "Qty": 4.5,
          "UnitPrice": 75,
          "ItemRef": {
            "name": "Design",
            "value": "4"
          }
        },
        "LineNum": 1,
        "Amount": 337.5,
        "Id": "1"
      },
      {
        "DetailType": "SubTotalLineDetail",
        "Amount": 337.5,
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
    "Id": "11",
    "TxnTaxDetail": {
      "TotalTax": 0
    },
    "MetaData": {
      "CreateTime": "2014-09-16T14:59:48-07:00",
      "LastUpdatedTime": "2015-07-29T09:45:55-07:00"
    }
  },
  "time": "2015-07-29T09:45:39.381-07:00"
}
```
