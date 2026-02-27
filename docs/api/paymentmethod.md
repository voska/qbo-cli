# PaymentMethod

The PaymentMethod resource provides the method of payment for received goods. Delete is achieved by setting the Active attribute to false in an object update request; thus, making it inactive. In this type of delete, the record is not permanently deleted, but is hidden for display purposes. References to inactive objects are left intact.

## The PaymentMethod Object

### Attributes

| Attribute | Required | Type | Description |
|---|---|---|---|
| Id | Required for update | String, filterable, sortable | Unique identifier for this object. Sort order is ASC by default. Read only. |
| Name | Required | String (max 31 chars) | User recognizable name for the payment method. |
| SyncToken | Required for update | String | Version number of the object. Used to lock an object for use by one app at a time. Read only. |
| Active | Optional | Boolean, filterable, sortable | If true, this entity is currently enabled for use by QuickBooks. |
| Type | Optional | String | Defines the type of payment. Valid values include `CREDIT_CARD` or `NON_CREDIT_CARD`. |
| MetaData | Optional | ModificationMetaData | Descriptive information about the object. Read only. |

### Sample Object

```json
{
  "PaymentMethod": {
    "SyncToken": "0",
    "domain": "QBO",
    "Name": "Diners Club",
    "sparse": false,
    "Active": true,
    "Type": "CREDIT_CARD",
    "Id": "7",
    "MetaData": {
      "CreateTime": "2014-09-11T14:42:05-07:00",
      "LastUpdatedTime": "2014-09-11T14:42:05-07:00"
    }
  },
  "time": "2015-07-24T15:29:33.401-07:00"
}
```

## Create a PaymentMethod

### Request

```
POST /v3/company/<realmID>/paymentmethod
Content-Type: application/json

Production Base URL: https://quickbooks.api.intuit.com
Sandbox Base URL: https://sandbox-quickbooks.api.intuit.com
```

### Request Body

```json
{
  "Name": "Business Check"
}
```

### Response

```json
{
  "PaymentMethod": {
    "SyncToken": "0",
    "domain": "QBO",
    "Name": "Business Check",
    "sparse": false,
    "Active": true,
    "Type": "NON_CREDIT_CARD",
    "Id": "10",
    "MetaData": {
      "CreateTime": "2015-07-24T15:37:44-07:00",
      "LastUpdatedTime": "2015-07-24T15:37:44-07:00"
    }
  },
  "time": "2015-07-24T15:37:44.79-07:00"
}
```

## Query a PaymentMethod

### Request

```
GET /v3/company/<realmID>/query?query=<selectStatement>
Content-Type: text/plain

Production Base URL: https://quickbooks.api.intuit.com
Sandbox Base URL: https://sandbox-quickbooks.api.intuit.com
```

### Sample Query

```sql
select * from PaymentMethod
```

### Response

```json
{
  "QueryResponse": {
    "startPosition": 1,
    "PaymentMethod": [
      {
        "SyncToken": "0",
        "domain": "QBO",
        "Name": "American Express",
        "sparse": false,
        "Active": true,
        "Type": "CREDIT_CARD",
        "Id": "6",
        "MetaData": {
          "CreateTime": "2014-09-11T14:42:05-07:00",
          "LastUpdatedTime": "2014-09-11T14:42:05-07:00"
        }
      },
      {
        "SyncToken": "0",
        "domain": "QBO",
        "Name": "Cash",
        "sparse": false,
        "Active": true,
        "Type": "NON_CREDIT_CARD",
        "Id": "1",
        "MetaData": {
          "CreateTime": "2014-09-11T14:42:05-07:00",
          "LastUpdatedTime": "2014-09-11T14:42:05-07:00"
        }
      },
      {
        "SyncToken": "0",
        "domain": "QBO",
        "Name": "Check",
        "sparse": false,
        "Active": true,
        "Type": "NON_CREDIT_CARD",
        "Id": "2",
        "MetaData": {
          "CreateTime": "2014-09-11T14:42:05-07:00",
          "LastUpdatedTime": "2014-09-11T14:42:05-07:00"
        }
      },
      {
        "SyncToken": "0",
        "domain": "QBO",
        "Name": "Diners Club",
        "sparse": false,
        "Active": true,
        "Type": "CREDIT_CARD",
        "Id": "7",
        "MetaData": {
          "CreateTime": "2014-09-11T14:42:05-07:00",
          "LastUpdatedTime": "2014-09-11T14:42:05-07:00"
        }
      },
      {
        "SyncToken": "0",
        "domain": "QBO",
        "Name": "Discover",
        "sparse": false,
        "Active": true,
        "Type": "CREDIT_CARD",
        "Id": "5",
        "MetaData": {
          "CreateTime": "2014-09-11T14:42:05-07:00",
          "LastUpdatedTime": "2014-09-11T14:42:05-07:00"
        }
      },
      {
        "SyncToken": "0",
        "domain": "QBO",
        "Name": "MasterCard",
        "sparse": false,
        "Active": true,
        "Type": "CREDIT_CARD",
        "Id": "4",
        "MetaData": {
          "CreateTime": "2014-09-11T14:42:05-07:00",
          "LastUpdatedTime": "2014-09-11T14:42:05-07:00"
        }
      },
      {
        "SyncToken": "0",
        "domain": "QBO",
        "Name": "Visa",
        "sparse": false,
        "Active": true,
        "Type": "CREDIT_CARD",
        "Id": "3",
        "MetaData": {
          "CreateTime": "2014-09-11T14:42:05-07:00",
          "LastUpdatedTime": "2014-09-11T14:42:05-07:00"
        }
      }
    ],
    "maxResults": 7
  },
  "time": "2015-07-24T15:26:51.916-07:00"
}
```

## Read a PaymentMethod

Retrieves the details of a paymentmethod object that has been previously created.

### Request

```
GET /v3/company/<realmID>/paymentmethod/<paymentmethodId>

Production Base URL: https://quickbooks.api.intuit.com
Sandbox Base URL: https://sandbox-quickbooks.api.intuit.com
```

### Response

```json
{
  "PaymentMethod": {
    "SyncToken": "0",
    "domain": "QBO",
    "Name": "Diners Club",
    "sparse": false,
    "Active": true,
    "Type": "CREDIT_CARD",
    "Id": "7",
    "MetaData": {
      "CreateTime": "2014-09-11T14:42:05-07:00",
      "LastUpdatedTime": "2014-09-11T14:42:05-07:00"
    }
  },
  "time": "2015-07-24T15:29:33.401-07:00"
}
```

## Full Update a PaymentMethod

Use this operation to update any of the writable fields of an existing paymentmethod object. The request body must include all writable fields of the existing object as returned in a read response. Writable fields omitted from the request body are set to NULL. The ID of the object to update is specified in the request body.

### Request

```
POST /v3/company/<realmID>/paymentmethod
Content-Type: application/json

Production Base URL: https://quickbooks.api.intuit.com
Sandbox Base URL: https://sandbox-quickbooks.api.intuit.com
```

### Request Body

```json
{
  "SyncToken": "2",
  "domain": "QBO",
  "Name": "Platinum Diners Club",
  "sparse": false,
  "Active": true,
  "Type": "CREDIT_CARD",
  "Id": "7",
  "MetaData": {
    "CreateTime": "2014-09-11T14:42:05-07:00",
    "LastUpdatedTime": "2014-09-11T14:42:05-07:00"
  }
}
```

### Response

```json
{
  "PaymentMethod": {
    "SyncToken": "2",
    "domain": "QBO",
    "Name": "Platinum Diners Club",
    "sparse": false,
    "Active": true,
    "Type": "CREDIT_CARD",
    "Id": "7",
    "MetaData": {
      "CreateTime": "2014-09-11T14:42:05-07:00",
      "LastUpdatedTime": "2015-07-24T15:33:48-07:00"
    }
  },
  "time": "2015-07-24T15:33:55.695-07:00"
}
```
