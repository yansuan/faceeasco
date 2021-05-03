package faceeasco

type WebsocketApiSettingsSetAdRequest struct {
	WebsocketRequestHeader
	Data WebsocketApiSettingsSetAdRequestData `json:"data"`
}

type WebsocketApiSettingsSetAdRequestData struct {
	WebsocketRequestData
	Advertising string `json:"advertising"`
}

type WebsocketApiSettingsSetAdResponse struct {
	WebsocketResponseHeader
	Data WebsocketResponseData `json:"data"`
}
