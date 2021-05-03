package faceeasco

type WebsocketApiSettingsSetHealthTreasureKeyRequest struct {
	WebsocketRequestHeader
	Data WebsocketApiSettingsSetHealthTreasureKeyRequestData `json:"data"`
}

type WebsocketApiSettingsSetHealthTreasureKeyRequestData struct {
	WebsocketRequestData
	Value string `json:"value"`
}

type WebsocketApiSettingsSetHealthTreasureKeyResponse struct {
	WebsocketResponseHeader
	Data WebsocketResponseData `json:"data"`
}
