package faceeasco

type HttpHead struct {
	SN    string `json:"sn"`    //
	Count int    `json:"Count"` //
}

type HttpResponse struct {
	Result  int    `json:"Result"`
	Content string `json:"Content"`
	Msg     string `json:"Msg"`
}
