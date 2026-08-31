// Package ntfy sends push notifications to your phone, desktop or
// browser via the ntfy.sh push service.
//
// The whole API is one function: call [NewTheme] with a topic name and
// send a message to it. Subscribing happens on the ntfy.sh side —
// install the app, subscribe to the topic, and you're done.
package ntfy

import (
	"fmt"

	"github.com/lazylib/request"
)

// NtfyTheme is a named topic on the ntfy.sh push service.
//
// Any client subscribed to the topic (phone app, browser, ntfy CLI)
// receives every message sent to it.
type NtfyTheme struct {
	theme string
}

// NewTheme creates a handle for the given topic.
//
// A topic is just a free string — pick anything unique enough to avoid
// collisions, for example "myapp_deploy_alerts".
func NewTheme(theme string) NtfyTheme {
	return NtfyTheme{theme: theme}
}

// Send pushes message to the topic as a plain-text notification.
//
// Messages are sent with a fire-and-forget HTTP PUT to ntfy.sh — the
// result is not returned and delivery errors are not surfaced. The ntfy
// service keeps them for 12 hours by default.
func (n NtfyTheme) Send(message string) {
	if n.theme == "" {
		return
	}

	request.Send[any](request.Options{
		Method:  "PUT",
		Url:     fmt.Sprintf("https://ntfy.sh/%s", n.theme),
		Body:    message,
		Headers: map[string]string{"Content-Type": "text/plain"},
	})
}
