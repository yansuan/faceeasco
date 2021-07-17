package faceeasco

type WebsocketApiUserGetRequest struct {
	WebsocketRequestHeader
	Data WebsocketApiUserAddRequestData `json:"data"`
}

type WebsocketApiUserGetRequestData struct {
	WebsocketRequestData
	Value int `json:"value"` //0: 读取人脸总数，1：读取所有人员用户的 user id，2 读取人证总数 ，3 读取所有人证用户的 user id
}

type WebsocketApiUserGetResponse struct {
	WebsocketResponse
	Data WebsocketApiUserGetResponseData `json:"data"`
}

type WebsocketApiUserGetResponseData struct {
	WebsocketResponseData
	UserIds []string `json:"userIds"`
	Total   int      `json:"total"`
}
