package faceeasco

type WebsocketApiSettingsSetTemperatureCorrectionRequest struct {
	WebsocketRequestHeader
	Data WebsocketApiSettingsSetTemperatureCorrectionRequestData `json:"data"`
}

type WebsocketApiSettingsSetTemperatureCorrectionRequestData struct {
	WebsocketRequestData
	Value float32 `json:"value"`
}

type WebsocketApiSettingsSetTemperatureCorrectionResponse struct {
	WebsocketResponseHeader
	Data WebsocketResponseData `json:"data"`
}
