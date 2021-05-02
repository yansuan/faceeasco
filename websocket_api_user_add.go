package faceeasco

type WebsocketApiUserAdd struct {
}

type WebsocketApiUserAddRequest struct {
	WebsocketRequestHeader
	Data WebsocketApiUserAddRequestData `json:"data"`
}

type WebsocketApiUserAddRequestData struct {
	WebsocketRequestData
	UserId          string                          `json:"user_id"`
	Name            string                          `json:"name"`             //
	IdCard          string                          `json:"id_card"`          //
	FaceTemplate    string                          `json:"face_template"`    //
	VlFaceTemplate  string                          `json:"vlface_template"`  //
	Ic              string                          `json:"Ic"`               //
	Password        string                          `json:"password"`         //
	UserType        int                             `json:"user_type"`        //
	ConfidenceLevel int                             `json:"confidence_level"` //
	Phone           string                          `json:"phone"`            //
	EffectTime      string                          `json:"effect_time"`      //
	IdValid         string                          `json:"id_valid"`         //
	ValidCycle      []WebsocketApiUserAddValidCycle `json:"valid_cycle"`      //
	StartTime       string                          `json:"start_time"`       //
	EndTime         string                          `json:"end_time"`         //
	passRuleId      string                          `json:"pass_rule_id"`     //
	Mode            int                             `json:"mode"`             //
}

type WebsocketApiUserAddValidCycle struct {
	StartTime string `json:"start_time"` //
	EndTime   string `json:"end_time"`   //
}

type WebsocketApiUserAddResponse struct {
	WebsocketResponse
}

func (this *WebsocketApiUserAdd) Response(r []byte) interface{} {
	return nil
}
