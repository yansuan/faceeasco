package faceeasco

import (
	"encoding/json"
	"net/http"
)

type HttpUpdateEnrollResponse struct {
	HttpHeader
	Logs []HttpStrangerData `json:"logs"`
}

type HttpUpdateEnrollData struct {
	UserId string `json:"user_id"` //
	Code   int    `json:"code"`    //
	Msg    string `json:"msg"`     //
}

func ParseUpdateEnroll(r *http.Request) (resp *HttpUpdateEnrollResponse, err error) {
	err = r.ParseForm()
	if err != nil {
		return
	}

	resp = &HttpUpdateEnrollResponse{}
	err = json.NewDecoder(r.Body).Decode(resp)
	if err != nil {
		return
	}
	return
}
