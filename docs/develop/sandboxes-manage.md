# Create and Test with a Sandbox Company

When you [create your developer profile](https://developer.intuit.com/app/developer/qbo/docs/get-started), you automatically get a sandbox company.

Sandbox companies are regionally-specific QuickBooks Online companies with sample data. They look and act just like a normal QuickBooks Online experience.

Use it as a testing environment during development to see your code in action. You can also connect them to external testing tools like [Postman](https://developer.intuit.com/app/developer/qbo/docs/develop/sandboxes/postman) or Insomnia.

> **Note**: Sandbox companies are for app development only. Intuit provides this environment, content, and sample data for non-commercial use and testing.

## Create a Sandbox Company

1. [Sign in](https://developer.intuit.com/dashboard) to your developer account.
2. Select **My Hub > Sandboxes** from the upper-right corner of the toolbar.
3. Select **Add**, located on the right.
4. Select **QuickBooks Online Plus** or **QuickBooks Online Advanced**.
5. If you select QuickBooks Online Plus, select a country from the **Country** dropdown.
   > **Note**: Sandbox companies are region-specific. You can't change this later on.
6. Select **Create**.

You can create up to 10 sandbox companies. They're valid and active for two years.

> **Tip**: Each [QuickBooks Online SKU](https://quickbooks.intuit.com/pricing/) comes with specific features. We recommend you test with the one your users use the most. Develop your app for features the majority of your users have access to.

## Test Code with a Sandbox Company

There are a few ways to use your sandbox company:

- Sandboxes automatically connect to our testing tools, like the [OAuth 2.0 Playground](https://developer.intuit.com/app/developer/qbo/docs/develop/authentication-and-authorization/oauth-2.0-playground)
- Connect them to third-party testing tools

Here's what you'll need to connect and use your sandbox company.

### Get a Sandbox's Credentials

1. [Sign in](https://developer.intuit.com/dashboard) to your developer account.
2. Select your app.
3. Select **Keys and credentials** from the left navigation pane.
4. Select **Development** and turn on **Show Credentials**.
5. Retrieve or copy the [Client ID and Client Secret](https://developer.intuit.com/app/developer/qbo/docs/get-started/get-client-id-and-client-secret).

### Get the Sandbox Base URL

Use this base URL for sandbox companies:

**`https://sandbox-quickbooks.api.intuit.com/v3`**

### Open a Sandbox Company to See Your Code in Action

1. [Sign in](https://developer.intuit.com/dashboard) to your developer account.
2. Select the name of the sandbox company.

> **Tip**: You can't open multiple sandboxes or QuickBooks Online companies in the same browser at once. Here are a few ways to [get around this](https://developer.intuit.com/app/developer/qbo/docs/develop/sandboxes/sandbox-faqs).

## Clear Data and Reset the Sandbox

You can delete all sample data and anything else you've entered since you started. This deletes all the data, but keeps the sandbox company.

1. [Sign in](https://developer.intuit.com/dashboard) to your developer account.
2. Find the sandbox company on the list.
3. Select **Clear data and reset** from the button in the **Action** column.

You can also completely delete a sandbox company if you no longer need it. To delete the sandbox entirely, select **Delete entire sandbox** from the button in the **Action** column.

## Learn More About Sandbox Companies

For a deeper dive, check out the [sandbox FAQ](https://developer.intuit.com/app/developer/qbo/docs/develop/sandboxes/sandbox-faqs). You can also connect with other developers on [the Intuit Developer community forum](https://developer.intuit.com/hub).

If you need more help, or run into errors, [reach out to our developer support team](https://help.developer.intuit.com/s/contactsupport).
