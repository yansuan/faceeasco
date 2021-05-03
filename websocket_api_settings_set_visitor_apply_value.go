package faceeasco

type WebsocketApiSettingsSetVisitorApplyValueRequest struct {
	WebsocketRequestHeader
	Data WebsocketApiSettingsSetVisitorApplyValueRequestData `json:"data"`
}

type WebsocketApiSettingsSetVisitorApplyValueRequestData struct {
	WebsocketRequestData
	Url   string `json:"url"`
	Photo string `json:"photo"`
}

type WebsocketApiSettingsSetVisitorApplyValueQualityResponse struct {
	WebsocketResponseHeader
	Data WebsocketResponseData `json:"data"`
}
