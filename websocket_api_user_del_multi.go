package faceeasco

type WebsocketApiUserDelMultiRequest struct {
	WebsocketRequestHeader
	Data WebsocketApiUserDelMultiRequestData `json:"data"`
}

type WebsocketApiUserDelMultiRequestData struct {
	WebsocketRequestData
	UserIds  []string `json:"user_ids"`  //
	UserType int      `json:"user_type"` //
}

type WebsocketApiUserDelMultiResponse struct {
	WebsocketResponseHeader
	Data WebsocketApiUserDelMultiDataResponse `json:"data"`
}

type WebsocketApiUserDelMultiDataResponse struct {
	WebsocketResponseData
	Failed []WebsocketApiUserDelMultiFailedResponse `json:"delFailed"`
}

type WebsocketApiUserDelMultiFailedResponse struct {
	UserId string `json:"user_id"` //
	Code   int    `json:"code"`    //0成功，其它失败
	Msg    string `json:"msg"`     //提示信息
}
