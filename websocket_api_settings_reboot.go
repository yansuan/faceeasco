package faceeasco

type WebsocketApiSettingsRebootRequest struct {
	WebsocketRequestHeader
	Data WebsocketApiSettingsRebootRequestData `json:"data"`
}

type WebsocketApiSettingsRebootRequestData struct {
	WebsocketRequestData
}

type WebsocketApiSettingsRebootResponse struct {
	WebsocketResponseHeader
	Data WebsocketResponseData `json:"data"`
}
