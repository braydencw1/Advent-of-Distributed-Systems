package main

import (
	"encoding/json"
	"log"

	"github.com/google/uuid"
	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

// maelstrom "github.com/jepsen-io/maelstrom/demo/go"
// challenge 0
//
//	func main() {
//		n := maelstrom.NewNode()
//		n.Handle("echo", func(msg maelstrom.Message) error {
//			var body map[string]any
//			if err := json.Unmarshal(msg.Body, &body); err != nil {
//				return nil
//			}
//			body["type"] = "echo_ok"
//
//			return n.Reply(msg, body)
//		})
//		if err := n.Run(); err != nil {
//			log.Fatal(err)
//		}
//	}

// challenge 1
func main() {
	n := maelstrom.NewNode()
	n.Handle("generate", func(msg maelstrom.Message) error {
		var body map[string]any
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return nil
		}
		body["type"] = "generate_ok"
		body["id"] = uuid.New()

		return n.Reply(msg, body)
	})
	if err := n.Run(); err != nil {
		log.Fatal(err)
	}
}
