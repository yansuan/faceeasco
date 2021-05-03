package faceeasco

type WebsocketApiSettingsSetVolumeRequest struct {
	WebsocketRequestHeader
	Data WebsocketApiSettingsSetVolumeRequestData `json:"data"`
}

type WebsocketApiSettingsSetVolumeRequestData struct {
	WebsocketRequestData
	Value int `json:"value"`
}

type WebsocketApiSettingsSetVolumeResponse struct {
	WebsocketResponseHeader
	Data WebsocketResponseData `json:"data"`
}
