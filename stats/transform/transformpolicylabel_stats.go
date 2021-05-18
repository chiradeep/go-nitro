package transform

type Transformpolicylabelstats struct {
	Clearstats            string      `json:"clearstats,omitempty"`
	Labelname             string      `json:"labelname,omitempty"`
	Pipolicylabelhits     int         `json:"pipolicylabelhits,omitempty"`
	Pipolicylabelhitsrate int         `json:"pipolicylabelhitsrate,omitempty"`
	Transformpolicy       interface{} `json:"transformpolicy,omitempty"`
}
