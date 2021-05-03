package faceeasco

type WebsocketApiSettingsSetCustomHomeLogoRequest struct {
	WebsocketRequestHeader
	Data WebsocketApiSettingsSetCustomHomeLogoRequestData `json:"data"`
}

type WebsocketApiSettingsSetCustomHomeLogoRequestData struct {
	WebsocketRequestData
	Type  int    `json:"type"`
	Text  string `json:"text"`
	Image struct {
		Type    int    `json:"type"`
		Width   int    `json:"width"`
		Height  int    `json:"height"`
		Content string `json:"content"`
	} `json:"image"`
}

type WebsocketApiSettingsSetCustomHomeLogoResponse struct {
	WebsocketResponseHeader
	Data WebsocketResponseData `json:"data"`
}
