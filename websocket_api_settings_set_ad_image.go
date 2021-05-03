package faceeasco

type WebsocketApiSettingsSetAdImageRequest struct {
	WebsocketRequestHeader
	Data WebsocketApiSettingsSetAdImageRequestData `json:"data"`
}

type WebsocketApiSettingsSetAdImageRequestData struct {
	WebsocketRequestData
	Id       string `json:"id"`
	Duration int    `json:"duration"`
	Priority int    `json:"priority"`
}

type WebsocketApiSettingsSetAdImageResponse struct {
	WebsocketResponseHeader
	Data WebsocketResponseData `json:"data"`
}
