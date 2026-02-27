# Estimate

The Estimate represents a proposal for a financial transaction from a business to a customer for goods or services proposed to be sold, including proposed pricing.

## Attributes

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| Id | String | Required for update | Unique identifier for this object. Sort order is ASC by default. |
| CustomerRef | ReferenceType | Required | Reference to a customer or job. Query the Customer name list resource to determine the appropriate Customer object for this reference. Use Customer.Id and Customer. |
| SyncToken | String | Required for update | Version number of the object. It is used to lock an object for use by one app at a time. As soon as an application modifies an object, its SyncToken is incremented. |
| ShipFromAddr | PhysicalAddress | Conditionally required | Identifies the address where the goods are shipped from. For transactions without shipping, it represents the address where the sale took place. If automated sales tax is enabled (Preferences. |
| CurrencyRef | CurrencyRefType | Conditionally required | Reference to the currency in which all amounts on the associated transaction are expressed. This must be defined if multicurrency is enabled for the company. |
| GlobalTaxCalculation | String | Conditionally required | Method in which tax is applied. Allowed values are: TaxExcluded, TaxInclusive, and NotApplicable. |
| ProjectRef | ReferenceType | Conditionally required | Reference to the Project ID associated with this transaction. When ProjectRef is passed in the US region for QuickBooks Online Advanced and Intuit Enterprise Suite SKUs, a project estimate is created ... |
| BillEmail | EmailAddress | Conditionally required | Identifies the e-mail address where the estimate is sent. If EmailStatus=NeedToSend, BillEmailis a required input. |
| TxnDate | Date | Optional | The date entered by the user when this transaction occurred. For posting transactions, this is the posting date that affects the financial statements. |
| ShipDate | Date | Optional | Date for delivery of goods or services. |
| ClassRef | ReferenceType | Optional | Reference to the Class associated with the transaction. Available if Preferences.AccountingInfoPrefs.ClassTrackingPerTxn is set to true. |
| PrintStatus | String | Optional | Printing status of the invoice. Valid values: NotSet, NeedToPrint, PrintComplete . |
| CustomField | CustomField | Optional | One of, up to three custom fields for the transaction. Available for custom fields so configured for the company. Check Preferences.SalesFormsPrefs.CustomField and Preferences. |
| SalesTermRef | ReferenceType | Optional | Reference to the sales term associated with the transaction. Query the Term name list resource to determine the appropriate Term object for this reference. Use Term.Id and Term. |
| TxnStatus | String | Optional | One of the following status settings: Accepted, Closed, Pending, Rejected, Converted |
| LinkedTxn | LinkedTxn | Optional | Zero or more Invoice objects related to this transaction. Use LinkedTxn.TxnId as the ID in a separate Invoice read request to retrieve details of the linked object. |
| AcceptedDate | Date | Optional | Date estimate was accepted. |
| ExpirationDate | Date | Optional | Date by which estimate must be accepted before invalidation. |
| TransactionLocationType | String | Optional | The account location. Valid values for France: WithinFrance, FranceOverseas, OutsideFranceWithEU, OutsideEU. |
| DueDate | Date | Optional | Date when the payment of the transaction is due. If date is not provided, the number of days specified in SalesTermRef added the transaction date will be used. |
| MetaData | ModificationMetaData | Read-only | Descriptive information about the object. The MetaData values are set by Data Services and are  for all applications. |
| DocNumber | String | Optional | Reference number for the transaction. |
| PrivateNote | String | Optional | User entered, organization-private note about the transaction. This note does not appear on the invoice to the customer. |
| Line | Line | Optional | Individual line items of a transaction. Valid Line types include: SalesItemLine, GroupLine, DescriptionOnlyLine (also used for inline Subtotal lines), DiscountLine and SubTotalLine (used... |
| CustomerMemo | MemoRef | Optional | User-entered message to the customer; this message is visible to end user on their transactions. |
| EmailStatus | String | Optional | Email status of the invoice. Valid values: NotSet, NeedToSend, EmailSent |
| TxnTaxDetail | TxnTaxDetail | Optional | This data type provides information for taxes charged on the transaction as a whole. |
| AcceptedBy | String | Optional | Name of customer who accepted the estimate. |
| ExchangeRate | Decimal | Optional | The number of home currency units it takes to equal one unit of currency specified by CurrencyRef. Applicable if multicurrency is enabled for the company. |
| ShipAddr | PhysicalAddress | Optional | Identifies the address where the goods must be shipped. |
| DepartmentRef | ReferenceType | Optional | A reference to a Department object specifying the location of the transaction. Available if Preferences.AccountingInfoPrefs.TrackDepartments is set to true. |
| ShipMethodRef | ReferenceType | Optional | Reference to the ShipMethod associated with the transaction. There is no shipping method list. Reference resolves to a string. |
| BillAddr | PhysicalAddress | Optional | Bill-to address of the Invoice. If BillAddris not specified, and a default Customer:BillingAddr is specified in QuickBooks for this customer, the default bill-to address is us... |
| ApplyTaxAfterDiscount | Boolean | Optional | If false or null, calculate the sales tax first, and then apply the discount. If true, subtract the discount first and then calculate the sales tax. |
| TotalAmt | BigDecimal | Read-only | Indicates the total amount of the transaction. This includes the total of all the charges, allowances, and taxes. |
| RecurDataRef | ReferenceType | Read-only | A reference to the Recurring Transaction. It captures what recurring transaction template the Estimate was created from. |
| TaxExemptionRef | ReferenceType | Read-only | Reference to the TaxExepmtion ID associated with this object. Available for companies that have |
| HomeTotalAmt | Decimal | Read-only | Total amount of the transaction in the home currency. Includes the total of all the charges, allowances and taxes. Calculated by QuickBooks business logic. |
| FreeFormAddress | Boolean | Optional | Denotes how ShipAddr is stored: formatted or unformatted. The value of this flag is  based on format of shipping address at object create time. |

## Sample Object

```json
{
  "Estimate": {
    "DocNumber": "1001",
    "SyncToken": "0",
    "domain": "QBO",
    "TxnStatus": "Pending",
    "BillEmail": {
      "Address": "Cool_Cars@intuit.com"
    },
    "TxnDate": "2015-03-26",
    "TotalAmt": 31.5,
    "CustomerRef": {
      "name": "Cool Cars",
      "value": "3"
    },
    "CustomerMemo": {
      "value": "Thank you for your business and have a great day!"
    },
    "ShipAddr": {
      "CountrySubDivisionCode": "CA",
      "City": "Half Moon Bay",
      "PostalCode": "94213",
      "Id": "104",
      "Line1": "65 Ocean Dr."
    },
    "ProjectRef": {
      "value": "39298034"
    },
    "PrintStatus": "NeedToPrint",
    "BillAddr": {
      "CountrySubDivisionCode": "CA",
      "City": "Half Moon Bay",
      "PostalCode": "94213",
      "Id": "103",
      "Line1": "65 Ocean Dr."
    },
    "sparse": false,
    "EmailStatus": "NotSet",
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
      },
      {
        "DetailType": "DiscountLineDetail",
        "Amount": 3.5,
        "DiscountLineDetail": {
          "DiscountAccountRef": {
            "name": "Discounts given",
            "value": "86"
          },
          "PercentBased": true,
          "DiscountPercent": 10
        }
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
    "Id": "177",
    "TxnTaxDetail": {
      "TotalTax": 0
    },
    "MetaData": {
      "CreateTime": "2015-03-26T13:25:05-07:00",
      "LastUpdatedTime": "2015-03-26T13:25:05-07:00"
    }
  },
  "time": "2015-03-26T13:25:05.473-07:00"
}
```

## Operations

### Create an estimate

- **Method**: POST
- **URL**: `/v3/company/{realmID}/estimate`

• An Estimate must have at least one line that describes an item. • An Estimate must have a reference to a customer.

**Request Body**:
```json
{
  "TotalAmt": 31.5,
  "BillEmail": {
    "Address": "Cool_Cars@intuit.com"
  },
  "CustomerMemo": {
    "value": "Thank you for your business and have a great day!"
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
  "PrintStatus": "NeedToPrint",
  "EmailStatus": "NotSet",
  "BillAddr": {
    "City": "Half Moon Bay",
    "Line1": "65 Ocean Dr.",
    "PostalCode": "94213",
    "Lat": "37.4300318",
    "Long": "-122.4336537",
    "CountrySubDivisionCode": "CA",
    "Id": "4"
  },
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
    },
    {
      "DetailType": "DiscountLineDetail",
      "Amount": 3.5,
      "DiscountLineDetail": {
        "DiscountAccountRef": {
          "name": "Discounts given",
          "value": "86"
        },
        "PercentBased": true,
        "DiscountPercent": 10
      }
    }
  ],
  "CustomerRef": {
    "name": "Cool Cars",
    "value": "3"
  },
  "TxnTaxDetail": {
    "TotalTax": 0
  },
  "ApplyTaxAfterDiscount": false
}
```

**Response**:
```json
{
  "Estimate": {
    "DocNumber": "1001",
    "SyncToken": "0",
    "domain": "QBO",
    "TxnStatus": "Pending",
    "BillEmail": {
      "Address": "Cool_Cars@intuit.com"
    },
    "TxnDate": "2015-03-26",
    "TotalAmt": 31.5,
    "CustomerRef": {
      "name": "Cool Cars",
      "value": "3"
    },
    "CustomerMemo": {
      "value": "Thank you for your business and have a great day!"
    },
    "ShipAddr": {
      "CountrySubDivisionCode": "CA",
      "City": "Half Moon Bay",
      "PostalCode": "94213",
      "Id": "104",
      "Line1": "65 Ocean Dr."
    },
    "ProjectRef": {
      "value": "39298034"
    },
    "PrintStatus": "NeedToPrint",
    "BillAddr": {
      "CountrySubDivisionCode": "CA",
      "City": "Half Moon Bay",
      "PostalCode": "94213",
      "Id": "103",
      "Line1": "65 Ocean Dr."
    },
    "sparse": false,
    "EmailStatus": "NotSet",
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
      },
      {
        "DetailType": "DiscountLineDetail",
        "Amount": 3.5,
        "DiscountLineDetail": {
          "DiscountAccountRef": {
            "name": "Discounts given",
            "value": "86"
          },
          "PercentBased": true,
          "DiscountPercent": 10
        }
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
    "Id": "177",
    "TxnTaxDetail": {
      "TotalTax": 0
    },
    "MetaData": {
      "CreateTime": "2015-03-26T13:25:05-07:00",
      "LastUpdatedTime": "2015-03-26T13:25:05-07:00"
    }
  },
  "time": "2015-03-26T13:25:05.473-07:00"
}
```

### Delete an estimate

- **Method**: POST
- **URL**: `/v3/company/{realmID}/estimate?operation=delete`

This operation deletes the estimate object specified in the request body. Include a minimum of Estimate.Id and Estimate.SyncToken in the request body.

**Request Body**:
```json
{
  "SyncToken": "3",
  "Id": "96"
}
```

**Response**:
```json
{
  "Estimate": {
    "status": "Deleted",
    "domain": "QBO",
    "Id": "96"
  },
  "time": "2013-03-14T15:05:14.981-07:00"
}
```

### Get an estimate as PDF

- **Method**: GET
- **URL**: `/v3/company/{realmID}/estimate/{estimateId}/pdf`

### Query an estimate

- **Method**: GET
- **URL**: `/v3/company/{realmID}/query?query={selectStatement}`

**Sample Query**: `select * from estimate where TxnDate < '2014-09-15'`

**Response**:
```json
{
  "QueryResponse": {
    "startPosition": 1,
    "Estimate": [
      {
        "TxnDate": "2014-09-07",
        "domain": "QBO",
        "PrintStatus": "NeedToPrint",
        "TxnStatus": "Closed",
        "TotalAmt": 582.5,
        "Line": [
          {
            "Description": "Rock Fountain",
            "DetailType": "SalesItemLineDetail",
            "SalesItemLineDetail": {
              "TaxCodeRef": {
                "value": "NON"
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
            "Description": "Custom Design",
            "DetailType": "SalesItemLineDetail",
            "SalesItemLineDetail": {
              "TaxCodeRef": {
                "value": "NON"
              },
              "Qty": 3.5,
              "UnitPrice": 75,
              "ItemRef": {
                "name": "Design",
                "value": "4"
              }
            },
            "LineNum": 2,
            "Amount": 262.5,
            "Id": "2"
          },
          {
            "Description": "Fountain Pump",
            "DetailType": "SalesItemLineDetail",
            "SalesItemLineDetail": {
              "TaxCodeRef": {
                "value": "NON"
              },
              "Qty": 2,
              "UnitPrice": 22.5,
              "ItemRef": {
                "name": "Pump",
                "value": "11"
              }
            },
            "LineNum": 3,
            "Amount": 45.0,
            "Id": "3"
          },
          {
            "DetailType": "SubTotalLineDetail",
            "Amount": 582.5,
            "SubTotalLineDetail": {}
          }
        ],
        "ApplyTaxAfterDiscount": false,
        "DocNumber": "1001",
        "sparse": false,
        "CustomerMemo": {
          "value": "Thank you for your business and have a great day!"
        },
        "ProjectRef": {
          "value": "39298035"
        },
        "CustomerRef": {
          "name": "Geeta Kalapatapu",
          "value": "10"
        },
        "TxnTaxDetail": {
          "TotalTax": 0
        },
        "SyncToken": "0",
        "LinkedTxn": [
          {
            "TxnId": "103",
            "TxnType": "Invoice"
          }
        ],
        "CustomField": [
          {
            "DefinitionId": "1",
            "Type": "StringType",
            "Name": "Crew #"
          }
        ],
        "ShipAddr": {
          "City": "Middlefield",
          "Line1": "1987 Main St.",
          "PostalCode": "94303",
          "Lat": "37.445013",
          "Long": "-122.1391443",
          "CountrySubDivisionCode": "CA",
          "Id": "10"
        },
        "EmailStatus": "NotSet",
        "BillAddr": {
          "Line3": "Middlefield, CA 94303",
          "Line2": "1987 Main St.",
          "Long": "-122.1178261",
          "Line1": "Geeta Kalapatapu",
          "Lat": "37.4530553",
          "Id": "59"
        },
        "MetaData": {
          "CreateTime": "2014-09-17T11:20:06-07:00",
          "LastUpdatedTime": "2014-09-18T13:41:59-07:00"
        },
        "BillEmail": {
          "Address": "Geeta@Kalapatapu.com"
        },
        "Id": "41"
      }
    ],
    "maxResults": 1
  },
  "time": "2015-07-24T14:00:04.902-07:00"
}
```

### Read an estimate

- **Method**: GET
- **URL**: `/v3/company/{realmID}/estimate/{estimateId}`

Retrieves the details of an estimate that has been previously created.

**Response**:
```json
{
  "Estimate": {
    "DocNumber": "1001",
    "SyncToken": "0",
    "domain": "QBO",
    "TxnStatus": "Pending",
    "BillEmail": {
      "Address": "Cool_Cars@intuit.com"
    },
    "TxnDate": "2015-03-26",
    "TotalAmt": 31.5,
    "CustomerRef": {
      "name": "Cool Cars",
      "value": "3"
    },
    "CustomerMemo": {
      "value": "Thank you for your business and have a great day!"
    },
    "ShipAddr": {
      "CountrySubDivisionCode": "CA",
      "City": "Half Moon Bay",
      "PostalCode": "94213",
      "Id": "104",
      "Line1": "65 Ocean Dr."
    },
    "ProjectRef": {
      "value": "39298034"
    },
    "PrintStatus": "NeedToPrint",
    "BillAddr": {
      "CountrySubDivisionCode": "CA",
      "City": "Half Moon Bay",
      "PostalCode": "94213",
      "Id": "103",
      "Line1": "65 Ocean Dr."
    },
    "sparse": false,
    "EmailStatus": "NotSet",
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
      },
      {
        "DetailType": "DiscountLineDetail",
        "Amount": 3.5,
        "DiscountLineDetail": {
          "DiscountAccountRef": {
            "name": "Discounts given",
            "value": "86"
          },
          "PercentBased": true,
          "DiscountPercent": 10
        }
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
    "Id": "177",
    "TxnTaxDetail": {
      "TotalTax": 0
    },
    "MetaData": {
      "CreateTime": "2015-03-26T13:25:05-07:00",
      "LastUpdatedTime": "2015-03-26T13:25:05-07:00"
    }
  },
  "time": "2015-03-26T13:25:05.473-07:00"
}
```

### Send an estimate

- **Method**: POST
- **URL**: `/v3/company/{realmID}/estimate/{estimateId}/send`

• The Estimate.EmailStatus parameter is set to EmailSent. • The Estimate.DeliveryInfo element is populated with sending information.

**Response**:
```json
{
  "Estimate": {
    "TxnDate": "2014-11-06",
    "domain": "QBO",
    "PrintStatus": "NeedToPrint",
    "DeliveryInfo": {
      "DeliveryType": "Email",
      "DeliveryTime": "2015-03-26T14:05:31-07:00"
    },
    "TxnStatus": "Closed",
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
    "ApplyTaxAfterDiscount": false,
    "DocNumber": "1001",
    "sparse": false,
    "CustomerMemo": {
      "value": "Thank you for your business and have a great day!"
    },
    "ProjectRef": {
      "value": "39298036"
    },
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
        "TxnId": "130",
        "TxnType": "Invoice"
      }
    ],
    "CustomField": [
      {
        "DefinitionId": "1",
        "Type": "StringType",
        "Name": "Crew #"
      }
    ],
    "ShipAddr": {
      "City": "Middlefield",
      "Line1": "5647 Cypress Hill Ave.",
      "PostalCode": "94303",
      "Lat": "37.4238562",
      "Long": "-122.1141681",
      "CountrySubDivisionCode": "CA",
      "Id": "25"
    },
    "EmailStatus": "EmailSent",
    "BillAddr": {
      "Lat": "37.4238562",
      "Long": "-122.1141681",
      "Id": "86",
      "Line1": "Russ Sonnenschein\\nSonnenschein Family Store\\n5647 Cypress Hill Ave.\\nMiddlefield, CA 94303"
    },
    "MetaData": {
      "CreateTime": "2014-11-08T13:37:55-08:00",
      "LastUpdatedTime": "2015-03-26T14:05:31-07:00"
    },
    "BillEmail": {
      "Address": "Familiystore@intuit.com"
    },
    "Id": "100"
  },
  "time": "2015-03-26T14:05:31.330-07:00"
}
```

### Sparse update an estimate

- **Method**: POST
- **URL**: `/v3/company/{realmID}/estimate`

Sparse updating provides the ability to update a subset of properties for a given object; only elements specified in the request are updated. Missing elements are left untouched. The ID of the object to update is specified in the request body.

**Request Body**:
```json
{
  "SyncToken": "3",
  "domain": "QBO",
  "CustomerMemo": {
    "value": "An updated memo via full update the second time."
  },
  "sparse": true,
  "Id": "41",
  "MetaData": {
    "CreateTime": "2014-09-17T11:20:06-07:00",
    "LastUpdatedTime": "2015-07-24T14:08:04-07:00"
  }
}
```

**Response**:
```json
{
  "Estimate": {
    "TxnDate": "2014-09-07",
    "domain": "QBO",
    "PrintStatus": "NeedToPrint",
    "TxnStatus": "Closed",
    "TotalAmt": 582.5,
    "Line": [
      {
        "Description": "Rock Fountain",
        "DetailType": "SalesItemLineDetail",
        "SalesItemLineDetail": {
          "TaxCodeRef": {
            "value": "NON"
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
        "Description": "Custom Design",
        "DetailType": "SalesItemLineDetail",
        "SalesItemLineDetail": {
          "TaxCodeRef": {
            "value": "NON"
          },
          "Qty": 3.5,
          "UnitPrice": 75,
          "ItemRef": {
            "name": "Design",
            "value": "4"
          }
        },
        "LineNum": 2,
        "Amount": 262.5,
        "Id": "2"
      },
      {
        "Description": "Fountain Pump",
        "DetailType": "SalesItemLineDetail",
        "SalesItemLineDetail": {
          "TaxCodeRef": {
            "value": "NON"
          },
          "Qty": 2,
          "UnitPrice": 22.5,
          "ItemRef": {
            "name": "Pump",
            "value": "11"
          }
        },
        "LineNum": 3,
        "Amount": 45.0,
        "Id": "3"
      },
      {
        "DetailType": "SubTotalLineDetail",
        "Amount": 582.5,
        "SubTotalLineDetail": {}
      }
    ],
    "ApplyTaxAfterDiscount": false,
    "DocNumber": "1001",
    "sparse": false,
    "CustomerMemo": {
      "value": "An updated memo via full update the second time."
    },
    "ProjectRef": {
      "value": "39298033"
    },
    "CustomerRef": {
      "name": "Geeta Kalapatapu",
      "value": "10"
    },
    "TxnTaxDetail": {
      "TotalTax": 0
    },
    "SyncToken": "4",
    "LinkedTxn": [
      {
        "TxnId": "103",
        "TxnType": "Invoice"
      }
    ],
    "CustomField": [
      {
        "DefinitionId": "1",
        "Type": "StringType",
        "Name": "Crew #"
      }
    ],
    "ShipAddr": {
      "CountrySubDivisionCode": "CA",
      "City": "Middlefield",
      "PostalCode": "94303",
      "Id": "119",
      "Line1": "1987 Main St."
    },
    "EmailStatus": "NotSet",
    "BillAddr": {
      "Line3": "Middlefield, CA 94303",
      "Id": "59",
      "Line1": "Geeta Kalapatapu",
      "Line2": "1987 Main St."
    },
    "MetaData": {
      "CreateTime": "2014-09-17T11:20:06-07:00",
      "LastUpdatedTime": "2015-07-24T14:17:23-07:00"
    },
    "BillEmail": {
      "Address": "Geeta@Kalapatapu.com"
    },
    "Id": "41"
  },
  "time": "2015-07-24T14:17:23.307-07:00"
}
```

### Full update an estimate

- **Method**: POST
- **URL**: `/v3/company/{realmID}/estimate`

Use this operation to update any of the writable fields of an existing estimate object. The request body must include all writable fields of the existing object as returned in a read response. Writable fields omitted from the request body are set to NULL. The ID of the object to update is specified ...

**Request Body**:
```json
{
  "TxnDate": "2014-09-07",
  "domain": "QBO",
  "PrintStatus": "NeedToPrint",
  "TxnStatus": "Closed",
  "TotalAmt": 582.5,
  "Line": [
    {
      "Description": "Rock Fountain",
      "DetailType": "SalesItemLineDetail",
      "SalesItemLineDetail": {
        "TaxCodeRef": {
          "value": "NON"
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
      "Description": "Custom Design",
      "DetailType": "SalesItemLineDetail",
      "SalesItemLineDetail": {
        "TaxCodeRef": {
          "value": "NON"
        },
        "Qty": 3.5,
        "UnitPrice": 75,
        "ItemRef": {
          "name": "Design",
          "value": "4"
        }
      },
      "LineNum": 2,
      "Amount": 262.5,
      "Id": "2"
    },
    {
      "Description": "Fountain Pump",
      "DetailType": "SalesItemLineDetail",
      "SalesItemLineDetail": {
        "TaxCodeRef": {
          "value": "NON"
        },
        "Qty": 2,
        "UnitPrice": 22.5,
        "ItemRef": {
          "name": "Pump",
          "value": "11"
        }
      },
      "LineNum": 3,
      "Amount": 45.0,
      "Id": "3"
    },
    {
      "DetailType": "SubTotalLineDetail",
      "Amount": 582.5,
      "SubTotalLineDetail": {}
    }
  ],
  "ApplyTaxAfterDiscount": false,
  "DocNumber": "1001",
  "sparse": false,
  "CustomerMemo": {
    "value": "An updated memo via full update."
  },
  "ProjectRef": {
    "value": "39298033"
  },
  "CustomerRef": {
    "name": "Geeta Kalapatapu",
    "value": "10"
  },
  "TxnTaxDetail": {
    "TotalTax": 0
  },
  "SyncToken": "2",
  "LinkedTxn": [
    {
      "TxnId": "103",
      "TxnType": "Invoice"
    }
  ],
  "CustomField": [
    {
      "DefinitionId": "1",
      "Type": "StringType",
      "Name": "Crew #"
    }
  ],
  "ShipAddr": {
    "CountrySubDivisionCode": "CA",
    "City": "Middlefield",
    "PostalCode": "94303",
    "Id": "119",
    "Line1": "1987 Main St."
  },
  "EmailStatus": "NotSet",
  "BillAddr": {
    "Line3": "Middlefield, CA 94303",
    "Id": "59",
    "Line1": "Geeta Kalapatapu",
    "Line2": "1987 Main St."
  },
  "MetaData": {
    "CreateTime": "2014-09-17T11:20:06-07:00",
    "LastUpdatedTime": "2015-07-24T14:08:04-07:00"
  },
  "BillEmail": {
    "Address": "Geeta@Kalapatapu.com"
  },
  "Id": "41"
}
```

**Response**:
```json
{
  "Estimate": {
    "TxnDate": "2014-09-07",
    "domain": "QBO",
    "PrintStatus": "NeedToPrint",
    "TxnStatus": "Closed",
    "TotalAmt": 582.5,
    "Line": [
      {
        "Description": "Rock Fountain",
        "DetailType": "SalesItemLineDetail",
        "SalesItemLineDetail": {
          "TaxCodeRef": {
            "value": "NON"
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
        "Description": "Custom Design",
        "DetailType": "SalesItemLineDetail",
        "SalesItemLineDetail": {
          "TaxCodeRef": {
            "value": "NON"
          },
          "Qty": 3.5,
          "UnitPrice": 75,
          "ItemRef": {
            "name": "Design",
            "value": "4"
          }
        },
        "LineNum": 2,
        "Amount": 262.5,
        "Id": "2"
      },
      {
        "Description": "Fountain Pump",
        "DetailType": "SalesItemLineDetail",
        "SalesItemLineDetail": {
          "TaxCodeRef": {
            "value": "NON"
          },
          "Qty": 2,
          "UnitPrice": 22.5,
          "ItemRef": {
            "name": "Pump",
            "value": "11"
          }
        },
        "LineNum": 3,
        "Amount": 45.0,
        "Id": "3"
      },
      {
        "DetailType": "SubTotalLineDetail",
        "Amount": 582.5,
        "SubTotalLineDetail": {}
      }
    ],
    "ApplyTaxAfterDiscount": false,
    "DocNumber": "1001",
    "sparse": false,
    "CustomerMemo": {
      "value": "An updated memo via full update."
    },
    "ProjectRef": {
      "value": "39298033"
    },
    "CustomerRef": {
      "name": "Geeta Kalapatapu",
      "value": "10"
    },
    "TxnTaxDetail": {
      "TotalTax": 0
    },
    "SyncToken": "3",
    "LinkedTxn": [
      {
        "TxnId": "103",
        "TxnType": "Invoice"
      }
    ],
    "CustomField": [
      {
        "DefinitionId": "1",
        "Type": "StringType",
        "Name": "Crew #"
      }
    ],
    "ShipAddr": {
      "CountrySubDivisionCode": "CA",
      "City": "Middlefield",
      "PostalCode": "94303",
      "Id": "119",
      "Line1": "1987 Main St."
    },
    "EmailStatus": "NotSet",
    "BillAddr": {
      "Line3": "Middlefield, CA 94303",
      "Id": "59",
      "Line1": "Geeta Kalapatapu",
      "Line2": "1987 Main St."
    },
    "MetaData": {
      "CreateTime": "2014-09-17T11:20:06-07:00",
      "LastUpdatedTime": "2015-07-24T14:15:10-07:00"
    },
    "BillEmail": {
      "Address": "Geeta@Kalapatapu.com"
    },
    "Id": "41"
  },
  "time": "2015-07-24T14:15:10.332-07:00"
}
```
