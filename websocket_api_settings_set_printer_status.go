package faceeasco

type WebsocketApiSettingsSetPrinterStatusRequest struct {
	WebsocketRequestHeader
	Data WebsocketApiSettingsSetPrinterStatusRequestData `json:"data"`
}

type WebsocketApiSettingsSetPrinterStatusRequestData struct {
	WebsocketRequestData
	Value int `json:"value"`
}

type WebsocketApiSettingsSetPrinterStatusResponse struct {
	WebsocketResponseHeader
	Data WebsocketResponseData `json:"data"`
}
