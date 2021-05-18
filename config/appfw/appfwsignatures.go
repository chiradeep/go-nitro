package appfw

type Appfwsignatures struct {
	Comment            string `json:"comment,omitempty"`
	Merge              bool   `json:"merge,omitempty"`
	Mergedefault       bool   `json:"mergedefault,omitempty"`
	Name               string `json:"name,omitempty"`
	Overwrite          bool   `json:"overwrite,omitempty"`
	Preservedefactions bool   `json:"preservedefactions,omitempty"`
	Response           string `json:"response,omitempty"`
	Sha1               string `json:"sha1,omitempty"`
	Src                string `json:"src,omitempty"`
	Vendortype         string `json:"vendortype,omitempty"`
	Xslt               string `json:"xslt,omitempty"`
}
