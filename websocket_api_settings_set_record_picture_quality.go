package faceeasco

type WebsocketApiSettingsSetRecordPictureQualityRequest struct {
	WebsocketRequestHeader
	Data WebsocketApiSettingsSetRecordPictureQualityRequestData `json:"data"`
}

type WebsocketApiSettingsSetRecordPictureQualityRequestData struct {
	WebsocketRequestData
	Value int `json:"value"`
}

type WebsocketApiSettingsSetRecordPictureQualityResponse struct {
	WebsocketResponseHeader
	Data WebsocketResponseData `json:"data"`
}
