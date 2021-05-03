package faceeasco

//体温检测开关
type WebsocketApiSettingsSetTemperatureModeRequest struct {
	WebsocketRequestHeader
	Data WebsocketApiSettingsSetTemperatureModeRequestData `json:"data"`
}

type WebsocketApiSettingsSetTemperatureModeRequestData struct {
	WebsocketRequestData
	Value int `json:"value"`
}

type WebsocketApiSettingsSetTemperatureModeResponse struct {
	WebsocketResponseHeader
	Data WebsocketResponseData `json:"data"`
}
