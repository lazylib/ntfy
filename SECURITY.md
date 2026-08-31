# Security Policy

## Supported versions

| Version  | Supported          |
| -------- | ------------------ |
| latest   | :white_check_mark: |
| < latest | :x:                |

This package has a tiny API and is maintained on a best-effort basis.
Bug fixes and security fixes are released as new minor versions.

## Reporting a vulnerability

Please **do not** open a public GitHub issue for security
vulnerabilities. Instead, email the maintainers (see the commit history
for contact info) with:

- a description of the issue
- a minimal reproducer
- the impact you believe it has

We will acknowledge within 72 hours and aim to ship a fix within 7 days
for confirmed issues.

## What this package does NOT do

Out of scope for security reports (please don't file these):

- **Message confidentiality.** Notifications are sent to a public
  topic on the public `ntfy.sh` service — anyone who knows the topic
  can read them. Never send secrets through it.
- **Authentication.** There is no auth in this package; the topic name
  is the only access control. Use long, unpredictable topic names.
- **TLS / certificate validation** — delegated to `crypto/tls` in the
  standard library.
