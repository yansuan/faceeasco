package faceeasco

type WebsocketApiSettingsSetTimeRequest struct {
	WebsocketRequestHeader
	Data WebsocketApiSettingsSetTimeRequestData `json:"data"`
}

type WebsocketApiSettingsSetTimeRequestData struct {
	WebsocketRequestData
	Value int `json:"value"`
}

type WebsocketApiSettingsSetTimeResponse struct {
	WebsocketResponseHeader
	Data WebsocketResponseData `json:"data"`
}
