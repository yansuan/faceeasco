package faceeasco

type WebsocketApiSettingsSetTemperatureDetectionRequest struct {
	WebsocketRequestHeader
	Data WebsocketApiSettingsSetTemperatureDetectionRequestData `json:"data"`
}

type WebsocketApiSettingsSetTemperatureDetectionRequestData struct {
	WebsocketRequestData
	Value int `json:"value"`
}

type WebsocketApiSettingsSetTemperatureDetectionResponse struct {
	WebsocketResponseHeader
	Data WebsocketResponseData `json:"data"`
}
