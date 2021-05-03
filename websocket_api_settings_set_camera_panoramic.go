package faceeasco

type WebsocketApiSettingsSetCameraPanoramicRequest struct {
	WebsocketRequestHeader
	Data WebsocketApiSettingsSetCameraPanoramicRequestData `json:"data"`
}

type WebsocketApiSettingsSetCameraPanoramicRequestData struct {
	WebsocketRequestData
	Value string `json:"value"`
}

type WebsocketApiSettingsSetCameraPanoramicResponse struct {
	WebsocketResponseHeader
	Data WebsocketResponseData `json:"data"`
}
