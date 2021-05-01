package faceeasco

import (
	"encoding/json"
	"net/http"
)

type HttpStrangerResponse struct {
	HttpHeader
	Logs []HttpStrangerData `json:"logs"`
}

type HttpStrangerData struct {
	UserId          string `json:"user_id"`    //
	RecogType       string `json:"recog_type"` //
	RecogTime       string `json:"recog_time"` //
	Photo           string `json:"photo"`      //
	Gender          int8   `json:"gender"`     //
	UserName        string `json:"user_name,omitempty"`
	BodyTemperature string `json:"body_temperature,omitempty"`
	Confidence      string `json:"confidence,omitempty"`
	Reflectivity    int    `json:"reflectivity,omitempty"`
	RoomTemperature string `json:"room_temperature,omitempty"`
	//HealthTreasureCode string `json:"health_treasure_code,omitempty"`
	Extra             string `json:"extra,omitempty"`
	CardNumber        string `json:"card_number,omitempty"`
	PanoramicPicture  string `json:"panoramic_picture,omitempty"`
	HealthCodeColor   string `json:"health_code_color,omitempty"`
	HealthCodePicture string `json:"health_code_picture,omitempty"`
}

func ParseStranger(r *http.Request) (resp *HttpStrangerResponse, err error) {
	err = r.ParseForm()
	if err != nil {
		return
	}

	resp = &HttpStrangerResponse{}
	err = json.NewDecoder(r.Body).Decode(resp)
	if err != nil {
		return
	}
	return
}
