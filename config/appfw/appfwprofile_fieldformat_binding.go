package appfw

type Appfwprofilefieldformatbinding struct {
	Alertonly            string `json:"alertonly,omitempty"`
	Comment              string `json:"comment,omitempty"`
	Fieldformat          string `json:"fieldformat,omitempty"`
	Fieldformatmaxlength int    `json:"fieldformatmaxlength,omitempty"`
	Fieldformatminlength int    `json:"fieldformatminlength,omitempty"`
	Fieldtype            string `json:"fieldtype,omitempty"`
	Formactionurlff      string `json:"formactionurl_ff,omitempty"`
	Isautodeployed       string `json:"isautodeployed,omitempty"`
	Isregexff            string `json:"isregex_ff,omitempty"`
	Name                 string `json:"name,omitempty"`
	Resourceid           string `json:"resourceid,omitempty"`
	State                string `json:"state,omitempty"`
}
