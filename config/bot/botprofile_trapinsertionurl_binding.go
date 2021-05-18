package bot

type Botprofiletrapinsertionurlbinding struct {
	Botbindcomment             string `json:"bot_bind_comment,omitempty"`
	Bottrapurl                 string `json:"bot_trap_url,omitempty"`
	Bottrapurlinsertionenabled string `json:"bot_trap_url_insertion_enabled,omitempty"`
	Logmessage                 string `json:"logmessage,omitempty"`
	Name                       string `json:"name,omitempty"`
	Trapinsertionurl           bool   `json:"trapinsertionurl,omitempty"`
}
