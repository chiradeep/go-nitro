package ha

type Hanodestats struct {
	Clearstats       string `json:"clearstats,omitempty"`
	Hacurmasterstate string `json:"hacurmasterstate,omitempty"`
	Hacurstate       string `json:"hacurstate,omitempty"`
	Hacurstatus      string `json:"hacurstatus,omitempty"`
	Haerrproptimeout int    `json:"haerrproptimeout,omitempty"`
	Haerrsyncfailure int    `json:"haerrsyncfailure,omitempty"`
	Hapktrxrate      int    `json:"hapktrxrate,omitempty"`
	Hapkttxrate      int    `json:"hapkttxrate,omitempty"`
	Hatotpktrx       int    `json:"hatotpktrx,omitempty"`
	Hatotpkttx       int    `json:"hatotpkttx,omitempty"`
	Transtime        string `json:"transtime,omitempty"`
}
