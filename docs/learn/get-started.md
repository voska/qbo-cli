# Create and Start Developing Your App

Let's get started. The QuickBooks Online Accounting API gives you incredible flexibility to build creative solutions for key business and accounting workflows.

- You can [browse the API Explorer](https://developer.intuit.com/app/developer/qbo/docs/api/accounting/most-commonly-used/account) to see our library of API entities and develop your app from scratch
- You can [start building your app around common business use cases](https://developer.intuit.com/app/developer/qbo/docs/workflows)
- If you want to integrate with QuickBooks Desktop, [follow these steps](https://developer.intuit.com/app/developer/qbdesktop/docs/get-started)
- If you want your app to also handle payments processing, here's how to [integrate the QuickBooks Payments API](https://developer.intuit.com/app/developer/qbpayments/docs/workflows/process-a-payment)

Here's what you need to set up to jump-start development.

## Step 1: Get to know the QuickBooks Online Accounting API platform

Already familiar with our APIs? Feel free to skip ahead. If you're new, review our platform's capabilities so you can develop with confidence:

- [Learn about our API platform](https://developer.intuit.com/app/developer/qbo/docs/learn/explore-the-quickbooks-online-api)
- [Learn how REST API works](https://developer.intuit.com/app/developer/qbo/docs/learn/rest-api-features)
- [Learn basic bookkeeping in QuickBooks](https://developer.intuit.com/app/developer/qbo/docs/learn/learn-basic-bookkeeping)
- [Learn about QuickBooks Online features](https://developer.intuit.com/app/developer/qbo/docs/learn/learn-quickbooks-online-basics)

## Step 2: Set up your developer account

Go to the Intuit Developer Portal and [create a developer account](https://developer.intuit.com/app/developer/myapps).

This profile is where you'll create apps, get app credentials, and set up sandbox companies for testing.

## Step 3: Create an app on the Intuit Developer Portal

The apps you create on our developer portal generate unique credentials and info you'll need during development.

1. [Sign in](https://developer.intuit.com/dashboard) to your developer account.
2. Select **My Hub > App dashboard** from the upper-right corner of the toolbar.
3. Select the app card with a **+** to create a new app.
4. Follow the on-screen steps.

When you create apps, you'll [pick your app's scopes](https://developer.intuit.com/app/developer/qbo/docs/learn/scopes). Scopes limit the type of data your app can access. If your app only needs access to accounting data, select the **accounting** scope. If it also needs to access payment data, also select the **payments** scope.

## Step 4: Create a sandbox QuickBooks Online company

When you create an Intuit Developer account, we give you a test company with sample data for you. We call this a sandbox company.

You can use this sandbox company to test your code. Here's how to [create additional sandbox companies](https://developer.intuit.com/app/developer/qbo/docs/develop/sandboxes/manage-your-sandboxes).

## Step 5: Learn how to get your app's credentials

During development, you'll need to [use your app's credentials](https://developer.intuit.com/app/developer/qbo/docs/get-started/get-client-id-and-client-secret) for various tasks such as authorization.

## Step 6: Authorize your app

Use your credentials (i.e., **Client ID** and **Client Secret**) to [connect your app to OAuth 2.0](https://developer.intuit.com/app/developer/qbo/docs/develop/authentication-and-authorization). This grants access tokens your app can use to send requests to our APIs.

### Want to make a quick sample API call?

You don't need to fully authenticate your app to test requests. Here's how to quickly call the `CompanyInfo` entity:

1. Visit the [OAuth 2.0 Playground](https://developer.intuit.com/app/developer/playground).
2. Select a sandbox company from the **Select app** dropdown.
3. Follow the on-screen steps to get an **access token**. This is a temporary token for testing only.
4. Copy the **Access token** and **Realm ID** (also known as the Company ID).

Next, create a sample request:

1. Copy and use the sample cURL **GET** request below.
2. Replace the `REPLACE_WITH` placeholders with the copied **Access token** and **Realm ID** values.
3. Send the request.

```bash
curl -X GET 'https://sandbox-quickbooks.api.intuit.com/v3/company/REPLACE_WITH_SANDBOX_COMPANY_ID/companyinfo/REPLACE_WITH_SANDBOX_COMPANY_ID?minorversion=12' \
  -H 'accept: application/json' \
  -H 'authorization:Bearer REPLACE_WITH_ACCESS_TOKEN' \
  -H 'content-type: application/json'
```

The response for your request should match the sample server response.

## Step 7: Test your code

Check out [QuickBooks and third-party testing tools](https://developer.intuit.com/app/developer/qbo/docs/develop/sandboxes) you can use during development. We've configured tools, including Postman, to make testing with your sandbox company a little easier.
