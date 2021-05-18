package stream

type Streamidentifierstats struct {
	Clearstats             string      `json:"clearstats,omitempty"`
	Name                   string      `json:"name,omitempty"`
	Pattern                interface{} `json:"pattern,omitempty"`
	Sortby                 string      `json:"sortby,omitempty"`
	Sortorder              string      `json:"sortorder,omitempty"`
	Streamobjbandw         int         `json:"streamobjbandw,omitempty"`
	Streamobjbreachcnt     int         `json:"streamobjbreachcnt,omitempty"`
	Streamobjconn          int         `json:"streamobjconn,omitempty"`
	Streamobjdroppedconns  int         `json:"streamobjdroppedconns,omitempty"`
	Streamobjpktcredits    int         `json:"streamobjpktcredits,omitempty"`
	Streamobjpktspersecond int         `json:"streamobjpktspersecond,omitempty"`
	Streamobjreq           int         `json:"streamobjreq,omitempty"`
	Streamobjresptime      int         `json:"streamobjresptime,omitempty"`
}
