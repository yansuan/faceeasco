package faceeasco

type WebsocketApiSettingsSetSwitchRequest struct {
	WebsocketRequestHeader
	Data WebsocketApiSettingsSetSwitchRequestData `json:"data"`
}

type WebsocketApiSettingsSetSwitchRequestData struct {
	WebsocketRequestData
	Value string `json:"value"`
}

type WebsocketApiSettingsSetSwitchResponse struct {
	WebsocketResponseHeader
	Data WebsocketResponseData `json:"data"`
}
