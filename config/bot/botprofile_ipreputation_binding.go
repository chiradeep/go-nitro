package bot

type Botprofileipreputationbinding struct {
	Botbindcomment  string      `json:"bot_bind_comment,omitempty"`
	Botiprepaction  interface{} `json:"bot_iprep_action,omitempty"`
	Botiprepenabled string      `json:"bot_iprep_enabled,omitempty"`
	Botipreputation bool        `json:"bot_ipreputation,omitempty"`
	Category        string      `json:"category,omitempty"`
	Logmessage      string      `json:"logmessage,omitempty"`
	Name            string      `json:"name,omitempty"`
}
