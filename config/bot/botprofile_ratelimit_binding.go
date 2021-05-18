package bot

type Botprofileratelimitbinding struct {
	Botbindcomment      string      `json:"bot_bind_comment,omitempty"`
	Botratelimit        bool        `json:"bot_ratelimit,omitempty"`
	Botratelimitaction  interface{} `json:"bot_rate_limit_action,omitempty"`
	Botratelimitenabled string      `json:"bot_rate_limit_enabled,omitempty"`
	Botratelimittype    string      `json:"bot_rate_limit_type,omitempty"`
	Botratelimiturl     string      `json:"bot_rate_limit_url,omitempty"`
	Cookiename          string      `json:"cookiename,omitempty"`
	Logmessage          string      `json:"logmessage,omitempty"`
	Name                string      `json:"name,omitempty"`
	Rate                int         `json:"rate,omitempty"`
	Timeslice           int         `json:"timeslice,omitempty"`
}
