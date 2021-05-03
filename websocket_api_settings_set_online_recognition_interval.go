package faceeasco

type WebsocketApiSettingsSetOnlineRecognitionIntervalRequest struct {
	WebsocketRequestHeader
	Data WebsocketApiSettingsSetOnlineRecognitionIntervalRequestData `json:"data"`
}

type WebsocketApiSettingsSetOnlineRecognitionIntervalRequestData struct {
	WebsocketRequestData
	Value int `json:"value"`
}

type WebsocketApiSettingsSetOnlineRecognitionIntervalResponse struct {
	WebsocketResponseHeader
	Data WebsocketResponseData `json:"data"`
}
