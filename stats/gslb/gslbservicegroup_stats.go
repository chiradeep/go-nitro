package gslb

type Gslbservicegroupstats struct {
	Clearstats             string      `json:"clearstats,omitempty"`
	Gslbservicegroupmember interface{} `json:"gslbservicegroupmember,omitempty"`
	Servicegroupname       string      `json:"servicegroupname,omitempty"`
	Servicetype            string      `json:"servicetype,omitempty"`
	State                  string      `json:"state,omitempty"`
}
