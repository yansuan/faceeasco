package faceeasco

import (
	"encoding/json"
	"net/http"
)

type HttpHeader struct {
	SN    string `json:"sn"`    //
	Count int    `json:"Count"` //
}

type HttpResponse struct {
	Result  int    `json:"Result"`
	Content string `json:"Content"`
	Msg     string `json:"Msg"`
}

func ParseRequest(r *http.Request, resp interface{}) (err error) {
	err = r.ParseForm()
	if err != nil {
		return
	}

	err = json.NewDecoder(r.Body).Decode(resp)
	if err != nil {
		return
	}
	return
}
