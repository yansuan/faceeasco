package faceeasco

import (
	"github.com/gorilla/websocket"
	"github.com/yansuan/faceeasco/queue"
	"log"
	"net/http"
	"time"
)

type Client struct {
	Queue *queue.Queue
	Hub   *Hub
}

func New() *Client {
	inst := &Client{}
	inst.Queue = queue.NewQueue(TIMEOUT)
	inst.Hub = newHub()
	return inst
}

type WebsocketRequestHeader struct {
	Cmd  string `json:"cmd"`          //to_device
	From string `json:"from"`         //deviceId
	To   string `json:"to"`           //requestId
	SN   string `json:"sn,omitempty"` //ping或declare时使用
}

type WebsocketRequestData struct {
	Cmd string `json:"cmd"`
}

type WebsocketResponseHeader struct {
	Cmd  string `json:"cmd"`  //ping、declare、to_client
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
	Data interface{} `json:"data"`
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

func (this *Client) ServeWebsocket(w http.ResponseWriter, r *http.Request, responseHeader http.Header) (err error) {
	conn, err := upgrader.Upgrade(w, r, responseHeader)
	if err != nil {
		log.Println(err)
		return
	}

	wsClient := &WebsocketClient{client: this, conn: conn, send: make(chan []byte, 256), connTime: time.Now()}
	this.Hub.register <- wsClient

	go wsClient.readPump()
	go wsClient.writePump()
	return
}
