package faceeasco

type WebsocketApiSettingsSetVersionGetInfoRequest struct {
	WebsocketRequestHeader
	Data WebsocketApiSettingsSetVersionGetInfoRequestData `json:"data"`
}

type WebsocketApiSettingsSetVersionGetInfoRequestData struct {
	WebsocketRequestData
}

type WebsocketApiSettingsSetVersionGetInfoResponse struct {
	WebsocketResponseHeader
	Data WebsocketResponseData `json:"data"`
}
