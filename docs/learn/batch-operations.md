# Batch Operation

Instead of sending requests individually, you can eliminate overhead and send "batches" of requests to the QuickBooks Online Accounting API.

## How the batch operation works

Use the batch operation to send up to 10 payloads with a single batch operation. You can send up to 40 requests per minute to an end-user's QuickBooks Online company (specified by the `realmId`).

Batching lets you streamline your app's code, and reduces network latency caused by individual, synchronous API calls.

## Use the batch operation

- Learn more about the [batch operation](https://developer.intuit.com/app/developer/qbo/docs/api/accounting/all-entities/batch)
- Learn more about batch operations for supported SDKs: [Java](https://developer.intuit.com/app/developer/qbo/docs/develop/sdks-and-samples-collections/java/asynchronous-calls#batch-process), [.NET](https://developer.intuit.com/app/developer/qbo/docs/develop/sdks-and-samples-collections/net/asynchronous-calls#batch-process), [PHP](https://developer.intuit.com/app/developer/qbo/docs/develop/sdks-and-samples-collections/php/synchronous-calls#batch-process)
