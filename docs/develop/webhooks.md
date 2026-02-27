# Webhooks

Use webhooks to receive event triggered callbacks for entities that your app needs to stay on top of. Webhooks automatically notify you whenever data changes in your end-user's QuickBooks Online company files.

Webhooks apply to all QuickBooks Online companies connected to your app.

You'll need to configure an endpoint our servers can call whenever user data changes trigger notifications. Once webhooks are active, we'll send the requested event data, changes, and notifications.

> **Note**: Information on this page related to the webhooks payload has been updated. Please read [this article](https://blogs.intuit.com/2025/11/12/upcoming-change-to-webhooks-payload-structure/) for more details.

## Step 1: Set up OAuth 2.0

If you haven't already, [set up OAuth 2.0](https://developer.intuit.com/app/developer/qbo/docs/develop/authentication-and-authorization/oauth-2.0) for your app.

Even if webhooks are active, you'll only receive change notifications for QuickBooks Online companies that are connected and authorized via OAuth 2.0.

**Tip**: If you plan to test webhooks for [sandbox environments](https://developer.intuit.com/app/developer/qbo/docs/develop/sandboxes), the QuickBooks Online companies you test with need to [complete the authentication flow via OAuth 2.0](https://developer.intuit.com/app/developer/qbo/docs/develop/authentication-and-authorization/oauth-2.0).

## Step 2: See the example implementation

Here is an example webhook implementation for [Java](https://github.com/IntuitDeveloper/SampleApp-Webhooks-Java-Cloudevents).

## Step 3: Review supported API entities

See which entities and operations support webhooks.

## Step 4: Configure webhook endpoints for an app

1. [Sign in](https://developer.intuit.com/dashboard) to your developer account.
2. Select **My Hub > App dashboard** from the upper-right corner of the toolbar.
3. Select and open the app you want to subscribe.

There are two sets of webhooks: one for live, in-production apps and a separate set for sandbox and testing environments.

You need to set up webhooks for production apps and sandbox environments separately.

### For live, in-production apps

1. From the left navigation pane, select **Webhooks**.
2. Select **Production**.
3. Enter the **Endpoint URL**. This is where the server will send notifications.
4. Select the **Show webhooks** dropdown.
5. Review the events and operations.
6. Select the notifications and actions you want to enable.
7. Select **Save**.

### For sandbox and developer environments

1. From the left navigation pane, select **Webhooks**.
2. Select **Development**.
3. Enter the **Endpoint URL**. This is where the server will send notifications.
4. Select the **Show webhooks** dropdown.
5. Review the events and operations.
6. Select the notifications and actions you want to enable.
7. Select **Save**.

**Tip**: It may take up to five minutes to get your first webhook notification.

## Step 5: Validate webhook notifications

After you configure webhook endpoints, we provide an app-specific verifier token. Use verifier tokens to validate the webhook notifications from the callback are from Intuit.

To see verifier tokens:

1. [Sign in](https://developer.intuit.com/dashboard) to your developer account.
2. Select **My Hub > App dashboard** from the upper-right corner of the toolbar.
3. Select and open the app you want to verify.
4. From the left navigation pane, select **Webhooks**.
5. Select **Development** or **Production**.
6. Review the **Verifier Token** field.

To use verifier tokens:

1. Hash the notification payload with **HMAC_SHA256_ALGORITHM** using the **\<verifier token\>** value as the key.
2. Compare the payload hash value with the **intuit-signature** header from the notification. The values should be identical.

Here's a sample header:

```
content-length:262
intuit-created-time:2016-02-02T16:25:00-0800
intuit-t-id:9cf50b60-8b0e-4fea-8327-e6a66099fe6f
proxy-connection:keep-alive
host:sample-endpoint.ilb.idg-notify-ppd.a.intuit.com:8443
intuit-notification-schema-version:0.1
content-type:application/json; charset=utf-8
intuit-signature:6kQBQtjwjupelRMwkyJsnpq80uhz2o+Rn92+m03GhKE=
accept:application/json
user-agent:intuit_notification_server/0.1
```

## Step 6: Review webhook notifications

Webhook notifications are POSTs with a JSON body. The notification payload contains the following fields:

| Field | Description |
|-------|-------------|
| `specversion` | The version of the CloudEvents specification. Value is always 1. Can be ignored. |
| `id` | Unique ID of the event. May be useful in troubleshooting or reporting any errors. |
| `source` | GUID value. Can be ignored. |
| `type` | Represents the entity and event for the notification, in the format: `namespace.entitytype.eventname.version` |
| `datacontenttype` | The content type of the data value: `application/json` |
| `time` | The timestamp of when the occurrence happened, provided in RFC 3339 (ISO 8601) format at UTC+0 (e.g., "2018-04-05T03:56:24Z"). |
| `intuitentityid` | The ID of the entity that was changed. |
| `intuitaccountid` | The QuickBooks Online company ID or realm ID. |
| `data` | Name-value pairs containing additional information about the occurrence. Can be ignored if empty. |

Here's an example notification payload:

```json
[
  {
    "specversion": "1.0",
    "id": "88cd52aa-33b6-4351-9aa4-47572edbd068",
    "source": "intuit.dsnBgbseACLLRZNxo2dfc4evmEJdxde58xeeYcZliOU=",
    "type": "qbo.account.created.v1",
    "datacontenttype": "application/json",
    "time": "2025-09-10T21:31:25.179851517Z",
    "intuitentityid": "1234",
    "intuitaccountid": "310687",
    "data": {}
  }
]
```

Webhook events are composed of an array of individual event notifications. Each of these event notifications correspond to unique information associated with a specific realm ID. A single notification can contain a list of events, each for a different QuickBooks Online company/realm ID. Your app must be updated to handle events for multiple companies per notification.

Using the above example, if your app is connected to multiple companies, you will receive an array of event notifications. Each element in this array represents a unique update for a specific company.

Each event notification can include multiple entities such as `Customer` or `Vendor`. These entities within the same realm ID denote changes or operations, such as `Create`, that happened at the given `lastUpdated` time.

**Tip**: There are no Intuit-imposed limits to payload size or number of events. Individual server architectures may impose their own limits (2MB is a common default size limit). Assume this limit is imposed by your server unless you know otherwise.

## Webhook best practices

**Reliability**: To compensate for the possibility of missed events, make a ChangeDataCapture (CDC) call for all required entities, dating back to the last known successfully processed webhook event for each entity.

Additionally, you can make a daily CDC call for all required entities to ensure your app has always processed the most up-to-date data.

**Respond promptly**: If your endpoint doesn't respond within three seconds, the transaction will time out and retry.

Make sure your app can always respond quickly. Don't process the notification payload or perform complex operations within the webhooks endpoint implementation. Do the processing on a separate thread asynchronously using a queue.

**Manage concurrency**: Event notifications are sent one realm ID at a time. When there are multiple rapid changes, your app may get frequent notifications. Process the queue linearly to avoid processing the same changes more than once.

**Notification ordering**: It's possible to receive events out of sequence. The `timestamp` field in the notification payload is always the source of truth for when events occur.

**Retry policy**: Your endpoint must respond to notifications with an HTTP 200 status within 3 seconds. When our webhooks delivery system does not receive a successful acknowledgement from your listener, we retry the failed event. The retry will happen at the following intervals: 10s, 20s, 30s, 5m, 20m, 2h, 4h, 6h and then stays at 6-hour intervals. You may not receive subsequent events from our system until you correctly acknowledge the first event.
