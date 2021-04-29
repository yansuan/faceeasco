package faceeasco

import (
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func NewWebsocket(w http.ResponseWriter, r *http.Request, responseHeader http.Header) (err error) {
	conn, err1 := upgrader.Upgrade(w, r, responseHeader)
	if err1 != nil {
		err = err1
		return
	}

	//接收发送过来的ping信息，缓存状态

	cache := &Cache{}
	cache.Add("sn", conn)

	return
}

func websocketPing() (err error) {
	return
}

