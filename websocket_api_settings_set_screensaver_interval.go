package faceeasco

//设置屏保显示时长
type WebsocketApiSettingsSetScreensaverIntervalRequest struct {
	WebsocketRequestHeader
	Data WebsocketApiSettingsSetScreensaverIntervalRequestData `json:"data"`
}

type WebsocketApiSettingsSetScreensaverIntervalRequestData struct {
	WebsocketRequestData
	Value int `json:"value"`
}

type WebsocketApiSettingsSetScreensaverIntervalResponse struct {
	WebsocketResponseHeader
	Data WebsocketResponseData `json:"data"`
}
