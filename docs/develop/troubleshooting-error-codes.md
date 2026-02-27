# Review QuickBooks Online Accounting API Error Codes

You may see errors during development or when your app is live and in production. Use the error codes to identify the error. Follow the error's message to narrow down the specific issue.

## HTTP Status Codes

Parse the response code to get more details for the request status.

| HTTP status code | Description |
|---|---|
| 200 OK | The request succeeded. However, the response body may contain a `<Fault>` element that identifies the error. Parse the response to get specific details. |
| 400 Bad request | The request can't be fulfilled due to bad syntax. In some cases, this response code may be for requests with bad authorization data. |
| 401 Unauthorized | Authentication or authorization failed. Usually, this means the token in use won't work for API calls since it's either expired or revoked. |
| 403 Forbidden | The URL exists, but it's restricted. External developers can't use or consume resources from this URL. |
| 404 Not Found | Couldn't find the requested resource or URL, or it doesn't exist. |
| 500 Internal Server Error | A server error occurred while processing the request. Resubmit the request once. If the problem persists, [contact our support team](https://help.developer.intuit.com/s/contactsupport). |
| 503 Service Unavailable | The service is temporarily unavailable. Here's how to [check the status of our API services](https://developer.intuit.com/app/developer/qbo/docs/develop/troubleshooting/api-status). |

## Fault and Error Elements

This is the general `<Fault>` message format:

```xml
<Fault type="faultType">
  <Error code="code" element="element">
    <Message>User-readable text that briefly describes the error.</Message>
    <Detail>Defines details of the exact error.</Detail>
  </Error>
</Fault>
```

| Field | Description |
|---|---|
| `error.code` | The error number. |
| `element` | The name of the field associated with the error, if applicable. |
| `faultType` | The error type. |
| `message` | Details about the error. |

The `<Fault>` element is in either the `<IntuitResponse>` or `<BatchItemResp>` element of the server response body:

| faultType | Description |
|---|---|
| `ValidationFault` | Indicates a potential error in your request. Review your request. There may be invalid code or data. Fix the request payload and header and it will succeed. |
| `SystemFault` | This means there's a problem with our servers. You can't fix this. Here's how to [check the status of our API services](https://developer.intuit.com/app/developer/qbo/docs/develop/troubleshooting/api-status). |
| `AuthenticationFault` | Review the authorization credentials in your request. They may not be correct. |
| `AuthorizationFault` | This means your user hasn't granted your app authorization for the specified resource. Since you don't have authorization, the request failed. We require users (who are admins in QuickBooks Online) to authorize access to all resources when they first connect to and authorize your app. As long as the authorization process is in place, this shouldn't happen for our APIs. |

Here's an example server response for `<Fault>`:

```xml
<IntuitResponse xmlns="http://schema.intuit.com/finance/v3" time="2013-03-29T15:58:08.976-07:00">
  <Fault type="ValidationFault">
    <Error code="2050" element="firstname">
      <Message>Length exceeds limit</Message>
      <Detail>Length of the field exceeds 21 chars</Detail>
    </Error>
    <Error code="2080" element="postalcode">
      <Message>Illegal number format</Message>
      <Detail>ZipCode should be a number with at least 5 digits</Detail>
    </Error>
  </Fault>
</IntuitResponse>
```

## Error Codes

The following tables describe the error codes contained in the `<Error>` element:

- The **Error Code** column lists values for `<Error code>`.
- The **Message** column lists values for `<Message>`.
- The **Detail** column gives possible solutions.

Here's an example error response:

```xml
<Error code="2050" element="firstname">
  <Message>Length exceeds limit</Message>
  <Detail>Length of the field exceeds 21 chars</Detail>
</Error>
```

### Success Response

| Error code | Message | Detail |
|---|---|---|
| 0 | Success | Success response |

### Authorization and Authentication Error Codes

| Error code | Message | Detail |
|---|---|---|
| 100 | General Authentication Error | Review the XXX error code and follow its recommendation. AuthenticationErrorGeneral: XXX |
| 110 | Authentication OAuth Error | Review the XXX error code and follow its recommendation. AuthenticationOAuthError: XXX |
| 120 | Authorization Failure | An accountant user was deleted from a QuickBooks Online company while the company was connected to your app. AuthorizationFailure: XXX |
| 130 | Accessing Wrong Cluster | Review the XXX error code and follow its recommendation. WrongClusterError: XXX, statusCode: XXX |
| 140 | Company Locked Out | Review the XXX error code and follow its recommendation. CompanyLockedOut: XXX, statusCode: XXX |
| 150 | Company Under Maintenance | Review the XXX error code and follow its recommendation. CompanyUnderMaintenance: XXX, statusCode: XXX |

### General Error Codes

| Error code | Message | Detail |
|---|---|---|
| 500 | Unsupported operation | The operation you're trying to perform for a given entity isn't supported. Operation XXX is not supported. |
| 600 | Duplicate Request ID | The request ID you're trying to use already exists in the QuickBooks Online Accounting API or QuickBooks Payments API. Our system detected the duplicate. |
| 610 | Object Not Found | The object ID you're requesting doesn't exist. |
| 620 | Txn ID Cannot Be Linked | The Txn ID (identified by XXX) you're trying to use can't be linked. |
| 630 | Duplicate Object | There's already an object with some properties (for example, duplicate reference IDs). This violates the unique constraint of the object. |
| 700 | Parent Reference Invalid | The parent isn't a valid reference. |
| 800 | Cannot delete object | The object can't be deleted due to existing dependencies. Delete those dependencies first. |
| 900 | Parent cannot be child | The same object can't be the parent or child of itself. |
| 1000 | Operation failed | The operation you're trying to perform failed for the entity (identified by XXX). Operation XXX failed. |
| 1010 | Create Failed | The object you tried to create wasn't created. Object creation failed, XXX. |
| 1020 | Update Failed | The object you tried to update wasn't updated. Object update failed, XXX. |
| 1040 | Batch size exceeds allowed limit | The batch size for the batch request or BatchItemRequest is too big. The batch limit size is stated by the XXX value. |
| 1050 | Invalid Content Type | This error occurs when messages are sent through the Java Message Service queue. We only support XML or JSON formats. |

### Validation Fault Error Codes

| Error code | Message | Detail |
|---|---|---|
| 2000 | Invalid or unsupported object name | The specified object name is unsupported or invalid. Object Name: XXX. |
| 2010 | Request has invalid or unsupported property | The specified property name (identified by XXX) is unsupported or invalid. Property Name: XXX. |
| 2020 | Required param missing | The required parameter is missing in the request. Required parameter XXX is missing. |
| 2030 | Invalid ID | IDs need to be a valid number. This error can be for any entity or operation. Supplied value: XXX. |
| 2040 | Invalid String | An element contains invalid characters. Review the string. It may contain unsupported characters. Element contains invalid characters: XXX. |
| 2050 | Invalid String Length | A specified string length is either too long or too short. Minimum and maximum length are defined in the error response. Min: XXX, Max: XXX, Supplied length: XXX. |
| 2060 | Invalid Date Format | The date format isn't valid. |
| 2070 | Invalid Date | The date value must be a valid value. Supplied value: XXX. |
| 2080 | Invalid Number Format | The number format supplied isn't valid. |
| 2090 | Invalid Number | The number supplied is the correct format, but the value isn't supported by the API. |
| 2100 | Invalid Decimal Format | The decimal format isn't valid. |
| 2110 | Invalid Decimal | The decimal supplied may be the correct format, but the value isn't supported by the API. |
| 2120 | Invalid Type | The specified type isn't compatible or supported by the API. This error usually applies to entities that have categories or types. For example, [items](https://developer.intuit.com/app/developer/qbo/docs/api/accounting/all-entities/item) can have a "product" or "service" type. |
| 2130 | Invalid Request Id Format | The format of the request ID is invalid and not supported by the API. |
| 2140 | Invalid Amount | The amount value isn't supported by the API. |
| 2150 | Invalid Percent | The percent value is invalid and not supported by the API. |
| 2160 | Invalid Quantity | The quantity value is invalid and not supported by the API. |
| 2170 | Invalid Enumeration | The enumeration value is invalid and not supported by the API. |
| 2180 | Invalid String Range | The string range isn't valid. |
| 2190 | Invalid Date Range | The date range isn't valid. |
| 2200 | Malformed Website Address Format | The URL isn't in the correct format. Supplied value: XXX. |
| 2210 | Invalid Email Address Format | The email address doesn't conform to RFC 822 syntax rules. Supplied value: XXX. |
| 2220 | Invalid Currency Type | The specified currency type isn't valid. |
| 2230 | Invalid Boolean | The specified boolean value isn't valid. |
| 2240 | Invalid Number Range | The specified number isn't within the valid range. Supplied value: XXX. |
| 2250 | Missing lines | You must enter at least one split line. |
| 2260 | MissingPostingType | A posting type is required. Valid values are **Credit** or **Debit**. |
| 2270 | MissingTaxApplicableOn | The TaxApplicableOn entity is required. Set it in the request payload. Valid values are **Credit** or **Debit**. |
| 2280 | MissingTaxAmount | A tax amount is required with TaxCodeRef applied. |
| 2290 | NegativeAmount | Negative amounts aren't supported on transaction lines. |
| 2300 | Amount on debits not equal to credits | A journal entry you created or updated unbalanced an account. Balance the debits and credits for the account. |
| 2500 | Invalid Reference Id | The reference ID is invalid: XXX. |
| 2600 | Invalid Request Conflict Element in Request | Conflicting elements found in the request. |
| 3000 | Invalid Custom Field | The custom field for the specified object isn't valid. |
| 3200 | ApplicationAuthenticationFailed | The signature in the OAuth request has been changed and is now invalid. SignatureBaseString: XXX. |
| 3202 | EmptyField | Review the error message for the cause of the issue. Related to oauth_token. |
| 4000 | The query cannot be parsed | Review the error message for the cause of the issue. Query Parser Error: XXX. |
| 4001 | The query is invalid | Review the error message for the cause of the issue. QueryValidationError: XXX. |
| 4002 | The query could not be processed | Review the error message for the cause of the issue. QueryProcessingError: XXX. |

### Specific Error Codes

| Error code | Message | Explanation |
|---|---|---|
| 5000 | Deprecated field | The request has a deprecated field. |
| 5010 | Stale Object | The requested update is for a stale object. Client needs to refresh the object. |
| 5020 | Permission Denied | You don't have permissions to access this feature. |
| 5030 | Feature Not Supported | The requested feature isn't supported. |
| 6000 | Business Validation Error | Review the error message for the cause of the issue. BusinessValidationError: XXX. |
| 6010 | Invalid Account.AccountType && Account.SubType | You must specify either an **Account.AccountType** or an **Account.AccountSubType**. |
| 6020 | Content length missing in request | Can't find the content-length header. |
| 6030 | Upload request size exceeds allowed limit | The file size exceeds the size limit. Upload requests shouldn't exceed (XXX). |
| 6040 | The file metadata must be of Attachable type | The submitted file metadata must be attachable. Supplied Type: XXX. Request parameter: XXX. |
| 6041 | Invalid Uploaded File | The uploaded file isn't valid. Learn more about [accepted attach images and note types](https://developer.intuit.com/app/developer/qbo/docs/workflows/attach-images-and-notes). |
| 6050 | The entity reference type is unsupported for Attachable | The entity reference type (specified by XXX) isn't supported for attachable. |
| 6060 | Account.OpenBalanceDate must be specified with Account.OpenBalance | A value for the `Account.OpenBalanceDate` field is required when specifying an `Account.OpenBalance`. |
| 6070 | Amount calculation incorrect in the request | The amount isn't equal to the Qty * UnitPrice. Supplied value: XXX. |
| 6080 | No Name Provided | Fill out values for at least one of the following: **Title**, **GivenName**, **MiddleName**, **FamilyName**, **DisplayName**, **Suffix**. |
| 6090 | Discount Line missing required info | A discount amount or percent is required. |
| 6100 | Invalid Line TaxCode in the request | The `TaxCode` value isn't valid. For US, specify **TAX** or **NON**. Supplied value: XXX. |
| 6110 | MAS Transaction cannot be updated | The requested transaction can't be updated since it's a MAS transaction. |
| 6120 | MAS Transaction cannot be deleted | The requested transaction can't be deleted since it's a MAS transaction. |
| 6130 | Cannot change who customer bills with | You can't change who this customer bills with because there are invoices connected to a parent account. Delete all payments or invoices for the parent account that are linked to this customer. |
| 6140 | Duplicate Doc Num | The specified number is already in use. Specify a different number. |
| 6150 | ParentRef Required Validation Error | A `ParentRef` is required depending if (XXX) is **true** or **false**. |
| 6160 | Invalid TaxCodeRef Error | The specified `CustomSalesTax` ID can't be used to create a transaction. |
| 6170 | Invalid TxnTaxDetail Error | The `TxnTaxDetail` attribute isn't allowed when `GlobalTaxCalculation` is set to **NotApplicable**. |
| 6180 | Sub Level Limit | Review the number of nested accounts and customers. You can nest up to 5 levels. |
| 6190 | Invalid Company Status Error | The subscription period for the QuickBooks Online company has ended, has a billing problem, or was canceled. Learn more [about QuickBooks Online subscription states](https://developer.intuit.com/app/developer/qbo/docs/develop/troubleshooting/subscription-states). |
| 6200 | Account Period Closed Error | The user has "closed the books" so the account period is closed for edits. |
| 6210 | Account Period Closed Required Password Error | The user has "closed the books" so the account period is closed for edits. A password is required to make changes and updates. |
| 6220 | Delete Entity Has Balance Error | The object or entity you're trying to delete has an open balance. |
| 6230 | Invalid TaxCodeRef Error | You can't update transactions to use the `CustomSalesTax` ID. |
| 6240 | Duplicate Name Exists Error | The specified name already exists. The customer ID is: XXX. |
| 6250 | InvalidSalesCustomer | The specified customer (identified by XXX) doesn't exist, or is inactive. |
| 6270 | InventoryTxnDatedBeforeItemStartDate | Transactions with inventory items that have a quantity on hand can't be dated earlier than the specified `InventoryStartDate`. |
| 6280 | Delete Customer Has Unbilled Expenses Error | This customer can't be deleted because they have unbilled charges. |
| 6290 | InvalidEmployeeOrVendor | You must specify either an "employee" or "vendor." You can't enter both. Their ID is: XXX. |
| 6300 | InvalidEmployeeOrVendorName | The employee or vendor's name is either missing or invalid. It should be **Employee** or **Vendor**. |
| 6310 | MissingCustomer | A customer is required if the transaction is billable. |
| 6320 | InvalidBillable | Invalid billable status: XXX. |
| 6330 | MissingBillingRate | Invalid hourly rate: XXX. |
| 6340 | InvalidTime | You must specify a time (hours:minutes): XXX. |
| 6350 | MissingTime | Time (hours:minutes) is missing: XXX. |
| 6360 | InvalidStartOrEndTime | Enter a valid start and end time. Make sure break times aren't longer than the elapsed time. |
| 6370 | StartTimeAfterEndTime | The start time is after end time: XXX. |
| 6380 | InvalidDuration | The time duration isn't within the valid range: XXX. |
| 6390 | InvalidMultipleDurations | Invalid multiple time durations. Specify a total elapsed time (hours:minutes) or the start and end time. |
| 6400 | Deleting linkedPurchase results in invalid Invoice | The purchase can't be deleted since it would invalidate a linked invoice. |
| 6420 | Email address is required | An email address is required for this customer since the delivery type is "email." |
| 6430 | Invalid account type used | Invalid account type: XXX |
| 6440 | Transaction Detail Information Required | You must select a product or service or an account for each split line that has an amount or a billable customer. |
| 6450 | Attempt to Void Transaction failed | Since the transaction is already settled, you can't void it. |
| 6460 | Invoice MAS Transaction Deposit Amount mismatch | The invoice deposit amount (XXX) and credit card amount (YYY) don't match. |
| 6470 | TaxLiabilityAccount or TaxRateError | You either selected a tax on a transaction that's not allowed, or haven't specified a tax rate. |
| 6480 | Matched Transaction Delete Error | You can't delete this transaction. It's matched to another transaction that was categorized into a different financial account. |
| 6490 | MultipleBudgetSecondaryListTypes | Review the details in the error response. |
| 6500 | MissingBudgetDetails | Review the details in the error response. |
| 6510 | BudgetAccountMismatch | Review the details in the error response. |
| 6520 | Missing Tracked Inventory Item Quantity | The tracked inventory item must have a quantity. |
| 6530 | Found raw Credit Card Numbering request | You can't enter raw credit card numbers. Manual credit card data isn't supported. Use a [token for the credit card number](https://developer.intuit.com/app/developer/qbpayments/docs/api/resources/all-entities/tokens). |
| 6540 | Deposited Transaction cannot be changed | This transaction is part of a deposit. If you want to change or delete it, you must edit and remove it from the deposit it appears on. |

### Important, Non-Recoverable Error Codes

| Error code | Message | Detail |
|---|---|---|
| 10000 | An application error has occurred while processing your request | Review the System Failure Error code (XXX) to identify the issue. |
| 10100 | Result Set Big | The result for the query is too large. |
| 10200 | Company Reset | Company has been reset. Request for CUD is stale. |
