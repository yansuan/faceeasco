package faceeasco

type WebsocketApiSettingsSetHttpTokenRequest struct {
	WebsocketRequestHeader
	Data WebsocketApiSettingsSetHttpTokenRequestData `json:"data"`
}

type WebsocketApiSettingsSetHttpTokenRequestData struct {
	WebsocketRequestData
	TokenKey   string `json:"token_key"`
	TokenValue string `json:"token_value"`
}

type WebsocketApiSettingsSetHttpTokenQualityResponse struct {
	WebsocketResponseHeader
	Data WebsocketResponseData `json:"data"`
}
