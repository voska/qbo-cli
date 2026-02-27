# Minor Versions of the QuickBooks Online API

The QuickBooks Online Accounting API supports past versions of our API.

We also maintain "minor versions" of our APIs to make incremental changes.

You can use minor version queries to access a specific version of an API entity that's stable with your app. Using minor versions is optional. Whenever we update the API schema, you can update your code to match the latest version, or keep your existing code with a minor version.

> **Important**: We will be discontinuing support for minor versions 1 through 74 of the QuickBooks Online Accounting API beginning August 1, 2025. We request that you review your app's usage of these minor versions and make any necessary updates before July 31, 2025. For details see [this article](https://blogs.intuit.com/2025/01/21/changes-to-our-accounting-api-that-may-impact-your-application/) on the Intuit Developer blog.

## How minor versions work

You can use the minor version of any API entity. We recommend working with only one minor version at a time, not multiple versions, to avoid schema conflicts.

You can access minor versions on a per resource basis. You don't need to use minor versions for all resources unless you want to.

- If you don't use a minor version, you'll use the base versions of APIs by default. The base version has no minor version changes applied to it.
- If you're using an SDK to develop your app, you're automatically using the latest version of the API schema.
- If you're not developing with an SDK, you can manually set the minor version in your code.

## How to use minor versions with SDKs

The SDK version automatically gives you the latest features of our APIs and all prior versions.

When we update our SDKs, you can update your code to take advantage of new features and version support. Apps using older versions of SDKs will continue to run.

If you're using an SDK, you need to set the minor version parameter as a part of the **serviceContext** object. You can set this in your config file so your app appends the correct minor version to request base URLs.

**For the .NET SDK:**

```csharp
ServiceContext context = new ServiceContext(appToken, realmId, intuitServiceType, reqvalidator);
context.IppConfiguration.MinorVersion.Qbo = "28";
```

**For the Java SDK:**

```java
context = new Context(oauth, appToken, ServiceType.QBO, realmID);
context.setMinorVersion("8");
```

> **Important**: We update our SDKs to support the latest API version by default. Unless you're referencing a specific minor version, you'll use the latest version associated with the SDK.

## How to use minor versions without an SDK

If you want to use a minor version of an API, use the minor version query parameter:

```
https://quickbooks.api.intuit.com/v3/company/<realmId>/<apiEntity>/entityId?minorversion=<#>
```

Replace the following in the parameter:

- **realmId**: The company ID of the QuickBooks Online company you're sending requests to
- **apiEntity**: The name of the entity
- **#**: The minor version number

Here's an example minor version query for minor version 1 of the **journalEntry** entity:

```
https://quickbooks.api.intuit.com/v3/company/<123456789>/journalentry/entityId?minorversion=1
```

> **Note**: XML Schema Definition (XSD) download is available when generating your own deserialized classes. See the table below.

## Minor Version History

| Minor Version | Release Date | New Features |
|--------------|-------------|-------------|
| 75 | Dec 16, 2024 | Added `DefaultTimeZone` field to `CompanyInfo` entity. [XSD](https://static.developer.intuit.com/resources/v3-minor-version-75.zip) |
| 74 | Dec 5, 2024 | API update to use the company's default service item if item ID is not provided in TimeActivity creation. |
| 73 | Jun 24, 2024 | API support for project estimate. [XSD](https://static.developer.intuit.com/resources/v3-minor-version-73.zip) |
| 71 | Mar 21, 2024 | Added support for TimeActivity `TimeChargeId` field. [XSD](https://static.developer.intuit.com/resources/v3-minor-version-71.zip) |
| 70 | Dec 12, 2023 | Added support for US locales to override sales tax without company level restrictions. |
| 69 | Jan 12, 2023 | Added query support for `projectRef` filter. [XSD](https://static.developer.intuit.com/resources/v3-minor-version-69.zip) |
| 68 | Oct 26, 2022 | MTD - Added support for `COA Detail types` for beta launch - `Contra types`. [XSD](https://static.developer.intuit.com/resources/v3-minor-version-68.zip) |
| 67 | Sep 14, 2022 | Added a new detail `account` type for other `debtors`. Fix provided for NPE thrown for `paymentMethodRef`. [XSD](https://static.developer.intuit.com/resources/v3-minor-version-67.zip) |
| 65 | May 11, 2022 | Added `RevenueRecognitionEnabled` and `RecognitionFrequencyType` under ProductAndServicePrefs. Support for parsing special characters, emojis on V3 query requests. [XSD](https://static.developer.intuit.com/resources/v3-minor-version-65.zip) |
| 63 | Oct 20, 2021 | Added support for `Cost Rate` for `Time Activity`, `Employee`, and `Vendor`. [XSD](https://static.developer.intuit.com/resources/v3-minor-version-63.zip) |
| 62 | Jul 14, 2021 | Added support for `OriginalTaxRateId` in `TaxRate`. Added seconds precision for `Start Time` and `End Time` in Time Activity API. [XSD](https://static.developer.intuit.com/resources/v3-minor-version-62.zip) |
| 59 | Mar 11, 2021 | Added support for source on list entities. [XSD](https://static.developer.intuit.com/resources/v3-minor-version-59.zip) |
| 57 | Oct 7, 2020 | Added support for `TaxSlipType` on Vendor entity. [XSD](https://static.developer.intuit.com/resources/v3-minor-version-57.zip) |
| 56 | Sep 23, 2020 | Added support for `LegalName` on Vendor entity. [XSD](https://static.developer.intuit.com/resources/v3-minor-version-56.zip) |
| 55 | Jul 15, 2020 | Added support for `Tax1099` on Vendor entity. [XSD](https://static.developer.intuit.com/resources/v3-minor-version-55.zip) |
| 54 | Jun 3, 2020 | Added support for `BillAddr` and `ShipAddr` override when updating `CustomerRef` for Invoice entity. [XSD](https://static.developer.intuit.com/resources/v3-minor-version-54.zip) |
| 53 | Not Applicable | No schema changes |
| 52 | Nov 13, 2019 | Added support for `RecurDataRef` on Invoice. [XSD](https://static.developer.intuit.com/resources/v3-minor-version-52.zip) |
| 51 | Sep 25, 2019 | Added `TaxCodeId` support on Item entity. Added `IncludeInAnnualTPAR` support on Vendor entity. [XSD](https://static.developer.intuit.com/resources/v3-minor-version-51.zip) |
| 50 | Not Applicable | No schema changes |
| 49 | Jul 31, 2019 | Support for General Ledger report. [XSD](https://static.developer.intuit.com/resources/v3-minor-version-49.zip) |
| 48 | Not Applicable | No schema changes |
| 47 | Jun 5, 2019 | Added support for `RecurringTransaction` entity. [XSD](https://static.developer.intuit.com/resources/v3-minor-version-47.zip) |
| 46 | Not Applicable | No schema changes |
| 45 | Not Applicable | No schema changes |
| 44 | Not Applicable | No schema changes |
| 43 | Not Applicable | No schema changes |
| 42 | Not Applicable | No schema changes |
| 41 | Not Applicable | No schema changes |
| 40 | Jul 18, 2018 | Added support for `TaxClassification` entity. [XSD](https://static.developer.intuit.com/resources/v3-minor-version-40.zip) |
| 39 | Not Applicable | No schema changes |
| 38 | Jun 20, 2018 | Added support for `CreditCardPayment` entity. [XSD](https://static.developer.intuit.com/resources/v3-minor-version-38.zip) |
| 37 | May 16, 2018 | Support for `Budget` entity. [XSD](https://static.developer.intuit.com/resources/v3-minor-version-37.zip) |
| 36 | Apr 18, 2018 | Added `InvoiceLink` attribute for Invoice entity. [XSD](https://static.developer.intuit.com/resources/v3-minor-version-36.zip) |
| 35 | Jan 10, 2018 | Added `ShipFromAddr` attribute for Estimate, Invoice, and SalesReceipt. [XSD](https://static.developer.intuit.com/resources/v3-minor-version-35.zip) |
| 34 | Not Applicable | No schema changes |
| 33 | Nov 8, 2017 | Support for `Exchangerate` entity. [XSD](https://static.developer.intuit.com/resources/v3-minor-version-33.zip) |
| 29 | Jun 14, 2017 | Support for `TransactionList` and `ChangeDataCapture`. [XSD](https://static.developer.intuit.com/resources/v3-minor-version-29.zip) |
| 28 | May 17, 2017 | Support for Customer `PrimaryTaxIdentifier`. [XSD](https://static.developer.intuit.com/resources/v3-minor-version-28.zip) |
| 21 | Jun 22, 2016 | Added `TaxExemptionRef` for automated sales tax. [XSD](https://static.developer.intuit.com/resources/v3-minor-version-21.zip) |
| 8 | Nov 18, 2015 | Added `BillEmailCc` and `BillEmailBcc` attributes for Invoice. [XSD](https://static.developer.intuit.com/resources/v3-minor-version-8.zip) |
| 4 | May 13, 2015 | Added `TransactionLocationType`. [XSD](https://static.developer.intuit.com/resources/v3-minor-version-4.zip) |
| 3 | Apr 8, 2015 | Added `HomeBalance` for multicurrency. [XSD](https://static.developer.intuit.com/resources/v3-minor-version-3.zip) |
| 1 | Jul 30, 2014 | Added `HomeTotalAmt` for multicurrency Invoice entity. [XSD](https://static.developer.intuit.com/resources/v3-minor-version-1.zip) |
