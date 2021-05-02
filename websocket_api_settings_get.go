package faceeasco

type WebsocketApiSettingsGetRequest struct {
	WebsocketRequestHeader
	Data WebsocketApiSettingsGetRequestData `json:"data"`
}

type WebsocketApiSettingsGetRequestData struct {
	WebsocketRequestData
	Settings []string `json:"settings"`
}

type WebsocketApiSettingsGetResponse struct {
	WebsocketResponseHeader
}
