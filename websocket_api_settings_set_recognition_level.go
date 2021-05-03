package faceeasco

type WebsocketApiSettingsSetRecognitionLevelRequest struct {
	WebsocketRequestHeader
	Data WebsocketApiSettingsSetRecognitionLevelRequestData `json:"data"`
}

type WebsocketApiSettingsSetRecognitionLevelRequestData struct {
	WebsocketRequestData
	Value float32 `json:"value"`
}

type WebsocketApiSettingsSetRecognitionLevelResponse struct {
	WebsocketResponseHeader
	Data WebsocketResponseData `json:"data"`
}
