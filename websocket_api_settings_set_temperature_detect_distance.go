package faceeasco

type WebsocketApiSettingsSetTemperatureDetectDistanceRequest struct {
	WebsocketRequestHeader
	Data WebsocketApiSettingsSetTemperatureDetectDistanceRequestData `json:"data"`
}

type WebsocketApiSettingsSetTemperatureDetectDistanceRequestData struct {
	WebsocketRequestData
	Value int `json:"value"`
}

type WebsocketApiSettingsSetTemperatureDetectDistanceResponse struct {
	WebsocketResponseHeader
	Data WebsocketResponseData `json:"data"`
}
