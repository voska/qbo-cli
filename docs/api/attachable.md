# Attachable

This page covers the Attachable, Upload, and Download resources used for attachment management. Attachments are supplemental information linked to a transaction or Item object. They can be files, notes, or a combination of both.

- In the case of file attachments, use an upload endpoint multipart request to upload the files to the QuickBooks attachment list and, optionally, to supply metadata for each via an attachable object. If metadata is not supplied with the upload request, the system creates it.
- In the case of a note, use the create attachable endpoint.

## Business Rules

An upload request may contain as many files as possible in a request, but the overall request size must not exceed 100 MB.

An attachable record can either contain a note only, a file attachment only, or both. When using FileName and Note attributes together, it's up to the app to manage how the note relates to the file name.

## The Attachable Object

### Attributes

| Attribute | Required | Type | Description |
|---|---|---|---|
| Id | Required for update | IdType, filterable, sortable | Unique Identifier for an Intuit entity (object). Required for the update operation. Read only. |
| SyncToken | Required for update | String | Version number of the object. Used to lock an object for use by one app at a time. Read only. |
| FileName | Conditionally required | String (max 1000 chars), filterable, sortable | FileName of the attachment. Required for file attachments. |
| Note | Conditionally required | String (max 2000 chars), filterable, sortable | This note is either related to the attachment specified by FileName or is a standalone note. Required for standalone notes. |
| Category | Optional | String (max 100 chars), filterable, sortable | Category of the attachment. Valid values (case sensitive): Contact Photo, Document, Image, Receipt, Signature, Sound, Other. |
| ContentType | Optional | String (max 100 chars), filterable, sortable | ContentType of the attachment. Returned for file attachments. |
| PlaceName | Optional | String (max 2000 chars), filterable, sortable | PlaceName from where the attachment was requested. |
| AttachableRef | Optional | AttachableRef | Specifies the transaction object to which this attachable file is to be linked. |
| Long | Optional | String (max 100 chars), filterable, sortable | Longitude from where the attachment was requested. |
| Tag | Optional | String (max 2000 chars), filterable, sortable | Tag name for the requested attachment. |
| Lat | Optional | String (max 100 chars), filterable, sortable | Latitude from where the attachment was requested. |
| MetaData | Optional | ModificationMetaData | Descriptive information about the entity. Read only. |
| FileAccessUri | Read only | String | FullPath FileAccess URI of the attachment. Returned for file attachments. |
| Size | Read only | Decimal, filterable, sortable | Size of the attachment. Returned for file attachments. |
| ThumbnailFileAccessUri | Read only | String | FullPath FileAccess URI of the attachment thumbnail if the content type supports thumbnails. Returned for file attachments. |
| TempDownloadUri | Read only | String | TempDownload URI which can be directly downloaded by clients. Returned for file attachments. |
| ThumbnailTempDownloadUri | Read only | String, filterable, sortable | Thumbnail TempDownload URI which can be directly downloaded by clients. Only available if the content type supports thumbnails. |

### Sample Object

```json
{
  "Attachable": {
    "SyncToken": "0",
    "domain": "QBO",
    "AttachableRef": [
      {
        "IncludeOnSend": false,
        "EntityRef": {
          "type": "Invoice",
          "value": "95"
        }
      }
    ],
    "Note": "This is an attached note.",
    "sparse": false,
    "Id": "200900000000000008541",
    "MetaData": {
      "CreateTime": "2015-11-17T11:05:15-08:00",
      "LastUpdatedTime": "2015-11-17T11:05:15-08:00"
    }
  },
  "time": "2015-11-17T11:05:15.797-08:00"
}
```

## Create a Note Attachment

Use this endpoint to attach a note to an object. Adjust values in the AttachableRef element for your specific target object. The value for `AttachableRef.EntityRef.value` is the Id of the target object as returned in its response body when queried.

### Request

```
POST /v3/company/<realmID>/attachable
Content-Type: application/json

Production Base URL: https://quickbooks.api.intuit.com
Sandbox Base URL: https://sandbox-quickbooks.api.intuit.com
```

### Request Body

```json
{
  "Note": "This is an attached note.",
  "AttachableRef": [
    {
      "IncludeOnSend": "false",
      "EntityRef": {
        "type": "Invoice",
        "value": "95"
      }
    }
  ]
}
```

### Response

```json
{
  "Attachable": {
    "SyncToken": "0",
    "domain": "QBO",
    "AttachableRef": [
      {
        "IncludeOnSend": false,
        "EntityRef": {
          "type": "Invoice",
          "value": "95"
        }
      }
    ],
    "Note": "This is an attached note.",
    "sparse": false,
    "Id": "200900000000000008541",
    "MetaData": {
      "CreateTime": "2015-11-17T11:05:15-08:00",
      "LastUpdatedTime": "2015-11-17T11:05:15-08:00"
    }
  },
  "time": "2015-11-17T11:05:15.797-08:00"
}
```

## Delete an Attachable

This operation deletes the attachable object specified in the request body. The request body must include the full payload of the attachable as returned in a read response.

### Request

```
POST /v3/company/<realmID>/attachable?operation=delete
Content-Type: application/json

Production Base URL: https://quickbooks.api.intuit.com
Sandbox Base URL: https://sandbox-quickbooks.api.intuit.com
```

### Request Body

```json
{
  "SyncToken": "0",
  "domain": "QBO",
  "AttachableRef": [
    {
      "IncludeOnSend": false,
      "EntityRef": {
        "type": "Invoice",
        "value": "95"
      }
    }
  ],
  "Note": "This is an attached note.",
  "sparse": false,
  "Id": "200900000000000008541",
  "MetaData": {
    "CreateTime": "2015-11-17T11:05:15-08:00",
    "LastUpdatedTime": "2015-11-17T11:05:15-08:00"
  }
}
```

### Response

```json
{
  "Attachable": {
    "status": "Deleted",
    "domain": "QBO",
    "Id": "200900000000000008541"
  },
  "time": "2015-03-15T11:27:41.514-07:00"
}
```

## Download an Attachment

Retrieves a temporary download URL to the specified attachableId as returned in the `attachable.Id` attribute from a read response. The application uses this URL to then download the file as a separate step. The URL expires after 15 minutes, after which time the app may obtain another.

### Request

```
GET /v3/company/<realmID>/download/<attachableId>
Content-Type: text/plain

Production Base URL: https://quickbooks.api.intuit.com
Sandbox Base URL: https://sandbox-quickbooks.api.intuit.com
```

### Response

```
"https://intuit-.../attachments/receipt_june10.pdf?Expires=...&AWSAccessKeyId=...&Signature=..."
```

## Query an Attachable

The sample query returns the attachable IDs for all attachments linked to the purchase object whose Id is 611. To formulate this query, you must first know the type of object and its object Id using the query endpoint for that resource.

- Set `AttachableRef.EntityRef.type` with the specific type of the target object. For this example, purchase is used.
- Set `AttachableRef.EntityRef.value` with the Id of the target object as returned in its response body when queried. For this example, 611 is used.

### Request

```
GET /v3/company/<realmID>/query?query=<selectStatement>
Content-Type: text/plain

Production Base URL: https://quickbooks.api.intuit.com
Sandbox Base URL: https://sandbox-quickbooks.api.intuit.com
```

### Sample Query

```sql
select Id from attachable where AttachableRef.EntityRef.Type = 'purchase' and AttachableRef.EntityRef.value = '611'
```

### Response

```json
{
  "QueryResponse": {
    "startPosition": 1,
    "Attachable": [
      {
        "Id": "100000000004062174",
        "sparse": true
      },
      {
        "Id": "100000000004158481",
        "sparse": true
      }
    ],
    "maxResults": 2
  },
  "time": "2015-11-24T10:18:31.289-08:00"
}
```

## Read an Attachable

Retrieves the details of an attachable item that has been previously created.

### Request

```
GET /v3/company/<realmID>/attachable/<attachableId>

Production Base URL: https://quickbooks.api.intuit.com
Sandbox Base URL: https://sandbox-quickbooks.api.intuit.com
```

### Response

```json
{
  "Attachable": {
    "SyncToken": "0",
    "domain": "QBO",
    "AttachableRef": [
      {
        "IncludeOnSend": false,
        "EntityRef": {
          "type": "Invoice",
          "value": "95"
        }
      }
    ],
    "Note": "This is an attached note.",
    "sparse": false,
    "Id": "5000000000000010341",
    "MetaData": {
      "CreateTime": "2015-11-17T11:05:15-08:00",
      "LastUpdatedTime": "2015-11-17T11:05:15-08:00"
    }
  },
  "time": "2015-11-17T11:09:34.216-08:00"
}
```

## Full Update an Attachable

Use this operation to update any of the writable fields of an existing attachable object. The request body must include all writable fields of the existing object as returned in a read response. Writable fields omitted from the request body are set to NULL. The ID of the object to update is specified in the request body.

### Request

```
POST /v3/company/<realmID>/attachable
Content-Type: application/json

Production Base URL: https://quickbooks.api.intuit.com
Sandbox Base URL: https://sandbox-quickbooks.api.intuit.com
```

### Request Body

```json
{
  "SyncToken": "1",
  "domain": "QBO",
  "AttachableRef": [
    {
      "IncludeOnSend": false,
      "EntityRef": {
        "type": "Invoice",
        "value": "95"
      }
    }
  ],
  "Note": "This is an updated attached note.",
  "sparse": false,
  "Id": "5000000000000010341",
  "MetaData": {
    "CreateTime": "2015-11-17T11:05:15-08:00",
    "LastUpdatedTime": "2015-11-17T11:05:15-08:00"
  }
}
```

### Response

```json
{
  "Attachable": {
    "SyncToken": "1",
    "domain": "QBO",
    "AttachableRef": [
      {
        "IncludeOnSend": false,
        "EntityRef": {
          "type": "Invoice",
          "value": "95"
        }
      }
    ],
    "Note": "This is an updated attached note.",
    "sparse": false,
    "Id": "5000000000000010341",
    "MetaData": {
      "CreateTime": "2015-11-17T11:05:15-08:00",
      "LastUpdatedTime": "2015-11-17T11:11:04-08:00"
    }
  },
  "time": "2015-11-17T11:11:21.679-08:00"
}
```

## Upload Attachments

The upload endpoint allows the app to upload one or more attachments, with or without associated metadata, via a multipart/form-data request. Each part has its own header that contains:

- A boundary separator.
- A header specifying content disposition and content type.

For your own request, adjust values in the AttachableRef element of the Attachable metadata part for your specific target object.

- Set `AttachableRef.EntityRef.value` with the Id of the target object as returned in its response body when queried.
- Set `AttachableRef.EntityRef.type` with the specific type of the target object. For example, invoice, bill, etc.

### Request

```
POST /v3/company/<realmID>/upload
Content-Type: multipart/form-data

Production Base URL: https://quickbooks.api.intuit.com
Sandbox Base URL: https://sandbox-quickbooks.api.intuit.com
```

### Request Body (Metadata Part)

```json
{
  "AttachableRef": [
    {
      "EntityRef": {
        "type": "Invoice",
        "value": "95"
      }
    }
  ],
  "ContentType": "image/jpg",
  "FileName": "receipt_nov15.jpg"
}
```

### Response

```json
{
  "AttachableResponse": [
    {
      "Attachable": {
        "SyncToken": "0",
        "domain": "QBO",
        "FileAccessUri": "...",
        "ThumbnailFileAccessUri": "...",
        "AttachableRef": [
          {
            "IncludeOnSend": false,
            "EntityRef": {
              "type": "Invoice",
              "value": "95"
            }
          }
        ],
        "TempDownloadUri": "https://...",
        "MetaData": {
          "CreateTime": "2015-11-16T10:59:02-08:00",
          "LastUpdatedTime": "2015-11-16T10:59:02-08:00"
        },
        "sparse": false,
        "ContentType": "image/jpeg",
        "FileName": "receipt_nov15.jpg",
        "Id": "100000000004190865",
        "Size": 1594261
      }
    }
  ],
  "time": "2015-11-16T10:58:58.100-08:00"
}
```
