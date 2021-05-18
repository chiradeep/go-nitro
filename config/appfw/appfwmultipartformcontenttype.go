package appfw

type Appfwmultipartformcontenttype struct {
	Builtin                       interface{} `json:"builtin,omitempty"`
	Feature                       string      `json:"feature,omitempty"`
	Isregex                       string      `json:"isregex,omitempty"`
	Multipartformcontenttypevalue string      `json:"multipartformcontenttypevalue,omitempty"`
}
