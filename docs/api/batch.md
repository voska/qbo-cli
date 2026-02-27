# Batch

The batch operation enables an application to perform multiple operations in a single request. For example, in a single batch request an application can create a customer, update an invoice, and read an account. Compared to multiple requests, a single batch request can improve an application's performance by decreasing network roundtrips and increasing throughput. The individual operations within a batch request are called BatchItemRequest objects.

## Business Rules

- The maximum number of payloads in a single BatchItemRequest is 30.
- The maximum number of requests to the batch endpoint per minute per realmID is 40.
- Execution order of BatchItemRequest objects should not be assumed.
- BatchItemRequest objects are treated independently; a given object cannot depend on another one within the same batch operation. For example, a newly created customer is not available for a subsequent invoice create operation within the same batch operation. You would need to create the customer object first, either autonomously or via a batch request, and then create the invoice object in a subsequent batch request.
- A batch request is authenticated once. This single authentication applies to all BatchItemRequest objects in the request.
- The maximum number of objects that can be returned in a response is 1000.

## Sample Batch Request

### Request Body

The request includes four BatchItemRequest objects.

**Attributes:**

| Attribute | Type | Required | Description |
|---|---|---|---|
| BatchItemRequest | batchitemrequest | Required | A wrapper around all request objects for this batch operation. |

### Returns

**Attributes:**

| Attribute | Type | Required | Description |
|---|---|---|---|
| BatchItemResponse | batchitemresponse | Required | A wrapper around all response objects for this batch operation. |

### Request URL

```
POST /v3/company/<realmID>/batch
Content-Type: application/json

Production Base URL: https://quickbooks.api.intuit.com
Sandbox Base URL: https://sandbox-quickbooks.api.intuit.com
```

### Request Body Example

```json
{
  "BatchItemRequest": [
    {
      "bId": "bid1",
      "Vendor": {
        "DisplayName": "Smith Family Store"
      },
      "operation": "create"
    },
    {
      "bId": "bid2",
      "operation": "delete",
      "Invoice": {
        "SyncToken": "0",
        "Id": "129"
      }
    },
    {
      "SalesReceipt": {
        "PrivateNote": "A private note.",
        "SyncToken": "0",
        "domain": "QBO",
        "Id": "11",
        "sparse": true
      },
      "bId": "bid3",
      "operation": "update"
    },
    {
      "Query": "select * from SalesReceipt where TotalAmt > '300.00'",
      "bId": "bid4"
    }
  ]
}
```

### Response Example

```json
{
  "BatchItemResponse": [
    {
      "Fault": {
        "type": "ValidationFault",
        "Error": [
          {
            "Message": "Duplicate Name Exists Error",
            "code": "6240",
            "Detail": "The name supplied already exists. : Another customer, vendor or employee is already using this name. Please use a different name.",
            "element": ""
          }
        ]
      },
      "bId": "bid1"
    },
    {
      "Fault": {
        "type": "ValidationFault",
        "Error": [
          {
            "Message": "Object Not Found",
            "code": "610",
            "Detail": "Object Not Found : Something you're trying to use has been made inactive. Check the fields with accounts, customers, items, vendors or employees.",
            "element": ""
          }
        ]
      },
      "bId": "bid2"
    },
    {
      "Fault": {
        "type": "ValidationFault",
        "Error": [
          {
            "Message": "Stale Object Error",
            "code": "5010",
            "Detail": "Stale Object Error : You and root were working on this at the same time. root finished before you did, so your work was not saved.",
            "element": ""
          }
        ]
      },
      "bId": "bid3"
    },
    {
      "bId": "bid4",
      "QueryResponse": {
        "SalesReceipt": [
          {
            "TxnDate": "2015-08-25",
            "domain": "QBO",
            "CurrencyRef": {
              "name": "United States Dollar",
              "value": "USD"
            },
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
            "PrivateNote": "A private note.",
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
              "CreateTime": "2015-08-27T14:59:48-07:00",
              "LastUpdatedTime": "2016-04-15T09:01:10-07:00"
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
      }
    }
  ],
  "time": "2016-04-15T09:01:18.141-07:00"
}
```
