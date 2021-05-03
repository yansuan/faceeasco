package faceeasco

type WebsocketApiSettingsSetRelayStateRequest struct {
	WebsocketRequestHeader
	Data WebsocketApiSettingsSetRelayStateRequestData `json:"data"`
}

type WebsocketApiSettingsSetRelayStateRequestData struct {
	WebsocketRequestData
	Value int `json:"value"`
}

type WebsocketApiSettingsSetRelayStateResponse struct {
	WebsocketResponseHeader
	Data WebsocketResponseData `json:"data"`
}
