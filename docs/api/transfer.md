# Transfer

A Transfer represents a transaction where funds are moved between two accounts from the company's QuickBooks chart of accounts.

## Business Rules


## Attributes

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| Id | String | Required for update | Unique identifier for this object. Sort order is ASC by default. |
| ToAccountRef | ReferenceType | Required | Identifies the asset account to which funds are transfered. Query the Account name list resource to determine the appropriate Account object for this reference. Use Account.Id and Account. |
| Amount | Decimal | Required | Indicates the total amount of the transaction. |
| FromAccountRef | ReferenceType | Required | Identifies the asset account from which funds are transfered. Query the Account name list resource to determine the appropriate Account object for this reference. Use Account.Id and Account. |
| SyncToken | String | Required for update | Version number of the object. It is used to lock an object for use by one app at a time. As soon as an application modifies an object, its SyncToken is incremented. |
| PrivateNote | String | Optional | User entered, organization-private note about the transaction. This note does not appear on the invoice to the customer. |
| TxnDate | Date | Optional | The date entered by the user when this transaction occurred. For posting transactions, this is the posting date that affects the financial statements. |
| TransactionLocationType | String | Optional | The account location. Valid values for France: WithinFrance, FranceOverseas, OutsideFranceWithEU, OutsideEU. |
| MetaData | ModificationMetaData | Read-only | Descriptive information about the object. The MetaData values are set by Data Services and are  for all applications. |
| RecurDataRef | ReferenceType | Read-only | A reference to the Recurring Transaction. It captures what recurring transaction template the Transfer was created from. |

## Sample Object

```json
{
  "Transfer": {
    "SyncToken": "0",
    "domain": "QBO",
    "TxnDate": "2015-02-06",
    "ToAccountRef": {
      "name": "Savings",
      "value": "36"
    },
    "Amount": 120.0,
    "sparse": false,
    "Id": "170",
    "FromAccountRef": {
      "name": "Checking",
      "value": "35"
    },
    "MetaData": {
      "CreateTime": "2015-02-06T11:06:12-08:00",
      "LastUpdatedTime": "2015-02-06T11:06:12-08:00"
    }
  },
  "time": "2015-02-06T11:06:12.017-08:00"
}
```

## Operations

### Create an transfer

- **Method**: POST
- **URL**: `/v3/company/{realmID}/transfer`

**Request Body**:
```json
{
  "Amount": "120.00",
  "ToAccountRef": {
    "name": "Savings",
    "value": "36"
  },
  "FromAccountRef": {
    "name": "Checking",
    "value": "35"
  }
}
```

**Response**:
```json
{
  "Transfer": {
    "SyncToken": "0",
    "domain": "QBO",
    "TxnDate": "2015-02-06",
    "ToAccountRef": {
      "name": "Savings",
      "value": "36"
    },
    "Amount": 120.0,
    "sparse": false,
    "Id": "170",
    "FromAccountRef": {
      "name": "Checking",
      "value": "35"
    },
    "MetaData": {
      "CreateTime": "2015-02-06T11:06:12-08:00",
      "LastUpdatedTime": "2015-02-06T11:06:12-08:00"
    }
  },
  "time": "2015-02-06T11:06:12.017-08:00"
}
```

### Delete an transfer

- **Method**: POST
- **URL**: `/v3/company/{realmID}/transfer?operation=delete`

This operation deletes the Transfer object specified in the request body. The request body must include the full payload of the transfer as returned in a read response.

**Request Body**:
```json
{
  "SyncToken": "2",
  "domain": "QBO",
  "TxnDate": "2015-02-06",
  "ToAccountRef": {
    "name": "Savings",
    "value": "36"
  },
  "Amount": 6600.0,
  "sparse": false,
  "Id": "170",
  "FromAccountRef": {
    "name": "Checking",
    "value": "35"
  }
}
```

**Response**:
```json
{
  "Transfer": {
    "status": "Deleted",
    "domain": "QBO",
    "Id": "170"
  },
  "time": "2015-02-06T11:22:26.347-08:00"
}
```

### Query an transfer

- **Method**: GET
- **URL**: `/v3/company/{realmID}/query?query={selectStatement}`

**Sample Query**: `select * from Transfer`

**Response**:
```json
{
  "QueryResponse": {
    "startPosition": 1,
    "Transfer": [
      {
        "SyncToken": "2",
        "domain": "QBO",
        "TxnDate": "2015-02-06",
        "ToAccountRef": {
          "name": "Savings",
          "value": "36"
        },
        "Amount": 660.0,
        "sparse": false,
        "Id": "170",
        "FromAccountRef": {
          "name": "Checking",
          "value": "35"
        },
        "MetaData": {
          "CreateTime": "2015-02-06T11:06:12-08:00",
          "LastUpdatedTime": "2015-02-06T11:16:06-08:00"
        }
      },
      {
        "SyncToken": "0",
        "domain": "QBO",
        "TxnDate": "2015-02-06",
        "ToAccountRef": {
          "name": "Savings",
          "value": "36"
        },
        "Amount": 120.0,
        "sparse": false,
        "Id": "169",
        "FromAccountRef": {
          "name": "Checking",
          "value": "35"
        },
        "MetaData": {
          "CreateTime": "2015-02-06T11:04:28-08:00",
          "LastUpdatedTime": "2015-02-06T11:04:28-08:00"
        }
      }
    ],
    "maxResults": 2,
    "totalCount": 2
  },
  "time": "2015-02-06T11:17:10.153-08:00"
}
```

### Read an transfer

- **Method**: GET
- **URL**: `/v3/company/{realmID}/transfer/{transferId}`

Retrieves the details of a Transfer object that has been previously created.

**Response**:
```json
{
  "Transfer": {
    "SyncToken": "0",
    "domain": "QBO",
    "TxnDate": "2015-02-06",
    "ToAccountRef": {
      "name": "Savings",
      "value": "36"
    },
    "Amount": 120.0,
    "sparse": false,
    "Id": "170",
    "FromAccountRef": {
      "name": "Checking",
      "value": "35"
    },
    "MetaData": {
      "CreateTime": "2015-02-06T11:06:12-08:00",
      "LastUpdatedTime": "2015-02-06T11:06:12-08:00"
    }
  },
  "time": "2015-02-06T11:06:12.017-08:00"
}
```

### Sparse update an transfer

- **Method**: POST
- **URL**: `/v3/company/{realmID}/transfer`

Sparse updating provides the ability to update a subset of properties for a given object; only elements specified in the request are updated. Missing elements are left untouched. The ID of the object to update is specified in the request body.

**Request Body**:
```json
{
  "SyncToken": "1",
  "domain": "QBO",
  "ToAccountRef": {
    "name": "Savings",
    "value": "36"
  },
  "Amount": 660.0,
  "sparse": true,
  "Id": "170",
  "FromAccountRef": {
    "name": "Checking",
    "value": "35"
  }
}
```

**Response**:
```json
{
  "Transfer": {
    "SyncToken": "2",
    "domain": "QBO",
    "TxnDate": "2015-02-06",
    "ToAccountRef": {
      "name": "Savings",
      "value": "36"
    },
    "Amount": 660.0,
    "sparse": false,
    "Id": "170",
    "FromAccountRef": {
      "name": "Checking",
      "value": "35"
    },
    "MetaData": {
      "CreateTime": "2015-02-06T11:06:12-08:00",
      "LastUpdatedTime": "2015-02-06T11:16:06-08:00"
    }
  },
  "time": "2015-02-06T11:16:06.672-08:00"
}
```

### Full update an transfer

- **Method**: POST
- **URL**: `/v3/company/{realmID}/transfer`

Use this operation to update any of the writable fields of an existing Transfer object. The request body must include all writable fields of the existing object as returned in a read response. Writable fields omitted from the request body are set to NULL. The ID of the object to update is specified ...

**Request Body**:
```json
{
  "SyncToken": "0",
  "domain": "QBO",
  "TxnDate": "2015-02-06",
  "ToAccountRef": {
    "name": "Savings",
    "value": "36"
  },
  "Amount": 550.0,
  "sparse": false,
  "Id": "170",
  "FromAccountRef": {
    "name": "Checking",
    "value": "35"
  },
  "MetaData": {
    "CreateTime": "2015-02-06T11:06:12-08:00",
    "LastUpdatedTime": "2015-02-06T11:06:12-08:00"
  }
}
```

**Response**:
```json
{
  "Transfer": {
    "SyncToken": "1",
    "domain": "QBO",
    "TxnDate": "2015-02-06",
    "ToAccountRef": {
      "name": "Savings",
      "value": "36"
    },
    "Amount": 550.0,
    "sparse": false,
    "Id": "170",
    "FromAccountRef": {
      "name": "Checking",
      "value": "35"
    },
    "MetaData": {
      "CreateTime": "2015-02-06T11:06:12-08:00",
      "LastUpdatedTime": "2015-02-06T11:11:36-08:00"
    }
  },
  "time": "2015-02-06T11:11:36.026-08:00"
}
```
