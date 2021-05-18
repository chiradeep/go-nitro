package gslb

type Gslbservicestats struct {
	Clearstats          string `json:"clearstats,omitempty"`
	Curclntconnections  int    `json:"curclntconnections,omitempty"`
	Curload             int    `json:"curload,omitempty"`
	Cursrvrconnections  int    `json:"cursrvrconnections,omitempty"`
	Establishedconn     int    `json:"establishedconn,omitempty"`
	Primaryipaddress    string `json:"primaryipaddress,omitempty"`
	Primaryport         int    `json:"primaryport,omitempty"`
	Requestbytesrate    int    `json:"requestbytesrate,omitempty"`
	Requestsrate        int    `json:"requestsrate,omitempty"`
	Responsebytesrate   int    `json:"responsebytesrate,omitempty"`
	Responsesrate       int    `json:"responsesrate,omitempty"`
	Servicename         string `json:"servicename,omitempty"`
	Servicetype         string `json:"servicetype,omitempty"`
	State               string `json:"state,omitempty"`
	Totalrequestbytes   int    `json:"totalrequestbytes,omitempty"`
	Totalrequests       int    `json:"totalrequests,omitempty"`
	Totalresponsebytes  int    `json:"totalresponsebytes,omitempty"`
	Totalresponses      int    `json:"totalresponses,omitempty"`
	Vsvrservicehits     int    `json:"vsvrservicehits,omitempty"`
	Vsvrservicehitsrate int    `json:"vsvrservicehitsrate,omitempty"`
}
