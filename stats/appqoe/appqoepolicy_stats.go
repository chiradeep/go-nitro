package appqoe

type Appqoepolicystats struct {
	Clearstats           string `json:"clearstats,omitempty"`
	Cltrequestsrate      int    `json:"cltrequestsrate,omitempty"`
	Jsbytessentrate      int    `json:"jsbytessentrate,omitempty"`
	Jssentrate           int    `json:"jssentrate,omitempty"`
	Name                 string `json:"name,omitempty"`
	Pipolicyhits         int    `json:"pipolicyhits,omitempty"`
	Pipolicyhitsrate     int    `json:"pipolicyhitsrate,omitempty"`
	Qosavgclientttlb     int    `json:"qosavgclientttlb,omitempty"`
	Qosavgclientttlbrate int    `json:"qosavgclientttlbrate,omitempty"`
	Qosavgserverttfb     int    `json:"qosavgserverttfb,omitempty"`
	Qosavgserverttfbrate int    `json:"qosavgserverttfbrate,omitempty"`
	Qosavgserverttlb     int    `json:"qosavgserverttlb,omitempty"`
	Qosavgserverttlbrate int    `json:"qosavgserverttlbrate,omitempty"`
	Requestbytesrate     int    `json:"requestbytesrate,omitempty"`
	Requestsrate         int    `json:"requestsrate,omitempty"`
	Responsebytesrate    int    `json:"responsebytesrate,omitempty"`
	Responserate         int    `json:"responserate,omitempty"`
	Svrlinkedtorate      int    `json:"svrlinkedtorate,omitempty"`
	Throughputrate       int    `json:"throughputrate,omitempty"`
	Totcltrequests       int    `json:"totcltrequests,omitempty"`
	Totjsbytessent       int    `json:"totjsbytessent,omitempty"`
	Totjssent            int    `json:"totjssent,omitempty"`
	Totrequestbytes      int    `json:"totrequestbytes,omitempty"`
	Totrequests          int    `json:"totrequests,omitempty"`
	Totresponse          int    `json:"totresponse,omitempty"`
	Totresponsebytes     int    `json:"totresponsebytes,omitempty"`
	Totsvrlinkedto       int    `json:"totsvrlinkedto,omitempty"`
	Totthroughput        int    `json:"totthroughput,omitempty"`
}
