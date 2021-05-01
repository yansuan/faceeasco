package faceeasco

import "errors"

type Hub struct {
	// Registered clients.
	clients map[*WebsocketClient]bool

	// Inbound messages from the clients.
	broadcast chan []byte

	// Register requests from the clients.
	register chan *WebsocketClient

	// Unregister requests from clients.
	unregister chan *WebsocketClient
}

func newHub() *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *WebsocketClient),
		unregister: make(chan *WebsocketClient),
		clients:    make(map[*WebsocketClient]bool),
	}
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}

func (h *Hub) send(sn string, message []byte) (err error) {
	isSend := false
	for client := range h.clients {
		if client.sn == sn {
			client.send <- message
			isSend = true
			break
		}
	}

	if !isSend {
		err = errors.New("websocket.send.err.10001")
	}

	return
}
