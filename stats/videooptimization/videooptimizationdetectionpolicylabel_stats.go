package videooptimization

type Videooptimizationdetectionpolicylabelstats struct {
	Clearstats                       string      `json:"clearstats,omitempty"`
	Labelname                        string      `json:"labelname,omitempty"`
	Pipolicylabelhits                int         `json:"pipolicylabelhits,omitempty"`
	Pipolicylabelhitsrate            int         `json:"pipolicylabelhitsrate,omitempty"`
	Videooptimizationdetectionpolicy interface{} `json:"videooptimizationdetectionpolicy,omitempty"`
}
