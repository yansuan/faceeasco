package faceeasco

type WebsocketApiSettingsSetDoorRequest struct {
	WebsocketRequestHeader
	Data WebsocketApiSettingsSetDoorRequestData `json:"data"`
}

type WebsocketApiSettingsSetDoorRequestData struct {
	WebsocketRequestData
	UserId string `json:"user_id"`
	Value  string `json:"value"`
}

type WebsocketApiSettingsSetDoorResponse struct {
	WebsocketResponseHeader
	Data struct {
		WebsocketResponseData
	}
}
