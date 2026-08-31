# Contributing

Thanks for your interest in `lazylib/ntfy`! This is a deliberately small
package — the whole point is "one function, zero config" — so the bar for
new features is high.

## Ground rules

1. **No new dependencies.** The entire API is `NewTheme` + `Send`.
2. **Tiny API surface.** Prefer improving the two existing functions over
   adding new ones.
3. **Backwards compatibility.** Any change to the signature of
   `NewTheme` or `Send` is breaking and needs a major version bump.
4. **Fire-and-forget stays fire-and-forget.** `Send` intentionally does
   not return errors — call sites should not need to grow error handling
   just to push a notification.

## Local development

```bash
go build ./...
go vet ./...
gofmt -l .              # should print nothing
go test ./...           # if tests are added
```

## Submitting a PR

1. Fork the repo and create a topic branch.
2. Make sure CI is green.
3. Squash commits; write a message in the imperative mood
   (`send notifications over HTTP`, not `sended …`).

## Reporting a bug

Open a **Bug report** issue. Include a minimal reproduction and your
`go version`.

## Suggesting a feature

Open a **Feature request**. Be explicit about whether it is a breaking
change, and what you considered instead. Many ideas (titles, priorities,
Markdown rendering, `click` actions) are better served by the ntfy.sh
service itself than by this package — check its
[documentation](https://docs.ntfy.sh) first.
