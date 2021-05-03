package faceeasco

//设置体温校准模式
type WebsocketApiSettingsSetTemperatureCalibrationModeRequest struct {
	WebsocketRequestHeader
	Data WebsocketApiSettingsSetTemperatureCalibrationModeRequestData `json:"data"`
}

type WebsocketApiSettingsSetTemperatureCalibrationModeRequestData struct {
	WebsocketRequestData
	Value int `json:"value"`
}

type WebsocketApiSettingsSetTemperatureCalibrationModeResponse struct {
	WebsocketResponseHeader
	Data WebsocketResponseData `json:"data"`
}
