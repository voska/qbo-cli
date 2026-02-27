# Fix Common QuickBooks Online Accounting API Errors

If you or your app users ever see an error, [look up the error code](https://developer.intuit.com/app/developer/qbo/docs/develop/troubleshooting/error-codes) to identify the issue.

Error responses follow this general format:

```json
{
  "Fault": {
    "Error": [
      {
        "Message": "<message>",
        "Detail": "<detail>",
        "code": "<code>",
        "element": ""
      }
    ],
    "type": "ValidationFault"
  },
  "time": "2017-09-29T09:17:31.892-07:00"
}
```

---

## 120: Authorization Failure

| Error code | Message | Detail |
|---|---|---|
| 120 | Authorization Failure | : -11014-You do not have access to use QuickBooks Online Plus. |

### Common Cause

One of the following conditions applies to an individual who connected their QuickBooks Online company to your app:

- They aren't an admin user
- Their admin status changed
- Their user profile was completely removed from the QuickBooks Online company

### Solution

Fix or disconnect the connection.

Reach out to a current admin user for the QuickBooks Online company. Ask if they still need to connect to your app.

If they don't, you can [revoke their access](https://developer.intuit.com/app/developer/qbo/docs/develop/authentication-and-authorization/oauth-2.0#revoking-access).

---

## 500: Unsupported Operation

| Error code | Message | Detail |
|---|---|---|
| 500 | Unsupported Operation | Operation com.sun.org.apache.xerces.internal.impl.io.MalformedByteSequenceException: Invalid byte 1 of 1-byte UTF-8 sequence is not supported. |

### Common Cause

There are non-UTF-8 characters in your request payload.

### Solution

Use any standard library or function to set UTF-8 encoding for your request payload.

---

## 610: Object Not Found

There are two possible reasons for error 610. Check the `detail` field in the response payload to see which one applies:

### Variant 1: Inactive Object

| Error code | Message | Detail |
|---|---|---|
| 610 | Object Not Found | Object Not Found: Something you're trying to use has been made inactive. Check the fields with accounts. |

#### Common Cause

This is usually caused by a reference to an inactive name list object (for example, a customer, vendor, or account) in a transaction request.

For instance, the customer object referenced in an invoice should have an `Invoice.CustomerRef` element. The likely cause is that the `Customer.Active` attribute is set to **false**.

#### Solution

To prevent this error:

- When your app first connects to a user's QuickBooks Online company, cache the `Id`, `SyncToken`, `DisplayName`, and `Active` states for all name list objects in your app's database.
- [Use webhooks](https://developer.intuit.com/app/developer/qbo/docs/develop/webhooks) or the [change data capture operation](https://developer.intuit.com/app/developer/qbo/docs/learn/explore-the-quickbooks-online-api/change-data-capture) to track changes to the name list objects' `Active` states.
- Always keep your app's internal database up to date.
- Validate the referenced object's `Active` state in transaction requests before sending them. Implement the appropriate logic in your app for `active` = **false**.

### Variant 2: Deleted Transaction

| Error code | Message | Detail |
|---|---|---|
| 610 | Object Not Found | Object Not Found : Another user has deleted this transaction. |

#### Common Cause

This is usually caused by a reference to a deleted transaction object in the request.

#### Solution

When users delete transactions in QuickBooks, they're permanently removed.

However, name list objects, such as customers or vendors, are not deleted. They're simply marked "inactive."

[Use webhooks](https://developer.intuit.com/app/developer/qbo/docs/develop/webhooks) or the [change data capture operation](https://developer.intuit.com/app/developer/qbo/docs/learn/explore-the-quickbooks-online-api/change-data-capture) to keep track of deleted and added transactions.

---

## 2020: Required Param Missing

There are two possible reasons for error 2020. Check the `detail` field in the response payload to see which one applies:

### Variant 1: Missing Asset Account

| Error code | Message | Detail |
|---|---|---|
| 2020 | Required Param Missing | Required parameter is missing in the request. An inventory assetaccount is required if you are tracking inventory quantities for this product. |

#### Common Cause

The `Item.AssetAccountRef` isn't set for an inventory object in your request.

#### Solution

You must include the `Item.AssetAccountRef` when creating an inventory item. Learn more about the [Item](https://developer.intuit.com/app/developer/qbo/docs/api/accounting/all-entities/item) entity.

### Variant 2: Missing Line Element

| Error code | Message | Detail |
|---|---|---|
| 2020 | Required Param Missing | Required parameter Line is missing in the request. |

#### Common Cause

A `Line` element is missing from the transaction object's request payload.

#### Solution

Learn more about required parameters for [transaction](https://developer.intuit.com/app/developer/qbo/docs/api/accounting/most-commonly-used/account) entities, such as invoice and sales receipt.

---

## 2170: Invalid Enumeration

| Error code | Message | Detail |
|---|---|---|
| 2170 | Invalid Enumeration | Invalid Enumeration : HAS_BEEN_BILLED |

### Common Cause

Apps can't directly set the `TimeActivity.BillableStatus` attribute to **HasBeenBilled**.

QuickBooks automatically sets the [TimeActivity](https://developer.intuit.com/app/developer/qbo/docs/api/accounting/all-entities/timeactivity) entity's billable status to **HasBeenBilled** (not by a user's action) when your app links it to an invoice object via `Invoice.LinkedTxn`.

### Solution

Don't manually set `TimeActivity.BillableStatus` to **HasBeenBilled**. Let the system handle it.

---

## 2500: Invalid Reference Id

| Error code | Message | Detail |
|---|---|---|
| 2500 | Invalid Reference Id | Invalid Reference Id : Something you're trying to use has been made inactive. Check the fields with accounts. |

### Common Cause

This is usually caused by references to inactive name list objects (for example, a customer, vendor, or account) in requests.

The `Active` attribute is likely set to **false** for the target name list object.

### Solution

To prevent this error:

- When your app first connects to a user's QuickBooks Online company, cache the `Id`, `SyncToken`, `DisplayName`, and `Active` states for all name list objects in your app's database.
- [Use webhooks](https://developer.intuit.com/app/developer/qbo/docs/develop/webhooks) or the [change data capture operation](https://developer.intuit.com/app/developer/qbo/docs/learn/explore-the-quickbooks-online-api/change-data-capture) to track changes to the name list objects' `Active` states.
- Always keep your app's internal database up to date.
- Validate the referenced object's `Active` state in transaction requests before sending them. Implement the appropriate logic in your app for `active` = **false**.

---

## 5010: Stale Object Error

| Error code | Message | Detail |
|---|---|---|
| 5010 | Stale Object Error | Stale Object Error : You and \<other user name\> were working on this at the same time. \<other user name\> finished before you did, so your work was not saved. Entity=\<entity name\> |

### Common Cause

This usually means you're working with a stale object.

### Solution

Always work with the latest version of objects. An object's version is stored in its `syncToken` attribute.

- When your app first connects to a user's QuickBooks Online company, cache the `Id`, `type`, and `syncToken` for objects in your app's database.
- [Use webhooks](https://developer.intuit.com/app/developer/qbo/docs/develop/webhooks) or the [change data capture operation](https://developer.intuit.com/app/developer/qbo/docs/learn/explore-the-quickbooks-online-api/change-data-capture) to track changes to objects.
- Before performing any update or delete operation, retrieve the latest version of the object.
- Use the `syncToken` in the response payload to get the latest version of an object. Use the returned `syncToken` for subsequent requests.
- Always keep your app's internal database up to date.

Performing read operations before updates makes sure you're always getting the latest data if you aren't using [webhooks](https://developer.intuit.com/app/developer/qbo/docs/develop/webhooks) or the [change data capture operation](https://developer.intuit.com/app/developer/qbo/docs/learn/explore-the-quickbooks-online-api/change-data-capture).

---

## 6000: A Business Validation Error Has Occurred

| Error code | Message | Detail |
|---|---|---|
| 6000 | A business validation error has occurred while processing your request | Business Validation Error: You must set a transaction amount. |

### Common Cause

The `TotalAmt` attribute isn't set for the [BillPayment](https://developer.intuit.com/app/developer/qbo/docs/api/accounting/all-entities/customerbalancedetail) or [Payment](https://developer.intuit.com/app/developer/qbo/docs/api/accounting/all-entities/payment) entities in your request. The `TotalAmt` attribute is required.

### Solution

Make sure the `TotalAmt` attribute has a valid transaction amount in requests.

---

## 6140: Duplicate Doc Num Error

| Error code | Message | Detail |
|---|---|---|
| 6140 | Duplicate DocNum Error | Duplicate Document Number Error : You must specify a different number. This number has already been used. |

### Common Cause

When you create a transaction, your app sends a string in `DocNumber` that happens to duplicate the number in an existing transaction object, and these conditions are met:

- The `Preferences.SalesFormsPrefs.CustomTxnNumbers` attribute is set to **true**.
- QuickBooks users enable the "**Warn if duplicate check/bill number is used**" message via the QuickBooks Online UI. Users can accept the warning and allow duplicate doc numbers. It's not possible to test if this setting is enabled, or accept warnings programmatically, via the QuickBooks Online Accounting API.

### Solution

Fetch and check the `Preferences.SalesFormsPrefs.CustomTxnNumbers` attribute before creating any transactions.

- If it's **true**, create the `DocNumber` with a custom pattern such as **\<your app name\>_docnumber**.
- If it's **false**, QuickBooks will automatically assign the `DocNumber`.

This prevents `DocNumber` duplications when users or other apps manually enter transactions that utilize QuickBooks data.

---

## 6190: Invalid Company Status

| Error code | Message | Detail |
|---|---|---|
| 6190 | Invalid Company Status | Subscription period has ended or canceled or there was a billing problem : You can't add data to QuickBooks Online Plus because your trial or subscription period ended. |

### Common Cause

The target user's [QuickBooks Online subscription status](https://developer.intuit.com/app/developer/qbo/docs/develop/troubleshooting/subscription-states) is preventing the operation.

For example, your app may be trying to use a write operation, but based on the user's subscription status, it can only do read-only operations.

### Solution

Use the `SubscriptionStatus` attribute of the [CompanyInfo](https://developer.intuit.com/app/developer/qbo/docs/api/accounting/all-entities/companyinfo) entity to check the company's subscription status.

Design your app so it can handle [changes to subscription status](https://developer.intuit.com/app/developer/qbo/docs/develop/troubleshooting/subscription-states). It's especially important to anticipate billing-driven issues. Consider whether your app can wait for subscriptions to become valid again, or letting users disconnect from your app if they end their QuickBooks subscription.

---

## 6210: Account Period Closed

| Error code | Message | Detail |
|---|---|---|
| 6210 | Account Period Closed | The account period has closed and the account books cannot be updated through the QBO Services API. Please use the QBO website to make these changes. |

### Common Cause

In QuickBooks Online, many users automatically [close the books](https://quickbooks.intuit.com/learn-support/en-us/customer-company-settings/close-your-books-to-lock-past-transactions/00/186384) at the end of their fiscal year.

This error usually occurs when you try to modify a transaction or an item within this closed accounting period.

### Solution

If a user automatically closes their books in QuickBooks Online, transactions and items within the closed year time period can't be modified.

Currently, it's not possible to check the end of the fiscal year, or if a user's books are closed, via the QuickBooks Online Accounting API.

It's also not possible to change the "close the books" settings via the QuickBooks Online Accounting API.

To anticipate this:

- Add logic to your app to query the QuickBooks Online company's admin user for the account period close date when they first connect to your app.
- Use the account period close date to determine if a target transaction is within the closed accounting period before sending your request.

---

## 6240: Duplicate Name Exists Error

| Error code | Message | Detail |
|---|---|---|
| 6240 | Duplicate Name Exists Error | The name supplied already exists. Another customer, vendor or employee is already using this name. Please use a different name. |

### Common Cause

The `DisplayName` attribute must be unique across all customer, vendor, and employee objects.

### Solution

- Fetch the `Id` and `DisplayName` attributes for all customer, vendor, and employee objects for the target QuickBooks Online company.
- Persist the `Id` and `DisplayName` attributes for name list entities (for example, customers, vendors, employees, class, or sales tax) in your app's database.
- Before creating new customer, vendor, or employee objects, check the database to determine if the `DisplayName` already exists.
- [Use webhooks](https://developer.intuit.com/app/developer/qbo/docs/develop/webhooks) or the [change data capture operation](https://developer.intuit.com/app/developer/qbo/docs/learn/explore-the-quickbooks-online-api/change-data-capture) to track changes to objects.
- Always keep your app's internal database up to date.
- For cases where multiple objects use the same name, create a naming scheme to distinguish them. For example, "John Smith" may be a customer and vendor. Store his name in the customer object as "John Smith (customer)" and in the vendor object as "John Smith (vendor)".

---

## 6540: Deposited Transaction Cannot Be Changed

| Error code | Message | Detail |
|---|---|---|
| 6540 | Deposited Transaction cannot be changed | This transaction has been deposited. If you want to change or delete it, you must edit the deposit it appears on and remove it first. |

### Common Cause

The target transaction is linked to a deposit object via the deposit's `Deposit.Line.LinkedTxn` attribute.

QuickBooks users record customer payments using payments (against an invoice) or sales receipts (if payment is received at the time of purchase). Learn more about [sales receipts and invoices](https://developer.intuit.com/app/developer/qbo/docs/learn/learn-basic-bookkeeping/invoicing).

Payments may go to what's known as an Undeposited Funds account. Users use a deposit to move money from the [Undeposited Funds Account](https://quickbooks.intuit.com/learn-support/en-us/chart-of-accounts/deposit-payments-into-the-undeposited-funds-account-in/00/185574) to another bank account. This subsequent deposit object links back to the original payment or salesReceipt object.

### Solution

You can't modify transactions (payments or sales receipts) that are already deposited and linked to a deposit.

Learn more about [bank deposits in QuickBooks Online](https://quickbooks.intuit.com/learn-support/en-us/bank-deposits/record-and-make-bank-deposits-in-quickbooks-online/00/185563).
