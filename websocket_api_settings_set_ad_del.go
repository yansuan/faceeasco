package faceeasco

type WebsocketApiSettingsSetDelAdRequest struct {
	WebsocketRequestHeader
	Data WebsocketApiSettingsSetDelAdRequestData `json:"data"`
}

type WebsocketApiSettingsSetDelAdRequestData struct {
	WebsocketRequestData
	IsDelAll int      `json:"isDelAll"`
	Id       []string `json:"id"`
}

type WebsocketApiSettingsSetDelAdResponse struct {
	WebsocketResponseHeader
	Data WebsocketResponseData `json:"data"`
}
