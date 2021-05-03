package faceeasco

type WebsocketApiSettingsSetTemperatureLowPassRequest struct {
	WebsocketRequestHeader
	Data WebsocketApiSettingsSetTemperatureLowPassRequestData `json:"data"`
}

type WebsocketApiSettingsSetTemperatureLowPassRequestData struct {
	WebsocketRequestData
	Value string `json:"value"`
}

type WebsocketApiSettingsSetTemperatureLowPassResponse struct {
	WebsocketResponseHeader
	Data WebsocketResponseData `json:"data"`
}
