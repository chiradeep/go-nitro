package basic

type Servicegroupstats struct {
	Clearstats         string      `json:"clearstats,omitempty"`
	Servicegroupmember interface{} `json:"servicegroupmember,omitempty"`
	Servicegroupname   string      `json:"servicegroupname,omitempty"`
	Servicetype        string      `json:"servicetype,omitempty"`
	State              string      `json:"state,omitempty"`
}
