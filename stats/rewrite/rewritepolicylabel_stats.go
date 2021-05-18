package rewrite

type Rewritepolicylabelstats struct {
	Clearstats            string      `json:"clearstats,omitempty"`
	Labelname             string      `json:"labelname,omitempty"`
	Pipolicylabelhits     int         `json:"pipolicylabelhits,omitempty"`
	Pipolicylabelhitsrate int         `json:"pipolicylabelhitsrate,omitempty"`
	Rewritepolicy         interface{} `json:"rewritepolicy,omitempty"`
}
