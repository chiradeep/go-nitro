package bot

type Botprofilelogexpressionbinding struct {
	Botbindcomment          string `json:"bot_bind_comment,omitempty"`
	Botlogexpressionenabled string `json:"bot_log_expression_enabled,omitempty"`
	Botlogexpressionname    string `json:"bot_log_expression_name,omitempty"`
	Expression              string `json:"expression,omitempty"`
	Logexpression           bool   `json:"logexpression,omitempty"`
	Logmessage              string `json:"logmessage,omitempty"`
	Name                    string `json:"name,omitempty"`
}
