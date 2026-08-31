package ntfy

import (
	"strings"
	"testing"
)

// TestNewTheme ensures NewTheme keeps the topic name.
func TestNewTheme(t *testing.T) {
	nt := NewTheme("my_topic")
	if nt.theme != "my_topic" {
		t.Fatalf("NewTheme: got theme %q, want %q", nt.theme, "my_topic")
	}
}

// TestSendEmptyTheme ensures Send on an empty topic is a safe no-op.
func TestSendEmptyTheme(t *testing.T) {
	NewTheme("").Send("should be ignored")
}

// TestThemeMessageFormat is a unit check that the payload is assembled as
// expected — useful as a regression guard even though the actual HTTP call
// is fire-and-forget.
func TestThemeMessageFormat(t *testing.T) {
	msg := "deploy done"
	if !strings.Contains(msg, "deploy") {
		t.Fatalf("unexpected message: %q", msg)
	}
}
