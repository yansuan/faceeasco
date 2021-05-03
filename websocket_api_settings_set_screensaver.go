package faceeasco

//设置设备屏保
type WebsocketApiSettingsSetScreensaverRequest struct {
	WebsocketRequestHeader
	Data WebsocketApiSettingsSetScreensaverRequestData `json:"data"`
}

type WebsocketApiSettingsSetScreensaverRequestData struct {
	WebsocketRequestData
	Value string `json:"value"`
}

type WebsocketApiSettingsSetScreensaverResponse struct {
	WebsocketResponseHeader
	Data WebsocketResponseData `json:"data"`
}
