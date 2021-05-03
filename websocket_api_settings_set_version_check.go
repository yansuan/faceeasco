package faceeasco

type WebsocketApiSettingsSetVersionCheckRequest struct {
	WebsocketRequestHeader
	Data WebsocketApiSettingsSetVersionCheckRequestData `json:"data"`
}

type WebsocketApiSettingsSetVersionCheckRequestData struct {
	WebsocketRequestData
}

type WebsocketApiSettingsSetVersionCheckResponse struct {
	WebsocketResponseHeader
	Data WebsocketResponseData `json:"data"`
}
