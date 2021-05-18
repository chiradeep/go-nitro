package bot

type Botprofilecaptchabinding struct {
	Botbindcomment    string      `json:"bot_bind_comment,omitempty"`
	Botcaptchaaction  interface{} `json:"bot_captcha_action,omitempty"`
	Botcaptchaenabled string      `json:"bot_captcha_enabled,omitempty"`
	Botcaptchaurl     string      `json:"bot_captcha_url,omitempty"`
	Captcharesource   bool        `json:"captcharesource,omitempty"`
	Graceperiod       int         `json:"graceperiod,omitempty"`
	Logmessage        string      `json:"logmessage,omitempty"`
	Muteperiod        int         `json:"muteperiod,omitempty"`
	Name              string      `json:"name,omitempty"`
	Requestsizelimit  int         `json:"requestsizelimit,omitempty"`
	Retryattempts     int         `json:"retryattempts,omitempty"`
	Waittime          int         `json:"waittime,omitempty"`
}
