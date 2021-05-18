package appfw

type Appfwprofilecmdinjectionbinding struct {
	Alertonly         string `json:"alertonly,omitempty"`
	Asscanlocationcmd string `json:"as_scan_location_cmd,omitempty"`
	Asvalueexprcmd    string `json:"as_value_expr_cmd,omitempty"`
	Asvaluetypecmd    string `json:"as_value_type_cmd,omitempty"`
	Cmdinjection      string `json:"cmdinjection,omitempty"`
	Comment           string `json:"comment,omitempty"`
	Formactionurlcmd  string `json:"formactionurl_cmd,omitempty"`
	Isautodeployed    string `json:"isautodeployed,omitempty"`
	Isregexcmd        string `json:"isregex_cmd,omitempty"`
	Isvalueregexcmd   string `json:"isvalueregex_cmd,omitempty"`
	Name              string `json:"name,omitempty"`
	Resourceid        string `json:"resourceid,omitempty"`
	State             string `json:"state,omitempty"`
}
