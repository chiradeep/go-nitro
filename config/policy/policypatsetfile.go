package policy

type Policypatsetfile struct {
	Bindstatus     string `json:"bindstatus,omitempty"`
	Bindstatuscode int    `json:"bindstatuscode,omitempty"`
	Boundpatterns  int    `json:"boundpatterns,omitempty"`
	Charset        string `json:"charset,omitempty"`
	Comment        string `json:"comment,omitempty"`
	Delimiter      string `json:"delimiter,omitempty"`
	Imported       bool   `json:"imported,omitempty"`
	Name           string `json:"name,omitempty"`
	Overwrite      bool   `json:"overwrite,omitempty"`
	Patsetname     string `json:"patsetname,omitempty"`
	Src            string `json:"src,omitempty"`
	Totalpatterns  int    `json:"totalpatterns,omitempty"`
}
