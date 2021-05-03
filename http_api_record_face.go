package faceeasco

type HttpRecordFaceResponse struct {
	HttpHeader
	Logs []HttpRecordFaceData `json:"logs"`
}

type HttpRecordFaceData struct {
	UserId             string `json:"user_id"`    //
	RecogType          string `json:"recog_type"` //
	RecogTime          string `json:"recog_time"` //
	Photo              string `json:"photo"`      //
	Gender             int8   `json:"gender"`     //
	UserName           string `json:"user_name,omitempty"`
	BodyTemperature    string `json:"body_temperature,omitempty"`
	Confidence         string `json:"confidence,omitempty"`
	Reflectivity       int    `json:"reflectivity,omitempty"`
	RoomTemperature    string `json:"room_temperature,omitempty"`
	HealthTreasureCode string `json:"health_treasure_code,omitempty"`
	Extra              string `json:"extra,omitempty"`
	CardNumber         string `json:"card_number,omitempty"`
	PanoramicPicture   string `json:"panoramic_picture,omitempty"`
	HealthCodeColor    string `json:"health_code_color,omitempty"`
	HealthCodePicture  string `json:"health_code_picture,omitempty"`
}
