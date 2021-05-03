package faceeasco

type WebsocketApiSettingsSetLivenessDetectionRequest struct {
	WebsocketRequestHeader
	Data WebsocketApiSettingsSetLivenessDetectionRequestData `json:"data"`
}

type WebsocketApiSettingsSetLivenessDetectionRequestData struct {
	WebsocketRequestData
	Value string `json:"value"`
}

type WebsocketApiSettingsSetLivenessDetectionResponse struct {
	WebsocketResponseHeader
	Data WebsocketResponseData `json:"data"`
}
