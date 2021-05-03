package faceeasco

type WebsocketApiSettingsSetVersionPushInfoRequest struct {
	WebsocketRequestHeader
	Data WebsocketApiSettingsSetVersionPushInfoRequestData `json:"data"`
}

type WebsocketApiSettingsSetVersionPushInfoRequestData struct {
	WebsocketRequestData
	ServerVersion struct {
		AppVersion   string `json:"app_version"`
		Declare      string `json:"declare"`
		RelationTime int64  `json:"relation_time"`
		Title        string `json:"title"`
	} `json:"server_version"`
	VersionCode  string `json:"version_code"`
	VersionName  string `json:"version_name"`
	VersionState int64  `json:"version_state"`
}

//type WebsocketApiSettingsSetVersionPushInfoResponse struct {
//	WebsocketResponseHeader
//	Data WebsocketResponseData `json:"data"`
//}
