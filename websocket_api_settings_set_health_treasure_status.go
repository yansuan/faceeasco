package faceeasco

type WebsocketApiSettingsSetHealthTreasureStatusRequest struct {
	WebsocketRequestHeader
	Data WebsocketApiSettingsSetHealthTreasureStatusRequestData `json:"data"`
}

type WebsocketApiSettingsSetHealthTreasureStatusRequestData struct {
	WebsocketRequestData
	Value int `json:"value"`
}

type WebsocketApiSettingsSetHealthTreasureStatusResponse struct {
	WebsocketResponseHeader
	Data WebsocketResponseData `json:"data"`
}
