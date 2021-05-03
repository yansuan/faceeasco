package faceeasco

//设备节能屏保开关
type WebsocketApiSettingsSetScreensaverStatusRequest struct {
	WebsocketRequestHeader
	Data WebsocketApiSettingsSetScreensaverStatusRequestData `json:"data"`
}

type WebsocketApiSettingsSetScreensaverStatusRequestData struct {
	WebsocketRequestData
	Value int `json:"value"`
}

type WebsocketApiSettingsSetScreensaverStatusResponse struct {
	WebsocketResponseHeader
	Data WebsocketResponseData `json:"data"`
}
