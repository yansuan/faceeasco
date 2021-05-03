package faceeasco

type WebsocketApiSettingsSetHealthQrCodeValueRequest struct {
	WebsocketRequestHeader
	Data WebsocketApiSettingsSetHealthQrCodeValueRequestData `json:"data"`
}

type WebsocketApiSettingsSetHealthQrCodeValueRequestData struct {
	WebsocketRequestData
	Photo string `json:"photo"`
	Url   string `json:"url"`
}

type WebsocketApiSettingsSetHealthQrCodeValueResponse struct {
	WebsocketResponseHeader
	Data WebsocketResponseData `json:"data"`
}
