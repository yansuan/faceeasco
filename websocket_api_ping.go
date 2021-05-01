package faceeasco

import "log"

type WebsocketApiPing struct {
}

type WebsocketApiPingRequest struct {
}

type WebsocketApiPingResponse struct {
	Cmd string `json:"cmd"`
}

func init() {
	log.Println("WebsocketApiPing init")
	apiPing := &WebsocketApiPing{}
	newWebsocketApi(WEBSOCKET_API_PING, apiPing.Response)
}

func (this *WebsocketApiPing) Response(r []byte) interface{} {
	resp := &WebsocketApiPingResponse{}
	resp.Cmd = "pong"

	return resp
}
