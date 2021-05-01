package faceeasco

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type WebsocketRequestHeader struct {
	Cmd  string `json:"cmd"` //ping或to_device
	From string `json:"from"`
	To   string `json:"to"`
	SN   string `json:"sn,omitempty"` //心跳时使用
}
type WebsocketRequestData struct {
	Cmd string `json:"cmd"`
}

type WebsocketResponseHeader struct {
	Cmd  string `json:"cmd"`  //to_client
	From string `json:"from"` //区别程序，自定义值
	To   string `json:"to"`   //设备编号
}

//数据里的信息
type WebsocketResponseData struct {
	Cmd  string `json:"cmd"`  //操作命令，如：addUser
	Code int    `json:"code"` //0成功，其它失败
	Msg  string `json:"msg"`  //提示信息
}

//数据包
type WebsocketRequest struct {
	WebsocketRequestHeader
	Data WebsocketRequestData `json:"data"`
}

//公共信息
type WebsocketResponse struct {
	WebsocketResponseHeader
	Data WebsocketResponseData `json:"data"`
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var hub *Hub

func init() {
	hub = newHub()
	go hub.run()
}
func ServeWebsocket(w http.ResponseWriter, r *http.Request, responseHeader http.Header) (err error) {
	conn, err := upgrader.Upgrade(w, r, responseHeader)
	if err != nil {
		log.Println(err)
		return
	}

	client := &WebsocketClient{hub: hub, conn: conn, send: make(chan []byte, 256)}
	client.hub.register <- client

	go client.readPump()
	go client.writePump()
	return
}

func SendWebsocketCommand(sn string, cmd interface{}) (err error) {
	bytes, err1 := json.Marshal(cmd)
	if err1 != nil {
		err = err1
		return
	}

	err = hub.send(sn, bytes)
	return
}
