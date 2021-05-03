package faceeasco

type WebsocketApiSettingsSetVisitorCallStatusRequest struct {
	WebsocketRequestHeader
	Data WebsocketApiSettingsSetVisitorCallStatusRequestData `json:"data"`
}

type WebsocketApiSettingsSetVisitorCallStatusRequestData struct {
	WebsocketRequestData
	Value string `json:"value"`
}

type WebsocketApiSettingsSetVisitorCallStatusResponse struct {
	WebsocketResponseHeader
	Data WebsocketResponseData `json:"data"`
}
