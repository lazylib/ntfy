# lazylib/ntfy

[![Go Reference](https://pkg.go.dev/badge/github.com/lazylib/ntfy.svg)](https://pkg.go.dev/github.com/lazylib/ntfy)
[![Go Report Card](https://goreportcard.com/badge/github.com/lazylib/ntfy)](https://goreportcard.com/report/github.com/lazylib/ntfy)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

Push notifications to your phone, desktop or browser in one line of Go —
no accounts, no API keys, no SDKs to configure.

```go
ntfy.NewTheme("myapp_deploy_alerts").Send("Deploy finished ✅")
```

That's it. One dependency, one exported constructor, one method. Zero
config.

---

## How it works

[ntfy.sh](https://ntfy.sh) is a free, open-source push notification
service. You send an HTTP request to a **topic** — any string you pick —
and every client subscribed to that topic receives it as a notification.

`lazylib/ntfy` wraps that HTTP call so you never write it by hand:

| You do | Happens under the hood |
| ------ | ---------------------- |
| `NewTheme("topic")` | creates a handle for the topic |
| `.Send("message")`  | `PUT https://ntfy.sh/topic` with the message as plain text |

## Install

```bash
go get github.com/lazylib/ntfy
```

## Quick start

```go
package main

import "github.com/lazylib/ntfy"

func main() {
	ntfy.NewTheme("myapp_deploy_alerts").Send("Deploy finished ✅")
}
```

Notify when something long-running finishes:

```go
err := deploy()
if err != nil {
	ntfy.NewTheme("myapp_deploy_alerts").Send("Deploy FAILED: " + err.Error())
} else {
	ntfy.NewTheme("myapp_deploy_alerts").Send("Deploy OK")
}
```

Triggered by a GitHub Actions workflow:

```go
ntfy.NewTheme("github_actions").Send("CI build #42 is green 🎉")
```

## Receiving notifications

The sending side is done. To actually get the notifications:

1. Install the app — [Android](https://play.google.com/store/apps/details?id=io.heckel.ntfy), [iOS](https://apps.apple.com/us/app/ntfy/id1625396347), or use the [web app](https://ntfy.sh/app).
2. Subscribe to your topic: tap "add subscription" and enter `myapp_deploy_alerts`.
3. Done. Every `Send` call pops up as a notification on the device.

You can also subscribe via the CLI:

```bash
ntfy sub myapp_deploy_alerts
```

## Choosing a topic

A topic is just a free string — there is no registration. Pick something
reasonably unique to avoid collisions, e.g. `yourname_appname_events`
instead of `deploy`. Anyone who knows the topic can subscribe to it, so
don't send anything sensitive through it.

## API

The whole public surface is two things:

- `ntfy.NewTheme(theme string) NtfyTheme` — create a handle for a topic.
- `(NtfyTheme) Send(message string)` — push a plain-text message to the
  topic. Fire-and-forget: it performs an HTTP `PUT` to `ntfy.sh` and
  does not return a result or surface delivery errors.

Messages are kept by ntfy.sh for **12 hours** by default, so a
notification sent while your phone is offline still arrives when it
reconnects.

## Project status

Active. The API is intentionally minimal — `NewTheme` + `Send`. The
build grows only if the ntfy service adds something worth wrapping, and
always additively, keeping existing call sites stable.

## Contributing

Bug reports and PRs welcome. Run `go test ./...` and `go vet ./...`
before submitting. See [CONTRIBUTING.md](CONTRIBUTING.md) and
[CODE_OF_CONDUCT.md](CODE_OF_CONDUCT.md).

## License

[MIT](LICENSE).
