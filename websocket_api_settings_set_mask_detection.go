package faceeasco

type WebsocketApiSettingsSetMaskDetectionRequest struct {
	WebsocketRequestHeader
	Data WebsocketApiSettingsSetMaskDetectionRequestData `json:"data"`
}

type WebsocketApiSettingsSetMaskDetectionRequestData struct {
	WebsocketRequestData
	Value string `json:"value"`
}

type WebsocketApiSettingsSetMaskDetectionResponse struct {
	WebsocketResponseHeader
	Data WebsocketResponseData `json:"data"`
}
