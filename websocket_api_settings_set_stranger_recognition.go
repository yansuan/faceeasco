package faceeasco

type WebsocketApiSettingsSetStrangerRecognitionRequest struct {
	WebsocketRequestHeader
	Data WebsocketApiSettingsSetStrangerRecognitionRequestData `json:"data"`
}

type WebsocketApiSettingsSetStrangerRecognitionRequestData struct {
	WebsocketRequestData
	Value string `json:"value"`
}

type WebsocketApiSettingsSetStrangerRecognitionResponse struct {
	WebsocketResponseHeader
	Data WebsocketResponseData `json:"data"`
}
