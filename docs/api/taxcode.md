# TaxCode

A TaxCode object is used to track the taxable or non-taxable status of products, services, and customers. You can assign a sales tax code to each of your products, services, and customers based on their taxable or non-taxable status. You can then use these codes to generate reports that provide information to the tax agencies about the taxable or non-taxable status of certain sales. See for more information about using TaxCode objects and the tax model in general.

## Operations

### Create a taxcode

- **Method**: POST
- **URL**: `/v3/company/{realmID}/taxcode`

Use the taxservice resource to create a tax code.

## Attributes

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| Id | String | Required for update | Unique identifier for this object. Sort order is ASC by default. |
| Name |  | Required | User recognizable name for the tax sales code. |
| SyncToken | String | Required for update | Version number of the object. It is used to lock an object for use by one app at a time. As soon as an application modifies an object, its SyncToken is incremented. |
| PurchaseTaxRateList |  | Conditionally required | TaxRateList List of references to tax rates that apply for purchase transactions when this tax code represents a group of tax rates. Required when TaxGroup is set to true |
| SalesTaxRateList |  | Conditionally required | TaxRateList List of references to tax rates that apply for sales transactions when this tax code represents a group of tax rates. Required when TaxGroup is set to true |
| TaxGroup | Boolean | Read-only | true—-this object represents a group of one or more tax rates. false—-this object represents pseudo-tax codes TAX and NON. |
| Taxable | Boolean | Read-only | False or null means meaning non-taxable. True means taxable. Always true, except for the pseudo taxcode NON. |
| Description |  | Optional | User entered description for the sales tax code. |
| Hidden | Boolean | Read-only | Read-only. Denotes whether active tax codes are displayed on transactions. true—-This tax code is hidden on transactions. false—-This tax code is displayed on transactions. |
| MetaData | ModificationMetaData | Read-only | Descriptive information about the object. The MetaData values are set by Data Services and are  for all applications. |
| TaxCodeConfigType | String | Read-only | Flag to identify whether the TaxCode is  by Automated Sales Tax engine or user generated. Valid values include USER_DEFINED, SYSTEM_GENERATEDSYSTEM_GENERATED. |

## Sample Object

```json
{
  "TaxCode": {
    "SyncToken": "0",
    "domain": "QBO",
    "TaxGroup": true,
    "Name": "California",
    "Taxable": true,
    "PurchaseTaxRateList": {
      "TaxRateDetail": []
    },
    "sparse": false,
    "Active": true,
    "Description": "California",
    "MetaData": {
      "CreateTime": "2014-09-18T12:17:04-07:00",
      "LastUpdatedTime": "2014-09-18T12:17:04-07:00"
    },
    "Id": "2",
    "SalesTaxRateList": {
      "TaxRateDetail": [
        {
          "TaxTypeApplicable": "TaxOnAmount",
          "TaxRateRef": {
            "name": "California",
            "value": "3"
          },
          "TaxOrder": 0
        }
      ]
    }
  },
  "time": "2015-07-27T12:37:22.733-07:00"
}
```

### Query a taxcode

- **Method**: GET
- **URL**: `/v3/company/{realmID}/query?query={selectStatement}`

**Sample Query**: `select * From TaxCode`

**Response**:
```json
{
  "QueryResponse": {
    "startPosition": 1,
    "totalCount": 5,
    "TaxCode": [
      {
        "TaxGroup": false,
        "Name": "TAX",
        "Taxable": true,
        "Description": "TAX",
        "Id": "TAX",
        "MetaData": {
          "CreateTime": "2014-10-15T11:28:33-07:00",
          "LastUpdatedTime": "2014-10-15T11:28:33-07:00"
        }
      },
      {
        "TaxGroup": false,
        "Name": "NON",
        "Taxable": false,
        "Description": "NON",
        "Id": "NON",
        "MetaData": {
          "CreateTime": "2014-10-15T11:28:33-07:00",
          "LastUpdatedTime": "2014-10-15T11:28:33-07:00"
        }
      },
      {
        "TaxGroup": true,
        "Name": "CustomSalesTax",
        "Taxable": true,
        "Description": "CustomSalesTax",
        "Id": "CustomSalesTax",
        "MetaData": {
          "CreateTime": "2014-10-15T11:28:33-07:00",
          "LastUpdatedTime": "2014-10-15T11:28:33-07:00"
        }
      },
      {
        "SyncToken": "0",
        "domain": "QBO",
        "TaxGroup": true,
        "Name": "California",
        "Taxable": true,
        "PurchaseTaxRateList": {
          "TaxRateDetail": []
        },
        "sparse": false,
        "Active": true,
        "Description": "California",
        "MetaData": {
          "CreateTime": "2014-09-18T12:17:04-07:00",
          "LastUpdatedTime": "2014-09-18T12:17:04-07:00"
        },
        "Id": "2",
        "SalesTaxRateList": {
          "TaxRateDetail": [
            {
              "TaxTypeApplicable": "TaxOnAmount",
              "TaxRateRef": {
                "name": "California",
                "value": "3"
              },
              "TaxOrder": 0
            }
          ]
        }
      },
      {
        "SyncToken": "0",
        "domain": "QBO",
        "TaxGroup": true,
        "Name": "Tucson",
        "Taxable": true,
        "PurchaseTaxRateList": {
          "TaxRateDetail": []
        },
        "sparse": false,
        "Active": true,
        "Description": "Tucson",
        "MetaData": {
          "CreateTime": "2014-09-18T12:17:04-07:00",
          "LastUpdatedTime": "2014-09-18T12:17:04-07:00"
        },
        "Id": "3",
        "SalesTaxRateList": {
          "TaxRateDetail": [
            {
              "TaxTypeApplicable": "TaxOnAmount",
              "TaxRateRef": {
                "name": "AZ State tax",
                "value": "1"
              },
              "TaxOrder": 0
            },
            {
              "TaxTypeApplicable": "TaxOnAmount",
              "TaxRateRef": {
                "name": "Tucson City",
                "value": "2"
              },
              "TaxOrder": 0
            }
          ]
        }
      }
    ],
    "maxResults": 5
  },
  "time": "2015-07-27T11:44:00.125-07:00"
}
```

### Read a taxcode

- **Method**: GET
- **URL**: `/v3/company/{realmID}/taxcode/{taxcodeId}`

Retrieves the details of a TaxCode object that has been previously created.

**Response**:
```json
{
  "TaxCode": {
    "SyncToken": "0",
    "domain": "QBO",
    "TaxGroup": true,
    "Name": "California",
    "Taxable": true,
    "PurchaseTaxRateList": {
      "TaxRateDetail": []
    },
    "sparse": false,
    "Active": true,
    "Description": "California",
    "MetaData": {
      "CreateTime": "2014-09-18T12:17:04-07:00",
      "LastUpdatedTime": "2014-09-18T12:17:04-07:00"
    },
    "Id": "2",
    "SalesTaxRateList": {
      "TaxRateDetail": [
        {
          "TaxTypeApplicable": "TaxOnAmount",
          "TaxRateRef": {
            "name": "California",
            "value": "3"
          },
          "TaxOrder": 0
        }
      ]
    }
  },
  "time": "2015-07-27T12:37:22.733-07:00"
}
```
