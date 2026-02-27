# CompanyInfo

The CompanyInfo object contains basic company information. In QuickBooks, company info and preferences are displayed in the same place under preferences, so it may be confusing to figure out from user interface which fields may belong to this object. But in general, properties such as company addresses or name are considered company information. Some attributes may exist in both CompanyInfo and Preferences objects.

## Attributes

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| Id | String | Read-only | Unique identifier for this object. Sort order is ASC by default. |
| SyncToken | String | Required for update | Version number of the object. It is used to lock an object for use by one app at a time. As soon as an application modifies an object, its SyncToken is incremented. |
| CompanyName |  | Required for update | The name of the company. |
| CompanyAddr | PhysicalAddress | Required for update | Company Address as described in preference. If a physical address is updated from within the transaction object, the QuickBooks Online API flows individual address components differently into the Line... |
| LegalAddr | PhysicalAddress | Optional | Legal Address given to the government for any government communication. |
| SupportedLanguages | String | Optional | Comma separated list of languages. |
| Country | String | Optional | Country name to which the company belongs for financial calculations. |
| WebAddr |  | Optional | 1000 chars WebSiteAddress Website address. |
| NameValue |  | Optional | NameValue pairs Any other preference not covered with the standard set of attributes. See Data Services Extensions, below, for special reserved name/value pairs. NameValue. |
| FiscalYearStartMonth |  | Optional | MonthEnum The start month of fiscal year. |
| CustomerCommunicationAddr | PhysicalAddress | Optional | Address of the company as given to their customer, sometimes the address given to the customer mail address is different from Company address. |
| PrimaryPhone | TelephoneNumber | Optional | Primary phone number. |
| LegalName |  | Optional | The legal name of the company. |
| EmployerId | String | Optional | If your QuickBooks company has defined an EIN in company settings, this value is returned. |
| MetaData | ModificationMetaData | Read-only | Descriptive information about the object. The MetaData values are set by Data Services and are  for all applications. |
| CompanyStartDate | DateTime | Read-only | DateTime when company file was created. This field and Metadata.CreateTimecontain the same value. |

## Sample Object

```json
{
  "CompanyInfo": {
    "SyncToken": "4",
    "domain": "QBO",
    "LegalAddr": {
      "City": "Mountain View",
      "Country": "US",
      "Line1": "2500 Garcia Ave",
      "PostalCode": "94043",
      "CountrySubDivisionCode": "CA",
      "Id": "1"
    },
    "SupportedLanguages": "en",
    "CompanyName": "Larry's Bakery",
    "Country": "US",
    "CompanyAddr": {
      "City": "Mountain View",
      "Country": "US",
      "Line1": "2500 Garcia Ave",
      "PostalCode": "94043",
      "CountrySubDivisionCode": "CA",
      "Id": "1"
    },
    "sparse": false,
    "Id": "1",
    "WebAddr": {},
    "FiscalYearStartMonth": "January",
    "CustomerCommunicationAddr": {
      "City": "Mountain View",
      "Country": "US",
      "Line1": "2500 Garcia Ave",
      "PostalCode": "94043",
      "CountrySubDivisionCode": "CA",
      "Id": "1"
    },
    "PrimaryPhone": {
      "FreeFormNumber": "(650)944-4444"
    },
    "LegalName": "Larry's Bakery",
    "CompanyStartDate": "2015-06-05",
    "EmployerId": "123456789",
    "Email": {
      "Address": "donotreply@intuit.com"
    },
    "NameValue": [
      {
        "Name": "NeoEnabled",
        "Value": "true"
      },
      {
        "Name": "IndustryType",
        "Value": "Bread and Bakery Product Manufacturing"
      },
      {
        "Name": "IndustryCode",
        "Value": "31181"
      },
      {
        "Name": "SubscriptionStatus",
        "Value": "PAID"
      },
      {
        "Name": "OfferingSku",
        "Value": "QuickBooks Online Plus"
      },
      {
        "Name": "PayrollFeature",
        "Value": "true"
      },
      {
        "Name": "AccountantFeature",
        "Value": "false"
      },
      {
        "Name": "IsQbdtMigrated",
        "Value": "true"
      },
      {
        "Name": "MigrationDate",
        "Value": "2024-09-14T01:47:34-07:00"
      },
      {
        "Name": "QBOIndustryType",
        "Value": "Manufacturing Businesses"
      },
      {
        "Name": "AssignedTime",
        "Value": "2024-09-14T01:47:34-07:00"
      }
    ],
    "MetaData": {
      "CreateTime": "2015-06-05T13:55:54-07:00",
      "LastUpdatedTime": "2015-07-06T08:51:50-07:00"
    }
  },
  "time": "2015-07-10T09:38:58.155-07:00"
}
```

## Operations

### Query companyinfo

- **Method**: GET
- **URL**: `/v3/company/{realmID}/query?query={selectStatement}`

**Sample Query**: `select * from CompanyInfo`

**Response**:
```json
{
  "CompanyInfo": {
    "SyncToken": "4",
    "domain": "QBO",
    "LegalAddr": {
      "City": "Mountain View",
      "Country": "US",
      "Line1": "2500 Garcia Ave",
      "PostalCode": "94043",
      "CountrySubDivisionCode": "CA",
      "Id": "1"
    },
    "SupportedLanguages": "en",
    "CompanyName": "Larry's Bakery",
    "Country": "US",
    "CompanyAddr": {
      "City": "Mountain View",
      "Country": "US",
      "Line1": "2500 Garcia Ave",
      "PostalCode": "94043",
      "CountrySubDivisionCode": "CA",
      "Id": "1"
    },
    "sparse": false,
    "Id": "1",
    "WebAddr": {},
    "FiscalYearStartMonth": "January",
    "CustomerCommunicationAddr": {
      "City": "Mountain View",
      "Country": "US",
      "Line1": "2500 Garcia Ave",
      "PostalCode": "94043",
      "CountrySubDivisionCode": "CA",
      "Id": "1"
    },
    "PrimaryPhone": {
      "FreeFormNumber": "(650)944-4444"
    },
    "LegalName": "Larry's Bakery",
    "CompanyStartDate": "2015-06-05",
    "EmployerId": "123456789",
    "Email": {
      "Address": "donotreply@intuit.com"
    },
    "NameValue": [
      {
        "Name": "NeoEnabled",
        "Value": "true"
      },
      {
        "Name": "IndustryType",
        "Value": "Bread and Bakery Product Manufacturing"
      },
      {
        "Name": "IndustryCode",
        "Value": "31181"
      },
      {
        "Name": "SubscriptionStatus",
        "Value": "PAID"
      },
      {
        "Name": "OfferingSku",
        "Value": "QuickBooks Online Plus"
      },
      {
        "Name": "PayrollFeature",
        "Value": "true"
      },
      {
        "Name": "AccountantFeature",
        "Value": "false"
      },
      {
        "Name": "IsQbdtMigrated",
        "Value": "true"
      },
      {
        "Name": "MigrationDate",
        "Value": "2024-09-14T01:47:34-07:00"
      },
      {
        "Name": "QBOIndustryType",
        "Value": "Manufacturing Businesses"
      },
      {
        "Name": "AssignedTime",
        "Value": "2024-09-14T01:47:34-07:00"
      }
    ],
    "MetaData": {
      "CreateTime": "2015-06-05T13:55:54-07:00",
      "LastUpdatedTime": "2015-07-06T08:51:50-07:00"
    }
  },
  "time": "2015-07-10T09:38:58.155-07:00"
}
```

### Read companyinfo

- **Method**: GET
- **URL**: `/v3/company/{realmID}/companyinfo/{realmID}`

Retrieves the details of the CompanyInfo object.

**Response**:
```json
{
  "CompanyInfo": {
    "SyncToken": "4",
    "domain": "QBO",
    "LegalAddr": {
      "City": "Mountain View",
      "Country": "US",
      "Line1": "2500 Garcia Ave",
      "PostalCode": "94043",
      "CountrySubDivisionCode": "CA",
      "Id": "1"
    },
    "SupportedLanguages": "en",
    "CompanyName": "Larry's Bakery",
    "Country": "US",
    "CompanyAddr": {
      "City": "Mountain View",
      "Country": "US",
      "Line1": "2500 Garcia Ave",
      "PostalCode": "94043",
      "CountrySubDivisionCode": "CA",
      "Id": "1"
    },
    "sparse": false,
    "Id": "1",
    "WebAddr": {},
    "FiscalYearStartMonth": "January",
    "CustomerCommunicationAddr": {
      "City": "Mountain View",
      "Country": "US",
      "Line1": "2500 Garcia Ave",
      "PostalCode": "94043",
      "CountrySubDivisionCode": "CA",
      "Id": "1"
    },
    "PrimaryPhone": {
      "FreeFormNumber": "(650)944-4444"
    },
    "LegalName": "Larry's Bakery",
    "CompanyStartDate": "2015-06-05",
    "EmployerId": "123456789",
    "Email": {
      "Address": "donotreply@intuit.com"
    },
    "NameValue": [
      {
        "Name": "NeoEnabled",
        "Value": "true"
      },
      {
        "Name": "IndustryType",
        "Value": "Bread and Bakery Product Manufacturing"
      },
      {
        "Name": "IndustryCode",
        "Value": "31181"
      },
      {
        "Name": "SubscriptionStatus",
        "Value": "PAID"
      },
      {
        "Name": "OfferingSku",
        "Value": "QuickBooks Online Plus"
      },
      {
        "Name": "PayrollFeature",
        "Value": "true"
      },
      {
        "Name": "AccountantFeature",
        "Value": "false"
      },
      {
        "Name": "IsQbdtMigrated",
        "Value": "true"
      },
      {
        "Name": "MigrationDate",
        "Value": "2024-09-14T01:47:34-07:00"
      },
      {
        "Name": "QBOIndustryType",
        "Value": "Manufacturing Businesses"
      },
      {
        "Name": "AssignedTime",
        "Value": "2024-09-14T01:47:34-07:00"
      }
    ],
    "MetaData": {
      "CreateTime": "2015-06-05T13:55:54-07:00",
      "LastUpdatedTime": "2015-07-06T08:51:50-07:00"
    }
  },
  "time": "2015-07-10T09:38:58.155-07:00"
}
```

### Full update companyinfo

- **Method**: POST
- **URL**: `/v3/company/{realmID}/companyinfo`

Available with minor version 11. Use this operation to update any of the writable fields of the companyinfo object. The request body must include all writable fields of the existing object as returned in a read response. Writable fields omitted from the request body are set to NULL. The ID of the ob...

**Request Body**:
```json
{
  "SyncToken": "3",
  "domain": "QBO",
  "LegalAddr": {
    "City": "Mountain View",
    "Country": "US",
    "Line1": "2500 Garcia Ave",
    "PostalCode": "94043",
    "CountrySubDivisionCode": "CA",
    "Id": "1"
  },
  "SupportedLanguages": "en",
  "CompanyName": "Larry's Bakery",
  "Country": "US",
  "CompanyAddr": {
    "City": "Mountain View",
    "Country": "US",
    "Line1": "2500 Garcia Ave",
    "PostalCode": "94043",
    "CountrySubDivisionCode": "CA",
    "Id": "1"
  },
  "sparse": false,
  "Id": "1",
  "WebAddr": {},
  "FiscalYearStartMonth": "January",
  "CustomerCommunicationAddr": {
    "City": "Mountain View",
    "Country": "US",
    "Line1": "2500 Garcia Ave",
    "PostalCode": "94043",
    "CountrySubDivisionCode": "CA",
    "Id": "1"
  },
  "PrimaryPhone": {
    "FreeFormNumber": "(650)944-4444"
  },
  "LegalName": "Larry's Bakery",
  "CompanyStartDate": "2015-06-05",
  "Email": {
    "Address": "donotreply@intuit.com"
  },
  "NameValue": [
    {
      "Name": "NeoEnabled",
      "Value": "true"
    },
    {
      "Name": "IndustryType",
      "Value": "Bread and Bakery Product Manufacturing"
    },
    {
      "Name": "IndustryCode",
      "Value": "31181"
    },
    {
      "Name": "SubscriptionStatus",
      "Value": "PAID"
    },
    {
      "Name": "OfferingSku",
      "Value": "QuickBooks Online Plus"
    },
    {
      "Name": "PayrollFeature",
      "Value": "true"
    },
    {
      "Name": "AccountantFeature",
      "Value": "false"
    },
    {
      "Name": "IsQbdtMigrated",
      "Value": "true"
    },
    {
      "Name": "MigrationDate",
      "Value": "2024-09-14T01:47:34-07:00"
    },
    {
      "Name": "QBOIndustryType",
      "Value": "Manufacturing Businesses"
    },
    {
      "Name": "AssignedTime",
      "Value": "2024-09-14T01:47:34-07:00"
    }
  ],
  "MetaData": {
    "CreateTime": "2015-06-05T13:55:54-07:00",
    "LastUpdatedTime": "2015-07-06T08:51:50-07:00"
  }
}
```

**Response**:
```json
{
  "CompanyInfo": {
    "SyncToken": "4",
    "domain": "QBO",
    "LegalAddr": {
      "City": "Mountain View",
      "Country": "US",
      "Line1": "2500 Garcia Ave",
      "PostalCode": "94043",
      "CountrySubDivisionCode": "CA",
      "Id": "1"
    },
    "SupportedLanguages": "en",
    "CompanyName": "Larry's Bakery",
    "Country": "US",
    "CompanyAddr": {
      "City": "Mountain View",
      "Country": "US",
      "Line1": "2500 Garcia Ave",
      "PostalCode": "94043",
      "CountrySubDivisionCode": "CA",
      "Id": "1"
    },
    "sparse": false,
    "Id": "1",
    "WebAddr": {},
    "FiscalYearStartMonth": "January",
    "CustomerCommunicationAddr": {
      "City": "Mountain View",
      "Country": "US",
      "Line1": "2500 Garcia Ave",
      "PostalCode": "94043",
      "CountrySubDivisionCode": "CA",
      "Id": "1"
    },
    "PrimaryPhone": {
      "FreeFormNumber": "(650)944-4444"
    },
    "LegalName": "Larry's Bakery",
    "CompanyStartDate": "2015-06-05",
    "EmployerId": "123456789",
    "Email": {
      "Address": "donotreply@intuit.com"
    },
    "NameValue": [
      {
        "Name": "NeoEnabled",
        "Value": "true"
      },
      {
        "Name": "IndustryType",
        "Value": "Bread and Bakery Product Manufacturing"
      },
      {
        "Name": "IndustryCode",
        "Value": "31181"
      },
      {
        "Name": "SubscriptionStatus",
        "Value": "PAID"
      },
      {
        "Name": "OfferingSku",
        "Value": "QuickBooks Online Plus"
      },
      {
        "Name": "PayrollFeature",
        "Value": "true"
      },
      {
        "Name": "AccountantFeature",
        "Value": "false"
      },
      {
        "Name": "IsQbdtMigrated",
        "Value": "true"
      },
      {
        "Name": "MigrationDate",
        "Value": "2024-09-14T01:47:34-07:00"
      },
      {
        "Name": "QBOIndustryType",
        "Value": "Manufacturing Businesses"
      },
      {
        "Name": "AssignedTime",
        "Value": "2024-09-14T01:47:34-07:00"
      }
    ],
    "MetaData": {
      "CreateTime": "2015-06-05T13:55:54-07:00",
      "LastUpdatedTime": "2015-07-06T08:51:50-07:00"
    }
  },
  "time": "2015-07-10T09:38:58.155-07:00"
}
```

### Sparse update companyinfo

- **Method**: POST
- **URL**: `/v3/company/{realmID}/companyinfo`

Sparse updating provides the ability to update a subset of properties for a given object; only elements specified in the request are updated. Missing elements are left untouched. The ID of the object to update is specified in the request body. Available with minor version 11.

**Request Body**:
```json
{
  "SyncToken": "3",
  "CompanyName": "Larry's Bakery",
  "CompanyAddr": {
    "City": "Mountain View",
    "Country": "US",
    "Line1": "2500 Garcia Ave",
    "PostalCode": "94043",
    "CountrySubDivisionCode": "CA",
    "Id": "1"
  },
  "sparse": true,
  "LegalName": "Larry Smith's Bakery",
  "Id": "1"
}
```

**Response**:
```json
{
  "CompanyInfo": {
    "SyncToken": "4",
    "domain": "QBO",
    "LegalAddr": {
      "City": "Mountain View",
      "Country": "US",
      "Line1": "2500 Garcia Ave",
      "PostalCode": "94043",
      "CountrySubDivisionCode": "CA",
      "Id": "1"
    },
    "SupportedLanguages": "en",
    "CompanyName": "Larry's Bakery",
    "Country": "US",
    "CompanyAddr": {
      "City": "Mountain View",
      "Country": "US",
      "Line1": "2500 Garcia Ave",
      "PostalCode": "94043",
      "CountrySubDivisionCode": "CA",
      "Id": "1"
    },
    "sparse": false,
    "Id": "1",
    "WebAddr": {},
    "FiscalYearStartMonth": "January",
    "CustomerCommunicationAddr": {
      "City": "Mountain View",
      "Country": "US",
      "Line1": "2500 Garcia Ave",
      "PostalCode": "94043",
      "CountrySubDivisionCode": "CA",
      "Id": "1"
    },
    "PrimaryPhone": {
      "FreeFormNumber": "(650)944-4444"
    },
    "LegalName": "Larry Smith's Bakery",
    "CompanyStartDate": "2015-06-05",
    "EmployerId": "123456789",
    "Email": {
      "Address": "donotreply@intuit.com"
    },
    "NameValue": [
      {
        "Name": "NeoEnabled",
        "Value": "true"
      },
      {
        "Name": "IndustryType",
        "Value": "Bread and Bakery Product Manufacturing"
      },
      {
        "Name": "IndustryCode",
        "Value": "31181"
      },
      {
        "Name": "SubscriptionStatus",
        "Value": "PAID"
      },
      {
        "Name": "OfferingSku",
        "Value": "QuickBooks Online Plus"
      },
      {
        "Name": "PayrollFeature",
        "Value": "true"
      },
      {
        "Name": "AccountantFeature",
        "Value": "false"
      },
      {
        "Name": "IsQbdtMigrated",
        "Value": "true"
      },
      {
        "Name": "MigrationDate",
        "Value": "2024-09-14T01:47:34-07:00"
      },
      {
        "Name": "QBOIndustryType",
        "Value": "Manufacturing Businesses"
      },
      {
        "Name": "AssignedTime",
        "Value": "2024-09-14T01:47:34-07:00"
      }
    ],
    "MetaData": {
      "CreateTime": "2015-06-05T13:55:54-07:00",
      "LastUpdatedTime": "2015-07-06T08:51:50-07:00"
    }
  },
  "time": "2015-07-10T09:38:58.155-07:00"
}
```
