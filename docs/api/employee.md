# Employee

An Employee object represents a person working for the company. If you are looking to create a Contractor via API, refer how to create a Vendor object, with Vendor1099 field set to true. • The DisplayName, Title, GivenName, MiddleName, FamilyName, Suffix, and PrintOnCheckName attributes must not contain colon (:), tab (\t), or newline (\n) characters. • The DisplayName attribute must be unique across all other customer, employee, and vendor objects. • The GivenName and FamilyName attributes are required. • The PrimaryEmailAddress attribute must contain an at sign (@) and dot (.). The full complement of read, create, delete via deactivation (active=false), and update operations are available both with and without QuickBooks Payroll enabled. However, when Payroll is enabled, support for some attributes is limited: • Title—Not supported when QuickBooks Payroll is enabled. • Suffix—Not supported when QuickBooks Payroll is enabled. • DisplayName —It’s read only when QuickBooks Payroll is enabled and a concatenation of GivenName MiddleName FamilyName. • PrintOnCheckName—Not supported when QuickBooks Payroll is enabled. • BillRate—Not supported when QuickBooks Payroll is enabled. • SSN—Masked SSNs, as is returned in a response, cannot be passed in a request when QuickBooks Payroll is enabled. Code for this field must be removed before submitting.

## Attributes

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| Id | String | Required for update | Unique identifier for this object. Sort order is ASC by default. |
| SyncToken | String | Required for update | Version number of the object. It is used to lock an object for use by one app at a time. As soon as an application modifies an object, its SyncToken is incremented. |
| PrimaryAddr |  | Conditionally required | 30 characters PhysicalAddress Represents the physical street address for this employee. If QuickBooks Payroll is enabled for the company, the following PhysicalAddress fields are required: |
| PrimaryEmailAddr | EmailAddress | Optional | Primary email address. |
| DisplayName |  | Optional | The name of the person or organization as displayed. |
| Title |  | Optional | 16 chars String Title of the person. This tag supports i18n, all locale. Not supported when QuickBooks Payroll is enabled. |
| BillableTime | Boolean | Optional | If true, this entity is currently enabled for use by QuickBooks. |
| GivenName |  | Optional | Given name or family name of a person. At least one of GivenName or FamilyName attributes is required. |
| BirthDate | Date | Optional | Birth date of the employee. |
| MiddleName |  | Optional | Middle name of the person. The person can have zero or more middle names. |
| SSN |  | Optional | 100 chars String Social security number (SSN) of the employee. If SSN is set, it is masked in the response with XXX-XX-XXXX. |
| PrimaryPhone |  | Optional | of 20 chars TelephoneNumber Primary phone number. |
| ReleasedDate | Date | Optional | Release date of the employee. |
| MetaData | ModificationMetaData | Read-only | Descriptive information about the object. The MetaData values are set by Data Services and are  for all applications. |
| CostRate | BigDecimal | Optional | Pay rate of the employee |
| Mobile |  | Optional | of 20 chars TelephoneNumber Mobile phone number. |
| Gender | String | Optional | Gender of the employee. To clear the gender value, set to Null in a full update request. Supported values include: Male or Female. |
| HiredDate | Date | Optional | Hire date of the employee. |
| BillRate | BigDecimal | Optional | This attribute can only be set if BillableTime is true. Not supported when QuickBooks Payroll is enabled. |
| Organization | Boolean | Optional | true--the object represents an organization. false--the object represents a person. |
| Suffix |  | Optional | Suffix of the name. For example, Jr. Not supported when QuickBooks Payroll is enabled. |
| FamilyName |  | Optional | Family name or the last name of the person. At least one of GivenName or FamilyName attributes is required. |
| PrintOnCheckName |  | Optional | Name of the person or organization as printed on a check. If not provided, this is populated from DisplayName. |
| EmployeeNumber |  | Optional | 100 chars String Specifies the ID number of the employee in the employer's directory. |
| V4IDPseudonym | String | Read-only | Employee reference number. For internal use only. |

## Sample Object

```json
{
  "Employee": {
    "SyncToken": "0",
    "domain": "QBO",
    "DisplayName": "Bill Miller",
    "PrimaryPhone": {
      "FreeFormNumber": "234-525-1234"
    },
    "PrintOnCheckName": "Bill Miller",
    "FamilyName": "Miller",
    "Active": true,
    "SSN": "XXX-XX-XXXX",
    "PrimaryAddr": {
      "CountrySubDivisionCode": "CA",
      "City": "Middlefield",
      "PostalCode": "93242",
      "Id": "116",
      "Line1": "45 N. Elm Street"
    },
    "sparse": false,
    "BillableTime": false,
    "GivenName": "Bill",
    "Id": "71",
    "MetaData": {
      "CreateTime": "2015-07-24T09:34:35-07:00",
      "LastUpdatedTime": "2015-07-24T09:34:35-07:00"
    }
  },
  "time": "2015-07-24T09:35:54.805-07:00"
}
```

## Operations

### Create an employee

- **Method**: POST
- **URL**: `/v3/company/{realmID}/employee`

**Request Body**:
```json
{
  "GivenName": "John",
  "SSN": "444-55-6666",
  "PrimaryAddr": {
    "CountrySubDivisionCode": "CA",
    "City": "Middlefield",
    "PostalCode": "93242",
    "Id": "50",
    "Line1": "45 N. Elm Street"
  },
  "PrimaryPhone": {
    "FreeFormNumber": "408-525-1234"
  },
  "FamilyName": "Meuller"
}
```

**Response**:
```json
{
  "Employee": {
    "SyncToken": "0",
    "domain": "QBO",
    "DisplayName": "John Meuller",
    "PrimaryPhone": {
      "FreeFormNumber": "408-525-1234"
    },
    "PrintOnCheckName": "John Meuller",
    "FamilyName": "Meuller",
    "Active": true,
    "SSN": "XXX-XX-XXXX",
    "PrimaryAddr": {
      "CountrySubDivisionCode": "CA",
      "City": "Middlefield",
      "PostalCode": "93242",
      "Id": "115",
      "Line1": "45 N. Elm Street"
    },
    "sparse": false,
    "BillableTime": false,
    "GivenName": "John",
    "Id": "70",
    "MetaData": {
      "CreateTime": "2015-07-24T09:24:57-07:00",
      "LastUpdatedTime": "2015-07-24T09:24:57-07:00"
    }
  },
  "time": "2015-07-24T09:24:57.907-07:00"
}
```

### Query an employee

- **Method**: GET
- **URL**: `/v3/company/{realmID}/query?query={selectStatement}`

**Sample Query**: `select * from Employee where DisplayName = 'Emily Platt'`

**Response**:
```json
{
  "QueryResponse": {
    "Employee": [
      {
        "SyncToken": "2",
        "domain": "QBO",
        "DisplayName": "Emily Platt",
        "MiddleName": "Jane",
        "FamilyName": "Platt",
        "Active": true,
        "PrintOnCheckName": "Emily Platt",
        "sparse": false,
        "BillableTime": false,
        "GivenName": "Emily",
        "Id": "55",
        "MetaData": {
          "CreateTime": "2014-09-17T11:21:48-07:00",
          "LastUpdatedTime": "2015-07-01T11:29:40-07:00"
        }
      }
    ],
    "startPosition": 1,
    "maxResults": 1
  },
  "time": "2015-07-24T08:56:55.808-07:00"
}
```

### Read an employee

- **Method**: GET
- **URL**: `/v3/company/{realmID}/employee/{employeeId}`

Retrieves the details of a Employee object that has been previously created.

**Response**:
```json
{
  "Employee": {
    "SyncToken": "0",
    "domain": "QBO",
    "DisplayName": "Bill Miller",
    "PrimaryPhone": {
      "FreeFormNumber": "234-525-1234"
    },
    "PrintOnCheckName": "Bill Miller",
    "FamilyName": "Miller",
    "Active": true,
    "SSN": "XXX-XX-XXXX",
    "PrimaryAddr": {
      "CountrySubDivisionCode": "CA",
      "City": "Middlefield",
      "PostalCode": "93242",
      "Id": "116",
      "Line1": "45 N. Elm Street"
    },
    "sparse": false,
    "BillableTime": false,
    "GivenName": "Bill",
    "Id": "71",
    "MetaData": {
      "CreateTime": "2015-07-24T09:34:35-07:00",
      "LastUpdatedTime": "2015-07-24T09:34:35-07:00"
    }
  },
  "time": "2015-07-24T09:35:54.805-07:00"
}
```

### Full update an employee

- **Method**: POST
- **URL**: `/v3/company/{realmID}/employee`

Use this operation to update any of the writable fields of an existing employee object. The request body must include all writable fields of the existing object as returned in a read response. Writable fields omitted from the request body are set to NULL. The ID of the object to update is specified ...

**Request Body**:
```json
{
  "SyncToken": "0",
  "domain": "QBO",
  "DisplayName": "Bill Miller",
  "PrimaryPhone": {
    "FreeFormNumber": "234-525-1234"
  },
  "PrintOnCheckName": "Bill Lee Miller",
  "FamilyName": "Miller",
  "Active": true,
  "SSN": "XXX-XX-XXXX",
  "PrimaryAddr": {
    "CountrySubDivisionCode": "CA",
    "City": "Middlefield",
    "PostalCode": "93242",
    "Id": "116",
    "Line1": "45 N. Elm Street"
  },
  "sparse": false,
  "BillableTime": false,
  "GivenName": "Bill",
  "Id": "71",
  "MetaData": {
    "CreateTime": "2015-07-24T09:34:35-07:00",
    "LastUpdatedTime": "2015-07-24T09:34:35-07:00"
  }
}
```

**Response**:
```json
{
  "Employee": {
    "SyncToken": "1",
    "domain": "QBO",
    "DisplayName": "Bill Miller",
    "PrimaryPhone": {
      "FreeFormNumber": "234-525-1234"
    },
    "PrintOnCheckName": "Bill Lee Miller",
    "FamilyName": "Miller",
    "Active": true,
    "SSN": "XXX-XX-XXXX",
    "PrimaryAddr": {
      "CountrySubDivisionCode": "CA",
      "City": "Middlefield",
      "PostalCode": "93242",
      "Id": "116",
      "Line1": "45 N. Elm Street"
    },
    "sparse": false,
    "BillableTime": false,
    "GivenName": "Bill",
    "Id": "71",
    "MetaData": {
      "CreateTime": "2015-07-24T09:34:35-07:00",
      "LastUpdatedTime": "2015-07-24T09:37:39-07:00"
    }
  },
  "time": "2015-07-24T09:37:39.399-07:00"
}
```
