package faceeasco

import (
	"encoding/json"
	"log"
)

//type IWebsocketApi interface {
//	GetCmd() string                  //得到当前的命令
//	Read(message []byte) (err error) //请示
//	Write() (err error)              //响应
//
//}

var (
	apis = make([]*WebsocketApi, 0)
)

//对接api的接口
type WebsocketApi struct {
	Cmd      string
	FN       func(message []byte) (err error) //注册数据接口
	Response func(message []byte) interface{} //响应接口
}

func newWebsocketApi(cmd string, write func(message []byte) interface{}) {
	apis = append(apis, &WebsocketApi{Cmd: cmd, Response: write})
}

//注册自己业务中的处理函数
func RegisterWebsocketFunc(cmd string, fn func(message []byte) (err error)) {
	for _, api := range apis {
		if api.Cmd == cmd {
			api.FN = fn
		}
	}
	return
}

func SendWebsocketMessage(sn string, data interface{}) (err error) {
	message, err1 := json.Marshal(data)
	if err1 != nil {
		err = err1
		return
	}

	err = hub.send(sn, message)
	return
}

func GetAliveClient(sn string) bool {
	for client := range hub.clients {
		if client.sn == sn {
			return true
		}
	}

	return false
}

func callApi(client *WebsocketClient, message []byte) (err error) {
	r := &WebsocketRequest{}
	err = json.Unmarshal(message, r)
	if err != nil {
		log.Println(err)
		return
	}

	client.sn = r.From

	////读取信息
	var api *WebsocketApi
	for _, a := range apis {
		cmd := r.Cmd
		//ping接口不一样
		if r.Cmd == "to_client" {
			cmd = r.Data.Cmd
		}
		if a.Cmd == cmd {
			api = a
			break
		}
	}

	//set client sn value
	if r.Cmd == WEBSOCKET_API_PING {
		client.sn = r.SN
	}

	//没有找到对应的api
	if api == nil {
		return
	}

	if api.FN != nil {
		err = api.FN(message)
		if err != nil {
			return
		}
	}

	resp := api.Response(message)

	if resp != nil {
		respBytes, err3 := json.Marshal(resp)
		if err3 != nil {
			err = err3
			return
		}

		if len(respBytes) > 0 {
			client.send <- respBytes
		}
	}

	return
}
