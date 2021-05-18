package network

type Vxlanstats struct {
	Clearstats       string `json:"clearstats,omitempty"`
	Id               int    `json:"id,omitempty"`
	Vxlanrxbytesrate int    `json:"vxlanrxbytesrate,omitempty"`
	Vxlanrxpktsrate  int    `json:"vxlanrxpktsrate,omitempty"`
	Vxlantotrxbytes  int    `json:"vxlantotrxbytes,omitempty"`
	Vxlantotrxpkts   int    `json:"vxlantotrxpkts,omitempty"`
	Vxlantottxbytes  int    `json:"vxlantottxbytes,omitempty"`
	Vxlantottxpkts   int    `json:"vxlantottxpkts,omitempty"`
	Vxlantxbytesrate int    `json:"vxlantxbytesrate,omitempty"`
	Vxlantxpktsrate  int    `json:"vxlantxpktsrate,omitempty"`
}
