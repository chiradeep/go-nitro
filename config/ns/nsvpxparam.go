package ns

type Nsvpxparam struct {
	Cloudproductcode    string `json:"cloudproductcode,omitempty"`
	Cpuyield            string `json:"cpuyield,omitempty"`
	Masterclockcpu1     string `json:"masterclockcpu1,omitempty"`
	Memorystatus        string `json:"memorystatus,omitempty"`
	Ownernode           int    `json:"ownernode,omitempty"`
	Technicalsupportpin string `json:"technicalsupportpin,omitempty"`
	Vpxenvironment      string `json:"vpxenvironment,omitempty"`
	Vpxoemcode          int    `json:"vpxoemcode,omitempty"`
}
