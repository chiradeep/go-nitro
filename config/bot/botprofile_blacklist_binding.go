package bot

type Botprofileblacklistbinding struct {
	Botbindcomment      string      `json:"bot_bind_comment,omitempty"`
	Botblacklist        bool        `json:"bot_blacklist,omitempty"`
	Botblacklistaction  interface{} `json:"bot_blacklist_action,omitempty"`
	Botblacklistenabled string      `json:"bot_blacklist_enabled,omitempty"`
	Botblacklisttype    string      `json:"bot_blacklist_type,omitempty"`
	Botblacklistvalue   string      `json:"bot_blacklist_value,omitempty"`
	Logmessage          string      `json:"logmessage,omitempty"`
	Name                string      `json:"name,omitempty"`
}
