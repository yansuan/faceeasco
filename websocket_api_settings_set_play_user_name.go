package faceeasco

type WebsocketApiSettingsSetPlayUserNameRequest struct {
	WebsocketRequestHeader
	Data WebsocketApiSettingsSetPlayUserNameRequestData `json:"data"`
}

type WebsocketApiSettingsSetPlayUserNameRequestData struct {
	WebsocketRequestData
	Value string `json:"value"`
}

type WebsocketApiSettingsSetPlayUserNameResponse struct {
	WebsocketResponseHeader
	Data WebsocketResponseData `json:"data"`
}
