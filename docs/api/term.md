# Term

A Term object represents the terms under which a sale is made, typically expressed in the form of days due after the goods are received. Optionally, a discount of the total amount may be applied if payment is made within a stipulated time. For example, net 30 indicates that payment is due within 30 days. A term of 2%/15 net 60 indicates that payment is due within 60 days, with a discount of 2% if payment is made within 15 days. This resource also supports:

- An absolute due date.
- A number of days from a start date.
- A percent discount.
- An absolute discount.

## The Term Object

### Attributes

| Attribute | Required | Type | Description |
|---|---|---|---|
| Id | Required for update | String, filterable, sortable | Unique identifier for this object. Sort order is ASC by default. Read only. |
| Name | Required | String (max 31 chars), filterable, sortable | User recognizable name for the term. For example, Net 30. |
| SyncToken | Required for update | String | Version number of the object. Used to lock an object for use by one app at a time. Read only. |
| DiscountPercent | Optional | Decimal (range 0-100) | Discount percentage available against an amount if paid within the days specified by DiscountDays. |
| DiscountDays | Optional | Integer (range 0-999) | Discount applies if paid within this number of days. Used only when DueDays is specified. |
| Active | Optional | Boolean, filterable, sortable | If true, this entity is currently enabled for use by QuickBooks. |
| Type | Optional | String | Type of the Sales Term. Valid values: `STANDARD` -- Used if DueDays is not null. `DATE_DRIVEN` -- Used if DueDays is null. System defined. |
| MetaData | Optional | ModificationMetaData | Descriptive information about the object. Read only. |
| DayOfMonthDue | Conditionally required | Integer (range 1-31) | Payment must be received by this day of the month. Used only if DueDays is not specified. Required if DueDays not present. |
| DiscountDayOfMonth | Conditionally required | Positive Integer (range 0-31) | Discount applies if paid before this day of month. Required if DueDays not present. |
| DueNextMonthDays | Conditionally required | Positive Integer (range 0-999) | Payment due next month if issued that many days before the DayOfMonthDue. Required if DueDays not present. |
| DueDays | Conditionally required | Integer (range 0-999) | Number of days from delivery of goods or services until the payment is due. Required if DayOfMonthDue not present. |

### Sample Object

```json
{
  "Term": {
    "SyncToken": "0",
    "domain": "QBO",
    "Name": "Net 60",
    "DiscountPercent": 0,
    "DiscountDays": 0,
    "Type": "STANDARD",
    "sparse": false,
    "Active": true,
    "DueDays": 60,
    "Id": "4",
    "MetaData": {
      "CreateTime": "2014-09-11T14:41:49-07:00",
      "LastUpdatedTime": "2014-09-11T14:41:49-07:00"
    }
  },
  "time": "2015-07-28T08:52:57.63-07:00"
}
```

## Create a Term

### Request

```
POST /v3/company/<realmID>/term
Content-Type: application/json

Production Base URL: https://quickbooks.api.intuit.com
Sandbox Base URL: https://sandbox-quickbooks.api.intuit.com
```

### Request Body

```json
{
  "DueDays": "120",
  "Name": "Term120"
}
```

### Response

```json
{
  "Term": {
    "SyncToken": "0",
    "domain": "QBO",
    "Name": "Term120",
    "DiscountPercent": 0,
    "Type": "STANDARD",
    "sparse": false,
    "Active": true,
    "DueDays": 120,
    "Id": "8",
    "MetaData": {
      "CreateTime": "2015-07-28T08:50:59-07:00",
      "LastUpdatedTime": "2015-07-28T08:50:59-07:00"
    }
  },
  "time": "2015-07-28T08:51:00.627-07:00"
}
```

## Query a Term

### Request

```
GET /v3/company/<realmID>/query?query=<selectStatement>
Content-Type: text/plain

Production Base URL: https://quickbooks.api.intuit.com
Sandbox Base URL: https://sandbox-quickbooks.api.intuit.com
```

### Sample Query

```sql
select * from Term
```

### Response

```json
{
  "QueryResponse": {
    "startPosition": 1,
    "Term": [
      {
        "SyncToken": "0",
        "domain": "QBO",
        "Name": "Due on receipt",
        "DiscountPercent": 0,
        "DiscountDays": 0,
        "Type": "STANDARD",
        "sparse": false,
        "Active": true,
        "DueDays": 0,
        "Id": "1",
        "MetaData": {
          "CreateTime": "2014-09-11T14:41:49-07:00",
          "LastUpdatedTime": "2014-09-11T14:41:49-07:00"
        }
      },
      {
        "SyncToken": "0",
        "domain": "QBO",
        "Name": "Net 15",
        "DiscountPercent": 0,
        "DiscountDays": 0,
        "Type": "STANDARD",
        "sparse": false,
        "Active": true,
        "DueDays": 15,
        "Id": "2",
        "MetaData": {
          "CreateTime": "2014-09-11T14:41:49-07:00",
          "LastUpdatedTime": "2014-09-11T14:41:49-07:00"
        }
      },
      {
        "SyncToken": "0",
        "domain": "QBO",
        "Name": "Net 30",
        "DiscountPercent": 0,
        "DiscountDays": 0,
        "Type": "STANDARD",
        "sparse": false,
        "Active": true,
        "DueDays": 30,
        "Id": "3",
        "MetaData": {
          "CreateTime": "2014-09-11T14:41:49-07:00",
          "LastUpdatedTime": "2014-09-11T14:41:49-07:00"
        }
      },
      {
        "SyncToken": "0",
        "domain": "QBO",
        "Name": "Net 60",
        "DiscountPercent": 0,
        "DiscountDays": 0,
        "Type": "STANDARD",
        "sparse": false,
        "Active": true,
        "DueDays": 60,
        "Id": "4",
        "MetaData": {
          "CreateTime": "2014-09-11T14:41:49-07:00",
          "LastUpdatedTime": "2014-09-11T14:41:49-07:00"
        }
      }
    ],
    "maxResults": 6
  },
  "time": "2015-07-28T08:26:23.942-07:00"
}
```

## Read a Term

Retrieves the details of a Term object that has been previously created.

### Request

```
GET /v3/company/<realmID>/term/<termId>

Production Base URL: https://quickbooks.api.intuit.com
Sandbox Base URL: https://sandbox-quickbooks.api.intuit.com
```

### Response

```json
{
  "Term": {
    "SyncToken": "0",
    "domain": "QBO",
    "Name": "Net 60",
    "DiscountPercent": 0,
    "DiscountDays": 0,
    "Type": "STANDARD",
    "sparse": false,
    "Active": true,
    "DueDays": 60,
    "Id": "4",
    "MetaData": {
      "CreateTime": "2014-09-11T14:41:49-07:00",
      "LastUpdatedTime": "2014-09-11T14:41:49-07:00"
    }
  },
  "time": "2015-07-28T08:52:57.63-07:00"
}
```

## Full Update a Term

Use this operation to update any of the writable fields of an existing Term object. The request body must include all writable fields of the existing object as returned in a read response. Writable fields omitted from the request body are set to NULL. The ID of the object to update is specified in the request body.

### Request

```
POST /v3/company/<realmID>/term
Content-Type: application/json

Production Base URL: https://quickbooks.api.intuit.com
Sandbox Base URL: https://sandbox-quickbooks.api.intuit.com
```

### Request Body

```json
{
  "SyncToken": "0",
  "domain": "QBO",
  "Name": "Net 30",
  "DiscountPercent": 0,
  "DiscountDays": 10,
  "Type": "STANDARD",
  "sparse": false,
  "Active": true,
  "DueDays": 30,
  "Id": "3",
  "MetaData": {
    "CreateTime": "2014-09-11T14:41:49-07:00",
    "LastUpdatedTime": "2014-09-11T14:41:49-07:00"
  }
}
```

### Response

```json
{
  "Term": {
    "SyncToken": "1",
    "domain": "QBO",
    "Name": "Net 30",
    "DiscountPercent": 0,
    "DiscountDays": 10,
    "Type": "STANDARD",
    "sparse": false,
    "Active": true,
    "DueDays": 30,
    "Id": "3",
    "MetaData": {
      "CreateTime": "2014-09-11T14:41:49-07:00",
      "LastUpdatedTime": "2015-07-28T08:55:59-07:00"
    }
  },
  "time": "2015-07-28T08:55:59.916-07:00"
}
```
