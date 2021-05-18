package bot

type Botsettings struct {
	Builtin             interface{} `json:"builtin,omitempty"`
	Defaultprofile      string      `json:"defaultprofile,omitempty"`
	Dfprequestlimit     int         `json:"dfprequestlimit,omitempty"`
	Feature             string      `json:"feature,omitempty"`
	Javascriptname      string      `json:"javascriptname,omitempty"`
	Proxyport           int         `json:"proxyport,omitempty"`
	Proxyserver         string      `json:"proxyserver,omitempty"`
	Sessioncookiename   string      `json:"sessioncookiename,omitempty"`
	Sessiontimeout      int         `json:"sessiontimeout,omitempty"`
	Signatureautoupdate string      `json:"signatureautoupdate,omitempty"`
	Signatureurl        string      `json:"signatureurl,omitempty"`
	Trapurlautogenerate string      `json:"trapurlautogenerate,omitempty"`
	Trapurlinterval     int         `json:"trapurlinterval,omitempty"`
	Trapurllength       int         `json:"trapurllength,omitempty"`
}
