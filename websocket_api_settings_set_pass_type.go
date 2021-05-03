package faceeasco

type WebsocketApiSettingsSetPassTypeRequest struct {
	WebsocketRequestHeader
	Data WebsocketApiSettingsSetPassTypeRequestData `json:"data"`
}

type WebsocketApiSettingsSetPassTypeRequestData struct {
	WebsocketRequestData
	Value int `json:"value"`
}

type WebsocketApiSettingsSetPassTypeResponse struct {
	WebsocketResponseHeader
	Data WebsocketResponseData `json:"data"`
}
