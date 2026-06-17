# Security Policy

## Supported versions

Only the latest released version of `qbo` receives security fixes. Please
upgrade (`brew upgrade qbo`, `scoop update qbo`, or `go install` the latest tag)
before reporting.

## Reporting a vulnerability

Please report security issues **privately** — do not open a public issue.

- Preferred: use GitHub's [private vulnerability reporting](https://github.com/voska/qbo-cli/security/advisories/new)
  (the "Report a vulnerability" button on the repository's **Security** tab).

We aim to acknowledge reports within 5 business days and to ship a fix or
mitigation for confirmed issues as quickly as is practical, crediting reporters
who wish to be named.

## Handling credentials

`qbo` touches financial data, so a few notes on how it treats secrets:

- OAuth access/refresh tokens are stored in the OS keyring (or an encrypted
  file fallback), never in plaintext config.
- `QBO_CLIENT_ID` / `QBO_CLIENT_SECRET` are read from the environment. Never
  commit them, and prefer a secrets manager over shell history.
- Data is written to stdout and hints/progress to stderr; avoid piping verbose
  output into shared logs that might capture tokens or company data.
