package faceeasco

type WebsocketApiVisitAddRequest struct {
	WebsocketRequestHeader
	Data WebsocketApiVisitAddRequestData `json:"data"`
}

type WebsocketApiVisitAddRequestData struct {
	WebsocketRequestData
	UserId     string `json:"user_id"`     //
	IdName     string `json:"id_name"`     //
	IdNumber   string `json:"id_number"`   //
	EffectTime string `json:"effect_time"` //
	Valid      string `json:"valid"`       //
}

type WebsocketApiVisitAddResponse struct {
	WebsocketResponseHeader
	Data WebsocketApiVisitAddResponseData `json:"data"`
}

type WebsocketApiVisitAddResponseData struct {
	WebsocketResponseData
	UserId string `json:"user_id"`
}
