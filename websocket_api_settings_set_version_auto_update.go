package faceeasco

type WebsocketApiSettingsSetVersionAutoUpdateRequest struct {
	WebsocketRequestHeader
	Data WebsocketApiSettingsSetVersionAutoUpdateRequestData `json:"data"`
}

type WebsocketApiSettingsSetVersionAutoUpdateRequestData struct {
	WebsocketRequestData
	Value string `json:"value"`
}

type WebsocketApiSettingsSetVersionAutoUpdateResponse struct {
	WebsocketResponseHeader
	Data WebsocketResponseData `json:"data"`
}
