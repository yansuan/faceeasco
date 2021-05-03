package faceeasco

type WebsocketApiSettingsSetTemperaturePlayRequest struct {
	WebsocketRequestHeader
	Data WebsocketApiSettingsSetTemperaturePlayRequestData `json:"data"`
}

type WebsocketApiSettingsSetTemperaturePlayRequestData struct {
	WebsocketRequestData
	Value string `json:"value"`
}

type WebsocketApiSettingsSetTemperaturePlayResponse struct {
	WebsocketResponseHeader
	Data WebsocketResponseData `json:"data"`
}
