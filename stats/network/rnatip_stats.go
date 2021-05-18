package network

type Rnatipstats struct {
	Clearstats        string `json:"clearstats,omitempty"`
	Iprnatcursessions int    `json:"iprnatcursessions,omitempty"`
	Iprnatrxbytesrate int    `json:"iprnatrxbytesrate,omitempty"`
	Iprnatrxpktsrate  int    `json:"iprnatrxpktsrate,omitempty"`
	Iprnattotrxbytes  int    `json:"iprnattotrxbytes,omitempty"`
	Iprnattotrxpkts   int    `json:"iprnattotrxpkts,omitempty"`
	Iprnattottxbytes  int    `json:"iprnattottxbytes,omitempty"`
	Iprnattottxpkts   int    `json:"iprnattottxpkts,omitempty"`
	Iprnattottxsyn    int    `json:"iprnattottxsyn,omitempty"`
	Iprnattxbytesrate int    `json:"iprnattxbytesrate,omitempty"`
	Iprnattxpktsrate  int    `json:"iprnattxpktsrate,omitempty"`
	Iprnattxsynrate   int    `json:"iprnattxsynrate,omitempty"`
	Iptd              int    `json:"iptd,omitempty"`
	Rnatip            string `json:"Rnatip,omitempty"`
}
