# CreditMemo

The CreditMemo object is a financial transaction representing a refund or credit of payment or part of a payment for goods or services that have been sold.

## Attributes

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| Id | String | Required for update | Unique identifier for this object. Sort order is ASC by default. |
| Line | Line | Required | Individual line items of a transaction. Valid Line types include:SalesItemLine, GroupLine, DescriptionOnlyLine, DiscountLine and SubTotalLine SalesItemLine GroupLine DescriptionOnlyLine DiscountLin... |
| CustomerRef | ReferenceType | Required | Reference to a customer or job. Query the Customer name list resource to determine the appropriate Customer object for this reference. Use Customer.Id and Customer. |
| SyncToken | String | Required for update | Version number of the object. It is used to lock an object for use by one app at a time. As soon as an application modifies an object, its SyncToken is incremented. |
| CurrencyRef | CurrencyRefType | Conditionally required | Reference to the currency in which all amounts on the associated transaction are expressed. This must be defined if multicurrency is enabled for the company. |
| GlobalTaxCalculation | String | Conditionally required | Method in which tax is applied. Allowed values are: TaxExcluded, TaxInclusive, and NotApplicable. |
| ProjectRef | ReferenceType | Conditionally required | Reference to the Project ID associated with this transaction. Available with Minor Version 69 and above |
| BillEmail | EmailAddress | Conditionally required | Identifies the e-mail address where the credit-memo is sent. If EmailStatus=NeedToSend, BillEmailis a required input. |
| TxnDate | Date | Optional | The date entered by the user when this transaction occurred. For posting transactions, this is the posting date that affects the financial statements. |
| CustomField | CustomField | Optional | One of, up to three custom fields for the transaction. Available for custom fields so configured for the company. Check Preferences.SalesFormsPrefs.CustomField and Preferences. |
| ClassRef | ReferenceType | Optional | Reference to the Class associated with the transaction. Available if Preferences.AccountingInfoPrefs.ClassTrackingPerLine is set to true. |
| PrintStatus | String | Optional | Printing status of the credit-memo. Valid values: NotSet, NeedToPrint, PrintComplete . |
| SalesTermRef | ReferenceType | Optional | Reference to the sales term associated with the transaction. Query the Term name list resource to determine the appropriate Term object for this reference. Use Term. |
| TotalAmt | BigDecimal | Optional | Indicates the total amount of the transaction. This includes the total of all the charges, allowances, and taxes. |
| InvoiceRef |  | Optional |  |
| ReferenceType |  | Optional | Reference to the Invoice for which Credit memo is issued. Needed for GST compliance. Use Invoice.Id and Invoice.Name from that object for InvoiceRef.value and InvoiceRef.name, respectively. |
| TransactionLocationType | String | Optional | The account location. Valid values for France: WithinFrance, FranceOverseas, OutsideFranceWithEU, OutsideEU. |
| ApplyTaxAfterDiscount | Boolean | Optional | If false or null, calculate the sales tax first, and then apply the discount. If true, subtract the discount first and then calculate the sales tax. |
| DocNumber | String | Optional | Reference number for the transaction. |
| PrivateNote | String | Optional | User entered, organization-private note about the transaction. This note does not appear on the deposit form. |
| CustomerMemo | MemoRef | Optional | User-entered message to the customer; this message is visible to end user on their transactions. |
| TxnTaxDetail | TxnTaxDetail | Optional | This data type provides information for taxes charged on the transaction as a whole. |
| PaymentMethodRef | ReferenceType | Read-only | Reference to a PaymentMethod associated with this transaction. Query the PaymentMethod name list resource to determine the appropriate PaymentMethod object for this reference. |
| ExchangeRate | Decimal | Optional | The number of home currency units it takes to equal one unit of currency specified by CurrencyRef. Applicable if multicurrency is enabled for the company. |
| ShipAddr | PhysicalAddress | Optional | Identifies the address where the goods must be shipped. |
| DepartmentRef | ReferenceType | Optional | A reference to a Department object specifying the location of the transaction. Available if Preferences.AccountingInfoPrefs.TrackDepartments is set to true. |
| EmailStatus | String | Optional | Email status of the credit-memo. Valid values: NotSet, NeedToSend, EmailSent |
| BillAddr | PhysicalAddress | Optional | Bill-to address of the credit memo. If BillAddris not specified, and a default Customer:BillingAddr is specified in QuickBooks for this customer, the default bill-to address i... |
| MetaData | ModificationMetaData | Read-only | Descriptive information about the object. The MetaData values are set by Data Services and are  for all applications. |
| HomeBalance | Decimal | Read-only | Convenience field containing the amount in Balance expressed in terms of the home currency. Calculated by QuickBooks business logic. |
| RemainingCredit | Decimal | Read-only | Indicates the total credit amount still available to apply towards the payment. |
| RecurDataRef | ReferenceType | Read-only | A reference to the Recurring Transaction. It captures what recurring transaction template the CreditMemo was created from. |
| TaxExemptionRef | ReferenceType | Read-only | Reference to the TaxExepmtion ID associated with this object. Available for companies that have |
| Balance | Decimal | Read-only | The balance reflecting any payments made against the transaction. Initially set to the value of TotalAmt. A Balance of 0 indicates the invoice is fully paid. |
| HomeTotalAmt | Decimal | Read-only | Total amount of the transaction in the home currency. Includes the total of all the charges, allowances and taxes. Calculated by QuickBooks business logic. |

## Sample Object

```json
{
  "CreditMemo": {
    "TxnDate": "2014-09-02",
    "domain": "QBO",
    "PrintStatus": "NeedToPrint",
    "TotalAmt": 100.0,
    "RemainingCredit": 0,
    "Line": [
      {
        "Description": "Pest Control Services",
        "DetailType": "SalesItemLineDetail",
        "SalesItemLineDetail": {
          "TaxCodeRef": {
            "value": "NON"
          },
          "Qty": 1,
          "UnitPrice": 100,
          "ItemRef": {
            "name": "Pest Control",
            "value": "10"
          }
        },
        "LineNum": 1,
        "Amount": 100.0,
        "Id": "1"
      },
      {
        "DetailType": "SubTotalLineDetail",
        "Amount": 100.0,
        "SubTotalLineDetail": {}
      }
    ],
    "ApplyTaxAfterDiscount": false,
    "DocNumber": "1026",
    "sparse": false,
    "CustomerMemo": {
      "value": "Updated customer memo."
    },
    "ProjectRef": {
      "value": "39298034"
    },
    "Balance": 0,
    "CustomerRef": {
      "name": "Amy's Bird Sanctuary",
      "value": "1"
    },
    "TxnTaxDetail": {
      "TotalTax": 0
    },
    "SyncToken": "3",
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
      "Id": "108",
      "Line1": "4581 Finch St."
    },
    "EmailStatus": "NotSet",
    "BillAddr": {
      "Line4": "Bayshore, CA 94326",
      "Line3": "4581 Finch St.",
      "Id": "79",
      "Line1": "Amy Lauterbach",
      "Line2": "Amy's Bird Sanctuary"
    },
    "MetaData": {
      "CreateTime": "2014-09-18T12:51:27-07:00",
      "LastUpdatedTime": "2015-07-01T09:16:28-07:00"
    },
    "BillEmail": {
      "Address": "Birds@Intuit.com"
    },
    "Id": "73"
  },
  "time": "2015-07-23T09:10:45.624-07:00"
}
```

## Operations

### Create a credit memo

- **Method**: POST
- **URL**: `/v3/company/{realmID}/creditmemo`

**Request Body**:
```json
{
  "Line": [
    {
      "DetailType": "SalesItemLineDetail",
      "Amount": 50,
      "SalesItemLineDetail": {
        "ItemRef": {
          "name": "Concrete",
          "value": "3"
        }
      }
    }
  ],
  "CustomerRef": {
    "name": "CoolCars",
    "value": "3"
  }
}
```

**Response**:
```json
{
  "CreditMemo": {
    "DocNumber": "1039",
    "SyncToken": "0",
    "domain": "QBO",
    "Balance": 50.0,
    "BillAddr": {
      "City": "Half Moon Bay",
      "Line1": "65 Ocean Dr.",
      "PostalCode": "94213",
      "Lat": "37.4300318",
      "Long": "-122.4336537",
      "CountrySubDivisionCode": "CA",
      "Id": "4"
    },
    "TxnDate": "2014-12-31",
    "TotalAmt": 50.0,
    "CustomerRef": {
      "name": "Cool Cars",
      "value": "3"
    },
    "ShipAddr": {
      "City": "Half Moon Bay",
      "Line1": "65 Ocean Dr.",
      "PostalCode": "94213",
      "Lat": "37.4300318",
      "Long": "-122.4336537",
      "CountrySubDivisionCode": "CA",
      "Id": "4"
    },
    "RemainingCredit": 50.0,
    "PrintStatus": "NeedToPrint",
    "ProjectRef": {
      "value": "39298034"
    },
    "EmailStatus": "NotSet",
    "sparse": false,
    "Line": [
      {
        "LineNum": 1,
        "Amount": 50.0,
        "SalesItemLineDetail": {
          "TaxCodeRef": {
            "value": "NON"
          },
          "ItemRef": {
            "name": "Concrete",
            "value": "3"
          }
        },
        "Id": "1",
        "DetailType": "SalesItemLineDetail"
      },
      {
        "DetailType": "SubTotalLineDetail",
        "Amount": 50.0,
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
    "Id": "150",
    "TxnTaxDetail": {
      "TotalTax": 0
    },
    "MetaData": {
      "CreateTime": "2014-12-31T09:44:40-08:00",
      "LastUpdatedTime": "2014-12-31T09:44:40-08:00"
    }
  },
  "time": "2014-12-31T09:44:40.726-08:00"
}
```

### Delete a credit memo

- **Method**: POST
- **URL**: `/v3/company/{realmID}/creditmemo?operation=delete`

This operation deletes the creditmemo object specified in the request body. Include a minimum of CreditMemo.Id and CreditMemo.SyncToken in the request body.

**Request Body**:
```json
{
  "SyncToken": "0",
  "Id": "73"
}
```

**Response**:
```json
{
  "CreditMemo": {
    "status": "Deleted",
    "domain": "QBO",
    "Id": "73"
  },
  "time": "2015-05-26T13:53:33.118-07:00"
}
```

### Get a credit memo as PDF

- **Method**: GET
- **URL**: `/v3/company/{realmID}/creditmemo/{creditmemoId}/pdf`

### Query a credit memo

- **Method**: GET
- **URL**: `/v3/company/{realmID}/query?query={selectStatement}`

**Sample Query**: `Select * from CreditMemo where TxnDate > '2014-04-15'`

**Response**:
```json
{
  "QueryResponse": {
    "startPosition": 1,
    "CreditMemo": [
      {
        "TxnDate": "2014-09-02",
        "domain": "QBO",
        "PrintStatus": "NeedToPrint",
        "TotalAmt": 100.0,
        "RemainingCredit": 0,
        "Line": [
          {
            "Description": "Pest Control Services",
            "DetailType": "SalesItemLineDetail",
            "SalesItemLineDetail": {
              "TaxCodeRef": {
                "value": "NON"
              },
              "Qty": 1,
              "UnitPrice": 100,
              "ItemRef": {
                "name": "Pest Control",
                "value": "10"
              }
            },
            "LineNum": 1,
            "Amount": 100.0,
            "Id": "1"
          },
          {
            "DetailType": "SubTotalLineDetail",
            "Amount": 100.0,
            "SubTotalLineDetail": {}
          }
        ],
        "ApplyTaxAfterDiscount": false,
        "DocNumber": "1026",
        "sparse": false,
        "CustomerMemo": {
          "value": "Updated customer memo."
        },
        "ProjectRef": {
          "value": "39298034"
        },
        "Balance": 0,
        "CustomerRef": {
          "name": "Amy's Bird Sanctuary",
          "value": "1"
        },
        "TxnTaxDetail": {
          "TotalTax": 0
        },
        "SyncToken": "3",
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
          "Id": "108",
          "Line1": "4581 Finch St."
        },
        "EmailStatus": "NotSet",
        "BillAddr": {
          "Line4": "Bayshore, CA 94326",
          "Line3": "4581 Finch St.",
          "Id": "79",
          "Line1": "Amy Lauterbach",
          "Line2": "Amy's Bird Sanctuary"
        },
        "MetaData": {
          "CreateTime": "2014-09-18T12:51:27-07:00",
          "LastUpdatedTime": "2015-07-01T09:16:28-07:00"
        },
        "BillEmail": {
          "Address": "Birds@Intuit.com"
        },
        "Id": "73"
      },
      {
        "DocNumber": "1039",
        "SyncToken": "0",
        "domain": "QBO",
        "Balance": 50.0,
        "BillAddr": {
          "City": "Half Moon Bay",
          "Line1": "65 Ocean Dr.",
          "PostalCode": "94213",
          "Lat": "37.4300318",
          "Long": "-122.4336537",
          "CountrySubDivisionCode": "CA",
          "Id": "4"
        },
        "TxnDate": "2015-01-13",
        "TotalAmt": 50.0,
        "CustomerRef": {
          "name": "Cool Cars",
          "value": "3"
        },
        "RemainingCredit": 50.0,
        "PrintStatus": "NeedToPrint",
        "ProjectRef": {
          "value": "36387497"
        },
        "EmailStatus": "NotSet",
        "sparse": false,
        "Line": [
          {
            "LineNum": 1,
            "Amount": 50.0,
            "SalesItemLineDetail": {
              "TaxCodeRef": {
                "value": "NON"
              },
              "ItemRef": {
                "name": "Concrete",
                "value": "3"
              }
            },
            "Id": "1",
            "DetailType": "SalesItemLineDetail"
          },
          {
            "DetailType": "SubTotalLineDetail",
            "Amount": 50.0,
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
        "Id": "158",
        "TxnTaxDetail": {
          "TotalTax": 0
        },
        "MetaData": {
          "CreateTime": "2015-01-13T10:32:35-08:00",
          "LastUpdatedTime": "2015-01-13T10:32:35-08:00"
        }
      }
    ],
    "maxResults": 2
  },
  "time": "2015-07-23T09:13:36.246-07:00"
}
```

### Read a credit memo

- **Method**: GET
- **URL**: `/v3/company/{realmID}/creditmemo/{creditmemoId}`

Retrieves the details of a creditmemo that has been previously created.

**Response**:
```json
{
  "CreditMemo": {
    "TxnDate": "2014-09-02",
    "domain": "QBO",
    "PrintStatus": "NeedToPrint",
    "TotalAmt": 100.0,
    "RemainingCredit": 0,
    "Line": [
      {
        "Description": "Pest Control Services",
        "DetailType": "SalesItemLineDetail",
        "SalesItemLineDetail": {
          "TaxCodeRef": {
            "value": "NON"
          },
          "Qty": 1,
          "UnitPrice": 100,
          "ItemRef": {
            "name": "Pest Control",
            "value": "10"
          }
        },
        "LineNum": 1,
        "Amount": 100.0,
        "Id": "1"
      },
      {
        "DetailType": "SubTotalLineDetail",
        "Amount": 100.0,
        "SubTotalLineDetail": {}
      }
    ],
    "ApplyTaxAfterDiscount": false,
    "DocNumber": "1026",
    "sparse": false,
    "CustomerMemo": {
      "value": "Updated customer memo."
    },
    "ProjectRef": {
      "value": "39298034"
    },
    "Balance": 0,
    "CustomerRef": {
      "name": "Amy's Bird Sanctuary",
      "value": "1"
    },
    "TxnTaxDetail": {
      "TotalTax": 0
    },
    "SyncToken": "3",
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
      "Id": "108",
      "Line1": "4581 Finch St."
    },
    "EmailStatus": "NotSet",
    "BillAddr": {
      "Line4": "Bayshore, CA 94326",
      "Line3": "4581 Finch St.",
      "Id": "79",
      "Line1": "Amy Lauterbach",
      "Line2": "Amy's Bird Sanctuary"
    },
    "MetaData": {
      "CreateTime": "2014-09-18T12:51:27-07:00",
      "LastUpdatedTime": "2015-07-01T09:16:28-07:00"
    },
    "BillEmail": {
      "Address": "Birds@Intuit.com"
    },
    "Id": "73"
  },
  "time": "2015-07-23T09:10:45.624-07:00"
}
```

### Send a credit memo

- **Method**: POST
- **URL**: `/v3/company/{realmID}/creditmemo/{creditmemoId}/send`

• The CreditMemo.EmailStatus parameter is set to EmailSent. • The CreditMemo.DeliveryInfo element is populated with sending information<./li>

**Response**:
```json
{
  "CreditMemo": {
    "TxnDate": "2014-09-02",
    "domain": "QBO",
    "PrintStatus": "NeedToPrint",
    "DeliveryInfo": {
      "DeliveryType": "Email",
      "DeliveryTime": "2019-09-19T10:43:46-07:00"
    },
    "TotalAmt": 100.0,
    "RemainingCredit": 0,
    "Line": [
      {
        "Description": "Pest Control Services",
        "DetailType": "SalesItemLineDetail",
        "SalesItemLineDetail": {
          "TaxCodeRef": {
            "value": "NON"
          },
          "Qty": 1,
          "UnitPrice": 100,
          "ItemRef": {
            "name": "Pest Control",
            "value": "10"
          }
        },
        "LineNum": 1,
        "Amount": 100.0,
        "Id": "1"
      },
      {
        "DetailType": "SubTotalLineDetail",
        "Amount": 100.0,
        "SubTotalLineDetail": {}
      }
    ],
    "ApplyTaxAfterDiscount": false,
    "DocNumber": "1026",
    "sparse": false,
    "CustomerMemo": {
      "value": "Updated customer memo."
    },
    "ProjectRef": {
      "value": "39298034"
    },
    "Balance": 0,
    "CustomerRef": {
      "name": "Amy's Bird Sanctuary",
      "value": "1"
    },
    "TxnTaxDetail": {
      "TotalTax": 0
    },
    "SyncToken": "3",
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
      "Id": "108",
      "Line1": "4581 Finch St."
    },
    "EmailStatus": "EmailSent",
    "BillAddr": {
      "Line4": "Bayshore, CA 94326",
      "Line3": "4581 Finch St.",
      "Id": "79",
      "Line1": "Amy Lauterbach",
      "Line2": "Amy's Bird Sanctuary"
    },
    "MetaData": {
      "CreateTime": "2014-09-18T12:51:27-07:00",
      "LastUpdatedTime": "2019-09-19T10:43:46-07:00"
    },
    "BillEmail": {
      "Address": "Birds@Intuit.com"
    },
    "Id": "73"
  },
  "time": "2019-09-19T10:43:46-07:00"
}
```

### Full update a credit memo

- **Method**: POST
- **URL**: `/v3/company/{realmID}/creditmemo`

Use this operation to update any of the writable fields of an existing creditmemo object. The request body must include all writable fields of the existing object as returned in a read response. Writable fields omitted from the request body are set to NULL. The ID of the object to update is specifie...

**Request Body**:
```json
{
  "TxnDate": "2014-09-02",
  "domain": "QBO",
  "PrintStatus": "NeedToPrint",
  "TotalAmt": 100.0,
  "RemainingCredit": 0,
  "Line": [
    {
      "Description": "Pest Control Services",
      "DetailType": "SalesItemLineDetail",
      "SalesItemLineDetail": {
        "TaxCodeRef": {
          "value": "NON"
        },
        "Qty": 1,
        "UnitPrice": 100,
        "ItemRef": {
          "name": "Pest Control",
          "value": "10"
        }
      },
      "LineNum": 1,
      "Amount": 100.0,
      "Id": "1"
    },
    {
      "DetailType": "SubTotalLineDetail",
      "Amount": 100.0,
      "SubTotalLineDetail": {}
    }
  ],
  "ApplyTaxAfterDiscount": false,
  "DocNumber": "1026",
  "sparse": false,
  "CustomerMemo": {
    "value": "Another memo update."
  },
  "ProjectRef": {
    "value": "39298045"
  },
  "Balance": 0,
  "CustomerRef": {
    "name": "Amy's Bird Sanctuary",
    "value": "1"
  },
  "TxnTaxDetail": {
    "TotalTax": 0
  },
  "SyncToken": "4",
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
    "Id": "108",
    "Line1": "4581 Finch St."
  },
  "EmailStatus": "NotSet",
  "BillAddr": {
    "Line4": "Bayshore, CA 94326",
    "Line3": "4581 Finch St.",
    "Id": "79",
    "Line1": "Amy Lauterbach",
    "Line2": "Amy's Bird Sanctuary"
  },
  "MetaData": {
    "CreateTime": "2014-09-18T12:51:27-07:00",
    "LastUpdatedTime": "2015-07-01T09:16:28-07:00"
  },
  "BillEmail": {
    "Address": "Birds@Intuit.com"
  },
  "Id": "73"
}
```

**Response**:
```json
{
  "CreditMemo": {
    "TxnDate": "2014-09-02",
    "domain": "QBO",
    "PrintStatus": "NeedToPrint",
    "TotalAmt": 100.0,
    "RemainingCredit": 0,
    "Line": [
      {
        "Description": "Pest Control Services",
        "DetailType": "SalesItemLineDetail",
        "SalesItemLineDetail": {
          "TaxCodeRef": {
            "value": "NON"
          },
          "Qty": 1,
          "UnitPrice": 100,
          "ItemRef": {
            "name": "Pest Control",
            "value": "10"
          }
        },
        "LineNum": 1,
        "Amount": 100.0,
        "Id": "1"
      },
      {
        "DetailType": "SubTotalLineDetail",
        "Amount": 100.0,
        "SubTotalLineDetail": {}
      }
    ],
    "ApplyTaxAfterDiscount": false,
    "DocNumber": "1026",
    "sparse": false,
    "CustomerMemo": {
      "value": "Another memo update."
    },
    "ProjectRef": {
      "value": "39298045"
    },
    "Balance": 0,
    "CustomerRef": {
      "name": "Amy's Bird Sanctuary",
      "value": "1"
    },
    "TxnTaxDetail": {
      "TotalTax": 0
    },
    "SyncToken": "5",
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
      "Id": "108",
      "Line1": "4581 Finch St."
    },
    "EmailStatus": "NotSet",
    "BillAddr": {
      "Line4": "Bayshore, CA 94326",
      "Line3": "4581 Finch St.",
      "Id": "79",
      "Line1": "Amy Lauterbach",
      "Line2": "Amy's Bird Sanctuary"
    },
    "MetaData": {
      "CreateTime": "2014-09-18T12:51:27-07:00",
      "LastUpdatedTime": "2015-07-23T09:23:52-07:00"
    },
    "BillEmail": {
      "Address": "Birds@Intuit.com"
    },
    "Id": "73"
  },
  "time": "2015-07-23T09:23:52.115-07:00"
}
```
