package ntfy

import (
	"fmt"

	"github.com/lazylib/request"
)

type NtfyTheme struct {
	theme string
}

func NewTheme(theme string) NtfyTheme {
	return NtfyTheme{theme: theme}
}

func (n *NtfyTheme) Send(message string) {
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
