package faceeasco

type WebsocketApiSettingsSetRecognitionIntervalRequest struct {
	WebsocketRequestHeader
	Data WebsocketApiSettingsSetRecognitionIntervalRequestData `json:"data"`
}

type WebsocketApiSettingsSetRecognitionIntervalRequestData struct {
	WebsocketRequestData
	Value int `json:"value"`
}

type WebsocketApiSettingsSetRecognitionIntervalResponse struct {
	WebsocketResponseHeader
	Data WebsocketResponseData `json:"data"`
}
