package bot

type Botpolicylabelstats struct {
	Botpolicy             interface{} `json:"botpolicy,omitempty"`
	Clearstats            string      `json:"clearstats,omitempty"`
	Labelname             string      `json:"labelname,omitempty"`
	Pipolicylabelhits     int         `json:"pipolicylabelhits,omitempty"`
	Pipolicylabelhitsrate int         `json:"pipolicylabelhitsrate,omitempty"`
}
