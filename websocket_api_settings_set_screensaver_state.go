package faceeasco

//设置屏保开关
type WebsocketApiSettingsSetScreensaverStateRequest struct {
	WebsocketRequestHeader
	Data WebsocketApiSettingsSetScreensaverStateRequestData `json:"data"`
}

type WebsocketApiSettingsSetScreensaverStateRequestData struct {
	WebsocketRequestData
	Value string `json:"value"`
}

type WebsocketApiSettingsSetScreensaverStateResponse struct {
	WebsocketResponseHeader
	Data WebsocketResponseData `json:"data"`
}
