package faceeasco

//体温检测开关
type WebsocketApiSettingsSetTemperaturePreliminaryScreeningStandardRequest struct {
	WebsocketRequestHeader
	Data WebsocketApiSettingsSetTemperaturePreliminaryScreeningStandardRequestData `json:"data"`
}

type WebsocketApiSettingsSetTemperaturePreliminaryScreeningStandardRequestData struct {
	WebsocketRequestData
	Value int `json:"value"`
}

type WebsocketApiSettingsSetTemperaturePreliminaryScreeningStandardResponse struct {
	WebsocketResponseHeader
	Data WebsocketResponseData `json:"data"`
}
