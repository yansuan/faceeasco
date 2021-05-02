package faceeasco

type WebsocketApiPing struct {
}

type WebsocketApiPingRequest struct {
}

type WebsocketApiPingResponse struct {
	Cmd string `json:"cmd"`
}

func (this *WebsocketApiPing) Response(r []byte) interface{} {
	resp := &WebsocketApiPingResponse{}
	resp.Cmd = "pong"

	return resp
}
