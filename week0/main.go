package main

import (
	"encoding/json"
	"log"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

// handleEcho processes the "echo" message and sends a reply.
func handleEcho(n *maelstrom.Node) func(maelstrom.Message) error {
	return func(msg maelstrom.Message) error {
		var body map[string]any
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}
		body["type"] = "echo_ok"
		return n.Reply(msg, body)
	}
}

func main() {
	n := maelstrom.NewNode()
	n.Handle("echo", handleEcho(n))

	if err := n.Run(); err != nil {
		log.Fatal(err)
	}
}
