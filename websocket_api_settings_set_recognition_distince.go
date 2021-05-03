package faceeasco

type WebsocketApiSettingsSetRecognitionDistanceRequest struct {
	WebsocketRequestHeader
	Data WebsocketApiSettingsSetRecognitionDistanceRequestData `json:"data"`
}

type WebsocketApiSettingsSetRecognitionDistanceRequestData struct {
	WebsocketRequestData
	Value float32 `json:"value"`
}

type WebsocketApiSettingsSetRecognitionDistanceResponse struct {
	WebsocketResponseHeader
	Data WebsocketResponseData `json:"data"`
}
