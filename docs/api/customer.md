# Customer

A customer is a consumer of the service or product that your business offers. An individual customer can have an underlying nested structure, with a parent customer (the top-level object) having zero or more sub-customers and jobs associated with it. • Sub-customer examples: • Members of a team or league: the team itself is the parent customer and the members are sub-customers. • Properties managed by a property management company: the management company is the parent customer and the properties are the sub-customers. • Job examples: • Tracking a kitchen remodel: the home owner is the parent customer and individual kitchen remodel tasks are jobs. • Tracking car repairs: the car owner is the parent customer and individual car repairs are jobs. Use the Customer resource to create parent customer objects, sub-customer objects, and job objects according to your business requirements. Use the ParentRef and Job attributes in the customer object to designate whether the object is a parent, nested job or nested sub-customer. • First, create parent customer objects: Set Job to false (default) and do not define ParentRef. • Then, create sub-customer and job objects: Set Job to true and set ParentRef to reference parent customer object. Going forward, specify an individual parent customer object, sub-customer object, or job object in sales transactions via the transaction's CustomerRef attribute, based on your business requirements.See QuickBooks product documentation for more about sub-customers and jobs.

## Business Rules

- The DisplayName, Title, GivenName, MiddleName, FamilyName, Suffix, and PrintOnCheckName attributes must not contain colon (:), tab (\t), or newline (\n) characters.
- The DisplayName attribute must be unique across all other customer, employee, and vendor objects.
- The PrimaryEmailAddress attribute must contain an at sign (@) and dot (.).
- Nested Customer objects can be used to define sub-customers, jobs, or a combination of both, under a parent.
- Up to four levels of nesting can be defined under a top-level parent Customer object.
- The Job attribute defines whether the object is a parent customer or nested sub-customer/job.
- The DisplayName attribute or at least one of Title, GivenName, MiddleName, FamilyName, or Suffix attributes is required during object create.

## Attributes

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| Id | String | Required for update | Unique identifier for this object. Sort order is ASC by default. |
| SyncToken | String | Required for update | Version number of the object. It is used to lock an object for use by one app at a time. As soon as an application modifies an object, its SyncToken is incremented. |
| DisplayName |  | Conditionally required | The name of the person or organization as displayed. Must be unique across all Customer, Vendor, and Employee objects. Cannot be removed with sparse update. |
| Title |  | Conditionally required | Title of the person. This tag supports i18n, all locales. The DisplayName attribute or at least one of Title, GivenName, MiddleName, FamilyName, or Suffix attributes is required. |
| GivenName |  | Conditionally required | Given name or first name of a person. The DisplayName attribute or at least one of Title, GivenName, MiddleName, FamilyName, or Suffix attributes is required. |
| MiddleName |  | Conditionally required | Middle name of the person. The person can have zero or more middle names. |
| Suffix |  | Conditionally required | Suffix of the name. For example, Jr. The DisplayName attribute or at least one of Title, GivenName, MiddleName, FamilyName, or Suffix attributes is required. |
| FamilyName |  | Conditionally required | Family name or the last name of the person. The DisplayName attribute or at least one of Title, GivenName, MiddleName, FamilyName, or Suffix attributes is required. |
| PrimaryEmailAddr | EmailAddress | Optional | Primary email address. |
| ResaleNum |  | Optional | chars String Resale number or some additional info about the customer. |
| SecondaryTaxIdentifier |  | Optional |  |
| ARAccountRef |  | Optional |  |
| ReferenceType |  | Optional | Identifies the accounts receivable account to be used for this customer. Each customer must have his own AR account. Applicable for France companies, only. |
| DefaultTaxCodeRef | ReferenceType | Optional | Reference to a default tax code associated with this Customer object. Reference is valid if Customer.Taxable is set to true; otherwise, it is ignored. |
| PreferredDeliveryMethod | String | Optional | Preferred delivery method. Values are Print, Email, or None. |
| GSTIN |  | Optional | of 15 chars |
| SalesTermRef | ReferenceType | Optional | Reference to a SalesTerm associated with this Customer object. Query the Term name list resource to determine the appropriate Term object for this reference. Use Term. |
| CustomerTypeRef | String | Optional | Reference to the customer type assigned to a customer. This field is only returned if the customer is assigned a customer type. |
| Fax |  | Optional | of 30 chars TelephoneNumber Fax number. |
| BusinessNumber |  | Optional | of 10 chars |
| BillWithParent | Boolean | Optional | If true, this Customer object is billed with its parent. If false, or null the customer is not to be billed with its parent. |
| CurrencyRef |  | Read-only | chars  CurrencyRef Reference to the currency in which all amounts associated with this customer are expressed. Once set, it cannot be changed. |
| Mobile |  | Optional | of 30 chars TelephoneNumber Mobile phone number. |
| Job | Boolean | Optional | If true, this is a Job or sub-customer. If false or null, this is a top level customer, not a Job or sub-customer. |
| BalanceWithJobs | Decimal | Optional | Cumulative open balance amount for the Customer (or Job) and all its sub-jobs. Cannot be written to QuickBooks. |
| PrimaryPhone |  | Optional | of 30 chars TelephoneNumber Primary phone number. |
| OpenBalanceDate | Date | Optional | Date of the Open Balance for the create operation. Write-on-create. |
| Taxable | Boolean | Optional | If true, transactions for this customer are taxable. Default behavior with minor version 10 and above: true, if DefaultTaxCodeRef is defined or false if TaxExemptionReasonId is set. |
| AlternatePhone |  | Optional | of 30 chars TelephoneNumber Alternate phone number. |
| MetaData | ModificationMetaData | Read-only | Descriptive information about the entity. The MetaData values are set by Data Services and are  for all applications. |
| ParentRef | ReferenceType | Optional | A reference to a Customer object that is the immediate parent of the Sub-Customer/Job in the hierarchical Customer:Job list. |
| Notes |  | Optional | Free form text describing the Customer. |
| WebAddr |  | Optional | of 1000 chars WebSiteAddress Website address. |
| CompanyName |  | Optional | The name of the company associated with the person or organization. |
| Balance | Decimal | Optional | Specifies the open balance amount or the amount unpaid by the customer. For the create operation, this represents the opening balance for the customer. |
| ShipAddr | PhysicalAddress | Optional | Default shipping address. If a physical address is updated from within the transaction object, the QuickBooks Online API flows individual address components differently into t... |
| PaymentMethodRef | ReferenceType | Optional | Reference to a PaymentMethod associated with this Customer object. |
| IsProject | Boolean | Read-only | If true, indicates this is a Project. |
| Source | String | Optional | The Source type of the transactions created by QuickBooks Commerce. Valid values include: QBCommerce |
| PrimaryTaxIdentifier |  | Optional |  |
| GSTRegistrationType |  | Optional | of 15 chars |
| PrintOnCheckName |  | Optional | Name of the person or organization as printed on a check. If not provided, this is populated from DisplayName. |
| BillAddr | PhysicalAddress | Optional | Default billing address. If a physical address is updated from within the transaction object, the QuickBooks Online API flows individual address components differently into th... |
| FullyQualifiedName | String | Read-only | Fully qualified name of the object. The fully qualified name prepends the topmost parent, followed by each sub element separated by colons. Takes the form of Customer:Job:Sub-job. System generated. |
| Level | Integer | Read-only | Specifies the level of the hierarchy in which the entity is located. Zero specifies the top level of the hierarchy; anything above will be level with respect to the parent. Constraints:up to 5 levels |
| TaxExemptionReasonId |  | Optional | Numeric Id The tax exemption reason associated with this customer object. Applicable if automated sales tax is enabled (Preferences.TaxPrefs.PartnerTaxEnabled is set to true) for the company. |

## Sample Object

```json
{
  "Customer": {
    "PrimaryEmailAddr": {
      "Address": "Surf@Intuit.com"
    },
    "SyncToken": "0",
    "domain": "QBO",
    "GivenName": "Bill",
    "DisplayName": "Bill's Windsurf Shop",
    "BillWithParent": false,
    "FullyQualifiedName": "Bill's Windsurf Shop",
    "CompanyName": "Bill's Windsurf Shop",
    "FamilyName": "Lucchini",
    "sparse": false,
    "PrimaryPhone": {
      "FreeFormNumber": "(415) 444-6538"
    },
    "Active": true,
    "Job": false,
    "BalanceWithJobs": 85.0,
    "BillAddr": {
      "City": "Half Moon Bay",
      "Line1": "12 Ocean Dr.",
      "PostalCode": "94213",
      "Lat": "37.4307072",
      "Long": "-122.4295234",
      "CountrySubDivisionCode": "CA",
      "Id": "3"
    },
    "PreferredDeliveryMethod": "Print",
    "Taxable": false,
    "PrintOnCheckName": "Bill's Windsurf Shop",
    "Balance": 85.0,
    "Id": "2",
    "MetaData": {
      "CreateTime": "2014-09-11T16:49:28-07:00",
      "LastUpdatedTime": "2014-09-18T12:56:01-07:00"
    }
  },
  "time": "2015-07-23T11:04:15.496-07:00"
}
```

## Operations

### Create a customer

- **Method**: POST
- **URL**: `/v3/company/{realmID}/customer`

The DisplayName attribute or at least one of Title, GivenName, MiddleName, FamilyName, or Suffix attributes is required during object create.

**Request Body**:
```json
{
  "FullyQualifiedName": "King Groceries",
  "PrimaryEmailAddr": {
    "Address": "jdrew@myemail.com"
  },
  "DisplayName": "King's Groceries",
  "Suffix": "Jr",
  "Title": "Mr",
  "MiddleName": "B",
  "Notes": "Here are other details.",
  "FamilyName": "King",
  "PrimaryPhone": {
    "FreeFormNumber": "(555) 555-5555"
  },
  "CompanyName": "King Groceries",
  "BillAddr": {
    "CountrySubDivisionCode": "CA",
    "City": "Mountain View",
    "PostalCode": "94042",
    "Line1": "123 Main Street",
    "Country": "USA"
  },
  "GivenName": "James"
}
```

**Response**:
```json
{
  "Customer": {
    "domain": "QBO",
    "PrimaryEmailAddr": {
      "Address": "jdrew@myemail.com"
    },
    "DisplayName": "King's Groceries",
    "CurrencyRef": {
      "name": "United States Dollar",
      "value": "USD"
    },
    "DefaultTaxCodeRef": {
      "value": "2"
    },
    "PreferredDeliveryMethod": "Print",
    "GivenName": "James",
    "FullyQualifiedName": "King's Groceries",
    "BillWithParent": false,
    "Title": "Mr",
    "Job": false,
    "BalanceWithJobs": 0,
    "PrimaryPhone": {
      "FreeFormNumber": "(555) 555-5555"
    },
    "Taxable": true,
    "MetaData": {
      "CreateTime": "2015-07-23T10:58:12-07:00",
      "LastUpdatedTime": "2015-07-23T10:58:12-07:00"
    },
    "BillAddr": {
      "City": "Mountain View",
      "Country": "USA",
      "Line1": "123 Main Street",
      "PostalCode": "94042",
      "CountrySubDivisionCode": "CA",
      "Id": "112"
    },
    "MiddleName": "B",
    "Notes": "Here are other details.",
    "Active": true,
    "Balance": 0,
    "SyncToken": "0",
    "Suffix": "Jr",
    "CompanyName": "King Groceries",
    "FamilyName": "King",
    "PrintOnCheckName": "King Groceries",
    "sparse": false,
    "Id": "67"
  },
  "time": "2015-07-23T10:58:12.099-07:00"
}
```

### Query a customer

- **Method**: GET
- **URL**: `/v3/company/{realmID}/query?query={selectStatement}`

**Sample Query**: `select * from Customer Where Metadata.LastUpdatedTime > '2015-03-01'`

**Response**:
```json
{
  "QueryResponse": {
    "Customer": [
      {
        "domain": "QBO",
        "FamilyName": "Lauterbach",
        "DisplayName": "Amy's Bird Sanctuary",
        "DefaultTaxCodeRef": {
          "value": "2"
        },
        "PrimaryEmailAddr": {
          "Address": "Birds@Intuit.com"
        },
        "PreferredDeliveryMethod": "Print",
        "GivenName": "Amy",
        "FullyQualifiedName": "Amy's Bird Sanctuary",
        "BillWithParent": false,
        "Job": false,
        "BalanceWithJobs": 274.0,
        "PrimaryPhone": {
          "FreeFormNumber": "(650) 555-3311"
        },
        "Active": true,
        "MetaData": {
          "CreateTime": "2014-09-11T16:48:43-07:00",
          "LastUpdatedTime": "2015-07-01T10:14:15-07:00"
        },
        "BillAddr": {
          "City": "Bayshore",
          "Line1": "4581 Finch St.",
          "PostalCode": "94326",
          "Lat": "INVALID",
          "Long": "INVALID",
          "CountrySubDivisionCode": "CA",
          "Id": "2"
        },
        "MiddleName": "Michelle",
        "Notes": "Note added via Update operation.",
        "Taxable": true,
        "Balance": 274.0,
        "SyncToken": "5",
        "CompanyName": "Amy's Bird Sanctuary",
        "ShipAddr": {
          "City": "Bayshore",
          "Line1": "4581 Finch St.",
          "PostalCode": "94326",
          "Lat": "INVALID",
          "Long": "INVALID",
          "CountrySubDivisionCode": "CA",
          "Id": "109"
        },
        "PrintOnCheckName": "Amy's Bird Sanctuary",
        "sparse": false,
        "Id": "1"
      },
      {
        "domain": "QBO",
        "PrimaryEmailAddr": {
          "Address": "Consulting@intuit.com"
        },
        "DisplayName": "Weiskopf Consulting",
        "FamilyName": "Weiskopf",
        "PreferredDeliveryMethod": "Print",
        "GivenName": "Nicola",
        "FullyQualifiedName": "Weiskopf Consulting",
        "BillWithParent": false,
        "Job": false,
        "BalanceWithJobs": 390.0,
        "PrimaryPhone": {
          "FreeFormNumber": "(650) 555-1423"
        },
        "Active": true,
        "MetaData": {
          "CreateTime": "2014-09-11T17:29:04-07:00",
          "LastUpdatedTime": "2015-06-24T15:54:02-07:00"
        },
        "BillAddr": {
          "City": "Bayshore",
          "Line1": "45612 Main St.",
          "PostalCode": "94326",
          "Lat": "45.256574",
          "Long": "-66.0943698",
          "CountrySubDivisionCode": "CA",
          "Id": "30"
        },
        "Taxable": false,
        "Balance": 390.0,
        "SyncToken": "0",
        "CompanyName": "Weiskopf Consulting",
        "ShipAddr": {
          "City": "Bayshore",
          "Line1": "45612 Main St.",
          "PostalCode": "94326",
          "Lat": "45.256574",
          "Long": "-66.0943698",
          "CountrySubDivisionCode": "CA",
          "Id": "30"
        },
        "PrintOnCheckName": "Weiskopf Consulting",
        "sparse": false,
        "Id": "29"
      }
    ],
    "startPosition": 1,
    "maxResults": 6
  },
  "time": "2015-07-23T11:02:25.149-07:00"
}
```

### Read a customer

- **Method**: GET
- **URL**: `/v3/company/{realmID}/customer/{customerId}`

Retrieves the details of a Customer object that has been previously created.

**Response**:
```json
{
  "Customer": {
    "PrimaryEmailAddr": {
      "Address": "Surf@Intuit.com"
    },
    "SyncToken": "0",
    "domain": "QBO",
    "GivenName": "Bill",
    "DisplayName": "Bill's Windsurf Shop",
    "BillWithParent": false,
    "FullyQualifiedName": "Bill's Windsurf Shop",
    "CompanyName": "Bill's Windsurf Shop",
    "FamilyName": "Lucchini",
    "sparse": false,
    "PrimaryPhone": {
      "FreeFormNumber": "(415) 444-6538"
    },
    "Active": true,
    "Job": false,
    "BalanceWithJobs": 85.0,
    "BillAddr": {
      "City": "Half Moon Bay",
      "Line1": "12 Ocean Dr.",
      "PostalCode": "94213",
      "Lat": "37.4307072",
      "Long": "-122.4295234",
      "CountrySubDivisionCode": "CA",
      "Id": "3"
    },
    "PreferredDeliveryMethod": "Print",
    "Taxable": false,
    "PrintOnCheckName": "Bill's Windsurf Shop",
    "Balance": 85.0,
    "Id": "2",
    "MetaData": {
      "CreateTime": "2014-09-11T16:49:28-07:00",
      "LastUpdatedTime": "2014-09-18T12:56:01-07:00"
    }
  },
  "time": "2015-07-23T11:04:15.496-07:00"
}
```

### Full update a customer

- **Method**: POST
- **URL**: `/v3/company/{realmID}/customer`

Use this operation to update any of the writable fields of an existing Customer object. The request body must include all writable fields of the existing object as returned in a read response. Writable fields omitted from the request body are set to NULL. The ID of the object to update is specified ...

**Request Body**:
```json
{
  "domain": "QBO",
  "PrimaryEmailAddr": {
    "Address": "Surf@Intuit.com"
  },
  "DisplayName": "Bill's Windsurf Shop",
  "PreferredDeliveryMethod": "Print",
  "GivenName": "Bill",
  "FullyQualifiedName": "Bill's Windsurf Shop",
  "BillWithParent": false,
  "Job": false,
  "BalanceWithJobs": 85.0,
  "PrimaryPhone": {
    "FreeFormNumber": "(415) 444-6538"
  },
  "Active": true,
  "MetaData": {
    "CreateTime": "2014-09-11T16:49:28-07:00",
    "LastUpdatedTime": "2015-07-23T11:07:55-07:00"
  },
  "BillAddr": {
    "City": "Half Moon Bay",
    "Line1": "12 Ocean Dr.",
    "PostalCode": "94213",
    "Lat": "37.4307072",
    "Long": "-122.4295234",
    "CountrySubDivisionCode": "CA",
    "Id": "3"
  },
  "MiddleName": "Mac",
  "Taxable": false,
  "Balance": 85.0,
  "SyncToken": "3",
  "CompanyName": "Bill's Windsurf Shop",
  "FamilyName": "Lucchini",
  "PrintOnCheckName": "Bill's Wind Surf Shop",
  "sparse": false,
  "Id": "2"
}
```

**Response**:
```json
{
  "Customer": {
    "domain": "QBO",
    "PrimaryEmailAddr": {
      "Address": "Surf@Intuit.com"
    },
    "DisplayName": "Bill's Windsurf Shop",
    "PreferredDeliveryMethod": "Print",
    "GivenName": "Bill",
    "FullyQualifiedName": "Bill's Windsurf Shop",
    "BillWithParent": false,
    "Job": false,
    "BalanceWithJobs": 85.0,
    "PrimaryPhone": {
      "FreeFormNumber": "(415) 444-6538"
    },
    "Active": true,
    "MetaData": {
      "CreateTime": "2014-09-11T16:49:28-07:00",
      "LastUpdatedTime": "2015-07-23T11:18:37-07:00"
    },
    "BillAddr": {
      "City": "Half Moon Bay",
      "Line1": "12 Ocean Dr.",
      "PostalCode": "94213",
      "Lat": "37.4307072",
      "Long": "-122.4295234",
      "CountrySubDivisionCode": "CA",
      "Id": "3"
    },
    "MiddleName": "Mac",
    "Taxable": false,
    "Balance": 85.0,
    "SyncToken": "4",
    "CompanyName": "Bill's Windsurf Shop",
    "FamilyName": "Lucchini",
    "PrintOnCheckName": "Bill's Wind Surf Shop",
    "sparse": false,
    "Id": "2"
  },
  "time": "2015-07-23T11:18:37.323-07:00"
}
```

### Sparse update a customer

- **Method**: POST
- **URL**: `/v3/company/{realmID}/customer`

Sparse updating provides the ability to update a subset of properties for a given object; only elements specified in the request are updated. Missing elements are left untouched. The ID of the object to update is specified in the request body.

**Request Body**:
```json
{
  "MiddleName": "Mark",
  "SyncToken": "0",
  "Id": "2",
  "sparse": true
}
```

**Response**:
```json
{
  "Customer": {
    "domain": "QBO",
    "PrimaryEmailAddr": {
      "Address": "Surf@Intuit.com"
    },
    "DisplayName": "Bill's Windsurf Shop",
    "PreferredDeliveryMethod": "Print",
    "GivenName": "Bill",
    "FullyQualifiedName": "Bill's Windsurf Shop",
    "BillWithParent": false,
    "Job": false,
    "BalanceWithJobs": 85.0,
    "PrimaryPhone": {
      "FreeFormNumber": "(415) 444-6538"
    },
    "Active": true,
    "MetaData": {
      "CreateTime": "2014-09-11T16:49:28-07:00",
      "LastUpdatedTime": "2015-07-23T11:07:55-07:00"
    },
    "BillAddr": {
      "City": "Half Moon Bay",
      "Line1": "12 Ocean Dr.",
      "PostalCode": "94213",
      "Lat": "37.4307072",
      "Long": "-122.4295234",
      "CountrySubDivisionCode": "CA",
      "Id": "3"
    },
    "MiddleName": "Mark",
    "Taxable": false,
    "Balance": 85.0,
    "SyncToken": "1",
    "CompanyName": "Bill's Windsurf Shop",
    "FamilyName": "Lucchini",
    "PrintOnCheckName": "Bill's Windsurf Shop",
    "sparse": false,
    "Id": "2"
  },
  "time": "2015-07-23T11:07:55.772-07:00"
}
```
