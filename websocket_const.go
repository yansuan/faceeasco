package faceeasco

const (
	WEBSOCKET_API_DECLARE = "declare" //declare
	WEBSOCKET_API_PING    = "ping"    //心跳
	//
	WEBSOCKET_API_USER_ADD       = "addUser"      //新增用户
	WEBSOCKET_API_USER_EDIT      = "editUser"     //编辑用户
	WEBSOCKET_API_USER_DEL       = "delUser"      //删除用户
	WEBSOCKET_API_USER_DEL_MULTI = "delMultiUser" //批量删除人员
	WEBSOCKET_API_USER_DEL_ALL   = "delAllUser"   //删除全部人员
	//
	WEBSOCKET_API_VISIT_ADD = "addVisit" //下发人证比对基本信息

	//设备设置接口
	WEBSOCKET_API_SETTINGS_GET                              = "getDeviceSettings"               //获取设备设置
	WEBSOCKET_API_SETTINGS_SET_customHomeLogo               = "customHomeLogo"                  //设置设备 logo
	WEBSOCKET_API_SETTINGS_SET_Ad                           = "setAd"                           //设置广告语
	WEBSOCKET_API_SETTINGS_SET_ScreensaverStatus            = "setScreensaverStatus"            //设备节能屏保开关
	WEBSOCKET_API_SETTINGS_SET_Time                         = "setTime"                         //设置设备时间
	WEBSOCKET_API_SETTINGS_SET_Door                         = "setDoor"                         //控制设备开门/关门
	WEBSOCKET_API_SETTINGS_SET_PassType                     = "setPassType"                     //设置可通行人员类型
	WEBSOCKET_API_SETTINGS_SET_RecognitionLevel             = "setRecognitionLevel"             //设置设备识别置信度
	WEBSOCKET_API_SETTINGS_SET_RecognitionInterval          = "setRecognitionInterval"          //设置设备识别间隔
	WEBSOCKET_API_SETTINGS_SET_RecognitionDistance          = "setRecognitionDistance"          //设置设备识别距离
	WEBSOCKET_API_SETTINGS_SET_MaskDetection                = "setMaskDetection"                //口罩检测开关
	WEBSOCKET_API_SETTINGS_SET_LivenessDetection            = "setLivenessDetection"            //活体检测
	WEBSOCKET_API_SETTINGS_SET_LivenessLevel                = "setLivenessLevel"                //活体检测等级
	WEBSOCKET_API_SETTINGS_SET_StrangerRecognition          = "setStrangerRecognition"          //陌生人识别
	WEBSOCKET_API_SETTINGS_SET_Volume                       = "setVolume"                       //设备音量
	WEBSOCKET_API_SETTINGS_SET_PlayUserName                 = "setPlayUserName"                 //播放用户名
	WEBSOCKET_API_SETTINGS_SET_Switch                       = "setSwitch"                       //开启/关闭设备识别
	WEBSOCKET_API_SETTINGS_SET_Voice                        = "setVoice"                        //设置设备提示语音
	WEBSOCKET_API_SETTINGS_SET_Password                     = "setPassword"                     //修改进入设置页面密码
	WEBSOCKET_API_SETTINGS_SET_reboot                       = "reboot"                          //重启设备
	WEBSOCKET_API_SETTINGS_SET_RelayState                   = "setRelayState"                   //设置继电器状态-常开/正常
	WEBSOCKET_API_SETTINGS_SET_RecordPictureQuality         = "setRecordPictureQuality"         //设置设备识别照片抓拍的质量
	WEBSOCKET_API_SETTINGS_SET_VisitorCallStatus            = "setVisitorCallStatus"            //访客呼叫开关
	WEBSOCKET_API_SETTINGS_SET_PrinterStatus                = "setPrinterStatus"                //打印机开关
	WEBSOCKET_API_SETTINGS_SET_OnlineRecognitionInterval    = "setOnlineRecognitionInterval"    //设置在线识别间隔
	WEBSOCKET_API_SETTINGS_SET_VisitorApplyValue            = "setVisitorApplyValue"            //首页访客申请入口内容
	WEBSOCKET_API_SETTINGS_SET_HttpToken                    = "setHttpToken"                    //设置 http 请求的 token
	WEBSOCKET_API_SETTINGS_SET_ScreensaverState             = "setScreensaverState"             //设置屏保开关
	WEBSOCKET_API_SETTINGS_SET_setScreensaver               = "setScreensaver"                  //设置设备屏保
	WEBSOCKET_API_SETTINGS_SET_ScreensaverInterval          = "setScreensaverInterval"          //设置屏保显示时长
	WEBSOCKET_API_SETTINGS_SET_TemperatureDetection         = "setTemperatureDetection"         //体温检测开关
	WEBSOCKET_API_SETTINGS_SET_TemperatureMode              = "setTemperatureMode"              //体温检测模式
	WEBSOCKET_API_SETTINGS_SET_PreliminaryScreeningStandard = "setPreliminaryScreeningStandard" //初筛模式标准
	WEBSOCKET_API_SETTINGS_SET_TemperatureCalibrationMode   = "setTemperatureCalibrationMode"   //设置体温校准模式
	WEBSOCKET_API_SETTINGS_SET_LowTemperaturePass           = "setLowTemperaturePass"           //低温可通行开关
	WEBSOCKET_API_SETTINGS_SET_PlayTemperature              = "setPlayTemperature"              //播报体温开关
	WEBSOCKET_API_SETTINGS_SET_Reflectivity                 = "setReflectivity"                 //设置体温反射率
	WEBSOCKET_API_SETTINGS_SET_TemperatureCorrection        = "setTemperatureCorrection"        //设置体温校准值
	WEBSOCKET_API_SETTINGS_SET_TemperatureDetectDistance    = "setTemperatureDetectDistance"    //设置设备体温检测距离
	WEBSOCKET_API_SETTINGS_SET_AutoUpdate                   = "setAutoUpdate"                   //开机自动更新
	WEBSOCKET_API_SETTINGS_SET_getVersionInfo               = "getVersionInfo"                  //获取设备版本信息
	WEBSOCKET_API_SETTINGS_SET_pushVersionInfo              = "pushVersionInfo"                 //发送设备版本信息给服务端
	WEBSOCKET_API_SETTINGS_SET_checkForUpdates              = "checkForUpdates"                 //通知设备在线检查更新
	WEBSOCKET_API_SETTINGS_SET_pushUpdateRet                = "pushUpdateRet"                   //发送升级结果给服务端
	WEBSOCKET_API_SETTINGS_SET_pushUpdateProgress           = "pushUpdateProgress"              //更新下载进度
	//
	WEBSOCKET_API_SETTINGS_SET_HealthTreasureStatus = "setHealthTreasureStatus" //设置健康码开关
	WEBSOCKET_API_SETTINGS_SET_HealthTreasureKey    = "setHealthTreasureKey"    //设置健康宝密钥
	WEBSOCKET_API_SETTINGS_SET_HealthQrCodeValue    = "setQrCodeValue"          //设置设备首页二维码内容
	//camera
	WEBSOCKET_API_SETTINGS_SET_PanoramicCamera       = "setPanoramicCamera"        //全景相机开关
	WEBSOCKET_API_SETTINGS_SET_PanoramicCameraParams = "setPanoramicCameraParams " //设置全景相机的参数

)
