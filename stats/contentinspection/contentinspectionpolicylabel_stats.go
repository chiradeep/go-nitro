package contentinspection

type Contentinspectionpolicylabelstats struct {
	Clearstats              string      `json:"clearstats,omitempty"`
	Contentinspectionpolicy interface{} `json:"contentinspectionpolicy,omitempty"`
	Labelname               string      `json:"labelname,omitempty"`
	Pipolicylabelhits       int         `json:"pipolicylabelhits,omitempty"`
	Pipolicylabelhitsrate   int         `json:"pipolicylabelhitsrate,omitempty"`
}
