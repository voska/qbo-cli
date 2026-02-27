# Department

The Department resource provides a way to track transactions based on physical locations such as stores, sales regions, or countries. As you create sales and expense transactions, consistently designate the department to which they belong. Delete is achieved by setting the Active attribute to false in an entity update request; thus, making it inactive. In this type of delete, the record is not permanently deleted, but is hidden for display purposes. References to inactive objects are left intact.

## The Department Object

### Attributes

| Attribute | Required | Type | Description |
|---|---|---|---|
| Id | Required for update | String, filterable, sortable | Unique identifier for this object. Sort order is ASC by default. Read only. |
| Name | Required | String (max 100 chars) | User recognizable name for the Department. |
| SyncToken | Required for update | String | Version number of the object. Used to lock an object for use by one app at a time. Read only. |
| ParentRef | Conditionally required | ReferenceType | The immediate parent of the SubDepartment. Required for the create operation if this object is a SubDepartment. |
| Active | Optional | Boolean, filterable, sortable | If true, this entity is currently enabled for use by QuickBooks. If set to false, this entity is not available. |
| MetaData | Optional | ModificationMetaData, filterable, sortable | Descriptive information about the object. Read only. |
| FullyQualifiedName | Read only | String, filterable, sortable | Fully qualified name of the entity. Prepends the topmost parent, followed by each sub element separated by colons. Takes the form of `Parent:Department1:SubDepartment1:SubDepartment2`. Limited to 5 levels. |
| SubDepartment | Read only | Boolean, sortable | Specifies whether this Department object is a SubDepartment. `true` -- SubDepartment. `false` or `null` -- top-level Department. |

### Sample Object

```json
{
  "Department": {
    "FullyQualifiedName": "Marketing Department",
    "domain": "QBO",
    "Name": "Marketing Department",
    "SyncToken": "0",
    "SubDepartment": false,
    "sparse": false,
    "Active": true,
    "Id": "2",
    "MetaData": {
      "CreateTime": "2013-08-13T11:52:48-07:00",
      "LastUpdatedTime": "2013-08-13T11:52:48-07:00"
    }
  },
  "time": "2013-08-13T11:54:48.026-07:00"
}
```

## Create a Department

### Request

```
POST /v3/company/<realmID>/department
Content-Type: application/json

Production Base URL: https://quickbooks.api.intuit.com
Sandbox Base URL: https://sandbox-quickbooks.api.intuit.com
```

### Request Body

```json
{
  "Name": "Marketing Department"
}
```

### Response

```json
{
  "Department": {
    "FullyQualifiedName": "Marketing Department",
    "domain": "QBO",
    "Name": "Marketing Department",
    "SyncToken": "0",
    "SubDepartment": false,
    "sparse": false,
    "Active": true,
    "Id": "3",
    "MetaData": {
      "CreateTime": "2015-07-23T12:54:44-07:00",
      "LastUpdatedTime": "2015-07-23T12:54:44-07:00"
    }
  },
  "time": "2015-07-23T12:54:44.248-07:00"
}
```

## Query a Department

### Request

```
GET /v3/company/<realmID>/query?query=<selectStatement>
Content-Type: text/plain

Production Base URL: https://quickbooks.api.intuit.com
Sandbox Base URL: https://sandbox-quickbooks.api.intuit.com
```

### Sample Query

```sql
select * from Department
```

### Response

```json
{
  "QueryResponse": {
    "Department": [
      {
        "FullyQualifiedName": "Sales Department",
        "domain": "QBO",
        "Name": "Sales Department",
        "SyncToken": "0",
        "SubDepartment": false,
        "sparse": false,
        "Active": false,
        "Id": "1",
        "MetaData": {
          "CreateTime": "2013-08-13T11:49:31-07:00",
          "LastUpdatedTime": "2013-08-13T11:49:31-07:00"
        }
      },
      {
        "FullyQualifiedName": "Support Department",
        "domain": "QBO",
        "Name": "Support Department",
        "SyncToken": "2",
        "SubDepartment": false,
        "sparse": false,
        "Active": false,
        "Id": "2",
        "MetaData": {
          "CreateTime": "2013-08-13T11:52:48-07:00",
          "LastUpdatedTime": "2013-08-13T11:58:58-07:00"
        }
      }
    ],
    "startPosition": 1,
    "maxResults": 2
  },
  "time": "2013-08-13T12:04:05.965-07:00"
}
```

## Read a Department

Retrieves the details of a Department object that has been previously created.

### Request

```
GET /v3/company/<realmID>/department/<departmentId>

Production Base URL: https://quickbooks.api.intuit.com
Sandbox Base URL: https://sandbox-quickbooks.api.intuit.com
```

### Response

```json
{
  "Department": {
    "FullyQualifiedName": "Marketing Department",
    "domain": "QBO",
    "Name": "Marketing Department",
    "SyncToken": "0",
    "SubDepartment": false,
    "sparse": false,
    "Active": true,
    "Id": "2",
    "MetaData": {
      "CreateTime": "2013-08-13T11:52:48-07:00",
      "LastUpdatedTime": "2013-08-13T11:52:48-07:00"
    }
  },
  "time": "2013-08-13T11:54:48.026-07:00"
}
```

## Full Update a Department

Use this operation to update any of the writable fields of an existing Department object. The request body must include all writable fields of the existing object as returned in a read response. Writable fields omitted from the request body are set to NULL. The ID of the object to update is specified in the request body.

### Request

```
POST /v3/company/<realmID>/department
Content-Type: application/json

Production Base URL: https://quickbooks.api.intuit.com
Sandbox Base URL: https://sandbox-quickbooks.api.intuit.com
```

### Request Body

```json
{
  "FullyQualifiedName": "Marketing Department",
  "domain": "QBO",
  "Name": "Support Department",
  "SyncToken": "1",
  "SubDepartment": false,
  "sparse": false,
  "Active": true,
  "Id": "2",
  "MetaData": {
    "CreateTime": "2013-08-13T11:52:48-07:00",
    "LastUpdatedTime": "2013-08-13T11:52:48-07:00"
  }
}
```

### Response

```json
{
  "Department": {
    "FullyQualifiedName": "Support Department",
    "domain": "QBO",
    "Name": "Support Department",
    "SyncToken": "2",
    "SubDepartment": false,
    "sparse": false,
    "Active": true,
    "Id": "2",
    "MetaData": {
      "CreateTime": "2013-08-13T11:52:48-07:00",
      "LastUpdatedTime": "2013-08-13T11:58:58-07:00"
    }
  },
  "time": "2013-08-13T11:58:58.925-07:00"
}
```
