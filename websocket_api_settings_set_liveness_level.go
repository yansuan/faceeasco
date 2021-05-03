package faceeasco

type WebsocketApiSettingsSetLivenessLevelRequest struct {
	WebsocketRequestHeader
	Data WebsocketApiSettingsSetLivenessLevelRequestData `json:"data"`
}

type WebsocketApiSettingsSetLivenessLevelRequestData struct {
	WebsocketRequestData
	Value int `json:"value"`
}

type WebsocketApiSettingsSetLivenessLevelResponse struct {
	WebsocketResponseHeader
	Data WebsocketResponseData `json:"data"`
}
