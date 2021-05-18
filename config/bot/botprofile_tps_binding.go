package bot

type Botprofiletpsbinding struct {
	Botbindcomment string      `json:"bot_bind_comment,omitempty"`
	Bottps         bool        `json:"bot_tps,omitempty"`
	Bottpsaction   interface{} `json:"bot_tps_action,omitempty"`
	Bottpstype     string      `json:"bot_tps_type,omitempty"`
	Logmessage     string      `json:"logmessage,omitempty"`
	Name           string      `json:"name,omitempty"`
	Percentage     int         `json:"percentage,omitempty"`
	Threshold      int         `json:"threshold,omitempty"`
}
