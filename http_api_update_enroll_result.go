package faceeasco

type HttpUpdateEnrollResponse struct {
	HttpHeader
	Logs []HttpStrangerData `json:"logs"`
}

type HttpUpdateEnrollData struct {
	UserId string `json:"user_id"` //
	Code   int    `json:"code"`    //
	Msg    string `json:"msg"`     //
}
