# Authorization FAQ

Here's a deeper dive into the [OAuth 2.0 and authorization process](https://developer.intuit.com/app/developer/qbo/docs/develop/authentication-and-authorization/oauth-2.0) to connect apps with the Intuit Financial Ecosystem.

## Do I need to authorize my app to integrate it with QuickBooks?

Yes. The first step is to [create an app](https://developer.intuit.com/app/developer/qbo/docs/get-started/start-developing-your-app) on the Intuit Developer Platform. This generates credentials you'll use to [connect to the Intuit OAuth 2.0 server](https://developer.intuit.com/app/developer/qbo/docs/develop/authentication-and-authorization/oauth-2.0), get access tokens, and make API calls.

OAuth 2.0 is our authorization standard. Only registered apps can connect to the Intuit OAuth2.0 server, call our API, and end-users' QuickBooks Online companies.

## What credentials do I need to connect to OAuth?

Here's what you need to get started. Keep in mind, different API entities and features may require additional info:

- **Client ID** and **Client secret**: The unique credentials that identify your app
- **Scope**: The level of data access of your app
- **Redirect URI**: The URI your app sends users to when they connect to your app
- **State/CSRF token**: This is required - the Intuit OAuth 2.0 Server won't accept requests without these parameters

## How do access tokens work?

Access tokens work behind the scenes. They allow the connection between your app, the Intuit OAuth 2.0 server, our API, and an end-user's device.

Tokening starts when users connect to your app. Your app sends users an authorization request. This request redirects users to sign in to their QuickBooks Online company. We ask if they're willing to give your app permission to access their data. This flow is called "user consent."

If a user grants permission, the Intuit OAuth 2.0 Server sends an authorization code to your app. Your app then sends the authorization code back to the server to get an access token. Use the access token to call our API and query objects or perform operations.

If the user doesn't grant the permission, the server returns an error and your app can't call our API.

## How long do access and refresh tokens last?

### Access tokens

Access tokens are valid for **3,600 seconds** (or one hour).

When it expires, use the latest `refresh_token` value from the most recent server response to "refresh" it.

If an API request returns a 401 unauthorized message, it means the access token has expired.

### Refresh tokens

Refresh tokens are valid for **100 days**. This expiry date is rolling and gets extended each time it's used to refresh an access token.

Refresh tokens are only for getting new access tokens.

As long as the refresh token itself hasn't expired, each time you refresh your access token, your app periodically updates the `refresh_token` value. This is for security. When you get a new refresh token, the previous refresh token value automatically expires.

**Tip**: We update the value of the `refresh_token` value every 24 hours, or the next time you refresh the access tokens after 24 hours. This is an additional security measure by Intuit to reduce risk of compromise.

We recommend you always store and use the latest `access_token` and `refresh_token` value from the most recent server response when you refresh access tokens.

If the refresh token hasn't been used in the last 100 days, meaning a user hasn't connected to your app and thus no API calls have been made, the refresh token expires. Your app's access to the QuickBooks company terminates. Users will need to sign in and authorize your app again.

## Why does the refresh_token value change after 24 hours?

Each time you make an API call and refresh your app's access token, our server also sends your app a new value for the `refresh_token`.

Keep in mind, we periodically update and change `refresh_token` values. This is an additional security measure by Intuit to reduce risk of compromise. While the refresh token itself doesn't expire, the value changes every 24 hours. This makes the previous `refresh_token` value stale.

We recommend you always store and use the latest `access_token` and `refresh_token` value from the most recent server response when you refresh access tokens.

## Why is my access token invalid?

Access tokens may stop working for a few reasons:

- **The access token expired**: If your app uses an access token that's 60 min or older to make an API call, you'll get a 401 error. You need to [use your refresh token to renew the access token](https://developer.intuit.com/app/developer/qbo/docs/develop/authentication-and-authorization/oauth-2.0#refresh-the-token).
- **The user revoked access**: If a user disconnects their QuickBooks Online company from your app, it invalidates the access token and refresh token. Your app needs to [ask users to reauthorize the connection](https://developer.intuit.com/app/developer/qbo/docs/develop/authentication-and-authorization/oauth-2.0).
- **The refresh token hasn't been used for 100 days**: Refresh tokens are valid for 100 days. After that, your app must [ask users to reauthorize the connection](https://developer.intuit.com/app/developer/qbo/docs/develop/authentication-and-authorization/oauth-2.0).

As a best practice, write your code to anticipate situations where granted access tokens become invalid.

## How do I fix "invalid_grant" errors?

You may see an **invalid_grant** error when exchanging your authorization code for access tokens or refreshing access tokens.

### When exchanging the authorization code for access tokens

- [Review the redirect URI](https://developer.intuit.com/app/developer/qbo/docs/develop/authentication-and-authorization/set-redirect-uri) for your app. It should match with the redirect URI used in queries to the Intuit OAuth 2.0 server. Redirect URIs shouldn't have any query parameters. If query parameters are required, pass them as `state` variables.
- If you have a multi-threaded app, make sure the developer's application is not using the same parameter name `code` for other purposes. `Code` can't be reused multiple times.
- [Check the Client ID and Client secret](https://developer.intuit.com/app/developer/qbo/docs/get-started/get-client-id-and-client-secret) for your app. Also make sure you use the correct set of keys: development keys should be used for development environments, production keys are for your live code.
- Only exchange access tokens one at a time. If two attempts are made to exchange for access tokens, the first attempt will succeed but the second will return `invalid_grant`.

Our servers may see this as a possible security issue and revoke your refresh tokens for the first successful call. Your next attempt to refresh tokens will return the `invalid_grant` error. At this point, you'll need to start the authorization process from the beginning.

### When refreshing an access token

- Make sure the refresh token hasn't expired. If you use an old refresh token, you'll get an `invalid_grant` error.
- Go to the [developer group status page](https://status.developer.intuit.com/) to see if there are any API service outages, or if our servers are undergoing maintenance. Service outages may return an `invalid_grant` error.
- Check your app's code to see if it's not sending stale or cached values, instead of the latest ones sent via responses from the Intuit OAuth 2.0 Server. If you're [using a QuickBooks Online SDK](https://developer.intuit.com/app/developer/qbo/docs/develop/sdks-and-samples), make sure the client object is updated with the latest token object.
- Make sure the refresh token hasn't been revoked. Using a revoked refresh token will result in an error.
- Requests to refresh tokens require the current `refresh_token`. Refresh tokens one at a time. If two attempts are made with the same parameter, the first attempt will succeed, but the second will return an `invalid_grant` error. Multiple attempts may invalidate your refresh token.

> **Important**: We periodically update `refresh_token` values. It's a good idea to always store and use the latest `refresh_token` value from the most recent server response.

## What's a redirect URI? Do I need one?

Redirect URIs are used for authorization. It's where the Intuit OAuth 2.0 Server sends users after they agree to authenticate your app during the authorization process.

It's also where your app sends requests from and our servers return responses after users connect to your app.

Code at the URI location needs to process the initial authorization server request, construct requests for access and refresh tokens, and manage tokens. The Intuit OAuth 2.0 Server can only redirect and send responses to registered URIs.

## How do I register a redirect URI?

Apps have two URIs: one for development (sandbox) environments and one for your production app.

Follow these steps to [set your app's URI](https://developer.intuit.com/app/developer/qbo/docs/develop/authentication-and-authorization/set-redirect-uri).

> **Important**: HTTP redirect URIs must be protected with TLS security. The Intuit OAuth 2.0 Server can only redirect to URIs beginning with **https**. IP addresses aren't allowed. This prevents access token interception during the authorization process.

During development, you also need to register redirect URIs for any third-party testing platforms (Postman, Insomnia, etc) or other apps when [testing with your sandbox company](https://developer.intuit.com/app/developer/qbo/docs/develop/sandboxes/manage-your-sandboxes).

## What are the current scopes for the QuickBooks API?

Scopes determine the type and level of access your app gets to QuickBooks Online company data. This also decides which API entities your app can use.

When you [set up OAuth 2.0](https://developer.intuit.com/app/developer/qbo/docs/develop/authentication-and-authorization/oauth-2.0), you'll list your app's scopes in the authorization request. We recommend requesting scopes incrementally, based on your current data requirements, rather than requesting all scopes up front.

Keep in mind, you must initiate the authorization workflow again with the new list of scopes and get new access and refresh tokens. Users also need to reauthorize your app each time you change your app's scopes.

## How can my app protect against cross-site request forgery (CSRF)?

You must protect the security of your users by preventing cross-site request forgery ([CSRF](https://en.wikipedia.org/wiki/Cross-site_request_forgery)) attacks. To mitigate attacks, include a unique session token in your app's authorization request (often referred to as cross-site request forgery [CSRF] tokens).

A good way is to create a string of 30 or so characters constructed using a high-quality random-number generator. Another is creating a hash and comparing it to a key kept secret on your back end.

Use the 'state' parameter in your authorization request to add your unique session token. Your app will match the unique session token with the one returned in the Intuit OAuth 2.0 server response. This verifies the user, not a malicious attacker or bot, started the authorization process.

## Available SDKs and sample integrations

| Language | Client Library | Integration Examples |
|----------|---------------|---------------------|
| Java | Yes | [Using client library](https://github.com/IntuitDeveloper/OAuth2-JavaWithSDK), [Without library](https://github.com/IntuitDeveloper/OAuth2-Java) |
| .Net | Yes | [Using SDK](https://github.com/IntuitDeveloper/OAuth2-Dotnet_UsingSDK), [MVC5](https://github.com/IntuitDeveloper/Oauth2-MVC5-DotnetSampleApp), [Without library](https://github.com/IntuitDeveloper/OAuth2-Dotnet-WithoutSDK) |
| PHP | Yes | [Using client library](https://github.com/IntuitDeveloper/OAuth2.0-demo-php), [Without library](https://github.com/IntuitDeveloper/OAuth2_PHP) |
| Ruby | No | [Yes](https://github.com/IntuitDeveloper/OAuth2_RubyOnRails), [Using client library](https://github.com/intuit/oauth-rubyclient) |
| Node.js | Yes | [Using client library](https://github.com/intuit/oauth-jsclient/tree/master/sample), [Using 3rd Party Library](https://github.com/IntuitDeveloper/OAuth2.0-demo-nodejs), [Without library](https://github.com/IntuitDeveloper/oauth2-nodejs) |
| Python | Yes | [Using client library](https://github.com/IntuitDeveloper/OAuth2PythonSampleApp), [Without library](https://github.com/IntuitDeveloper/OAuth2-Python) |
