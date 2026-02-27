# Vendor

The Vendor object represents the seller from whom your company purchases any service or product.

## Business Rules

- The DisplayName, Title, GivenName, MiddleName, FamilyName, Suffix, and PrintOnCheckName attributes must not contain colon (:), tab (\t), or newline (\n) characters.
- The DisplayName attribute must be unique across all other Customer, Employee, and Vendor objects.
- The PrimaryEmailAddress attribute must contain an at sign (@) and dot (.).
- The DisplayName attribute or at least one of Title, GivenName, MiddleName, FamilyName, or Suffix attributes is required during object create.

## Attributes

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| Id | String | Required for update | Unique identifier for this object. Sort order is ASC by default. |
| SyncToken | String | Required for update | Version number of the object. It is used to lock an object for use by one app at a time. As soon as an application modifies an object, its SyncToken is incremented. |
| Title |  | Conditionally required | Title of the person. This tag supports i18n, all locales. |
| GivenName |  | Conditionally required | Given name or first name of a person. The DisplayName attribute or at least one of Title, GivenName, MiddleName, FamilyName, or Suffix attributes is required for object create. |
| MiddleName |  | Conditionally required | Middle name of the person. The person can have zero or more middle names. |
| Suffix |  | Conditionally required | Suffix of the name. For example, Jr. The DisplayName attribute or at least one of Title, GivenName, MiddleName, FamilyName, or Suffix attributes is required for object create. |
| FamilyName |  | Conditionally required | Family name or the last name of the person. |
| PrimaryEmailAddr | EmailAddress | Optional | Primary email address. |
| DisplayName |  | Optional | The name of the vendor as displayed. Must be unique across all Vendor, Customer, and Employee objects. Cannot be removed with sparse update. |
| OtherContactInfo |  | Optional | ContactInfo List of ContactInfo entities of any contact info type. |
| APAccountRef |  | Optional |  |
| ReferenceType |  | Optional | Identifies the accounts payable account to be used for this supplier. Each supplier must have his own AP account. Applicable for France companies, only. |
| TermRef | ReferenceType | Optional | Reference to a default Term associated with this Vendor object. Query the Term name list resource to determine the appropriate Term object for this reference. Use Term. |
| Source | String | Optional | The Source type of the transactions created by QuickBooks Commerce. Valid values include: QBCommerce |
| GSTIN |  | Optional | of 15 chars |
| T4AEligible |  | Optional |  |
| Fax | TelephoneNumber | Optional | Fax number. |
| BusinessNumber |  | Optional | of 10 chars |
| CurrencyRef |  | Read-only | CurrencyRef Reference to the currency in which all amounts associated with this vendor are expressed. Once set, it cannot be changed. |
| HasTPAR |  | Optional |  |
| TaxReportingBasis |  | Optional |  |
| Mobile | TelephoneNumber | Optional | Mobile phone number. |
| PrimaryPhone | TelephoneNumber | Optional | Primary phone number. |
| AlternatePhone | TelephoneNumber | Optional | Alternate phone number. |
| MetaData | ModificationMetaData | Read-only | Descriptive information about the object. The MetaData values are set by Data Services and are  for all applications. |
| Vendor1099 | Boolean | Optional | This vendor is an independent contractor; someone who is given a 1099-MISC form at the end of the year. |
| CostRate | BigDecimal | Optional | Pay rate of the vendor |
| BillRate | Decimal | Optional | BillRate can be set to specify this vendor's hourly billing rate. |
| WebAddr | WebSiteAddress | Optional | Website address. |
| T5018Eligible |  | Optional |  |
| CompanyName |  | Optional | The name of the company associated with the person or organization. |
| VendorPaymentBankDetail |  | Optional |  |
| TaxIdentifier |  | Optional | 20 characters String The tax ID of the Person or Organization. The value is masked in responses, exposing only last four characters. |
| AcctNum |  | Optional | Name or number of the account associated with this vendor. |
| GSTRegistrationType |  | Optional | of 15 chars |
| PrintOnCheckName |  | Optional | Name of the person or organization as printed on a check. If not provided, this is populated from DisplayName. |
| BillAddr | PhysicalAddress | Optional | Default billing address. If a physical address is updated from within the transaction object, the QuickBooks Online API flows individual address components differently into th... |
| Balance | Decimal | Read-only | Specifies the open balance amount or the amount unpaid by the customer. For the create operation, this represents the opening balance for the customer. |

## Sample Object

```json
{
  "Vendor": {
    "PrimaryEmailAddr": {
      "Address": "Books@Intuit.com"
    },
    "Vendor1099": false,
    "domain": "QBO",
    "GivenName": "Bessie",
    "DisplayName": "Books by Bessie",
    "BillAddr": {
      "City": "Palo Alto",
      "Line1": "15 Main St.",
      "PostalCode": "94303",
      "Lat": "37.445013",
      "Long": "-122.1391443",
      "CountrySubDivisionCode": "CA",
      "Id": "31"
    },
    "SyncToken": "0",
    "PrintOnCheckName": "Books by Bessie",
    "FamilyName": "Williams",
    "PrimaryPhone": {
      "FreeFormNumber": "(650) 555-7745"
    },
    "AcctNum": "1345",
    "CompanyName": "Books by Bessie",
    "WebAddr": {
      "URI": "http://www.booksbybessie .co"
    },
    "sparse": false,
    "Active": true,
    "Balance": 0,
    "Id": "30",
    "MetaData": {
      "CreateTime": "2014-09-12T10:07:56-07:00",
      "LastUpdatedTime": "2014-09-17T11:13:46-07:00"
    }
  },
  "time": "2015-07-28T13:33:09.453-07:00"
}
```

## Operations

### Create a vendor

- **Method**: POST
- **URL**: `/v3/company/{realmID}/vendor`

Either the DisplayName attribute or at least one of Title, GivenName, MiddleName, FamilyName, or Suffix attributes are required during create.

**Request Body**:
```json
{
  "PrimaryEmailAddr": {
    "Address": "dbradley@myemail.com"
  },
  "WebAddr": {
    "URI": "http://DiannesAutoShop.com"
  },
  "PrimaryPhone": {
    "FreeFormNumber": "(650) 555-2342"
  },
  "DisplayName": "Dianne's Auto Shop",
  "Suffix": "Sr.",
  "Title": "Ms.",
  "Mobile": {
    "FreeFormNumber": "(650) 555-2000"
  },
  "FamilyName": "Bradley",
  "TaxIdentifier": "99-5688293",
  "AcctNum": "35372649",
  "CompanyName": "Dianne's Auto Shop",
  "BillAddr": {
    "City": "Millbrae",
    "Country": "U.S.A",
    "Line3": "29834 Mustang Ave.",
    "Line2": "Dianne Bradley",
    "Line1": "Dianne's Auto Shop",
    "PostalCode": "94030",
    "CountrySubDivisionCode": "CA"
  },
  "GivenName": "Dianne",
  "PrintOnCheckName": "Dianne's Auto Shop"
}
```

**Response**:
```json
{
  "Vendor": {
    "domain": "QBO",
    "PrimaryEmailAddr": {
      "Address": "dbradley@myemail.com"
    },
    "DisplayName": "Dianne's Auto Shop",
    "CurrencyRef": {
      "name": "United States Dollar",
      "value": "USD"
    },
    "GivenName": "Dianne",
    "Title": "Ms.",
    "PrimaryPhone": {
      "FreeFormNumber": "(650) 555-2342"
    },
    "Active": true,
    "MetaData": {
      "CreateTime": "2015-07-28T12:51:21-07:00",
      "LastUpdatedTime": "2015-07-28T12:51:21-07:00"
    },
    "Vendor1099": false,
    "BillAddr": {
      "City": "Millbrae",
      "Country": "U.S.A",
      "Line3": "29834 Mustang Ave.",
      "Line2": "Dianne Bradley",
      "Line1": "Dianne's Auto Shop",
      "PostalCode": "94030",
      "CountrySubDivisionCode": "CA",
      "Id": "423"
    },
    "Mobile": {
      "FreeFormNumber": "(650) 555-2000"
    },
    "WebAddr": {
      "URI": "http://DiannesAutoShop.com"
    },
    "Balance": 0,
    "SyncToken": "0",
    "Suffix": "Sr.",
    "CompanyName": "Dianne's Auto Shop",
    "FamilyName": "Bradley",
    "TaxIdentifier": "99-5688293",
    "AcctNum": "35372649",
    "PrintOnCheckName": "Dianne's Auto Shop",
    "sparse": false,
    "Id": "137"
  },
  "time": "2015-07-28T12:51:21.326-07:00"
}
```

### Query a vendor

- **Method**: GET
- **URL**: `/v3/company/{realmID}/query?query={selectStatement}`

**Sample Query**: `select * from vendor where MetaData.LastUpdatedTime > '2014-09-17T15:28:48-07:00'`

**Response**:
```json
{
  "QueryResponse": {
    "startPosition": 1,
    "Vendor": [
      {
        "Vendor1099": false,
        "domain": "QBO",
        "DisplayName": "Bob's Burger Joint",
        "SyncToken": "0",
        "PrintOnCheckName": "Bob's Burger Joint",
        "sparse": false,
        "Active": true,
        "Balance": 390.0,
        "Id": "56",
        "MetaData": {
          "CreateTime": "2014-10-03T14:28:52-07:00",
          "LastUpdatedTime": "2015-07-14T12:37:57-07:00"
        }
      },
      {
        "Vendor1099": false,
        "domain": "QBO",
        "DisplayName": "Cal Telephone",
        "BillAddr": {
          "City": "Palo Alto",
          "Line1": "10 Main St.",
          "PostalCode": "94303",
          "Lat": "37.445013",
          "Long": "-122.1391443",
          "CountrySubDivisionCode": "CA",
          "Id": "33"
        },
        "SyncToken": "0",
        "CompanyName": "Cal Telephone",
        "TermRef": {
          "value": "1"
        },
        "PrimaryPhone": {
          "FreeFormNumber": "(650) 555-1616"
        },
        "PrintOnCheckName": "Cal Telephone",
        "sparse": false,
        "Active": true,
        "Balance": 0,
        "Id": "32",
        "MetaData": {
          "CreateTime": "2014-09-12T10:13:24-07:00",
          "LastUpdatedTime": "2014-09-19T12:55:23-07:00"
        }
      },
      {
        "Vendor1099": false,
        "domain": "QBO",
        "GivenName": "Melanie",
        "DisplayName": "Hall Properties",
        "sparse": false,
        "SyncToken": "0",
        "Mobile": {
          "FreeFormNumber": "(973) 888-6222"
        },
        "PrintOnCheckName": "Hall Properties",
        "PrimaryPhone": {
          "FreeFormNumber": "(973) 555-3827"
        },
        "FamilyName": "Hall",
        "TaxIdentifier": "XXXXXXXX2222",
        "AcctNum": "55642",
        "CompanyName": "Hall Properties",
        "WebAddr": {
          "URI": "http://www .hallproperties.intuit .org"
        },
        "BillAddr": {
          "City": "South Orange",
          "Line1": "P.O.Box 357",
          "PostalCode": "07079",
          "Lat": "40.7489277",
          "Long": "-74.2609903",
          "CountrySubDivisionCode": "NJ",
          "Id": "36"
        },
        "Active": true,
        "Balance": 0,
        "Id": "40",
        "MetaData": {
          "CreateTime": "2014-09-12T10:24:28-07:00",
          "LastUpdatedTime": "2014-09-18T13:43:08-07:00"
        }
      },
      {
        "Vendor1099": false,
        "domain": "QBO",
        "GivenName": "Geoff",
        "DisplayName": "Hicks Hardware",
        "BillAddr": {
          "City": "Middlefield",
          "Line1": "42 Main St.",
          "PostalCode": "94303",
          "Lat": "37.445013",
          "Long": "-122.1391443",
          "CountrySubDivisionCode": "CA",
          "Id": "37"
        },
        "SyncToken": "0",
        "Mobile": {
          "FreeFormNumber": "(650) 445-6666"
        },
        "PrintOnCheckName": "Hicks Hardware",
        "FamilyName": "Hicks",
        "PrimaryPhone": {
          "FreeFormNumber": "(650) 554-1973"
        },
        "AcctNum": "556223",
        "CompanyName": "Hicks Hardware",
        "WebAddr": {
          "URI": "http://Hickshardware .co"
        },
        "sparse": false,
        "Active": true,
        "Balance": 0,
        "Id": "41",
        "MetaData": {
          "CreateTime": "2014-09-12T10:26:56-07:00",
          "LastUpdatedTime": "2014-09-18T13:01:57-07:00"
        }
      },
      {
        "PrimaryEmailAddr": {
          "Address": "Materials@intuit.com"
        },
        "Vendor1099": false,
        "domain": "QBO",
        "GivenName": "Julie",
        "DisplayName": "Norton Lumber and Building Materials",
        "BillAddr": {
          "City": "Middlefield",
          "Line1": "4528 Country Road",
          "PostalCode": "94303",
          "Lat": "37.3752919",
          "Long": "-122.1692159",
          "CountrySubDivisionCode": "CA",
          "Id": "40"
        },
        "SyncToken": "0",
        "PrintOnCheckName": "Norton Lumber and Building Materials",
        "FamilyName": "Norton",
        "PrimaryPhone": {
          "FreeFormNumber": "(650) 363-6578"
        },
        "AcctNum": "32980256",
        "CompanyName": "Norton Lumber and Building Materials",
        "sparse": false,
        "Active": true,
        "Balance": 0,
        "Id": "46",
        "MetaData": {
          "CreateTime": "2014-09-12T10:32:55-07:00",
          "LastUpdatedTime": "2015-01-16T16:00:29-08:00"
        }
      },
      {
        "Vendor1099": false,
        "domain": "QBO",
        "PrimaryEmailAddr": {
          "Address": "utilities@noemail.com"
        },
        "DisplayName": "PG&E",
        "BillAddr": {
          "City": "Palo Alto",
          "Line1": "4 Main St.",
          "PostalCode": "94303",
          "Lat": "37.445013",
          "Long": "-122.1391443",
          "CountrySubDivisionCode": "CA",
          "Id": "42"
        },
        "SyncToken": "1",
        "CompanyName": "PG&E",
        "PrimaryPhone": {
          "FreeFormNumber": "(888) 555-9465"
        },
        "AcctNum": "00649587213",
        "PrintOnCheckName": "PG&E",
        "sparse": false,
        "Active": true,
        "Balance": 0,
        "Id": "48",
        "MetaData": {
          "CreateTime": "2014-09-12T10:36:57-07:00",
          "LastUpdatedTime": "2015-01-16T15:36:20-08:00"
        }
      },
      {
        "Vendor1099": false,
        "domain": "QBO",
        "DisplayName": "QuickBooks Payments",
        "SyncToken": "0",
        "PrintOnCheckName": "QuickBooks Payments",
        "sparse": false,
        "Active": true,
        "Balance": 0,
        "Id": "63",
        "MetaData": {
          "CreateTime": "2015-04-13T13:42:23-07:00",
          "LastUpdatedTime": "2015-04-13T13:42:23-07:00"
        }
      },
      {
        "Vendor1099": false,
        "domain": "QBO",
        "GivenName": "Jenny",
        "DisplayName": "Robertson & Associates",
        "BillAddr": {
          "City": "Bayshore",
          "Line1": "P.O. Box 147",
          "PostalCode": "94326",
          "Lat": "45.2720537",
          "Long": "-79.7935909",
          "CountrySubDivisionCode": "CA",
          "Id": "43"
        },
        "SyncToken": "0",
        "PrintOnCheckName": "Robertson & Associates",
        "FamilyName": "Robertson",
        "PrimaryPhone": {
          "FreeFormNumber": "(650) 557-1111"
        },
        "AcctNum": "000005641",
        "CompanyName": "Robertson & Associates",
        "sparse": false,
        "Active": true,
        "Balance": 95.0,
        "Id": "49",
        "MetaData": {
          "CreateTime": "2014-09-12T10:38:12-07:00",
          "LastUpdatedTime": "2015-06-30T15:09:07-07:00"
        }
      },
      {
        "Vendor1099": false,
        "domain": "QBO",
        "DisplayName": "Squeaky Kleen Car Wash",
        "SyncToken": "0",
        "PrintOnCheckName": "Squeaky Kleen Car Wash",
        "sparse": false,
        "Active": true,
        "Balance": 0,
        "Id": "57",
        "MetaData": {
          "CreateTime": "2014-10-03T14:29:35-07:00",
          "LastUpdatedTime": "2014-10-03T14:29:35-07:00"
        }
      },
      {
        "PrimaryEmailAddr": {
          "Address": "tim .philip@timphilipmasonry.com"
        },
        "Vendor1099": false,
        "domain": "QBO",
        "GivenName": "Tim",
        "DisplayName": "Tim Philip Masonry",
        "sparse": false,
        "SyncToken": "0",
        "Mobile": {
          "FreeFormNumber": "(650) 555-1549"
        },
        "PrintOnCheckName": "Tim Philip Masonry",
        "PrimaryPhone": {
          "FreeFormNumber": "(800) 556-1254"
        },
        "FamilyName": "Philip",
        "TaxIdentifier": "XXXXXXXX5555",
        "AcctNum": "0078965",
        "CompanyName": "Tim Philip Masonry",
        "WebAddr": {
          "URI": "http://www .bricksbytim4less.co"
        },
        "BillAddr": {
          "City": "Middlefield",
          "Line1": "3948 Elm St.",
          "PostalCode": "94482",
          "Lat": "37.4604972",
          "Long": "-122.1547528",
          "CountrySubDivisionCode": "CA",
          "Id": "45"
        },
        "Active": true,
        "Balance": 0,
        "Id": "51",
        "MetaData": {
          "CreateTime": "2014-09-12T10:42:31-07:00",
          "LastUpdatedTime": "2014-09-18T13:06:58-07:00"
        }
      }
    ],
    "maxResults": 10
  },
  "time": "2015-07-28T13:29:10.643-07:00"
}
```

### Read a vendor

- **Method**: GET
- **URL**: `/v3/company/{realmID}/vendor/{vendorId}`

Retrieves the details of a Vendor object that has been previously created.

**Response**:
```json
{
  "Vendor": {
    "PrimaryEmailAddr": {
      "Address": "Books@Intuit.com"
    },
    "Vendor1099": false,
    "domain": "QBO",
    "GivenName": "Bessie",
    "DisplayName": "Books by Bessie",
    "BillAddr": {
      "City": "Palo Alto",
      "Line1": "15 Main St.",
      "PostalCode": "94303",
      "Lat": "37.445013",
      "Long": "-122.1391443",
      "CountrySubDivisionCode": "CA",
      "Id": "31"
    },
    "SyncToken": "0",
    "PrintOnCheckName": "Books by Bessie",
    "FamilyName": "Williams",
    "PrimaryPhone": {
      "FreeFormNumber": "(650) 555-7745"
    },
    "AcctNum": "1345",
    "CompanyName": "Books by Bessie",
    "WebAddr": {
      "URI": "http://www.booksbybessie .co"
    },
    "sparse": false,
    "Active": true,
    "Balance": 0,
    "Id": "30",
    "MetaData": {
      "CreateTime": "2014-09-12T10:07:56-07:00",
      "LastUpdatedTime": "2014-09-17T11:13:46-07:00"
    }
  },
  "time": "2015-07-28T13:33:09.453-07:00"
}
```

### Full update a vendor

- **Method**: POST
- **URL**: `/v3/company/{realmID}/vendor`

Use this operation to update any of the writable fields of an existing Vendor object. The request body must include all writable fields of the existing object as returned in a read response. Writable fields omitted from the request body are set to NULL. The ID of the object to update is specified in...

**Request Body**:
```json
{
  "PrimaryEmailAddr": {
    "Address": "Books@Intuit.com"
  },
  "Vendor1099": false,
  "domain": "QBO",
  "GivenName": "Bessie",
  "DisplayName": "Books by Bessie",
  "BillAddr": {
    "City": "Palo Alto",
    "Line1": "15 Main St.",
    "PostalCode": "94303",
    "Lat": "37.445013",
    "Long": "-122.1391443",
    "CountrySubDivisionCode": "CA",
    "Id": "31"
  },
  "SyncToken": "1",
  "PrintOnCheckName": "Books by Bessie and Joan",
  "FamilyName": "Williams",
  "PrimaryPhone": {
    "FreeFormNumber": "(650) 555-7745"
  },
  "AcctNum": "13451234",
  "CompanyName": "Books by Bessie",
  "WebAddr": {
    "URI": "http://www.booksbybessie .co"
  },
  "sparse": false,
  "Active": true,
  "Balance": 0,
  "Id": "30",
  "MetaData": {
    "CreateTime": "2014-09-12T10:07:56-07:00",
    "LastUpdatedTime": "2015-07-28T13:34:38-07:00"
  }
}
```

**Response**:
```json
{
  "Vendor": {
    "PrimaryEmailAddr": {
      "Address": "Books@Intuit.com"
    },
    "Vendor1099": false,
    "domain": "QBO",
    "GivenName": "Bessie",
    "DisplayName": "Books by Bessie",
    "BillAddr": {
      "City": "Palo Alto",
      "Line1": "15 Main St.",
      "PostalCode": "94303",
      "Lat": "37.445013",
      "Long": "-122.1391443",
      "CountrySubDivisionCode": "CA",
      "Id": "31"
    },
    "SyncToken": "2",
    "PrintOnCheckName": "Books by Bessie and Joan",
    "FamilyName": "Williams",
    "PrimaryPhone": {
      "FreeFormNumber": "(650) 555-7745"
    },
    "AcctNum": "13451234",
    "CompanyName": "Books by Bessie",
    "WebAddr": {
      "URI": "http://www.booksbybessie .co"
    },
    "sparse": false,
    "Active": true,
    "Balance": 0,
    "Id": "30",
    "MetaData": {
      "CreateTime": "2014-09-12T10:07:56-07:00",
      "LastUpdatedTime": "2015-07-28T13:37:05-07:00"
    }
  },
  "time": "2015-07-28T13:37:07.196-07:00"
}
```
