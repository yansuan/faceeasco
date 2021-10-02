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

func (this *Client) SendWebsocketMessage(requestId, sn string, data interface{}) (result []byte, err error) {
	var message []byte
	if bs, ok := data.([]byte); ok {
		message = bs
	} else {
		message, err = json.Marshal(data)
		if err != nil {
			return
		}
	}

	queueClient := this.Queue.Connect(requestId)
	defer this.Queue.Disconnect(requestId)

	err = this.Hub.send(sn, message)

	body := make([]byte, 0)
	isTimeOut := true
	select {
	case body = <-queueClient.Message:
		isTimeOut = false
	case <-time.After(time.Second * time.Duration(TIMEOUT)):
		isTimeOut = true
		break
	}

	if isTimeOut {
		err = ErrorTimeout
		return
	}

	result = body

	return
}

func (this *Client) GetAliveClient(sn string) *WebsocketClient {
	for client := range this.Hub.clients {
		if client.sn == sn {
			return client
		}
	}

	return nil
}

func (this *Client) GetMessageData() map[string]*queue.Message {
	return this.Queue.GetData()
}

func (this *Client) GetClientList() (result []*WebsocketClient) {
	//var result = make([]string, 0)
	//for client := range hub.clients {
	//	if client.sn != "" {
	//		result = append(result, client.sn)
	//	}
	//}

	for k, _ := range this.Hub.clients {
		result = append(result, k)
	}

	return
}

func (this *Client) callApi(client *WebsocketClient, message []byte) (err error) {
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
		this.Hub.register <- client
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

	this.Queue.Push(msg)

	return
}
