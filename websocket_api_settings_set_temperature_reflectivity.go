package faceeasco

type WebsocketApiSettingsSetTemperatureReflectivityRequest struct {
	WebsocketRequestHeader
	Data WebsocketApiSettingsSetTemperatureReflectivityRequestData `json:"data"`
}

type WebsocketApiSettingsSetTemperatureReflectivityRequestData struct {
	WebsocketRequestData
	Value int `json:"value"`
}

type WebsocketApiSettingsSetTemperatureReflectivityResponse struct {
	WebsocketResponseHeader
	Data WebsocketResponseData `json:"data"`
}
