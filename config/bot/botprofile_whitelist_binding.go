package bot

type Botprofilewhitelistbinding struct {
	Botbindcomment      string `json:"bot_bind_comment,omitempty"`
	Botwhitelist        bool   `json:"bot_whitelist,omitempty"`
	Botwhitelistenabled string `json:"bot_whitelist_enabled,omitempty"`
	Botwhitelisttype    string `json:"bot_whitelist_type,omitempty"`
	Botwhitelistvalue   string `json:"bot_whitelist_value,omitempty"`
	Log                 string `json:"log,omitempty"`
	Logmessage          string `json:"logmessage,omitempty"`
	Name                string `json:"name,omitempty"`
}
