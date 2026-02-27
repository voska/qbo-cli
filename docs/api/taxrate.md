# TaxRate

A TaxRate object represents rate applied to calculate tax liability. Use the TaxService entity to create a taxrate. See for more information about using TaxRate objects and the tax model in general

## Operations

### Create a taxrate

- **Method**: POST
- **URL**: `/v3/company/{realmID}/taxrate`

Use the TaxService resource to create a tax rate.

## Attributes

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| Id | String | Read-only | Unique identifier for this object. Sort order is ASC by default. |
| SyncToken | String | Required for update | Version number of the object. It is used to lock an object for use by one app at a time. As soon as an application modifies an object, its SyncToken is incremented. |
| RateValue | String | Read-only | Value of the tax rate. |
| Name |  | Read-only | User recognizable name for the tax rate. |
| AgencyRef | ReferenceType | Read-only | Reference to the tax agency associated with this object. |
| SpecialTaxType |  | Read-only | Sting Special tax type to handle zero rate taxes. Used with VAT registered Businesses who receive goods/services (acquisitions) from other EU countries, will need to calculate the VAT due, b... |
| EffectiveTaxRate |  | Read-only | EffectiveTaxRateData List of EffectiveTaxRate. An EffectiveTaxRate is used to know which taxrate is applicable on any date. |
| DisplayType |  | Read-only | Sting TaxRate DisplayType enum which acts as display config. |
| TaxReturnLineRef | ReferenceType | Read-only | Reference to the tax return line associated with this object. |
| MetaData | ModificationMetaData | Read-only | Descriptive information about the object. The MetaData values are set by Data Services and are  for all applications. |
| OriginalTaxRate |  | Read-only |  |
| Description |  | Read-only | User entered description for the tax rate. |

## Sample Object

```json
{
  "TaxRate": {
    "RateValue": 2,
    "AgencyRef": {
      "value": "1"
    },
    "domain": "QBO",
    "Name": "Tucson City",
    "SyncToken": "0",
    "SpecialTaxType": "NONE",
    "DisplayType": "ReadOnly",
    "sparse": false,
    "Active": true,
    "MetaData": {
      "CreateTime": "2014-09-18T12:17:04-07:00",
      "LastUpdatedTime": "2014-09-18T12:17:04-07:00"
    },
    "Id": "2",
    "Description": "Sales Tax"
  },
  "time": "2015-07-27T13:29:41.836-07:00"
}
```

### Query a taxrate

- **Method**: GET
- **URL**: `/v3/company/{realmID}/query?query={selectStatement}`

**Sample Query**: `Select * From TaxRate`

**Response**:
```json
{
  "QueryResponse": {
    "startPosition": 1,
    "totalCount": 3,
    "TaxRate": [
      {
        "RateValue": 7.1,
        "AgencyRef": {
          "value": "1"
        },
        "domain": "QBO",
        "Name": "AZ State tax",
        "SyncToken": "0",
        "SpecialTaxType": "NONE",
        "DisplayType": "ReadOnly",
        "sparse": false,
        "Active": true,
        "MetaData": {
          "CreateTime": "2014-09-18T12:17:04-07:00",
          "LastUpdatedTime": "2014-09-18T12:17:04-07:00"
        },
        "Id": "1",
        "Description": "Sales Tax"
      },
      {
        "RateValue": 8,
        "AgencyRef": {
          "value": "2"
        },
        "domain": "QBO",
        "Name": "California",
        "SyncToken": "0",
        "SpecialTaxType": "NONE",
        "DisplayType": "ReadOnly",
        "sparse": false,
        "Active": true,
        "MetaData": {
          "CreateTime": "2014-09-18T12:17:04-07:00",
          "LastUpdatedTime": "2014-09-18T12:17:04-07:00"
        },
        "Id": "3",
        "Description": "Sales Tax"
      },
      {
        "RateValue": 2,
        "AgencyRef": {
          "value": "1"
        },
        "domain": "QBO",
        "Name": "Tucson City",
        "SyncToken": "0",
        "SpecialTaxType": "NONE",
        "DisplayType": "ReadOnly",
        "sparse": false,
        "Active": true,
        "MetaData": {
          "CreateTime": "2014-09-18T12:17:04-07:00",
          "LastUpdatedTime": "2014-09-18T12:17:04-07:00"
        },
        "Id": "2",
        "Description": "Sales Tax"
      }
    ],
    "maxResults": 3
  },
  "time": "2015-07-27T13:32:06.76-07:00"
}
```

### Read a taxrate

- **Method**: GET
- **URL**: `/v3/company/{realmID}/taxrate/{taxrateId}`

Retrieves the details of a TaxRate object.

**Response**:
```json
{
  "TaxRate": {
    "RateValue": 2,
    "AgencyRef": {
      "value": "1"
    },
    "domain": "QBO",
    "Name": "Tucson City",
    "SyncToken": "0",
    "SpecialTaxType": "NONE",
    "DisplayType": "ReadOnly",
    "sparse": false,
    "Active": true,
    "MetaData": {
      "CreateTime": "2014-09-18T12:17:04-07:00",
      "LastUpdatedTime": "2014-09-18T12:17:04-07:00"
    },
    "Id": "2",
    "Description": "Sales Tax"
  },
  "time": "2015-07-27T13:29:41.836-07:00"
}
```
