# ChangeDataCapture

The change data capture (CDC) operation returns a list of objects that have changed since a specified time. This operation is for an app that periodically polls data services in order to refresh its local copy of object data. The app calls the CDC operation, specifying a comma-separated list of object types and a date-time stamp specifying how far to look back. Data services returns all objects specified by entityList that have changed since the specified date-time. Look-back time can be up to 30 days.

## Business Rules

- This operation is supported for all objects except JournalCode, TimeActivity, TaxAgency, TaxCode, and TaxRate.
- Objects are grouped by type and then in order of last updated time within the group. Objects deleted within the look-back period are returned after active objects.
- A given CDC request returns a maximum of 1000 objects. It is suggested to query with a look-back time shorter than 30 days that can ensure full data is returned.
- The full payload for each object is returned.

## The Change Data Capture (CDC) Object

### Sample Object

```json
{
  "CDCResponse": [
    {
      "QueryResponse": [
        {
          "Customer": [
            { "Id": "63" },
            { "Id": "99" }
          ],
          "startPosition": 1,
          "maxResults": 3
        },
        {
          "startPosition": 1,
          "totalCount": 5,
          "Estimate": [
            { "Id": "34" },
            { "Id": "123" },
            {
              "status": "Deleted",
              "domain": "QBO",
              "Id": "979",
              "MetaData": {
                "LastUpdatedTime": "2015-12-23T12:55:50-08:00"
              }
            }
          ],
          "maxResults": 5
        }
      ]
    }
  ],
  "time": "2015-12-23T12:56:06.196-08:00"
}
```

## Get a List of Changed Entities

`<entityList>` is a comma-separated list of entity names. `<dateTime>` is a date-time stamp for a date within 30 days of today.

### Request

```
GET /v3/company/<realmID>/cdc?entities=<entityList>&changedSince=<dateTime>
Content-Type: text/plain

Production Base URL: https://quickbooks.api.intuit.com
Sandbox Base URL: https://sandbox-quickbooks.api.intuit.com
```

In the query, realmID, entityList, and changedSince date should be replaced with your own values.

### Request Body Example

```
entities=Customer,Estimate&changedSince=2015-11-28
```

### Response Example

```json
{
  "CDCResponse": [
    {
      "QueryResponse": [
        {
          "Customer": [
            { "Id": "63" },
            { "Id": "99" }
          ],
          "startPosition": 1,
          "maxResults": 3
        },
        {
          "startPosition": 1,
          "totalCount": 5,
          "Estimate": [
            { "Id": "34" },
            { "Id": "123" },
            {
              "status": "Deleted",
              "domain": "QBO",
              "Id": "979",
              "MetaData": {
                "LastUpdatedTime": "2015-12-23T12:55:50-08:00"
              }
            }
          ],
          "maxResults": 5
        }
      ]
    }
  ],
  "time": "2015-12-23T12:56:06.196-08:00"
}
```
