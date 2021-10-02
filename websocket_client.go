package faceeasco

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"time"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 30 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 10 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

type WebsocketClient struct {
	client *Client
	sn       string
	conn     *websocket.Conn
	send     chan []byte
	connTime time.Time
}

func (c *WebsocketClient) GetSN() string {
	return c.sn
}
func (c *WebsocketClient) GetConnTime() time.Time {
	return c.connTime
}

func (c *WebsocketClient) readPump() {
	defer func() {
		c.client.Hub.unregister <- c
		c.conn.Close()
	}()
	//c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error {
		c.conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("websocket IsUnexpectedCloseError: %v", err)
			} else {
				log.Printf("websocket Error: %v", err)
			}
			break
		}

		err = c.client.callApi(c, message)
		if err != nil {
			log.Println("readPump", err)
		}

		if Debug {
			log.Println("readPump", string(message))
		}

		//c.hub.broadcast <- message
	}
}

func (c *WebsocketClient) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			if Debug {
				resp := &WebsocketResponse{}
				errJson := json.Unmarshal(message, resp)
				dumpBody, _ := json.Marshal(resp)
				log.Println("writePump", errJson, string(dumpBody))
			}

			// Add queued chat messages to the current websocket message.
			n := len(c.send)
			for i := 0; i < n; i++ {
				//w.Write(newline)
				w.Write(<-c.send)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}

func (c *WebsocketClient) SendMessage(message interface{}) (err error) {
	if message == nil {
		return
	}

	if bytes, ok := message.([]byte); ok {
		c.send <- bytes
		return
	}

	var bytes []byte
	bytes, err = json.Marshal(message)
	if err != nil {
		return
	}

	if Debug {
		log.Println(string(bytes))
	}

	c.send <- bytes
	return
}
