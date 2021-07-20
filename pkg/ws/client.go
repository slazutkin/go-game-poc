package ws

import (
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
	ID   string
	Conn *websocket.Conn
	Pool *Pool
}

func (c *Client) Read() {
	defer func() {
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()

	for {
		_, p, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		j := json.RawMessage{}
		if err := j.UnmarshalJSON(p); err != nil {
			log.Println(err)
			return
		}

		message := InboundMessage{ClientID: c.ID, Data: j}
		c.Pool.Incoming <- message
	}
}
