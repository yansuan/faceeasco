package faceeasco

type WebsocketApiSettingsSetPasswordRequest struct {
	WebsocketRequestHeader
	Data WebsocketApiSettingsSetPasswordRequestData `json:"data"`
}

type WebsocketApiSettingsSetPasswordRequestData struct {
	WebsocketRequestData
	OldPasswordValue int `json:"old_password"`
	NewPassword      int `json:"new_password"`
}

type WebsocketApiSettingsSetPasswordResponse struct {
	WebsocketResponseHeader
	Data WebsocketResponseData `json:"data"`
}
