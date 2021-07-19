package faceeasco

import (
	"encoding/json"
	"github.com/yansuan/faceeasco/queue"
	"log"
	"time"
)

//对接api的接口
type WebsocketApi struct {
	Cmd      string
	FN       func(message []byte) (err error) //注册数接口
	Response func(message []byte) interface{} //响应接口
}

func SendWebsocketMessage(requestId, sn string, data interface{}) (result []byte, err error) {
	var message []byte
	if bs, ok := data.([]byte); ok {
		message = bs
	} else {
		message, err = json.Marshal(data)
		if err != nil {
			return
		}
	}

	queueClient := queue.Connect(requestId)
	err = hub.send(sn, message)

	body := make([]byte, 0)
	isTimeOut := true
	select {
	case body = <-queueClient.Message:
		isTimeOut = false
	case <-time.After(time.Second * queue.TIMEOUT):
		isTimeOut = true
		break
	}
	
	queue.Debug()

	if isTimeOut {
		err = ErrorTimeout
		return
	}

	result = body

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

func GetClientList() []string {
	var result = make([]string, 0)
	for client := range hub.clients {
		if client.sn != "" {
			result = append(result, client.sn)
		}
	}

	return result
}

func callApi(client *WebsocketClient, message []byte) (err error) {
	rData := &WebsocketRequest{}
	err = json.Unmarshal(message, rData)
	if err != nil {
		log.Println(err)
		return
	}

	if rData.Cmd == WEBSOCKET_API_PING {
		//client.sn = rData.SN
		api := WebsocketApiPing{}
		respMessage := api.Response(message)
		err = client.SendMessage(respMessage)
		return
	}

	if rData.Cmd == WEBSOCKET_API_DECLARE {
		client.sn = rData.SN
		client.hub.register <- client
		api := WebsocketApiDeclare{}
		respMessage := api.Response(message)
		err = client.SendMessage(respMessage)
		return
	}

	//to_client,add queue
	msg := &queue.Message{}
	msg.Time = time.Now().Unix()
	msg.RequestId = rData.To
	msg.SN = rData.From
	msg.Body = message

	queue.Push(msg)

	return
}
