package faceeasco

type WebsocketApiSettingsSetVoiceRequest struct {
	WebsocketRequestHeader
	Data WebsocketApiSettingsSetVoiceRequestData `json:"data"`
}

type WebsocketApiSettingsSetVoiceRequestData struct {
	WebsocketRequestData
	Type      int `json:"type"`
	VoiceCode int `json:"voice_code"`
	VoiceText int `json:"voice_text"`
}

type WebsocketApiSettingsSetVoiceResponse struct {
	WebsocketResponseHeader
	Data WebsocketResponseData `json:"data"`
}
