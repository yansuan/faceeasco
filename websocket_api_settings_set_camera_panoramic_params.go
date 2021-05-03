package faceeasco

type WebsocketApiSettingsSetCameraPanoramicParamsRequest struct {
	WebsocketRequestHeader
	Data WebsocketApiSettingsSetCameraPanoramicParamsRequestData `json:"data"`
}

type WebsocketApiSettingsSetCameraPanoramicParamsRequestData struct {
	WebsocketRequestData
	IP       string `json:"ip"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Channel  int    `json:"channel"`
}

type WebsocketApiSettingsSetCameraPanoramicParamsResponse struct {
	WebsocketResponseHeader
	Data WebsocketResponseData `json:"data"`
}
