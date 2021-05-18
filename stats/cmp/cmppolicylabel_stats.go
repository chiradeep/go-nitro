package cmp

type Cmppolicylabelstats struct {
	Clearstats            string      `json:"clearstats,omitempty"`
	Cmppolicy             interface{} `json:"cmppolicy,omitempty"`
	Labelname             string      `json:"labelname,omitempty"`
	Pipolicylabelhits     int         `json:"pipolicylabelhits,omitempty"`
	Pipolicylabelhitsrate int         `json:"pipolicylabelhitsrate,omitempty"`
}
