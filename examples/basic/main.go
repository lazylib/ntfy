// Command basic is a complete runnable example: it sends a push
// notification to a topic on the public ntfy.sh service.
//
// Run it with:
//
//	go run ./examples/basic
//
// Then subscribe to the topic (e.g. with the ntfy CLI:
// `ntfy sub myapp_demo_alerts`) to see the notification arrive.
package main

import (
	"fmt"

	"github.com/lazylib/ntfy"
)

func main() {
	topic := "myapp_demo_alerts"

	ntfy.NewTheme(topic).Send("Hello from lazylib/ntfy! 👋")

	fmt.Printf("Notification sent to topic %q.\n", topic)
	fmt.Println("Subscribe to it (e.g. `ntfy sub " + topic + "`) to receive it.")
}
