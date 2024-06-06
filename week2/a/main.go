package main

import (
	"encoding/json"
	"log"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

type BroadcastMsg struct {
	Type    string `json:"type"`
	Message int    `json:"message"`
}

type ReadMsg struct {
	Type     string `json:"type"`
	Messages []int  `json:"messages"`
}
type TopoMsg struct {
	Type     string              `json:"type"`
	Topology map[string][]string `json:"topology"`
}

var MessageCounting []int

// handleEcho processes the "echo" message and sends a reply.
func handleBroad(n *maelstrom.Node) func(maelstrom.Message) error {
	return func(msg maelstrom.Message) error {
		var body BroadcastMsg
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}
		log.Printf("%v", body)

		MessageCounting = append(MessageCounting, body.Message)
		replyMsg := map[string]string{"type": "broadcast_ok"}
		return n.Reply(msg, replyMsg)
	}
}
func readHandle(n *maelstrom.Node) func(maelstrom.Message) error {
	return func(msg maelstrom.Message) error {
		var body ReadMsg
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}
		replyMsg := ReadMsg{
			Type:     "read_ok",
			Messages: append([]int(nil), MessageCounting...),
		}
		return n.Reply(msg, replyMsg)
	}
}

func handleTopo(n *maelstrom.Node) func(maelstrom.Message) error {
	return func(msg maelstrom.Message) error {
		var body TopoMsg
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}
		replyMsg := map[string]string{"type": "topology_ok"}
		return n.Reply(msg, replyMsg)
	}
}

func main() {
	n := maelstrom.NewNode()
	n.Handle("broadcast", handleBroad(n))
	n.Handle("read", readHandle(n))
	n.Handle("topology", handleTopo(n))
	if err := n.Run(); err != nil {
		log.Fatal(err)
	}
}

// func main() {
// n := maelstrom.NewNode()
// msgCount := make(chan int)
// var mySlice []int
// var mu sync.Mutex
// n.Handle("broadcast", func(msg maelstrom.Message) error {
// var body map[string]any
// if err := json.Unmarshal(msg.Body, &body); err != nil {
// return nil
// }
// if mCount, ok := body["message"].(int); ok {
// msgCount <- int(mCount)
// }

// delete(body, "message")
// body["type"] = "broadcast_ok"
// close(msgCount)
// return n.Reply(msg, body)
// })
// go func() {
// for c := range msgCount {
// mu.Lock()
// mySlice = append(mySlice, c)
// mu.Unlock()
// }
// }()
// n.Handle("read", func(msg maelstrom.Message) error {
// var response map[string]any
// if err := json.Unmarshal(msg.Body, &response); err != nil {
// return nil
// }
// response["message"] = mySlice
// response["type"] = "read_ok"
// return n.Reply(msg, response)
// })

// n.Handle("topology", func(msg maelstrom.Message) error {
// var response map[string]any
// if err := json.Unmarshal(msg.Body, &response); err != nil {
// return nil
// }
// response["type"] = "topology_ok"
// delete(response, "topology")
// return n.Reply(msg, response)

// })
// if err := n.Run(); err != nil {
// log.Fatal(err)
// }
// }
