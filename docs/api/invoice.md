# Invoice

An Invoice represents a sales form where the customer pays for a product or service later.

## Business Rules

- An invoice must have at least one Line for either a sales item or an inline subtotal.
- An invoice must have CustomerRef populated.
- The DocNumber attribute is populated automatically by the data service if not supplied.
- If ShipAddr, BillAddr, or both are not provided, the appropriate customer address from the referenced Customer object is used to fill those values.
- If you have a large number of invoice and corresponding payment records that you wish to import to the QuickBooks Online company, sort the invoice and payment records in chronological order and use the batch resource to send invoice and payments batches of 10, one after the other, to ensure any open invoices get credited with their payments.
- If an invoice is taxable, there is a limit of 750 lines per invoice.

## Attributes

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| Id | String | Required for update | Unique identifier for this object. Sort order is ASC by default. |
| Line | Line | Required | Individual line items of a transaction. Valid Line types include SalesItemLine, GroupLine, DescriptionOnlyLine (also used for inline Subtotal lines), DiscountLine and SubTotalLine (used for the overal... |
| CustomerRef | ReferenceType | Required | Reference to a customer or job. Query the Customer name list resource to determine the appropriate Customer object for this reference. Use Customer.Id and Customer. |
| SyncToken | String | Required for update | Version number of the object. It is used to lock an object for use by one app at a time. As soon as an application modifies an object, its SyncToken is incremented. |
| ShipFromAddr | PhysicalAddress | Conditionally required | Identifies the address where the goods are shipped from. For transactions without shipping, it represents the address where the sale took place. If automated sales tax is enabled (Preferences. |
| CurrencyRef | CurrencyRefType | Conditionally required | Reference to the currency in which all amounts on the associated transaction are expressed. This must be defined if multicurrency is enabled for the company. |
| GlobalTaxCalculation | String | Conditionally required | Method in which tax is applied. Allowed values are: TaxExcluded, TaxInclusive, and NotApplicable. |
| DocNumber | String | Conditionally required | Reference number for the transaction. If not explicitly provided at create time, this field is populated based on the setting of Preferences:CustomTxnNumber as follows: |
| ProjectRef | ReferenceType | Conditionally required | Reference to the Project ID associated with this transaction. Available with Minor Version 69 and above |
| BillEmail | EmailAddress | Conditionally required | Identifies the e-mail address where the invoice is sent. If EmailStatus=NeedToSend, BillEmailis a required input. |
| TxnDate | Date | Optional | The date entered by the user when this transaction occurred. |
| ShipDate | Date | Optional | Date for delivery of goods or services. |
| TrackingNum | String | Optional | Shipping provider's tracking number for the delivery of the goods associated with the transaction. |
| ClassRef | ReferenceType | Optional | Reference to the Class associated with the transaction. Available if Preferences.AccountingInfoPrefs.ClassTrackingPerTxn is set to true. |
| PrintStatus | String | Optional | Printing status of the invoice. Valid values: NotSet, NeedToPrint, PrintComplete . |
| SalesTermRef | ReferenceType | Optional | Reference to the sales term associated with the transaction. Query the Term name list resource to determine the appropriate Term object for this reference. Use Term. |
| TxnSource | String | Optional | Used internally to specify originating source of a credit card transaction. |
| LinkedTxn | LinkedTxn | Optional | Zero or more related transactions to this Invoice object. The following linked relationships are supported: |
| DepositToAccountRef | ReferenceType | Optional | Account to which money is deposited. Query the Account name list resource to determine the appropriate Account object for this reference, where Account. |
| AllowOnlineACHPayment | Boolean | Optional | Specifies if this invoice can be paid with online bank transfers. |
| TransactionLocationType | String | Optional | The account location. Valid values for France: WithinFrance, FranceOverseas, OutsideFranceWithEU, OutsideEU. |
| DueDate | Date | Optional | Date when the payment of the transaction is due. If date is not provided, the number of days specified in SalesTermRef added the transaction date will be used. |
| MetaData | ModificationMetaData | Read-only | Descriptive information about the object. The MetaData values are set by Data Services and are  for all applications. |
| PrivateNote | String | Optional | User entered, organization-private note about the transaction. This note does not appear on the invoice to the customer. |
| BillEmailCc | EmailAddress | Optional | Identifies the carbon copy e-mail address where the invoice is sent. If not specified, this field is populated from that defined in Preferences.SalesFormsPrefs.SalesEmailCc. |
| CustomerMemo | MemoRef | Optional | User-entered message to the customer; this message is visible to end user on their transactions. |
| EmailStatus | String | Optional | Email status of the invoice. Valid values: NotSet, NeedToSend, EmailSent |
| ExchangeRate | Decimal | Optional | The number of home currency units it takes to equal one unit of currency specified by CurrencyRef. Applicable if multicurrency is enabled for the company. |
| Deposit | Decimal | Optional | The deposit made towards this invoice. |
| TxnTaxDetail | TxnTaxDetail | Optional | This data type provides information for taxes charged on the transaction as a whole. |
| AllowOnlineCreditCardPayment | Boolean | Optional | Specifies if online credit card payments are allowed for this invoice. |
| CustomField | CustomField | Optional | One of, up to three custom fields for the transaction. Available for custom fields so configured for the company. Check Preferences.SalesFormsPrefs.CustomField and Preferences. |
| ShipAddr | PhysicalAddress | Optional | Identifies the address where the goods must be shipped. |
| DepartmentRef | ReferenceType | Optional | A reference to a Department object specifying the location of the transaction. Available if Preferences.AccountingInfoPrefs.TrackDepartments is set to true. |
| BillEmailBcc | EmailAddress | Optional | Identifies the blind carbon copy e-mail address where the invoice is sent. If not specified, this field is populated from that defined in Preferences.SalesFormsPrefs. |
| ShipMethodRef | ReferenceType | Optional | Reference to the ShipMethod associated with the transaction. There is no shipping method list. Reference resolves to a string. |
| BillAddr | PhysicalAddress | Optional | Bill-to address of the Invoice. If BillAddris not specified, and a default Customer:BillingAddr is specified in QuickBooks for this customer, the default bill-to address is us... |
| ApplyTaxAfterDiscount | Boolean | Optional | If false or null, calculate the sales tax first, and then apply the discount. If true, subtract the discount first and then calculate the sales tax. |
| HomeBalance | Decimal | Read-only | Convenience field containing the amount in Balance expressed in terms of the home currency. Calculated by QuickBooks business logic. |
| DeliveryInfo | DeliveryInfo | Read-only | Email delivery information. Returned when a request has been made to deliver email with the send operation. |
| TotalAmt | BigDecimal | Read-only | Indicates the total amount of the transaction. This includes the total of all the charges, allowances, and taxes. |
| InvoiceLink | String | Read-only | Sharable link for the invoice sent to external customers. The link is generated only for invoices with online payment enabled and having a valid customer email address. |
| RecurDataRef | ReferenceType | Read-only | A reference to the Recurring Transaction. It captures what recurring transaction template the Invoice was created from. |
| TaxExemptionRef | ReferenceType | Read-only | Reference to the TaxExepmtion ID associated with this object. Available for companies that have |
| Balance | Decimal | Read-only | The balance reflecting any payments made against the transaction. Initially set to the value of TotalAmt. A Balance of 0 indicates the invoice is fully paid. |
| HomeTotalAmt | Decimal | Read-only | Total amount of the transaction in the home currency. Includes the total of all the charges, allowances and taxes. Calculated by QuickBooks business logic. |
| FreeFormAddress | Boolean | Optional | Denotes how ShipAddr is stored: formatted or unformatted. The value of this flag is  based on format of shipping address at object create time. |
| AllowOnlinePayment | Boolean | Deprecated |  deprecated Boolean Flag to allow online payments. In use before AllowOnlineCreditCardPayment and AllowOnlineACHPayment flags existed and provided to maintain backward compatibility. |
| AllowIPNPayment | Boolean | Deprecated |  deprecated Boolean Flag to allow payments from legacy Intuit Payment Network (IPN). Provided to maintain backward compatibility and must always be set to false. Do not modify |

## Sample Object

```json
{
  "Invoice": {
    "TxnDate": "2014-09-19",
    "domain": "QBO",
    "PrintStatus": "NeedToPrint",
    "SalesTermRef": {
      "value": "3"
    },
    "TotalAmt": 362.07,
    "Line": [
      {
        "Description": "Rock Fountain",
        "DetailType": "SalesItemLineDetail",
        "SalesItemLineDetail": {
          "TaxCodeRef": {
            "value": "TAX"
          },
          "Qty": 1,
          "UnitPrice": 275,
          "ItemRef": {
            "name": "Rock Fountain",
            "value": "5"
          }
        },
        "LineNum": 1,
        "Amount": 275.0,
        "Id": "1"
      },
      {
        "Description": "Fountain Pump",
        "DetailType": "SalesItemLineDetail",
        "SalesItemLineDetail": {
          "TaxCodeRef": {
            "value": "TAX"
          },
          "Qty": 1,
          "UnitPrice": 12.75,
          "ItemRef": {
            "name": "Pump",
            "value": "11"
          }
        },
        "LineNum": 2,
        "Amount": 12.75,
        "Id": "2"
      },
      {
        "Description": "Concrete for fountain installation",
        "DetailType": "SalesItemLineDetail",
        "SalesItemLineDetail": {
          "TaxCodeRef": {
            "value": "TAX"
          },
          "Qty": 5,
          "UnitPrice": 9.5,
          "ItemRef": {
            "name": "Concrete",
            "value": "3"
          }
        },
        "LineNum": 3,
        "Amount": 47.5,
        "Id": "3"
      },
      {
        "DetailType": "SubTotalLineDetail",
        "Amount": 335.25,
        "SubTotalLineDetail": {}
      }
    ],
    "DueDate": "2014-10-19",
    "ApplyTaxAfterDiscount": false,
    "DocNumber": "1037",
    "sparse": false,
    "CustomerMemo": {
      "value": "Thank you for your business and have a great day!"
    },
    "ProjectRef": {
      "value": "39298045"
    },
    "Deposit": 0,
    "Balance": 362.07,
    "CustomerRef": {
      "name": "Sonnenschein Family Store",
      "value": "24"
    },
    "TxnTaxDetail": {
      "TxnTaxCodeRef": {
        "value": "2"
      },
      "TotalTax": 26.82,
      "TaxLine": [
        {
          "DetailType": "TaxLineDetail",
          "Amount": 26.82,
          "TaxLineDetail": {
            "NetAmountTaxable": 335.25,
            "TaxPercent": 8,
            "TaxRateRef": {
              "value": "3"
            },
            "PercentBased": true
          }
        }
      ]
    },
    "SyncToken": "0",
    "LinkedTxn": [
      {
        "TxnId": "100",
        "TxnType": "Estimate"
      }
    ],
    "BillEmail": {
      "Address": "Familiystore@intuit.com"
    },
    "ShipAddr": {
      "City": "Middlefield",
      "Line1": "5647 Cypress Hill Ave.",
      "PostalCode": "94303",
      "Lat": "37.4238562",
      "Long": "-122.1141681",
      "CountrySubDivisionCode": "CA",
      "Id": "25"
    },
    "EmailStatus": "NotSet",
    "BillAddr": {
      "Line4": "Middlefield, CA 94303",
      "Line3": "5647 Cypress Hill Ave.",
      "Line2": "Sonnenschein Family Store",
      "Line1": "Russ Sonnenschein",
      "Long": "-122.1141681",
      "Lat": "37.4238562",
      "Id": "95"
    },
    "MetaData": {
      "CreateTime": "2014-09-19T13:16:17-07:00",
      "LastUpdatedTime": "2014-09-19T13:16:17-07:00"
    },
    "CustomField": [
      {
        "DefinitionId": "1",
        "StringValue": "102",
        "Type": "StringType",
        "Name": "Crew #"
      }
    ],
    "Id": "130"
  },
  "time": "2015-07-24T10:48:27.082-07:00"
}
```

## Operations

### Create an invoice

- **Method**: POST
- **URL**: `/v3/company/{realmID}/invoice`

• Have at least one Line a sales item or inline subtotal. • Have a populated CustomerRef element.

**Request Body**:
```json
{
  "Line": [
    {
      "DetailType": "SalesItemLineDetail",
      "Amount": 100.0,
      "SalesItemLineDetail": {
        "ItemRef": {
          "name": "Services",
          "value": "1"
        }
      }
    }
  ],
  "CustomerRef": {
    "value": "1"
  }
}
```

**Response**:
```json
{
  "Invoice": {
    "TxnDate": "2015-07-24",
    "domain": "QBO",
    "PrintStatus": "NeedToPrint",
    "TotalAmt": 100.0,
    "Line": [
      {
        "LineNum": 1,
        "Amount": 100.0,
        "SalesItemLineDetail": {
          "TaxCodeRef": {
            "value": "NON"
          },
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
        "Amount": 100.0,
        "SubTotalLineDetail": {}
      }
    ],
    "DueDate": "2015-08-23",
    "ApplyTaxAfterDiscount": false,
    "DocNumber": "1069",
    "sparse": false,
    "ProjectRef": {
      "value": "39298034"
    },
    "Deposit": 0,
    "Balance": 100.0,
    "CustomerRef": {
      "name": "Amy's Bird Sanctuary",
      "value": "1"
    },
    "TxnTaxDetail": {
      "TotalTax": 0
    },
    "SyncToken": "0",
    "LinkedTxn": [],
    "ShipAddr": {
      "City": "Bayshore",
      "Line1": "4581 Finch St.",
      "PostalCode": "94326",
      "Lat": "INVALID",
      "Long": "INVALID",
      "CountrySubDivisionCode": "CA",
      "Id": "109"
    },
    "EmailStatus": "NotSet",
    "BillAddr": {
      "City": "Bayshore",
      "Line1": "4581 Finch St.",
      "PostalCode": "94326",
      "Lat": "INVALID",
      "Long": "INVALID",
      "CountrySubDivisionCode": "CA",
      "Id": "2"
    },
    "MetaData": {
      "CreateTime": "2015-07-24T10:33:39-07:00",
      "LastUpdatedTime": "2015-07-24T10:33:39-07:00"
    },
    "CustomField": [
      {
        "DefinitionId": "1",
        "Type": "StringType",
        "Name": "Crew #"
      }
    ],
    "Id": "238"
  },
  "time": "2015-07-24T10:33:39.11-07:00"
}
```

### Delete an invoice

- **Method**: POST
- **URL**: `/v3/company/{realmID}/invoice?operation=delete`

This operation deletes the invoice object specified in the request body. Include a minimum of Invoice.Id and Invoice.SyncToken in the request body. You must unlink any linked transactions associated with the invoice object before deleting it.

**Request Body**:
```json
{
  "SyncToken": "3",
  "Id": "33"
}
```

**Response**:
```json
{
  "Invoice": {
    "status": "Deleted",
    "domain": "QBO",
    "Id": "33"
  },
  "time": "2013-03-15T00:18:15.322-07:00"
}
```

### Void an invoice

- **Method**: POST
- **URL**: `/v3/company/{realmID}/invoice?operation=void`

Use this operation to void an existing invoice object; include a minimum of Invoice.Id and the current Invoice.SyncToken. The transaction remains active but all amounts and quantities are zeroed and the string, Voided, is injected into Invoice.PrivateNote, prepended to existing text if present.

**Request Body**:
```json
{
  "SyncToken": "0",
  "Id": "129"
}
```

**Response**:
```json
{
  "Invoice": {
    "AllowOnlineACHPayment": false,
    "domain": "QBO",
    "TxnDate": "2014-11-09",
    "PrintStatus": "NEED_TO_PRINT",
    "SalesTermRef": {
      "value": "3"
    },
    "TotalAmt": 0,
    "Line": [
      {
        "Description": "Sod",
        "DetailType": "SALES_ITEM_LINE_DETAIL",
        "SalesItemLineDetail": {
          "TaxCodeRef": {
            "value": "TAX"
          },
          "Qty": 0,
          "ItemRef": {
            "name": "Sod",
            "value": "14"
          }
        },
        "LineNum": 1,
        "Amount": 0,
        "Id": "1"
      },
      {
        "Description": "2 cubic ft. bag",
        "DetailType": "SALES_ITEM_LINE_DETAIL",
        "SalesItemLineDetail": {
          "TaxCodeRef": {
            "value": "TAX"
          },
          "Qty": 0,
          "ItemRef": {
            "name": "Soil",
            "value": "15"
          }
        },
        "LineNum": 2,
        "Amount": 0,
        "Id": "2"
      },
      {
        "Description": "Weekly Gardening Service",
        "DetailType": "SALES_ITEM_LINE_DETAIL",
        "SalesItemLineDetail": {
          "TaxCodeRef": {
            "value": "NON"
          },
          "Qty": 0,
          "ItemRef": {
            "name": "Gardening",
            "value": "6"
          }
        },
        "LineNum": 3,
        "Amount": 0,
        "Id": "3"
      },
      {
        "Description": "Rock Fountain",
        "DetailType": "SALES_ITEM_LINE_DETAIL",
        "SalesItemLineDetail": {
          "TaxCodeRef": {
            "value": "TAX"
          },
          "Qty": 0,
          "ItemRef": {
            "name": "Rock Fountain",
            "value": "5"
          }
        },
        "LineNum": 4,
        "Amount": 0,
        "Id": "4"
      },
      {
        "Description": "Fountain Pump",
        "DetailType": "SALES_ITEM_LINE_DETAIL",
        "SalesItemLineDetail": {
          "TaxCodeRef": {
            "value": "TAX"
          },
          "Qty": 0,
          "ItemRef": {
            "name": "Pump",
            "value": "11"
          }
        },
        "LineNum": 5,
        "Amount": 0,
        "Id": "5"
      },
      {
        "DetailType": "SUB_TOTAL_LINE_DETAIL",
        "Amount": 0,
        "SubTotalLineDetail": {}
      }
    ],
    "DueDate": "2014-12-09",
    "MetaData": {
      "CreateTime": "2014-11-09T13:15:36-08:00",
      "LastUpdatedTime": "2016-03-16T12:27:10-07:00"
    },
    "DocNumber": "1036",
    "PrivateNote": "Voided",
    "sparse": false,
    "CustomerMemo": {
      "value": "Thank you for your business and have a great day!"
    },
    "ProjectRef": {
      "value": "39298045"
    },
    "Deposit": 0,
    "Balance": 0,
    "CustomerRef": {
      "name": "0969 Ocean View Road",
      "value": "8"
    },
    "TxnTaxDetail": {
      "TotalTax": 0
    },
    "AllowOnlineCreditCardPayment": false,
    "SyncToken": "1",
    "LinkedTxn": [],
    "BillEmail": {
      "Address": "Sporting_goods@intuit.com"
    },
    "ShipAddr": {
      "City": "Middlefield",
      "Line1": "370 Easy St.",
      "PostalCode": "94482",
      "Lat": "37.4031672",
      "Long": "-122.0642815",
      "CountrySubDivisionCode": "CA",
      "Id": "8"
    },
    "EmailStatus": "NOT_SET",
    "BillAddr": {
      "Line4": "Middlefield, CA 94482",
      "Line3": "370 Easy St.",
      "Line2": "Freeman Sporting Goods",
      "Line1": "Sasha Tillou",
      "Long": "INVALID",
      "Lat": "INVALID",
      "Id": "94"
    },
    "ApplyTaxAfterDiscount": false,
    "CustomField": [
      {
        "DefinitionId": "1",
        "StringValue": "105",
        "Type": "STRING_TYPE",
        "Name": "Crew #"
      }
    ],
    "Id": "129",
    "AllowOnlinePayment": false,
    "AllowIPNPayment": false
  },
  "time": "2016-03-16T12:27:10.711-07:00"
}
```

### Get an invoice as PDF

- **Method**: GET
- **URL**: `/v3/company/{realmID}/invoice/{invoiceId}/pdf`

### Query an invoice

- **Method**: GET
- **URL**: `/v3/company/{realmID}/query?query={selectStatement}`

**Sample Query**: `select * from Invoice where id = '239'`

**Response**:
```json
{
  "QueryResponse": {
    "startPosition": 1,
    "totalCount": 1,
    "maxResults": 1,
    "Invoice": [
      {
        "TxnDate": "2015-07-24",
        "domain": "QBO",
        "PrintStatus": "NeedToPrint",
        "TotalAmt": 150.0,
        "Line": [
          {
            "LineNum": 1,
            "Amount": 150.0,
            "SalesItemLineDetail": {
              "TaxCodeRef": {
                "value": "NON"
              },
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
            "Amount": 150.0,
            "SubTotalLineDetail": {}
          }
        ],
        "DueDate": "2015-08-23",
        "ApplyTaxAfterDiscount": false,
        "DocNumber": "1070",
        "sparse": false,
        "ProjectRef": {
          "value": "39298034"
        },
        "Deposit": 0,
        "Balance": 150.0,
        "CustomerRef": {
          "name": "Amy's Bird Sanctuary",
          "value": "1"
        },
        "TxnTaxDetail": {
          "TotalTax": 0
        },
        "SyncToken": "0",
        "LinkedTxn": [],
        "ShipAddr": {
          "City": "Bayshore",
          "Line1": "4581 Finch St.",
          "PostalCode": "94326",
          "Lat": "INVALID",
          "Long": "INVALID",
          "CountrySubDivisionCode": "CA",
          "Id": "109"
        },
        "EmailStatus": "NotSet",
        "BillAddr": {
          "City": "Bayshore",
          "Line1": "4581 Finch St.",
          "PostalCode": "94326",
          "Lat": "INVALID",
          "Long": "INVALID",
          "CountrySubDivisionCode": "CA",
          "Id": "2"
        },
        "MetaData": {
          "CreateTime": "2015-07-24T10:35:08-07:00",
          "LastUpdatedTime": "2015-07-24T10:35:08-07:00"
        },
        "CustomField": [
          {
            "DefinitionId": "1",
            "Type": "StringType",
            "Name": "Crew #"
          }
        ],
        "Id": "239"
      }
    ]
  },
  "time": "2015-07-24T10:38:50.01-07:00"
}
```

### Read an invoice

- **Method**: GET
- **URL**: `/v3/company/{realmID}/invoice/{invoiceId}`

Retrieves the details of an invoice that has been previously created.

**Response**:
```json
{
  "Invoice": {
    "TxnDate": "2014-09-19",
    "domain": "QBO",
    "PrintStatus": "NeedToPrint",
    "SalesTermRef": {
      "value": "3"
    },
    "TotalAmt": 362.07,
    "Line": [
      {
        "Description": "Rock Fountain",
        "DetailType": "SalesItemLineDetail",
        "SalesItemLineDetail": {
          "TaxCodeRef": {
            "value": "TAX"
          },
          "Qty": 1,
          "UnitPrice": 275,
          "ItemRef": {
            "name": "Rock Fountain",
            "value": "5"
          }
        },
        "LineNum": 1,
        "Amount": 275.0,
        "Id": "1"
      },
      {
        "Description": "Fountain Pump",
        "DetailType": "SalesItemLineDetail",
        "SalesItemLineDetail": {
          "TaxCodeRef": {
            "value": "TAX"
          },
          "Qty": 1,
          "UnitPrice": 12.75,
          "ItemRef": {
            "name": "Pump",
            "value": "11"
          }
        },
        "LineNum": 2,
        "Amount": 12.75,
        "Id": "2"
      },
      {
        "Description": "Concrete for fountain installation",
        "DetailType": "SalesItemLineDetail",
        "SalesItemLineDetail": {
          "TaxCodeRef": {
            "value": "TAX"
          },
          "Qty": 5,
          "UnitPrice": 9.5,
          "ItemRef": {
            "name": "Concrete",
            "value": "3"
          }
        },
        "LineNum": 3,
        "Amount": 47.5,
        "Id": "3"
      },
      {
        "DetailType": "SubTotalLineDetail",
        "Amount": 335.25,
        "SubTotalLineDetail": {}
      }
    ],
    "DueDate": "2014-10-19",
    "ApplyTaxAfterDiscount": false,
    "DocNumber": "1037",
    "sparse": false,
    "CustomerMemo": {
      "value": "Thank you for your business and have a great day!"
    },
    "ProjectRef": {
      "value": "39298045"
    },
    "Deposit": 0,
    "Balance": 362.07,
    "CustomerRef": {
      "name": "Sonnenschein Family Store",
      "value": "24"
    },
    "TxnTaxDetail": {
      "TxnTaxCodeRef": {
        "value": "2"
      },
      "TotalTax": 26.82,
      "TaxLine": [
        {
          "DetailType": "TaxLineDetail",
          "Amount": 26.82,
          "TaxLineDetail": {
            "NetAmountTaxable": 335.25,
            "TaxPercent": 8,
            "TaxRateRef": {
              "value": "3"
            },
            "PercentBased": true
          }
        }
      ]
    },
    "SyncToken": "0",
    "LinkedTxn": [
      {
        "TxnId": "100",
        "TxnType": "Estimate"
      }
    ],
    "BillEmail": {
      "Address": "Familiystore@intuit.com"
    },
    "ShipAddr": {
      "City": "Middlefield",
      "Line1": "5647 Cypress Hill Ave.",
      "PostalCode": "94303",
      "Lat": "37.4238562",
      "Long": "-122.1141681",
      "CountrySubDivisionCode": "CA",
      "Id": "25"
    },
    "EmailStatus": "NotSet",
    "BillAddr": {
      "Line4": "Middlefield, CA 94303",
      "Line3": "5647 Cypress Hill Ave.",
      "Line2": "Sonnenschein Family Store",
      "Line1": "Russ Sonnenschein",
      "Long": "-122.1141681",
      "Lat": "37.4238562",
      "Id": "95"
    },
    "MetaData": {
      "CreateTime": "2014-09-19T13:16:17-07:00",
      "LastUpdatedTime": "2014-09-19T13:16:17-07:00"
    },
    "CustomField": [
      {
        "DefinitionId": "1",
        "StringValue": "102",
        "Type": "StringType",
        "Name": "Crew #"
      }
    ],
    "Id": "130"
  },
  "time": "2015-07-24T10:48:27.082-07:00"
}
```

### Send an invoice

- **Method**: POST
- **URL**: `/v3/company/{realmID}/invoice/<invoiceId >/send`

• The Invoice.EmailStatus parameter is set to EmailSent. • The Invoice.DeliveryInfo element is populated with sending information<./li>

**Response**:
```json
{
  "Invoice": {
    "TxnDate": "2013-03-14",
    "domain": "QBO",
    "CurrencyRef": {
      "name": "United States Dollar",
      "value": "USD"
    },
    "ShipDate": "2013-03-01",
    "TrackingNum": "123456789",
    "ClassRef": {
      "name": "Class 1",
      "value": "200900000000000003901"
    },
    "PrintStatus": "NeedToPrint",
    "SalesTermRef": {
      "value": "4"
    },
    "DeliveryInfo": {
      "DeliveryType": "Email",
      "DeliveryTime": "2014-12-17T11:50:52-08:00"
    },
    "TotalAmt": 52.0,
    "Line": [
      {
        "Description": "Sample invoice create request",
        "DetailType": "SalesItemLineDetail",
        "SalesItemLineDetail": {
          "TaxCodeRef": {
            "value": "TAX"
          },
          "Qty": 1,
          "UnitPrice": 50,
          "ServiceDate": "2013-03-04",
          "ItemRef": {
            "name": "Hours",
            "value": "2"
          }
        },
        "LineNum": 1,
        "Amount": 50.0,
        "Id": "1"
      },
      {
        "DetailType": "SubTotalLineDetail",
        "Amount": 50.0,
        "SubTotalLineDetail": {}
      },
      {
        "DetailType": "DiscountLineDetail",
        "Amount": 5.0,
        "DiscountLineDetail": {
          "DiscountAccountRef": {
            "name": "Discounts given",
            "value": "30"
          },
          "PercentBased": true,
          "DiscountPercent": 10
        }
      },
      {
        "DetailType": "SalesItemLineDetail",
        "Amount": 2.0,
        "SalesItemLineDetail": {
          "ItemRef": {
            "value": "SHIPPING_ITEM_ID"
          }
        }
      }
    ],
    "DueDate": "2013-05-13",
    "MetaData": {
      "CreateTime": "2013-03-14T01:42:16-07:00",
      "LastUpdatedTime": "2014-12-17T11:50:58-08:00"
    },
    "DocNumber": "Sample_Inv#2",
    "PrivateNote": "Summary for sample invoice",
    "sparse": false,
    "DepositToAccountRef": {
      "name": "Undeposited Funds",
      "value": "4"
    },
    "CustomerMemo": {
      "value": "This is the customer message"
    },
    "EmailStatus": "EmailSent",
    "ProjectRef": {
      "value": "39298037"
    },
    "Deposit": 12.0,
    "Balance": 40.0,
    "CustomerRef": {
      "name": "Mr V3 Service Customer Jr2",
      "value": "15"
    },
    "TxnTaxDetail": {
      "TxnTaxCodeRef": {
        "value": "5"
      },
      "TotalTax": 5.0,
      "TaxLine": [
        {
          "DetailType": "TaxLineDetail",
          "Amount": 5.0,
          "TaxLineDetail": {
            "NetAmountTaxable": 50.0,
            "TaxPercent": 10,
            "TaxRateRef": {
              "value": "2"
            },
            "PercentBased": true
          }
        }
      ]
    },
    "SyncToken": "0",
    "BillEmail": {
      "Address": "test@intuit.com"
    },
    "ShipAddr": {
      "City": "San Jose",
      "Country": "USA",
      "Line5": "Cube 999",
      "Line4": "Dept 12",
      "Line3": "123 street",
      "Line2": "Building 1",
      "Line1": "Intuit",
      "PostalCode": "95123",
      "Lat": "37.2374847",
      "Long": "-121.8277925",
      "CountrySubDivisionCode": "CA",
      "Id": "36"
    },
    "DepartmentRef": {
      "name": "Mountain View",
      "value": "1"
    },
    "ShipMethodRef": {
      "name": "UPS",
      "value": "UPS"
    },
    "BillAddr": {
      "City": "Mountain View",
      "Country": "USA",
      "Line5": "Cube 999",
      "Line4": "Dept 12",
      "Line3": "123 street",
      "Line2": "Building 1",
      "Line1": "Google",
      "PostalCode": "95123",
      "Lat": "37.2374847",
      "Long": "-121.8277925",
      "CountrySubDivisionCode": "CA",
      "Id": "35"
    },
    "ApplyTaxAfterDiscount": false,
    "CustomField": [
      {
        "StringValue": "Custom1",
        "Type": "StringType",
        "Name": "Custom 1"
      },
      {
        "StringValue": "Custom2",
        "Type": "StringType",
        "Name": "Custom 2"
      },
      {
        "StringValue": "Custom3",
        "Type": "StringType",
        "Name": "Custom 3"
      }
    ],
    "Id": "96"
  },
  "time": "2013-03-14T13:32:04.895-07:00"
}
```

### Sparse update an invoice

- **Method**: POST
- **URL**: `/v3/company/{realmID}/invoice`

Sparse updating provides the ability to update a subset of properties for a given object; only elements specified in the request are updated. Missing elements are left untouched. The ID of the object to update is specified in the request body.

**Request Body**:
```json
{
  "SyncToken": "0",
  "Id": "238",
  "sparse": true,
  "DueDate": "2015-09-30"
}
```

**Response**:
```json
{
  "Invoice": {
    "TxnDate": "2015-07-24",
    "domain": "QBO",
    "PrintStatus": "NeedToPrint",
    "TotalAmt": 100.0,
    "Line": [
      {
        "LineNum": 1,
        "Amount": 100.0,
        "SalesItemLineDetail": {
          "TaxCodeRef": {
            "value": "NON"
          },
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
        "Amount": 100.0,
        "SubTotalLineDetail": {}
      }
    ],
    "DueDate": "2015-09-30",
    "ApplyTaxAfterDiscount": false,
    "DocNumber": "1069",
    "sparse": false,
    "ProjectRef": {
      "value": "39298045"
    },
    "Deposit": 0,
    "Balance": 100.0,
    "CustomerRef": {
      "name": "Amy's Bird Sanctuary",
      "value": "1"
    },
    "TxnTaxDetail": {
      "TotalTax": 0
    },
    "SyncToken": "1",
    "LinkedTxn": [],
    "ShipAddr": {
      "City": "Bayshore",
      "Line1": "4581 Finch St.",
      "PostalCode": "94326",
      "Lat": "INVALID",
      "Long": "INVALID",
      "CountrySubDivisionCode": "CA",
      "Id": "109"
    },
    "EmailStatus": "NotSet",
    "BillAddr": {
      "City": "Bayshore",
      "Line1": "4581 Finch St.",
      "PostalCode": "94326",
      "Lat": "INVALID",
      "Long": "INVALID",
      "CountrySubDivisionCode": "CA",
      "Id": "2"
    },
    "MetaData": {
      "CreateTime": "2015-07-24T10:33:39-07:00",
      "LastUpdatedTime": "2015-07-24T11:03:26-07:00"
    },
    "CustomField": [
      {
        "DefinitionId": "1",
        "Type": "StringType",
        "Name": "Crew #"
      }
    ],
    "Id": "238"
  },
  "time": "2015-07-24T11:03:26.674-07:00"
}
```

### Full update an invoice

- **Method**: POST
- **URL**: `/v3/company/{realmID}/invoice`

Use this operation to update any of the writable fields of an existing invoice object. The request body must include all writable fields of the existing object as returned in a read response. Writable fields omitted from the request body are set to NULL. The ID of the object to update is specified i...

**Request Body**:
```json
{
  "TxnDate": "2015-07-24",
  "domain": "QBO",
  "PrintStatus": "NeedToPrint",
  "TotalAmt": 150.0,
  "Line": [
    {
      "LineNum": 1,
      "Amount": 150.0,
      "SalesItemLineDetail": {
        "TaxCodeRef": {
          "value": "NON"
        },
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
      "Amount": 150.0,
      "SubTotalLineDetail": {}
    }
  ],
  "DueDate": "2015-08-23",
  "ApplyTaxAfterDiscount": false,
  "DocNumber": "1070",
  "sparse": false,
  "CustomerMemo": {
    "value": "Added customer memo."
  },
  "ProjectRef": {
    "value": "39298045"
  },
  "Balance": 150.0,
  "CustomerRef": {
    "name": "Amy's Bird Sanctuary",
    "value": "1"
  },
  "TxnTaxDetail": {
    "TotalTax": 0
  },
  "SyncToken": "0",
  "LinkedTxn": [],
  "ShipAddr": {
    "City": "Bayshore",
    "Line1": "4581 Finch St.",
    "PostalCode": "94326",
    "Lat": "INVALID",
    "Long": "INVALID",
    "CountrySubDivisionCode": "CA",
    "Id": "109"
  },
  "EmailStatus": "NotSet",
  "BillAddr": {
    "City": "Bayshore",
    "Line1": "4581 Finch St.",
    "PostalCode": "94326",
    "Lat": "INVALID",
    "Long": "INVALID",
    "CountrySubDivisionCode": "CA",
    "Id": "2"
  },
  "MetaData": {
    "CreateTime": "2015-07-24T10:35:08-07:00",
    "LastUpdatedTime": "2015-07-24T10:35:08-07:00"
  },
  "CustomField": [
    {
      "DefinitionId": "1",
      "Type": "StringType",
      "Name": "Crew #"
    }
  ],
  "Id": "239"
}
```

**Response**:
```json
{
  "Invoice": {
    "TxnDate": "2015-07-24",
    "domain": "QBO",
    "PrintStatus": "NeedToPrint",
    "TotalAmt": 150.0,
    "Line": [
      {
        "LineNum": 1,
        "Amount": 150.0,
        "SalesItemLineDetail": {
          "TaxCodeRef": {
            "value": "NON"
          },
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
        "Amount": 150.0,
        "SubTotalLineDetail": {}
      }
    ],
    "DueDate": "2015-08-23",
    "ApplyTaxAfterDiscount": false,
    "DocNumber": "1070",
    "sparse": false,
    "CustomerMemo": {
      "value": "Added customer memo."
    },
    "ProjectRef": {
      "value": "39298045"
    },
    "Deposit": 0,
    "Balance": 150.0,
    "CustomerRef": {
      "name": "Amy's Bird Sanctuary",
      "value": "1"
    },
    "TxnTaxDetail": {
      "TotalTax": 0
    },
    "SyncToken": "1",
    "LinkedTxn": [],
    "ShipAddr": {
      "CountrySubDivisionCode": "CA",
      "City": "Bayshore",
      "PostalCode": "94326",
      "Id": "118",
      "Line1": "4581 Finch St."
    },
    "EmailStatus": "NotSet",
    "BillAddr": {
      "CountrySubDivisionCode": "CA",
      "City": "Bayshore",
      "PostalCode": "94326",
      "Id": "117",
      "Line1": "4581 Finch St."
    },
    "MetaData": {
      "CreateTime": "2015-07-24T10:35:08-07:00",
      "LastUpdatedTime": "2015-07-24T10:53:39-07:00"
    },
    "CustomField": [
      {
        "DefinitionId": "1",
        "Type": "StringType",
        "Name": "Crew #"
      }
    ],
    "Id": "239"
  },
  "time": "2015-07-24T10:53:39.287-07:00"
}
```
