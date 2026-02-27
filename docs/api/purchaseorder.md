# PurchaseOrder

The PurchaseOrder object is a non-posting transaction representing a request to purchase goods or services from a third party.

## Attributes

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| Id | String | Required for update | Unique identifier for this object. Sort order is ASC by default. |
| APAccountRef | ReferenceType | Required | Specifies to which AP account the bill is credited. Query the Account name list resource to determine the appropriate Account object for this reference. Use Account.Id and Account. |
| VendorRef | ReferenceType | Required | Reference to the vendor for this transaction. Query the Vendor name list resource to determine the appropriate Vendor object for this reference. Use Vendor.Id and Vendor. |
| Line | Line | Required | Individual line items of a transaction. Valid Line types include: ItemBasedExpenseLine and AccountBasedExpenseLine ItemBasedExpenseLine AccountBasedExpenseLine |
| SyncToken | String | Required for update | Version number of the object. It is used to lock an object for use by one app at a time. As soon as an application modifies an object, its SyncToken is incremented. |
| CurrencyRef | CurrencyRefType | Conditionally required | Reference to the currency in which all amounts on the associated transaction are expressed. This must be defined if multicurrency is enabled for the company. |
| GlobalTaxCalculation | String | Conditionally required | Method in which tax is applied. Allowed values are: TaxExcluded, TaxInclusive, and NotApplicable. |
| TxnDate | Date | Optional | The date entered by the user when this transaction occurred. For posting transactions, this is the posting date that affects the financial statements. |
| CustomField | CustomField | Optional | One of, up to three custom fields for the transaction. Available for custom fields so configured for the company. Check Preferences.SalesFormsPrefs.CustomField and Preferences. |
| POEmail | EmailAddress | Optional | Used to specify the vendor e-mail address where the purchase req is sent. |
| ClassRef | ReferenceType | Optional | Reference to the Class associated with the transaction. Available if Preferences.AccountingInfoPrefs.ClassTrackingPerTxn is set to true. |
| SalesTermRef | ReferenceType | Optional | Reference to the sales term associated with the transaction. Query the Term name list resource to determine the appropriate Term object for this reference. Use Term.Id and Term. |
| LinkedTxn | LinkedTxn | Optional | Zero or more Bill objects linked to this purchase order; LinkedTxn.TxnType is set to Bill. |
| Memo |  | Optional | A message for the vendor. This text appears on the Purchase Order object sent to the vendor. |
| POStatus | String | Optional | Purchase order status. Valid values are: Open and Closed. |
| TransactionLocationType | String | Optional | The account location. Valid values for France: WithinFrance, FranceOverseas, OutsideFranceWithEU, OutsideEU. |
| DueDate | Date | Optional | Date when the payment of the transaction is due. If date is not provided, the number of days specified in SalesTermRef added the transaction date will be used. |
| MetaData | ModificationMetaData | Read-only | Descriptive information about the object. The MetaData values are set by Data Services and are  for all applications. |
| DocNumber | String | Optional | Reference number for the transaction. |
| Throws |  | Optional | an error when duplicate DocNumber is sent in the request. Recommended best practice: check the setting of Preferences:OtherPrefs:NameValue.Name = VendorAndPurchasesPrefs. |
| PrivateNote | String | Optional | User entered, organization-private note about the transaction. This note does not appear on the purchase order to the vendor. |
| ShipMethodRef | ReferenceType | Optional | Reference to the user-defined ShipMethod associated with the transaction. Store shipping method string in both ShipMethodRef.value and ShipMethodRef.name. |
| TxnTaxDetail | TxnTaxDetail | Optional | This data type provides information for taxes charged on the transaction as a whole. |
| ShipTo | ReferenceType | Optional | Reference to the customer to whose shipping address the order will be shipped to. |
| ExchangeRate | Decimal | Optional | The number of home currency units it takes to equal one unit of currency specified by CurrencyRef. Applicable if multicurrency is enabled for the company. |
| ShipAddr | PhysicalAddress | Optional | Address to which the vendor shipped or will ship any goods associated with the purchase. |
| VendorAddr | PhysicalAddress | Optional | Address to which the payment should be sent. |
| EmailStatus | String | Optional | Email status of the purchase order. Valid values: NotSet, NeedToSend, EmailSent |
| TotalAmt | BigDecimal | Read-only | Indicates the total amount of the transaction. This includes the total of all the charges, allowances, and taxes. |
| RecurDataRef | ReferenceType | Read-only | A reference to the Recurring Transaction. It captures what recurring transaction template the PurchaseOrder was created from. |

## Sample Object

```json
{
  "PurchaseOrder": {
    "DocNumber": "1005",
    "SyncToken": "0",
    "POEmail": {
      "Address": "send_email@intuit.com"
    },
    "APAccountRef": {
      "name": "Accounts Payable (A/P)",
      "value": "33"
    },
    "CurrencyRef": {
      "name": "United States Dollar",
      "value": "USD"
    },
    "TxnDate": "2015-07-28",
    "TotalAmt": 25.0,
    "ShipAddr": {
      "Line4": "Half Moon Bay, CA 94213",
      "Line3": "65 Ocean Dr.",
      "Id": "121",
      "Line1": "Grace Pariente",
      "Line2": "Cool Cars"
    },
    "domain": "QBO",
    "Id": "257",
    "POStatus": "Open",
    "sparse": false,
    "EmailStatus": "NotSet",
    "VendorRef": {
      "name": "Hicks Hardware",
      "value": "41"
    },
    "Line": [
      {
        "DetailType": "ItemBasedExpenseLineDetail",
        "Amount": 25.0,
        "ProjectRef": {
          "value": "39298034"
        },
        "Id": "1",
        "ItemBasedExpenseLineDetail": {
          "ItemRef": {
            "name": "Garden Supplies",
            "value": "38"
          },
          "CustomerRef": {
            "name": "Cool Cars",
            "value": "3"
          },
          "Qty": 1,
          "TaxCodeRef": {
            "value": "NON"
          },
          "BillableStatus": "NotBillable",
          "UnitPrice": 25
        }
      }
    ],
    "CustomField": [
      {
        "DefinitionId": "1",
        "Type": "StringType",
        "Name": "Crew #"
      },
      {
        "DefinitionId": "2",
        "Type": "StringType",
        "Name": "Sales Rep"
      }
    ],
    "VendorAddr": {
      "Line4": "Middlefield, CA 94303",
      "Line3": "42 Main St.",
      "Id": "120",
      "Line1": "Geoff Hicks",
      "Line2": "Hicks Hardware"
    },
    "MetaData": {
      "CreateTime": "2015-07-28T16:01:47-07:00",
      "LastUpdatedTime": "2015-07-28T16:01:47-07:00"
    }
  },
  "time": "2015-07-28T16:04:49.874-07:00"
}
```

## Operations

### Create a purchase order

- **Method**: POST
- **URL**: `/v3/company/{realmID}/purchaseorder`

**Request Body**:
```json
{
  "TotalAmt": 25.0,
  "Line": [
    {
      "DetailType": "ItemBasedExpenseLineDetail",
      "Amount": 25.0,
      "ProjectRef": {
        "value": "39298034"
      },
      "Id": "1",
      "ItemBasedExpenseLineDetail": {
        "ItemRef": {
          "name": "Pump",
          "value": "11"
        },
        "CustomerRef": {
          "name": "Cool Cars",
          "value": "3"
        },
        "Qty": 1,
        "TaxCodeRef": {
          "value": "NON"
        },
        "BillableStatus": "NotBillable",
        "UnitPrice": 25
      }
    }
  ],
  "APAccountRef": {
    "name": "Accounts Payable (A/P)",
    "value": "33"
  },
  "VendorRef": {
    "name": "Hicks Hardware",
    "value": "41"
  },
  "ShipTo": {
    "name": "Jeff's Jalopies",
    "value": "12"
  }
}
```

**Response**:
```json
{
  "PurchaseOrder": {
    "DocNumber": "1007",
    "SyncToken": "0",
    "domain": "QBO",
    "VendorRef": {
      "name": "Hicks Hardware",
      "value": "41"
    },
    "TxnDate": "2015-07-28",
    "TotalAmt": 25.0,
    "APAccountRef": {
      "name": "Accounts Payable (A/P)",
      "value": "33"
    },
    "EmailStatus": "NotSet",
    "sparse": false,
    "Line": [
      {
        "DetailType": "ItemBasedExpenseLineDetail",
        "Amount": 25.0,
        "ProjectRef": {
          "value": "39298034"
        },
        "Id": "1",
        "ItemBasedExpenseLineDetail": {
          "ItemRef": {
            "name": "Pump",
            "value": "11"
          },
          "CustomerRef": {
            "name": "Cool Cars",
            "value": "3"
          },
          "Qty": 1,
          "TaxCodeRef": {
            "value": "NON"
          },
          "BillableStatus": "NotBillable",
          "UnitPrice": 25
        }
      }
    ],
    "CustomField": [
      {
        "DefinitionId": "1",
        "Type": "StringType",
        "Name": "Crew #"
      },
      {
        "DefinitionId": "2",
        "Type": "StringType",
        "Name": "Sales Rep"
      }
    ],
    "Id": "259",
    "MetaData": {
      "CreateTime": "2015-07-28T16:06:03-07:00",
      "LastUpdatedTime": "2015-07-28T16:06:03-07:00"
    }
  },
  "time": "2015-07-28T16:06:04.864-07:00"
}
```

### Delete a purchase order

- **Method**: POST
- **URL**: `/v3/company/{realmID}/purchaseorder?operation=delete`

This operation deletes the purchaseorder object specified in the request body. Include a minimum of PurchaseOrder.Id and PurchaseOrder.SyncToken in the request body. You must unlink any linked transactions associated with the purchaseorder object before deleting it.

**Request Body**:
```json
{
  "SyncToken": "0",
  "Id": "125"
}
```

**Response**:
```json
{
  "PurchaseOrder": {
    "status": "Deleted",
    "domain": "QBO",
    "Id": "125"
  },
  "time": "2015-05-26T14:08:39.858-07:00"
}
```

### Get a purchase order as PDF

- **Method**: GET
- **URL**: `/v3/company/{realmID}/purchaseorder/{purchaseorderId}/pdf`

### Query a purchase order

- **Method**: GET
- **URL**: `/v3/company/{realmID}/query?query={selectStatement}`

**Sample Query**: `select * from PurchaseOrder where Id = '259'`

**Response**:
```json
{
  "QueryResponse": {
    "startPosition": 1,
    "totalCount": 1,
    "PurchaseOrder": [
      {
        "DocNumber": "1007",
        "SyncToken": "0",
        "domain": "QBO",
        "VendorRef": {
          "name": "Hicks Hardware",
          "value": "41"
        },
        "TxnDate": "2015-07-28",
        "TotalAmt": 25.0,
        "APAccountRef": {
          "name": "Accounts Payable (A /P)",
          "value": "33"
        },
        "sparse": false,
        "Line": [
          {
            "DetailType": "ItemBasedExpenseLineDeta il",
            "Amount": 25.0,
            "ProjectRef": {
              "value": "39298034"
            },
            "Id": "1",
            "ItemBasedExpenseLineDeta il": {
              "ItemRef": {
                "name": "Garden Supplies",
                "value": "38"
              },
              "CustomerRef": {
                "name": "Cool Cars",
                "value": "3"
              },
              "Qty": 1,
              "TaxCodeRef": {
                "value": "NON"
              },
              "BillableStatus": "NotBillable",
              "UnitPrice": 25
            }
          }
        ],
        "CustomField": [
          {
            "DefinitionId": "1",
            "Type": "StringType",
            "Name": "Crew #"
          },
          {
            "DefinitionId": "2",
            "Type": "StringType",
            "Name": "Sales Rep"
          }
        ],
        "Id": "259",
        "MetaData": {
          "CreateTime": "2015-07-28T16:06:03-07:00",
          "LastUpdatedTime": "2015-07-28T16:06:03-07:00"
        }
      }
    ],
    "maxResults": 1
  },
  "time": "2015-07-28T16:09:26.277-07:00"
}
```

### Read a purchase order

- **Method**: GET
- **URL**: `/v3/company/{realmID}/purchaseorder/{purchaseorderId}`

Retrieves the details of a purchase order that has been previously created.

**Response**:
```json
{
  "PurchaseOrder": {
    "DocNumber": "1005",
    "SyncToken": "0",
    "POEmail": {
      "Address": "send_email@intuit.com"
    },
    "APAccountRef": {
      "name": "Accounts Payable (A/P)",
      "value": "33"
    },
    "CurrencyRef": {
      "name": "United States Dollar",
      "value": "USD"
    },
    "TxnDate": "2015-07-28",
    "TotalAmt": 25.0,
    "ShipAddr": {
      "Line4": "Half Moon Bay, CA 94213",
      "Line3": "65 Ocean Dr.",
      "Id": "121",
      "Line1": "Grace Pariente",
      "Line2": "Cool Cars"
    },
    "domain": "QBO",
    "Id": "257",
    "POStatus": "Open",
    "sparse": false,
    "EmailStatus": "NotSet",
    "VendorRef": {
      "name": "Hicks Hardware",
      "value": "41"
    },
    "Line": [
      {
        "DetailType": "ItemBasedExpenseLineDetail",
        "Amount": 25.0,
        "ProjectRef": {
          "value": "39298034"
        },
        "Id": "1",
        "ItemBasedExpenseLineDetail": {
          "ItemRef": {
            "name": "Garden Supplies",
            "value": "38"
          },
          "CustomerRef": {
            "name": "Cool Cars",
            "value": "3"
          },
          "Qty": 1,
          "TaxCodeRef": {
            "value": "NON"
          },
          "BillableStatus": "NotBillable",
          "UnitPrice": 25
        }
      }
    ],
    "CustomField": [
      {
        "DefinitionId": "1",
        "Type": "StringType",
        "Name": "Crew #"
      },
      {
        "DefinitionId": "2",
        "Type": "StringType",
        "Name": "Sales Rep"
      }
    ],
    "VendorAddr": {
      "Line4": "Middlefield, CA 94303",
      "Line3": "42 Main St.",
      "Id": "120",
      "Line1": "Geoff Hicks",
      "Line2": "Hicks Hardware"
    },
    "MetaData": {
      "CreateTime": "2015-07-28T16:01:47-07:00",
      "LastUpdatedTime": "2015-07-28T16:01:47-07:00"
    }
  },
  "time": "2015-07-28T16:04:49.874-07:00"
}
```

### Send a purchase order

- **Method**: POST
- **URL**: `/v3/company/{realmID}/purchaseorder/{purchaseorderId}/send`

• The PurchaseOrder.EmailStatus parameter is set to EmailSent • The PurchaseOrder.POEmail.Address parameter is updated to the address specified with the value of the sendTo query parameter, if specified and if the request's minor version is 17 and above.

**Response**:
```json
{
  "PurchaseOrder": {
    "DocNumber": "1005",
    "SyncToken": "0",
    "POEmail": {
      "Address": "send_email@intuit.com"
    },
    "APAccountRef": {
      "name": "Accounts Payable (A/P)",
      "value": "33"
    },
    "CurrencyRef": {
      "name": "United States Dollar",
      "value": "USD"
    },
    "TxnDate": "2015-07-28",
    "TotalAmt": 25.0,
    "ShipAddr": {
      "Line4": "Half Moon Bay, CA 94213",
      "Line3": "65 Ocean Dr.",
      "Id": "121",
      "Line1": "Grace Pariente",
      "Line2": "Cool Cars"
    },
    "domain": "QBO",
    "Id": "257",
    "POStatus": "Open",
    "sparse": false,
    "EmailStatus": "EmailSent",
    "VendorRef": {
      "name": "Hicks Hardware",
      "value": "41"
    },
    "Line": [
      {
        "DetailType": "ItemBasedExpenseLineDetail",
        "Amount": 25.0,
        "ProjectRef": {
          "value": "39298034"
        },
        "Id": "1",
        "ItemBasedExpenseLineDetail": {
          "ItemRef": {
            "name": "Garden Supplies",
            "value": "38"
          },
          "CustomerRef": {
            "name": "Cool Cars",
            "value": "3"
          },
          "Qty": 1,
          "TaxCodeRef": {
            "value": "NON"
          },
          "BillableStatus": "NotBillable",
          "UnitPrice": 25
        }
      }
    ],
    "CustomField": [
      {
        "DefinitionId": "1",
        "Type": "StringType",
        "Name": "Crew #"
      },
      {
        "DefinitionId": "2",
        "Type": "StringType",
        "Name": "Sales Rep"
      }
    ],
    "VendorAddr": {
      "Line4": "Middlefield, CA 94303",
      "Line3": "42 Main St.",
      "Id": "120",
      "Line1": "Geoff Hicks",
      "Line2": "Hicks Hardware"
    },
    "MetaData": {
      "CreateTime": "2015-07-28T16:01:47-07:00",
      "LastUpdatedTime": "2019-09-19T10:43:46-07:00"
    }
  },
  "time": "2019-09-19T10:43:46-07:00"
}
```

### Full update a purchase order

- **Method**: POST
- **URL**: `/v3/company/{realmID}/purchaseorder`

Use this operation to update any of the writable fields of an existing purchase order object. The request body must include all writable fields of the existing object as returned in a read response. Writable fields omitted from the request body are set to NULL. The ID of the object to update is spec...

**Request Body**:
```json
{
  "DocNumber": "1005",
  "SyncToken": "0",
  "POEmail": {
    "Address": "send_email@intuit.com"
  },
  "APAccountRef": {
    "name": "Accounts Payable (A/P)",
    "value": "33"
  },
  "CurrencyRef": {
    "name": "United States Dollar",
    "value": "USD"
  },
  "sparse": false,
  "TxnDate": "2015-07-28",
  "TotalAmt": 25.0,
  "ShipAddr": {
    "Line4": "Half Moon Bay, CA 94213",
    "Line3": "65 Ocean Dr.",
    "Id": "121",
    "Line1": "Grace Pariente",
    "Line2": "Cool Cars"
  },
  "PrivateNote": "This is a private note added during update.",
  "Id": "257",
  "POStatus": "Open",
  "domain": "QBO",
  "VendorRef": {
    "name": "Hicks Hardware",
    "value": "41"
  },
  "Line": [
    {
      "DetailType": "ItemBasedExpenseLineDetail",
      "Amount": 25.0,
      "ProjectRef": {
        "value": "39298034"
      },
      "Id": "1",
      "ItemBasedExpenseLineDetail": {
        "ItemRef": {
          "name": "Garden Supplies",
          "value": "38"
        },
        "CustomerRef": {
          "name": "Cool Cars",
          "value": "3"
        },
        "Qty": 1,
        "TaxCodeRef": {
          "value": "NON"
        },
        "BillableStatus": "NotBillable",
        "UnitPrice": 25
      }
    }
  ],
  "CustomField": [
    {
      "DefinitionId": "1",
      "Type": "StringType",
      "Name": "Crew #"
    },
    {
      "DefinitionId": "2",
      "Type": "StringType",
      "Name": "Sales Rep"
    }
  ],
  "VendorAddr": {
    "Line4": "Middlefield, CA 94303",
    "Line3": "42 Main St.",
    "Id": "120",
    "Line1": "Geoff Hicks",
    "Line2": "Hicks Hardware"
  },
  "MetaData": {
    "CreateTime": "2015-07-28T16:01:47-07:00",
    "LastUpdatedTime": "2015-07-28T16:01:47-07:00"
  }
}
```

**Response**:
```json
{
  "PurchaseOrder": {
    "DocNumber": "1005",
    "SyncToken": "1",
    "domain": "QBO",
    "APAccountRef": {
      "name": "Accounts Payable (A/P)",
      "value": "33"
    },
    "CurrencyRef": {
      "name": "United States Dollar",
      "value": "USD"
    },
    "TxnDate": "2015-07-28",
    "TotalAmt": 25.0,
    "ShipAddr": {
      "Line4": "Half Moon Bay, CA 94213",
      "Line3": "65 Ocean Dr.",
      "Id": "121",
      "Line1": "Grace Pariente",
      "Line2": "Cool Cars"
    },
    "PrivateNote": "This is a private note added during update.",
    "VendorAddr": {
      "Line4": "Middlefield, CA 94303",
      "Line3": "42 Main St.",
      "Id": "120",
      "Line1": "Geoff Hicks",
      "Line2": "Hicks Hardware"
    },
    "POStatus": "Open",
    "sparse": false,
    "VendorRef": {
      "name": "Hicks Hardware",
      "value": "41"
    },
    "Line": [
      {
        "DetailType": "ItemBasedExpenseLineDetail",
        "Amount": 25.0,
        "ProjectRef": {
          "value": "39298034"
        },
        "Id": "1",
        "ItemBasedExpenseLineDetail": {
          "ItemRef": {
            "name": "Garden Supplies",
            "value": "38"
          },
          "CustomerRef": {
            "name": "Cool Cars",
            "value": "3"
          },
          "Qty": 1,
          "TaxCodeRef": {
            "value": "NON"
          },
          "BillableStatus": "NotBillable",
          "UnitPrice": 25
        }
      }
    ],
    "CustomField": [
      {
        "DefinitionId": "1",
        "Type": "StringType",
        "Name": "Crew #"
      },
      {
        "DefinitionId": "2",
        "Type": "StringType",
        "Name": "Sales Rep"
      }
    ],
    "Id": "257",
    "MetaData": {
      "CreateTime": "2015-07-28T16:01:47-07:00",
      "LastUpdatedTime": "2015-07-28T16:17:41-07:00"
    }
  },
  "time": "2015-07-28T16:17:42.952-07:00"
}
```
