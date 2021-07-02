package faceeasco

type WebsocketApiOnlineAuthorizationRequest struct {
	WebsocketRequestHeader
	Data WebsocketApiOnlineAuthorizationRequestData `json:"data"`
}

type WebsocketApiOnlineAuthorizationRequestData struct {
	WebsocketRequestData
}

type WebsocketApiOnlineAuthorizationResponse struct {
	WebsocketResponseHeader
	Data WebsocketApiOnlineAuthorizationResponseData `json:"data"`
}

type WebsocketApiOnlineAuthorizationResponseData struct {
	WebsocketResponseData
	IrFaceTemplate string `json:"ir_face_template"`
	VlFaceTemplate string `json:"vl_face_template"`
}
