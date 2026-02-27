# Set up OAuth 2.0

> **Note about deprecation of TLS1.0 and TLS1.1**: As we mentioned in our previous blog posts, we're no longer supporting TLS1.0 and TLS1.1 starting June 1st, 2022. We'll only support TLS1.2 or higher going forward. Refer to our blog for more details:
> - [TLS 1.0 and 1.1 Disablement for the Intuit Developer Group](https://blogs.intuit.com/blog/2017/07/11/tls-1-0-1-1-disablement-intuit-developer-group/)
> - [Upgrading your apps to support TLS 1.2](https://blogs.intuit.com/blog/2017/08/03/upgrading-your-apps-to-support-tls-1-2/)

Use the [OAuth 2.0 protocol](https://tools.ietf.org/html/rfc6749) to implement authentication and authorization. Authorization is essential for both testing via sandbox companies and production apps.

We'll show you how to set up the authorization flow so users can authorize to your app and give it permission to connect to their QuickBooks Online company.

If users grant permission, our Intuit OAuth 2.0 Server sends an authorization code back to your app. Your app exchanges this code for access tokens. These tokens are tied to your users' now authorized QuickBooks Online company (identified by the **realmID**).

Your app needs access tokens to [make API calls](https://developer.intuit.com/app/developer/qbo/docs/learn/explore-the-quickbooks-online-api) and interact with QuickBooks Online data.

## Step 1: Create your app on the Intuit Developer Portal

Start by [signing in](https://developer.intuit.com/dashboard) to your developer account and [creating an app on the Intuit Developer Portal](https://developer.intuit.com/app/developer/qbo/docs/get-started/start-developing-your-app).

When you create your app, select the [QuickBooks Online Accounting scope](https://developer.intuit.com/app/developer/qbo/docs/learn/scopes).

This app provides the credentials you'll need for authorization requests.

## Step 2: Practice authorization in the OAuth Playground

[Check out the OAuth Playground](https://developer.intuit.com/app/developer/playground) to preview each step of the authorization flow. We've provided sample data, like the redirect URI, so you can focus on the overall flow.

Using the OAuth Playground isn't required, but we recommend it.

## Step 3: Start developing with an SDK

Our SDKs come with a built-in OAuth 2.0 Client Library and handle many parts of the authorization implementation for you. For instance, our SDKs implement handlers that automatically exchange the authorization codes sent from the Intuit OAuth 2.0 Server for access tokens.

Select a link to download an SDK:

- [OAuth 2.0 Client Library for .Net](https://github.com/intuit/QuickBooks-V3-DotNET-SDK)
- [OAuth 2.0 Client Library for Java](https://github.com/intuit/QuickBooks-V3-Java-SDK)
- [OAuth 2.0 Client Library for PHP](https://github.com/intuit/QuickBooks-V3-PHP-SDK)
- [OAuth 2.0 Client Library for Node.js](https://github.com/intuit/oauth-jsclient)
- [OAuth 2.0 Client Library for Python](https://github.com/intuit/oauth-pythonclient)
- [OAuth 2.0 Client Library for Ruby](https://github.com/intuit/oauth-rubyclient)

**Note**: Our SDKs are only for OAuth 2.0 and QuickBooks Online. SDKs aren't required to develop an awesome app. However, given the importance of OAuth implementation, we strongly encourage you use them.

## Step 4: Understand the end-to-end authorization flow

While authorization is a simple step for app users, it involves several tasks on the backend. Here's a brief overview:

**Creating authorization requests**

- [Review the scopes](https://developer.intuit.com/app/developer/qbo/docs/learn/scopes) you selected for your app. These determine what data and API it can access. You may want to use the **openID** scope to [implement OpenID Connect](https://developer.intuit.com/app/developer/qbo/docs/develop/authentication-and-authorization/openid-connect) in addition to others.
- Configure your app so it uses the correct credentials (i.e., **Client ID** and **Client secret**) and redirect URIs.
- Review the base URLs in our [authorization discovery documents](https://developer.intuit.com/app/developer/qbo/docs/develop/authentication-and-authorization/oauth-openid-discovery-doc).
- Use this info to create authorization requests.

**Managing the authorization flow**

- When a user connects to your app, it sends an authorization request to the Intuit OAuth 2.0 Server.
- The user gets redirected to an authorization page where they can give your app permission to access their QuickBooks Online company and its data. This is the "user consent" step of the process.
- If the user authorizes your app, the Intuit OAuth 2.0 Server sends an authorization code back to your app.

**Getting access and refresh tokens**

- Your app sends the authorization code back to the Intuit OAuth 2.0 Server and exchanges it for access and refresh tokens.
- Your app extracts the latest access and refresh tokens from the server response.

**Making API calls**

- Use access tokens to call specific APIs and interact with users' QuickBooks Online company data.
- When access tokens expire, use a refresh token to "refresh" the access token.
- If a refresh token expires, users need to go through the authorization flow again.

## Step 5: Get your app's credentials

Sign in to your Intuit Developer account and [get your app's Client ID and Client secret](https://developer.intuit.com/app/developer/qbo/docs/get-started/get-client-id-and-client-secret).

If you're implementing authorization for a live, in-production app, go to the **Production section** and select **Keys & OAuth** to get your credentials.

If you're implementing authorization for testing with a sandbox company, go to the **Development section** and select **Keys & OAuth** to get your credentials.

## Step 6: Learn about discovery documents

OAuth 2.0 requires multiple URLs for authentication and requests for resources like tokens, user info, and credentials.

[Use our discovery documents](https://developer.intuit.com/app/developer/qbo/docs/develop/authentication-and-authorization/oauth-openid-discovery-doc) to simplify the implementation. These JSON documents, found at a well known location, contain field:value pairs and URL endpoints for `authorization`, `token`, `userinfo`, and other data points.

## Step 7: Add your app's redirect URIs

URIs handle responses from the Intuit OAuth 2.0 Server during the authorization flow. Basically, they're your app's endpoints.

[Add at least one redirect URI](https://developer.intuit.com/app/developer/qbo/docs/develop/authentication-and-authorization/set-redirect-uri) for your app.

If you're developing with an SDK, use the URI value generated by the SDK.

**Tip:** As a best practice, design your app's endpoints so they don't expose authorization codes to other resources on the page. Learn more about [URIs and OAuth 2.0](https://developer.intuit.com/app/developer/qbo/docs/develop/authentication-and-authorization/faq).

## Step 8: Create an authorization request

Create the authorization request your app will send to the Intuit OAuth 2.0 Server when users connect to your app.

### Authorization request parameters

Request parameters should identify your app and [include the required scopes](https://developer.intuit.com/app/developer/qbo/docs/learn/scopes).

| Field | Description | Required |
|-------|-------------|----------|
| `client_id` | Identifies the app making the request. You got the `client_id` value in Step 5. | **Yes** |
| `scope` | Lists the [scopes your app uses](https://developer.intuit.com/app/developer/qbo/docs/learn/scopes). Enter [one or more scopes](https://developer.intuit.com/app/developer/qbo/docs/learn/scopes). The list should be space-delimited. The `scope` value defines the type of data an app can utilize. This information appears on the authorization page users see when they connect to your app. **Tip**: We recommend apps request scopes incrementally based on your feature requirements, rather than all scopes up front. | **Yes** |
| `redirect_uri` | Determines where the Intuit OAuth 2.0 Server redirects users to if they authorize your app. The `redirect` value must match the URI you listed in Step 7, including casing, http scheme, and trailing "/." | **Yes** |
| `response_type` | States if the Intuit OAuth 2.0 endpoint returns an authorization code. Always set the value to "**code**". | **Yes** |
| `state` | Defines the state between your authorization request and the Intuit OAuth 2.0 Server response. The `state` field is used for validation. It checks if the client (i.e., your app) gets the data back that it sent in the original request. **Tip**: We strongly recommend you include an anti-forgery token for the `state` and confirm it in the response. This prevents cross-site request forgery. Learn more [about CSRF](https://developer.intuit.com/app/developer/qbo/docs/develop/authentication-and-authorization/faq). | **Yes** |
| `claims` | **Optional**. A JSON Object containing a set of user claims that your application is requesting via OpenID Connect. Available claims include: `realmId` -- The realm ID of the specific QuickBooks company to which the user is associated. | **No** |

**Important**: You must create authorization requests in a browser modal. If your app doesn't have browser support, use the [OAuth Playground](https://developer.intuit.com/app/developer/qbo/docs/develop/authentication-and-authorization/oauth-2.0-playground), [Postman](https://developer.intuit.com/app/developer/qbo/docs/develop/sandboxes), or another web component to set up the authorization flow.

### If you're developing with a QuickBooks SDK

You can create and configure an object that defines the authorization request with the required parameters. Here are example authorization requests for [supported SDKs](https://developer.intuit.com/app/developer/qbo/docs/develop/authentication-and-authorization/example-sdk-oauth-2.0-integrations):

```csharp
// Instantiate object
public static OAuth2Client oauthClient = new OAuth2Client("clientid", "clientsecret", "redirectUrl", "environment");
// environment is "sandbox" or "production"

//Prepare scopes
List<OidcScopes> scopes = new List<OidcScopes>();
scopes.Add(OidcScopes.Accounting);

//Get the authorization URL
string authorizeUrl = oauthClient.GetAuthorizationURL(scopes);
```

Depending on the SDK, you may need to set up the configuration file with your **Client ID**, **Client secret**, and **Redirect URI**. Simply reuse the values from previous steps.

### If you're creating an HTTPS/REST request manually

Manually create an authorization request in your app's language. You'll need to call the Intuit OAuth 2.0 Server endpoint, generate a URL, and define the URL's parameters.

Get the [base URI from the discovery document](https://developer.intuit.com/app/developer/qbo/docs/develop/authentication-and-authorization/oauth-openid-discovery-doc). You can also follow these links:

- For sandboxes and testing environments: [Click here](https://developer.api.intuit.com/.well-known/openid_sandbox_configuration)
- For production apps: [Click here](https://developer.api.intuit.com/.well-known/openid_configuration)

Use the values from the `authorization_endpoint`.

## Step 9: Redirect users to the authorization page

Your app needs to redirect users to the authorization page via the Intuit OAuth 2.0 Server. This starts the "user consent" step of the process. We'll call this the "authorization flow."

The authorization flow is required when users connect to your app for the first time. It's also required any time you change your app's scopes (i.e., incremental authorization). Users need to grant permission again since your app needs to request access additional data.

Use the authorization URL from your authorization request, or use the following examples.

### Example authorization request URL

Here's an example authorization request URL. It specifies the QuickBooks Online Accounting scope (i.e., **com.intuit.quickbooks.accounting**) and adds line breaks for readability:

```
https://appcenter.intuit.com/connect/oauth2?
  client_id=Q3ylJatCvnkYqVKLmkxxxxxxxxxxxxxxxkYB36b5mws7HkKUEv9aI&response_type=code&
  scope=com.intuit.quickbooks.accounting&
  redirect_uri=https://www.mydemoapp.com/oauth-redirect&
  state=security_token%3D138r5719ru3e1%26url%3Dhttps://www.mydemoapp.com/oauth-redirect
```

### Review manual requests with cURL

```bash
curl -X POST '<Authorization URI from previous step>'
```

## Step 10: Create the UI that redirects users to the authorization page

There are a few ways users can connect to your app:

- Add a "**Connect to**" button somewhere in your app. When users select it, your app should redirect them and start the authorization flow. Here's a quick guide to [creating button UI](https://developer.intuit.com/app/developer/qbo/docs/develop/authentication-and-authorization/faq#in-app-connection).
- If you [list your app in the QuickBooks App Store](https://developer.intuit.com/app/developer/qbo/docs/go-live/list-on-the-app-store), there are a few options. You can simply link to your app's website where users can get your app. You can also [set up Intuit Single Sign-on](https://developer.intuit.com/app/developer/qbo/docs/develop/authentication-and-authorization/single-sign-on-models) so users can connect to your app directly from your app page.

Since designing this UI may come later in development, we won't focus on it here. However you decide to set up the UI, it needs to redirect users to the Intuit OAuth 2.0 Server and open the authorization page.

This authorization page is where users authorize your app and give it permission to access their data. They'll see your app's name and the QuickBooks Online company they're connecting with.

If they authorize your app, the server will redirect them back to the redirect URI you set.

## Step 11: Get the authorization code from server response

At this point, your app is waiting for a response from the Intuit OAuth 2.0 Server.

If users authorize your app, the Intuit OAuth 2.0 Server sends a response to the redirect URI you specified in Step 7. The response contains an authorization code in the `code` field. Copy the `code` value.

### Example server response for authorization codes

```
https://www.mydemoapp.com/oauth-redirect?
  code=4/P7q7W91a-oMsCeLvIaQm6bTrgtp7&
  state=security_token%3D138r5719ru3e1%26url%3Dhttps://www.mydemoapp.com/oauth-redirect&
  realmId=1231434565226279
```

| Parameter | Description |
|-----------|-------------|
| `code` | The authorization code sent by the Intuit OAuth 2.0 Server. Max length: 512 characters. |
| `realmId` | The unique ID of the connected user's QuickBooks Online company. It's also sometimes called the "company ID." Use the `realmId` for subsequent API endpoint URLs to get data from QuickBooks Online companies. |
| `state` | The `state` value sent from the Intuit OAuth 2.0 Server. It should match the `state` sent in the original authorization request. |

If users don't authorize your app, the server sends an `access_denied` error.

If the authorization request has a scope issue, the server sends an `invalid_scope` error.

**Note**: If your response endpoint renders an HTML page, other resources on that page can see the authorization code in the URL. Scripts can read the URL directly, and all resources may be sent to the URL in the caller and Referer HTTP header. Carefully consider if you want to send authorization credentials this way. This is especially important for third-party scripts, such as social plugins and analytics. To avoid issues, we recommend the server first handle the request, then redirect to another URL that doesn't include the response parameters.

## Step 12: Exchange the authorization code for access tokens

Your app should send the authorization code (i.e., the value of the `code` parameter) back to the Intuit OAuth 2.0 server to exchange it for access and refresh tokens.

### If you're developing with a QuickBooks SDK

Use the example code to create an object named **tokenResponse**. This automatically exchanges the authorization code for access and refresh tokens:

```csharp
// Get OAuth2 Bearer token
var tokenResponse = await auth2Client.GetBearerTokenAsync(code);

//retrieve access_token and refresh_token
tokenResponse.AccessToken
tokenResponse.RefreshToken
```

### If you're creating an HTTPS/REST request manually

Get the [base URI from the discovery document](https://developer.intuit.com/app/developer/qbo/docs/develop/authentication-and-authorization/oauth-openid-discovery-doc). You can also follow these links:

- For sandboxes and testing environments: [Click here](https://developer.api.intuit.com/.well-known/openid_sandbox_configuration)
- For production apps: [Click here](https://developer.api.intuit.com/.well-known/openid_configuration)

Create a POST request to exchange the authorization code for access and refresh tokens.

Send requests to the `token_endpoint` (available in the [discovery document](https://developer.intuit.com/app/developer/qbo/docs/develop/authentication-and-authorization/oauth-openid-discovery-doc)) using the following parameters:

**Important**: Only send one request to exchange the authorization code. Multiple requests may invalidate tokens.

### Request parameters for token exchange

| Field | Description | Required |
|-------|-------------|----------|
| `code` | The authorization code your app received from the Intuit OAuth 2.0 response. | **Yes** |
| `redirect_uri` | The redirect URI listed for your app. You set this in Step 7. | **Yes** |
| `grant_type` | The type defined by the OAuth 2.0 server specification. It must have the value **authorization_code**. | **Yes** |

Here's an example token exchange request:

```http
POST https://oauth.platform.intuit.com/oauth2/v1/tokens/bearer HTTP/1.1
Accept: application/json
Authorization: Basic UTM0dVBvRDIwanp2OUdxNXE1dmlMemppcTlwM1d2
  NzRUdDNReGkwZVNTTDhFRWwxb0g6VEh0WEJlR3dheEtZSlVNaFhzeGxma1l
  XaFg3ZlFlRzFtN2szTFRwbw==
Content-Type: application/x-www-form-urlencoded
Host: oauth.platform.intuit.com
Body:
  grant_type=authorization_code&
  code=L3114709614564VSU8JSEiPkXx1xhV8D9mv4xbv6sZJycibMUI&
  redirect_uri=https://www.mydemoapp.com/oauth-redirect
```

### cURL example for token exchange

The Authorization header should follow this format:

```
"Basic " + base64encode(client_id + ":" + client_secret)
```

```bash
curl -X POST 'https://oauth.platform.intuit.com/oauth2/v1/tokens/bearer' \
  -H 'Accept: application/json' \
  -H 'Content-Type: application/x-www-form-urlencoded' \
  -H 'Authorization: REPLACE_WITH_AUTHORIZATION_HEADER (details below)' \
  -d 'grant_type=authorization_code' \
  -d 'code=REPLACE_WITH_AUTHORIZATION_CODE' \
  -d 'redirect_uri=REPLACE_WITH_REDIRECT_URI'
```

### Example server response for exchanged tokens

The server returns a JSON object. The access code is the `access_token` field value.

```json
{
  "token_type": "bearer",
  "expires_in": 3600,
  "refresh_token": "Q311488394272qbajGfLBwGmVsbF6VoNpUKaIO5oL49aXLVJUB",
  "x_refresh_token_expires_in": 15551893,
  "access_token": "eJlbmMiOiJBMTI4Q0JDLUhTMjU2Iiw..."
}
```

| Field | Description |
|-------|-------------|
| `access_token` | The token used to access our API. Max length: 4096 characters. |
| `refresh_token` | The token used for refreshing the access token. Max length: 512 characters. |
| `expires_in` | The remaining lifetime of the access token. The value begins at 3600. This is in seconds (one hour). |
| `x_refresh_token_expires_in` | The remaining lifetime of the current refresh token. This is in seconds. When this expires, users must reauthorize your app. |
| `token_type` | The type of token returned. This always has the value **bearer**. |

**Tip**: The response may have other parameters. Don't treat them as errors. This only lists the minimum set.

## Step 13: Decide if you want to implement OpenID Connect

Your app is now set up with OAuth 2.0.

At this point, you can also [implement Intuit Single Sign-on and OpenID Connect](https://developer.intuit.com/app/developer/qbo/docs/develop/authentication-and-authorization/openid-connect) to further enhance authorization features.

## Step 14: Use access tokens to make API calls

Access tokens let your app send requests to our APIs.

Pass the `access_token` value in the Authorization header of requests each time your app calls an API. The value should always be: **Authorization: bearer {AccessToken}**

Access tokens are valid for 60 minutes (3,600 seconds). When they expire, use refresh tokens to refresh them. Learn more [about access and refresh tokens](https://developer.intuit.com/app/developer/qbo/docs/develop/authentication-and-authorization/faq).

## Step 15: Refresh access tokens

Use refresh tokens to "refresh" expired access tokens. You can refresh access tokens without prompting users for permission.

### If you're developing with a QuickBooks SDK

```csharp
// Instantiate object
public static OAuth2Client oauthClient = new OAuth2Client("clientid", "clientsecret", "redirectUrl", "environment");
// environment is "sandbox" or "production"

//Refresh token endpoint
var tokenResp = await oauthClient.RefreshTokenAsync("refreshToken");
```

### If you're creating an HTTPS/REST request manually

Create a POST request and use the latest `refresh_token` value from the most recent API server response.

Send requests to the `token_endpoint` (available in the [discovery document](https://developer.intuit.com/app/developer/qbo/docs/develop/authentication-and-authorization/oauth-openid-discovery-doc)) using the following parameters:

### Request parameters for refresh tokens

| Field | Description | Required |
|-------|-------------|----------|
| `grant_type` | This is defined in the OAuth 2.0 server specification. It must have the value **refresh_token**. | **Required** |
| `refresh_token` | The `refresh_token` value you exchanged the authorization code for. | **Required** |

Here's an example request:

```http
POST /oauth2/v1/tokens/bearer HTTP/1.1
Accept: application/json
Authorization: Basic UTM0dVBvRDIwanp2OUdxNXE1dmlMemppcTlwM1d2
  NzRUdDNReGkwZVNTTDhFRWwxb0g6VEh0WEJlR3dheEtZSlVNaFhzeGxma1l
  XaFg3ZlFlRzFtN2szTFRwbw==
Content-Type: application/x-www-form-urlencoded
Body:
  grant_type=refresh_token&
  refresh_token=Q311488394272qbajGfLBwGmVsbF6VoNpUKaIO5oL49aXLVJUB
```

### cURL example for refresh tokens

```bash
curl -X POST 'https://oauth.platform.intuit.com/oauth2/v1/tokens/bearer' \
  -H 'Accept: application/json' \
  -H 'Content-Type: application/x-www-form-urlencoded' \
  -H 'Authorization: REPLACE_WITH_AUTHORIZATION_HEADER (details below)' \
  -d 'grant_type=refresh_token' \
  -d 'refresh_token=REPLACE_WITH_REFRESH_TOKEN'
```

**Tip**: Since apps use access and refresh tokens frequently, store them in a secure location that's accessible between different invocations of your app. If access or refresh tokens aren't working, here are a [few areas to troubleshoot](https://developer.intuit.com/app/developer/qbo/docs/develop/authentication-and-authorization/faq).

## Step 16: Revoke access tokens

If users disconnect from your app, it needs to automatically revoke access and refresh tokens. For example, this process can start when users select a **Disconnect** link or button somewhere in your app.

### If you're developing with a QuickBooks SDK

```csharp
// Instantiate object
public static OAuth2Client oauthClient = new OAuth2Client("clientid", "clientsecret", "redirectUrl", "environment");
// environment is "sandbox" or "production"

//Revoke token endpoint
var tokenResp = await oauthClient.RevokeTokenAsync("refreshToken");
```

### If you're creating an HTTPS/REST request manually

Create a POST request and include the `refresh_token` value for the token parameter.

Send the request to the `revocation_endpoint` (available in the [discovery document](https://developer.intuit.com/app/developer/qbo/docs/develop/authentication-and-authorization/oauth-openid-discovery-doc)). Here's an example request:

```http
POST https://developer.api.intuit.com/v2/oauth2/tokens/revoke HTTP/1.1
Accept: application/json
Authorization: Basic UTM0dVBvRDIwanp2OUdxNXE1dmlMemppcTlwM1d2
  NzRUdDNReGkwZVNTTDhFRWwxb0g6VEh0WEJlR3dheEtZSlVNaFhzeGxma1l
  XaFg3ZlFlRzFtN2szTFRwbw==
Content-Type: application/json
{
  "token": "{bearerToken or refreshToken}"
}
```

### cURL example for revoking tokens

```bash
curl -X POST 'https://developer.api.intuit.com/v2/oauth2/tokens/revoke' \
  -H 'Accept: application/json' \
  -H 'Content-Type: application/x-www-form-urlencoded' \
  -H 'Authorization: REPLACE_WITH_AUTHORIZATION_HEADER (details below)' \
  -d 'token:REPLACE_WITH_REFRESH_TOKEN/REPLACE_WITH_ACCESS_TOKEN'
```

### Server responses for revoked tokens

If the app successfully revoked access, the server sends a response code: `status_code 200`.

If it didn't, or there was an error, you'll see `status_code 400`. Review the error message and follow its instructions.

### Determine who disconnects from your app

When a user disconnects, you can identify their company by including `realmId` as a query parameter in the `revoke` endpoint. You can use this information to show the user options to reconnect. For example: `https://myappsite.com/disconnect?realmId=`

## Step 17: Get new refresh tokens

Refresh tokens have a rolling expiry of 100 days.

As long as refresh tokens are valid, you can use them to obtain new access tokens. Always store the latest `refresh_token` value from the most recent API server response. Use it to make requests and obtain new access tokens.

If 100 days pass, or your refresh token expires, users need to go through the authorization flow again and reauthorize your app.

**Tip**: You can also get new refresh tokens programmatically using the Refresh Token API before the 100 days expire.
