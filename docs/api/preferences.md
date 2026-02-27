# Preferences

The Preferences resource represents a set of company preferences that control application behavior in QuickBooks Online. They are mostly exposed as read-only through the Preferences endpoint with only a very small subset of them available as writable. Preferences are not necessarily honored when making requests via the QuickBooks API because a lot of them control UI behavior in the application and may not be applicable for apps.

## Business Rules

- The create operation is not supported.
- The read request retrieves all preferences. There is no notion of preference objects or object IDs.
- Update operations are supported for a limited subset of preferences, which are not marked as readonly.
- The Delete operation is not supported.
- Query is supported with sorting and filtering enabled for Metadata timestamp attributes. Pagination is not supported.
- OtherPrefs type is used as an extension mechanism to contain additional attributes as Name/Value pairs.

## Attributes

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| Id | String | Required for update | Unique identifier for this object. Sort order is ASC by default. |
| SyncToken | String | Required for update | Version number of the object. It is used to lock an object for use by one app at a time. As soon as an application modifies an object, its SyncToken is incremented. |
| MetaData | ModificationMetaData | Read-only | Descriptive information about the object. The MetaData values are set by Data Services and are  for all applications. |
| EmailMessagesPrefs |  | Optional |  |
| ProductAndServicesPrefs |  | Optional |  |
| ReportPrefs |  | Optional |  |
| AccountingInfoPrefs |  | Optional | The following settings are available for QuickBooks Online Plus editions, only. To determine this edition type, query the value of the OfferingSku CustomerInfo. |
| SalesFormsPrefs |  | Optional |  |
| VendorAndPurchasesPrefs |  | Optional |  |
| TaxPrefs |  | Optional |  |
| OtherPrefs |  | Optional | Specifies extension of Preference resource to allow extension of Name-Value pair based extension at the top level. |
| TimeTrackingPrefs |  | Optional |  |
| CurrencyPrefs |  | Optional |  |

## Sample Object

```json
{
  "Preferences": {
    "EmailMessagesPrefs": {
      "InvoiceMessage": {
        "Message": "Your invoice is attached. Please remit payment at your earliest convenience.\\nThank you for your business- we appreciate it very much .\\n\\nSincerely,\\nCraig's Design and Landscaping Services",
        "Subject": "Invoice from Craig's Design and Landscaping Services"
      },
      "EstimateMessage": {
        "Message": "Please review the estimate below. Feel free to contact us if you have any questions.\\nWe look forward to working with you .\\n\\nSincerely,\\nCraig's Design and Landscaping Services",
        "Subject": "Estimate from Craig's Design and Landscaping Services"
      },
      "SalesReceiptMessage": {
        "Message": "Your sales receipt is attached.\\nThank you for your business- we appreciate it very much .\\n\\nSincerely,\\nCraig's Design and Landscaping Services",
        "Subject": "Sales Receipt from Craig's Design and Landscaping Services"
      },
      "StatementMessage": {
        "Message": "Your statement is attached. Please remit payment at your earliest convenience.\\nThank you for your business- we appreciate it very much .\\n\\nSincerely,\\nCraig's Design and Landscaping Services",
        "Subject": "Statement from Craig's Design and Landscaping Services"
      }
    },
    "ProductAndServicesPrefs": {
      "QuantityWithPriceAndRate": true,
      "ForPurchase": true,
      "QuantityOnHand": true,
      "ForSales": true
    },
    "domain": "QBO",
    "SyncToken": "6",
    "ReportPrefs": {
      "ReportBasis": "Accrual",
      "CalcAgingReportFromTxnDate": false
    },
    "AccountingInfoPrefs": {
      "FirstMonthOfFiscalYear": "January",
      "UseAccountNumbers": true,
      "TaxYearMonth": "January",
      "ClassTrackingPerTxn": false,
      "TrackDepartments": true,
      "TaxForm": "6",
      "CustomerTerminology": "Customers",
      "BookCloseDate": "2018-12-31",
      "DepartmentTerminology": "Location",
      "ClassTrackingPerTxnLine": true
    },
    "SalesFormsPrefs": {
      "ETransactionPaymentEnabled": false,
      "CustomTxnNumbers": false,
      "AllowShipping": false,
      "AllowServiceDate": false,
      "ETransactionEnabledStatus": "NotApplicable",
      "DefaultCustomerMessage": "Thank you for your business and have a great day!",
      "EmailCopyToCompany": false,
      "AllowEstimates": true,
      "DefaultTerms": {
        "value": "3"
      },
      "AllowDiscount": true,
      "DefaultDiscountAccount": "86",
      "AllowDeposit": true,
      "AutoApplyPayments": true,
      "IPNSupportEnabled": false,
      "AutoApplyCredit": true,
      "CustomField": [
        {
          "CustomField": [
            {
              "BooleanValue": false,
              "Type": "BooleanType",
              "Name": "SalesFormsPrefs .UseSalesCustom3"
            },
            {
              "BooleanValue": false,
              "Type": "BooleanType",
              "Name": "SalesFormsPrefs .UseSalesCustom2"
            },
            {
              "BooleanValue": true,
              "Type": "BooleanType",
              "Name": "SalesFormsPrefs .UseSalesCustom1"
            }
          ]
        },
        {
          "CustomField": [
            {
              "StringValue": "Crew #",
              "Type": "StringType",
              "Name": "SalesFormsPrefs .SalesCustomName1"
            }
          ]
        }
      ],
      "UsingPriceLevels": false,
      "ETransactionAttachPDF": false
    },
    "VendorAndPurchasesPrefs": {
      "BillableExpenseTracking": true,
      "TrackingByCustomer": true,
      "POCustomField": [
        {
          "CustomField": [
            {
              "BooleanValue": false,
              "Type": "BooleanType",
              "Name": "PurchasePrefs .UsePurchaseCustom3"
            },
            {
              "BooleanValue": true,
              "Type": "BooleanType",
              "Name": "PurchasePrefs .UsePurchaseCustom2"
            },
            {
              "BooleanValue": true,
              "Type": "BooleanType",
              "Name": "PurchasePrefs .UsePurchaseCustom1"
            }
          ]
        },
        {
          "CustomField": [
            {
              "StringValue": "Sales Rep",
              "Type": "StringType",
              "Name": "PurchasePrefs .PurchaseCustomName2"
            },
            {
              "StringValue": "Crew #",
              "Type": "StringType",
              "Name": "PurchasePrefs .PurchaseCustomName1"
            }
          ]
        }
      ]
    },
    "TaxPrefs": {
      "TaxGroupCodeRef": {
        "value": "2"
      },
      "UsingSalesTax": true
    },
    "OtherPrefs": {
      "NameValue": [
        {
          "Name": "SalesFormsPrefs .DefaultCustomerMessage",
          "Value": "Thank you for your business and have a great day!"
        },
        {
          "Name": "SalesFormsPrefs .DefaultItem",
          "Value": "1"
        },
        {
          "Name": "DTXCopyMemo",
          "Value": "false"
        },
        {
          "Name": "UncategorizedAssetAccount Id",
          "Value": "32"
        },
        {
          "Name": "UncategorizedIncomeAccoun tId",
          "Value": "30"
        },
        {
          "Name": "UncategorizedExpenseAccou ntId",
          "Value": "31"
        },
        {
          "Name": "SFCEnabled",
          "Value": "true"
        },
        {
          "Name": "DataPartner",
          "Value": "false"
        },
        {
          "Name": "Vendor1099Enabled",
          "Value": "true"
        },
        {
          "Name": "TimeTrackingFeatureEnable d",
          "Value": "true"
        },
        {
          "Name": "FDPEnabled",
          "Value": "false"
        },
        {
          "Name": "ProjectsEnabled",
          "Value": "false"
        },
        {
          "Name": "DateFormat",
          "Value": "Month Date Year separated by a slash"
        },
        {
          "Name": "DateFormatMnemonic",
          "Value": "MMDDYYYY_SEP_SLASH"
        },
        {
          "Name": "NumberFormat",
          "Value": "US Number Format"
        },
        {
          "Name": "NumberFormatMnemonic",
          "Value": "US_NB"
        },
        {
          "Name": "WarnDuplicateCheckNumber",
          "Value": "true"
        },
        {
          "Name": "WarnDuplicateBillNumber",
          "Value": "false"
        },
        {
          "Name": "SignoutInactiveMinutes",
          "Value": "60"
        },
        {
          "Name": "AccountingInfoPrefs .ShowAccountNumbers",
          "Value": "false"
        }
      ]
    },
    "sparse": false,
    "TimeTrackingPrefs": {
      "WorkWeekStartDate": "Monday",
      "MarkTimeEntriesBillable": true,
      "ShowBillRateToAll": false,
      "UseServices": true,
      "BillCustomers": true
    },
    "CurrencyPrefs": {
      "HomeCurrency": {
        "value": "USD"
      },
      "MultiCurrencyEnabled": false
    },
    "Id": "1",
    "MetaData": {
      "CreateTime": "2017-10-25T01:05:43-07:00",
      "LastUpdatedTime": "2018-03-08T13:24:26-08:00"
    }
  },
  "time": "2018-03-12T08:22:43.280-07:00"
}
```

## Operations

### Query preferences

- **Method**: GET
- **URL**: `/v3/company/{realmID}/query?query={selectStatement}`

**Sample Query**: `select * from Preferences`

**Response**:
```json
{
  "Preferences": {
    "EmailMessagesPrefs": {
      "InvoiceMessage": {
        "Message": "Your invoice is attached. Please remit payment at your earliest convenience.\\nThank you for your business- we appreciate it very much .\\n\\nSincerely,\\nCraig's Design and Landscaping Services",
        "Subject": "Invoice from Craig's Design and Landscaping Services"
      },
      "EstimateMessage": {
        "Message": "Please review the estimate below. Feel free to contact us if you have any questions.\\nWe look forward to working with you .\\n\\nSincerely,\\nCraig's Design and Landscaping Services",
        "Subject": "Estimate from Craig's Design and Landscaping Services"
      },
      "SalesReceiptMessage": {
        "Message": "Your sales receipt is attached.\\nThank you for your business- we appreciate it very much .\\n\\nSincerely,\\nCraig's Design and Landscaping Services",
        "Subject": "Sales Receipt from Craig's Design and Landscaping Services"
      },
      "StatementMessage": {
        "Message": "Your statement is attached. Please remit payment at your earliest convenience.\\nThank you for your business- we appreciate it very much .\\n\\nSincerely,\\nCraig's Design and Landscaping Services",
        "Subject": "Statement from Craig's Design and Landscaping Services"
      }
    },
    "ProductAndServicesPrefs": {
      "QuantityWithPriceAndRate": true,
      "ForPurchase": true,
      "QuantityOnHand": true,
      "ForSales": true
    },
    "domain": "QBO",
    "SyncToken": "6",
    "ReportPrefs": {
      "ReportBasis": "Accrual",
      "CalcAgingReportFromTxnDate": false
    },
    "AccountingInfoPrefs": {
      "FirstMonthOfFiscalYear": "January",
      "UseAccountNumbers": true,
      "TaxYearMonth": "January",
      "ClassTrackingPerTxn": false,
      "TrackDepartments": true,
      "TaxForm": "6",
      "CustomerTerminology": "Customers",
      "BookCloseDate": "2018-12-31",
      "DepartmentTerminology": "Location",
      "ClassTrackingPerTxnLine": true
    },
    "SalesFormsPrefs": {
      "ETransactionPaymentEnabled": false,
      "CustomTxnNumbers": false,
      "AllowShipping": false,
      "AllowServiceDate": false,
      "ETransactionEnabledStatus": "NotApplicable",
      "DefaultCustomerMessage": "Thank you for your business and have a great day!",
      "EmailCopyToCompany": false,
      "AllowEstimates": true,
      "DefaultTerms": {
        "value": "3"
      },
      "AllowDiscount": true,
      "DefaultDiscountAccount": "86",
      "AllowDeposit": true,
      "AutoApplyPayments": true,
      "IPNSupportEnabled": false,
      "AutoApplyCredit": true,
      "CustomField": [
        {
          "CustomField": [
            {
              "BooleanValue": false,
              "Type": "BooleanType",
              "Name": "SalesFormsPrefs .UseSalesCustom3"
            },
            {
              "BooleanValue": false,
              "Type": "BooleanType",
              "Name": "SalesFormsPrefs .UseSalesCustom2"
            },
            {
              "BooleanValue": true,
              "Type": "BooleanType",
              "Name": "SalesFormsPrefs .UseSalesCustom1"
            }
          ]
        },
        {
          "CustomField": [
            {
              "StringValue": "Crew #",
              "Type": "StringType",
              "Name": "SalesFormsPrefs .SalesCustomName1"
            }
          ]
        }
      ],
      "UsingPriceLevels": false,
      "ETransactionAttachPDF": false
    },
    "VendorAndPurchasesPrefs": {
      "BillableExpenseTracking": true,
      "TrackingByCustomer": true,
      "POCustomField": [
        {
          "CustomField": [
            {
              "BooleanValue": false,
              "Type": "BooleanType",
              "Name": "PurchasePrefs .UsePurchaseCustom3"
            },
            {
              "BooleanValue": true,
              "Type": "BooleanType",
              "Name": "PurchasePrefs .UsePurchaseCustom2"
            },
            {
              "BooleanValue": true,
              "Type": "BooleanType",
              "Name": "PurchasePrefs .UsePurchaseCustom1"
            }
          ]
        },
        {
          "CustomField": [
            {
              "StringValue": "Sales Rep",
              "Type": "StringType",
              "Name": "PurchasePrefs .PurchaseCustomName2"
            },
            {
              "StringValue": "Crew #",
              "Type": "StringType",
              "Name": "PurchasePrefs .PurchaseCustomName1"
            }
          ]
        }
      ]
    },
    "TaxPrefs": {
      "TaxGroupCodeRef": {
        "value": "2"
      },
      "UsingSalesTax": true
    },
    "OtherPrefs": {
      "NameValue": [
        {
          "Name": "SalesFormsPrefs .DefaultCustomerMessage",
          "Value": "Thank you for your business and have a great day!"
        },
        {
          "Name": "SalesFormsPrefs .DefaultItem",
          "Value": "1"
        },
        {
          "Name": "DTXCopyMemo",
          "Value": "false"
        },
        {
          "Name": "UncategorizedAssetAccount Id",
          "Value": "32"
        },
        {
          "Name": "UncategorizedIncomeAccoun tId",
          "Value": "30"
        },
        {
          "Name": "UncategorizedExpenseAccou ntId",
          "Value": "31"
        },
        {
          "Name": "SFCEnabled",
          "Value": "true"
        },
        {
          "Name": "DataPartner",
          "Value": "false"
        },
        {
          "Name": "Vendor1099Enabled",
          "Value": "true"
        },
        {
          "Name": "TimeTrackingFeatureEnable d",
          "Value": "true"
        },
        {
          "Name": "FDPEnabled",
          "Value": "false"
        },
        {
          "Name": "ProjectsEnabled",
          "Value": "false"
        },
        {
          "Name": "DateFormat",
          "Value": "Month Date Year separated by a slash"
        },
        {
          "Name": "DateFormatMnemonic",
          "Value": "MMDDYYYY_SEP_SLASH"
        },
        {
          "Name": "NumberFormat",
          "Value": "US Number Format"
        },
        {
          "Name": "NumberFormatMnemonic",
          "Value": "US_NB"
        },
        {
          "Name": "WarnDuplicateCheckNumber",
          "Value": "true"
        },
        {
          "Name": "WarnDuplicateBillNumber",
          "Value": "false"
        },
        {
          "Name": "SignoutInactiveMinutes",
          "Value": "60"
        },
        {
          "Name": "AccountingInfoPrefs .ShowAccountNumbers",
          "Value": "false"
        }
      ]
    },
    "sparse": false,
    "TimeTrackingPrefs": {
      "WorkWeekStartDate": "Monday",
      "MarkTimeEntriesBillable": true,
      "ShowBillRateToAll": false,
      "UseServices": true,
      "BillCustomers": true
    },
    "CurrencyPrefs": {
      "HomeCurrency": {
        "value": "USD"
      },
      "MultiCurrencyEnabled": false
    },
    "Id": "1",
    "MetaData": {
      "CreateTime": "2017-10-25T01:05:43-07:00",
      "LastUpdatedTime": "2018-03-08T13:24:26-08:00"
    }
  },
  "time": "2018-03-12T08:45:52.965-07:00"
}
```

### Read preferences

- **Method**: GET
- **URL**: `/v3/company/{realmID}/preferences`

Retrieves the Preferences details for the specified company.

**Response**:
```json
{
  "Preferences": {
    "EmailMessagesPrefs": {
      "InvoiceMessage": {
        "Message": "Your invoice is attached. Please remit payment at your earliest convenience.\\nThank you for your business- we appreciate it very much .\\n\\nSincerely,\\nCraig's Design and Landscaping Services",
        "Subject": "Invoice from Craig's Design and Landscaping Services"
      },
      "EstimateMessage": {
        "Message": "Please review the estimate below. Feel free to contact us if you have any questions.\\nWe look forward to working with you .\\n\\nSincerely,\\nCraig's Design and Landscaping Services",
        "Subject": "Estimate from Craig's Design and Landscaping Services"
      },
      "SalesReceiptMessage": {
        "Message": "Your sales receipt is attached.\\nThank you for your business- we appreciate it very much .\\n\\nSincerely,\\nCraig's Design and Landscaping Services",
        "Subject": "Sales Receipt from Craig's Design and Landscaping Services"
      },
      "StatementMessage": {
        "Message": "Your statement is attached. Please remit payment at your earliest convenience.\\nThank you for your business- we appreciate it very much .\\n\\nSincerely,\\nCraig's Design and Landscaping Services",
        "Subject": "Statement from Craig's Design and Landscaping Services"
      }
    },
    "ProductAndServicesPrefs": {
      "QuantityWithPriceAndRate": true,
      "ForPurchase": true,
      "QuantityOnHand": true,
      "ForSales": true
    },
    "domain": "QBO",
    "SyncToken": "6",
    "ReportPrefs": {
      "ReportBasis": "Accrual",
      "CalcAgingReportFromTxnDate": false
    },
    "AccountingInfoPrefs": {
      "FirstMonthOfFiscalYear": "January",
      "UseAccountNumbers": true,
      "TaxYearMonth": "January",
      "ClassTrackingPerTxn": false,
      "TrackDepartments": true,
      "TaxForm": "6",
      "CustomerTerminology": "Customers",
      "BookCloseDate": "2018-12-31",
      "DepartmentTerminology": "Location",
      "ClassTrackingPerTxnLine": true
    },
    "SalesFormsPrefs": {
      "ETransactionPaymentEnabled": false,
      "CustomTxnNumbers": false,
      "AllowShipping": false,
      "AllowServiceDate": false,
      "ETransactionEnabledStatus": "NotApplicable",
      "DefaultCustomerMessage": "Thank you for your business and have a great day!",
      "EmailCopyToCompany": false,
      "AllowEstimates": true,
      "DefaultTerms": {
        "value": "3"
      },
      "AllowDiscount": true,
      "DefaultDiscountAccount": "86",
      "AllowDeposit": true,
      "AutoApplyPayments": true,
      "IPNSupportEnabled": false,
      "AutoApplyCredit": true,
      "CustomField": [
        {
          "CustomField": [
            {
              "BooleanValue": false,
              "Type": "BooleanType",
              "Name": "SalesFormsPrefs .UseSalesCustom3"
            },
            {
              "BooleanValue": false,
              "Type": "BooleanType",
              "Name": "SalesFormsPrefs .UseSalesCustom2"
            },
            {
              "BooleanValue": true,
              "Type": "BooleanType",
              "Name": "SalesFormsPrefs .UseSalesCustom1"
            }
          ]
        },
        {
          "CustomField": [
            {
              "StringValue": "Crew #",
              "Type": "StringType",
              "Name": "SalesFormsPrefs .SalesCustomName1"
            }
          ]
        }
      ],
      "UsingPriceLevels": false,
      "ETransactionAttachPDF": false
    },
    "VendorAndPurchasesPrefs": {
      "BillableExpenseTracking": true,
      "TrackingByCustomer": true,
      "POCustomField": [
        {
          "CustomField": [
            {
              "BooleanValue": false,
              "Type": "BooleanType",
              "Name": "PurchasePrefs .UsePurchaseCustom3"
            },
            {
              "BooleanValue": true,
              "Type": "BooleanType",
              "Name": "PurchasePrefs .UsePurchaseCustom2"
            },
            {
              "BooleanValue": true,
              "Type": "BooleanType",
              "Name": "PurchasePrefs .UsePurchaseCustom1"
            }
          ]
        },
        {
          "CustomField": [
            {
              "StringValue": "Sales Rep",
              "Type": "StringType",
              "Name": "PurchasePrefs .PurchaseCustomName2"
            },
            {
              "StringValue": "Crew #",
              "Type": "StringType",
              "Name": "PurchasePrefs .PurchaseCustomName1"
            }
          ]
        }
      ]
    },
    "TaxPrefs": {
      "TaxGroupCodeRef": {
        "value": "2"
      },
      "UsingSalesTax": true
    },
    "OtherPrefs": {
      "NameValue": [
        {
          "Name": "SalesFormsPrefs .DefaultCustomerMessage",
          "Value": "Thank you for your business and have a great day!"
        },
        {
          "Name": "SalesFormsPrefs .DefaultItem",
          "Value": "1"
        },
        {
          "Name": "DTXCopyMemo",
          "Value": "false"
        },
        {
          "Name": "UncategorizedAssetAccount Id",
          "Value": "32"
        },
        {
          "Name": "UncategorizedIncomeAccoun tId",
          "Value": "30"
        },
        {
          "Name": "UncategorizedExpenseAccou ntId",
          "Value": "31"
        },
        {
          "Name": "SFCEnabled",
          "Value": "true"
        },
        {
          "Name": "DataPartner",
          "Value": "false"
        },
        {
          "Name": "Vendor1099Enabled",
          "Value": "true"
        },
        {
          "Name": "TimeTrackingFeatureEnable d",
          "Value": "true"
        },
        {
          "Name": "FDPEnabled",
          "Value": "false"
        },
        {
          "Name": "ProjectsEnabled",
          "Value": "false"
        },
        {
          "Name": "DateFormat",
          "Value": "Month Date Year separated by a slash"
        },
        {
          "Name": "DateFormatMnemonic",
          "Value": "MMDDYYYY_SEP_SLASH"
        },
        {
          "Name": "NumberFormat",
          "Value": "US Number Format"
        },
        {
          "Name": "NumberFormatMnemonic",
          "Value": "US_NB"
        },
        {
          "Name": "WarnDuplicateCheckNumber",
          "Value": "true"
        },
        {
          "Name": "WarnDuplicateBillNumber",
          "Value": "false"
        },
        {
          "Name": "SignoutInactiveMinutes",
          "Value": "60"
        },
        {
          "Name": "AccountingInfoPrefs .ShowAccountNumbers",
          "Value": "false"
        }
      ]
    },
    "sparse": false,
    "TimeTrackingPrefs": {
      "WorkWeekStartDate": "Monday",
      "MarkTimeEntriesBillable": true,
      "ShowBillRateToAll": false,
      "UseServices": true,
      "BillCustomers": true
    },
    "CurrencyPrefs": {
      "HomeCurrency": {
        "value": "USD"
      },
      "MultiCurrencyEnabled": false
    },
    "Id": "1",
    "MetaData": {
      "CreateTime": "2017-10-25T01:05:43-07:00",
      "LastUpdatedTime": "2018-03-08T13:24:26-08:00"
    }
  },
  "time": "2018-03-12T08:22:43.280-07:00"
}
```

### Full update preferences

- **Method**: POST
- **URL**: `/v3/company/{realmID}/preferences`

Use this operation to update any of the writable preference fields. The request body must include all writable fields of the existing object as returned in a read response. Writable fields omitted from the request body are set to NULL or reverted to a default value. The ID of the object to update is...

**Request Body**:
```json
{
  "ProductAndServicesPrefs": {
    "ForPurchase": true,
    "ForSales": true
  },
  "SyncToken": "20",
  "sparse": false,
  "SalesFormsPrefs": {
    "AllowEstimates": true
  },
  "Id": "1"
}
```

**Response**:
```json
{
  "Preferences": {
    "EmailMessagesPrefs": {
      "InvoiceMessage": {
        "Message": "Your invoice is attached. Please remit payment at your earliest convenience.\\nThank you for your business- we appreciate it very much .\\n\\nSincerely,\\nCraig's Design and Landscaping Services",
        "Subject": "Invoice from Craig's Design and Landscaping Services"
      },
      "EstimateMessage": {
        "Message": "Please review the estimate below. Feel free to contact us if you have any questions.\\nWe look forward to working with you .\\n\\nSincerely,\\nCraig's Design and Landscaping Services",
        "Subject": "Estimate from Craig's Design and Landscaping Services"
      },
      "SalesReceiptMessage": {
        "Message": "Your sales receipt is attached.\\nThank you for your business- we appreciate it very much .\\n\\nSincerely,\\nCraig's Design and Landscaping Services",
        "Subject": "Sales Receipt from Craig's Design and Landscaping Services"
      },
      "StatementMessage": {
        "Message": "Your statement is attached. Please remit payment at your earliest convenience.\\nThank you for your business- we appreciate it very much .\\n\\nSincerely,\\nCraig's Design and Landscaping Services",
        "Subject": "Statement from Craig's Design and Landscaping Services"
      }
    },
    "ProductAndServicesPrefs": {
      "QuantityWithPriceAndRate": true,
      "ForPurchase": true,
      "QuantityOnHand": true,
      "ForSales": true
    },
    "domain": "QBO",
    "SyncToken": "6",
    "ReportPrefs": {
      "ReportBasis": "Accrual",
      "CalcAgingReportFromTxnDate": false
    },
    "AccountingInfoPrefs": {
      "FirstMonthOfFiscalYear": "January",
      "UseAccountNumbers": true,
      "TaxYearMonth": "January",
      "ClassTrackingPerTxn": false,
      "TrackDepartments": true,
      "TaxForm": "6",
      "CustomerTerminology": "Customers",
      "BookCloseDate": "2018-12-31",
      "DepartmentTerminology": "Location",
      "ClassTrackingPerTxnLine": true
    },
    "SalesFormsPrefs": {
      "ETransactionPaymentEnabled": false,
      "CustomTxnNumbers": false,
      "AllowShipping": false,
      "AllowServiceDate": false,
      "ETransactionEnabledStatus": "NotApplicable",
      "DefaultCustomerMessage": "Thank you for your business and have a great day!",
      "EmailCopyToCompany": false,
      "AllowEstimates": true,
      "DefaultTerms": {
        "value": "3"
      },
      "AllowDiscount": true,
      "DefaultDiscountAccount": "86",
      "AllowDeposit": true,
      "AutoApplyPayments": true,
      "IPNSupportEnabled": false,
      "AutoApplyCredit": true,
      "CustomField": [
        {
          "CustomField": [
            {
              "BooleanValue": false,
              "Type": "BooleanType",
              "Name": "SalesFormsPrefs .UseSalesCustom3"
            },
            {
              "BooleanValue": false,
              "Type": "BooleanType",
              "Name": "SalesFormsPrefs .UseSalesCustom2"
            },
            {
              "BooleanValue": true,
              "Type": "BooleanType",
              "Name": "SalesFormsPrefs .UseSalesCustom1"
            }
          ]
        },
        {
          "CustomField": [
            {
              "StringValue": "Crew #",
              "Type": "StringType",
              "Name": "SalesFormsPrefs .SalesCustomName1"
            }
          ]
        }
      ],
      "UsingPriceLevels": false,
      "ETransactionAttachPDF": false
    },
    "VendorAndPurchasesPrefs": {
      "BillableExpenseTracking": true,
      "TrackingByCustomer": true,
      "POCustomField": [
        {
          "CustomField": [
            {
              "BooleanValue": false,
              "Type": "BooleanType",
              "Name": "PurchasePrefs .UsePurchaseCustom3"
            },
            {
              "BooleanValue": true,
              "Type": "BooleanType",
              "Name": "PurchasePrefs .UsePurchaseCustom2"
            },
            {
              "BooleanValue": true,
              "Type": "BooleanType",
              "Name": "PurchasePrefs .UsePurchaseCustom1"
            }
          ]
        },
        {
          "CustomField": [
            {
              "StringValue": "Sales Rep",
              "Type": "StringType",
              "Name": "PurchasePrefs .PurchaseCustomName2"
            },
            {
              "StringValue": "Crew #",
              "Type": "StringType",
              "Name": "PurchasePrefs .PurchaseCustomName1"
            }
          ]
        }
      ]
    },
    "TaxPrefs": {
      "TaxGroupCodeRef": {
        "value": "2"
      },
      "UsingSalesTax": true
    },
    "OtherPrefs": {
      "NameValue": [
        {
          "Name": "SalesFormsPrefs .DefaultCustomerMessage",
          "Value": "Thank you for your business and have a great day!"
        },
        {
          "Name": "SalesFormsPrefs .DefaultItem",
          "Value": "1"
        },
        {
          "Name": "DTXCopyMemo",
          "Value": "false"
        },
        {
          "Name": "UncategorizedAssetAccount Id",
          "Value": "32"
        },
        {
          "Name": "UncategorizedIncomeAccoun tId",
          "Value": "30"
        },
        {
          "Name": "UncategorizedExpenseAccou ntId",
          "Value": "31"
        },
        {
          "Name": "SFCEnabled",
          "Value": "true"
        },
        {
          "Name": "DataPartner",
          "Value": "false"
        },
        {
          "Name": "Vendor1099Enabled",
          "Value": "true"
        },
        {
          "Name": "TimeTrackingFeatureEnable d",
          "Value": "true"
        },
        {
          "Name": "FDPEnabled",
          "Value": "false"
        },
        {
          "Name": "ProjectsEnabled",
          "Value": "false"
        },
        {
          "Name": "DateFormat",
          "Value": "Month Date Year separated by a slash"
        },
        {
          "Name": "DateFormatMnemonic",
          "Value": "MMDDYYYY_SEP_SLASH"
        },
        {
          "Name": "NumberFormat",
          "Value": "US Number Format"
        },
        {
          "Name": "NumberFormatMnemonic",
          "Value": "US_NB"
        },
        {
          "Name": "WarnDuplicateCheckNumber",
          "Value": "true"
        },
        {
          "Name": "WarnDuplicateBillNumber",
          "Value": "false"
        },
        {
          "Name": "SignoutInactiveMinutes",
          "Value": "60"
        },
        {
          "Name": "AccountingInfoPrefs .ShowAccountNumbers",
          "Value": "false"
        }
      ]
    },
    "sparse": false,
    "TimeTrackingPrefs": {
      "WorkWeekStartDate": "Monday",
      "MarkTimeEntriesBillable": true,
      "ShowBillRateToAll": false,
      "UseServices": true,
      "BillCustomers": true
    },
    "CurrencyPrefs": {
      "HomeCurrency": {
        "value": "USD"
      },
      "MultiCurrencyEnabled": false
    },
    "Id": "1",
    "MetaData": {
      "CreateTime": "2017-10-25T01:05:43-07:00",
      "LastUpdatedTime": "2018-03-08T13:24:26-08:00"
    }
  },
  "time": "2018-03-12T08:45:52.965-07:00"
}
```
