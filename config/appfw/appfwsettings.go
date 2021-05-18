package appfw

type Appfwsettings struct {
	Builtin                  interface{} `json:"builtin,omitempty"`
	Ceflogging               string      `json:"ceflogging,omitempty"`
	Centralizedlearning      string      `json:"centralizedlearning,omitempty"`
	Clientiploggingheader    string      `json:"clientiploggingheader,omitempty"`
	Cookiepostencryptprefix  string      `json:"cookiepostencryptprefix,omitempty"`
	Defaultprofile           string      `json:"defaultprofile,omitempty"`
	Entitydecoding           string      `json:"entitydecoding,omitempty"`
	Feature                  string      `json:"feature,omitempty"`
	Geolocationlogging       string      `json:"geolocationlogging,omitempty"`
	Importsizelimit          int         `json:"importsizelimit,omitempty"`
	Learning                 string      `json:"learning,omitempty"`
	Learnratelimit           int         `json:"learnratelimit,omitempty"`
	Logmalformedreq          string      `json:"logmalformedreq,omitempty"`
	Malformedreqaction       interface{} `json:"malformedreqaction,omitempty"`
	Sessioncookiename        string      `json:"sessioncookiename,omitempty"`
	Sessionlifetime          int         `json:"sessionlifetime,omitempty"`
	Sessionlimit             int         `json:"sessionlimit,omitempty"`
	Sessiontimeout           int         `json:"sessiontimeout,omitempty"`
	Signatureautoupdate      string      `json:"signatureautoupdate,omitempty"`
	Signatureurl             string      `json:"signatureurl,omitempty"`
	Undefaction              string      `json:"undefaction,omitempty"`
	Useconfigurablesecretkey string      `json:"useconfigurablesecretkey,omitempty"`
}
