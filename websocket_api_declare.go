package faceeasco

type WebsocketApiDeclare struct {
}

type WebsocketApiDeclareRequest struct {
	Cmd         string `json:"cmd"`  //
	Type        string `json:"type"` //
	SN          string `json:"sn"`   //
	VersionCode string `json:"version_code"`
	VersionName string `json:"version_name"`
}

type WebsocketApiDeclareResponse struct {
	Cmd      string `json:"cmd"`
	ClientId string `json:"client_id"`
}

func (this *WebsocketApiDeclare) Response(r []byte) interface{} {
	resp := &WebsocketApiDeclareResponse{}
	resp.Cmd = "declare"
	resp.ClientId = "test"

	return resp
}
