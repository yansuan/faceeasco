package faceeasco

type WebsocketApiSettingsGetRequest struct {
	WebsocketRequestHeader
	Data WebsocketApiSettingsGetRequestData `json:"data"`
}

type WebsocketApiSettingsGetRequestData struct {
	WebsocketRequestData
	Settings []string `json:"settings"`
}

type WebsocketApiSettingsGetResponse struct {
	WebsocketResponseHeader
	Data WebsocketApiSettingsGetResponseData `json:"data"`
}

type WebsocketApiSettingsGetResponseData struct {
	WebsocketResponseData
	Settings WebsocketApiSettingsGetResponseDataSettings `json:"settings"`
}

type WebsocketApiSettingsGetResponseDataSettings struct {
	Advertising    string `json:"advertising"`
	CustomHomeLogo struct {
		Image struct {
			Content string `json:"content"`
			Height  int  `json:"height"`
			Type    int  `json:"type"`
			Width   int  `json:"width"`
		} `json:"image"`
		Text string `json:"text"`
		Type int64  `json:"type"`
	} `json:"custom_home_logo"`
	Disable struct {
		Offtime string `json:"offtime"`
		Value   string `json:"value"`
	} `json:"disable"`
	HealthTreasureKey    string `json:"health_treasure_key"`
	HealthTreasureStatus int  `json:"health_treasure_status"`
	LivenessDetection    string `json:"liveness_detection"`
	LivenessLevel        int  `json:"liveness_level"`
	LowTemperaturePass   string `json:"low_temperature_pass"`
	MaskDetection        string `json:"mask_detection"`
	NonLivingVoice       struct {
		VoiceCode int  `json:"voice_code"`
		VoiceText string `json:"voice_text"`
	} `json:"non_living_voice"`
	OnlineRecognitionInterval int  `json:"online_recognition_interval"`
	PassType                  int  `json:"pass_type"`
	Password                  string `json:"password"`
	PlayTemperature           string `json:"play_temperature"`
	PlayUsername              string `json:"play_username"`
	PreliminaryScreeningMode  int  `json:"preliminary_screening_mode"`
	RecognitionDistance       int  `json:"recognition_distance"`
	RecognitionInterval       int  `json:"recognition_interval"`
	RecognitionLevel          int  `json:"recognition_level"`
	RecognizeVoice            struct {
		VoiceCode int64  `json:"voice_code"`
		VoiceText string `json:"voice_text"`
	} `json:"recognize_voice"`
	RecordPictureQuality int  `json:"record_picture_quality"`
	Reflectivity         int  `json:"reflectivity"`
	RelayState           int  `json:"relay_state"`
	ScreensaverStatus    int  `json:"screensaver_status"`
	StrangerRecognition  string `json:"stranger_recognition"`
	StrangerVoice        struct {
		VoiceCode int64  `json:"voice_code"`
		VoiceText string `json:"voice_text"`
	} `json:"stranger_voice"`
	TemperatureCalibrationMode int   `json:"temperature_calibration_mode"`
	TemperatureCorrection      float64 `json:"temperature_correction"`
	TemperatureDetection       int   `json:"temperature_detection"`
	TemperatureMode            int   `json:"temperature_mode"`
	Time                       int64   `json:"time"`
	Volume                     string  `json:"volume"`
}

//type WebsocketApiSettingsGetResponseDataSettingsCustomHomeLogo struct {
//	Text  string `json:"text"`
//	Type  int    `json:"type"`
//	Image struct {
//		Content string `json:"content"`
//		Height  int    `json:"height"`
//		Type    int    `json:"type"`
//		Width   int    `json:"width"`
//	} `json:"image"`
//}
