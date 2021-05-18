package bot

type Botprofile struct {
	Botenableblacklist                     string      `json:"bot_enable_black_list,omitempty"`
	Botenableipreputation                  string      `json:"bot_enable_ip_reputation,omitempty"`
	Botenableratelimit                     string      `json:"bot_enable_rate_limit,omitempty"`
	Botenabletps                           string      `json:"bot_enable_tps,omitempty"`
	Botenablewhitelist                     string      `json:"bot_enable_white_list,omitempty"`
	Builtin                                interface{} `json:"builtin,omitempty"`
	Clientipexpression                     string      `json:"clientipexpression,omitempty"`
	Comment                                string      `json:"comment,omitempty"`
	Devicefingerprint                      string      `json:"devicefingerprint,omitempty"`
	Devicefingerprintaction                interface{} `json:"devicefingerprintaction,omitempty"`
	Devicefingerprintmobile                interface{} `json:"devicefingerprintmobile,omitempty"`
	Errorurl                               string      `json:"errorurl,omitempty"`
	Feature                                string      `json:"feature,omitempty"`
	Name                                   string      `json:"name,omitempty"`
	Signature                              string      `json:"signature,omitempty"`
	Signaturemultipleuseragentheaderaction interface{} `json:"signaturemultipleuseragentheaderaction,omitempty"`
	Signaturenouseragentheaderaction       interface{} `json:"signaturenouseragentheaderaction,omitempty"`
	Trap                                   string      `json:"trap,omitempty"`
	Trapaction                             interface{} `json:"trapaction,omitempty"`
	Trapurl                                string      `json:"trapurl,omitempty"`
}
