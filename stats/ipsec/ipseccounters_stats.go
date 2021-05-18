package ipsec

type Ipseccountersstats struct {
	Clearstats       string `json:"clearstats,omitempty"`
	Ipsecrxbytesrate int    `json:"ipsecrxbytesrate,omitempty"`
	Ipsecrxpktsrate  int    `json:"ipsecrxpktsrate,omitempty"`
	Ipsectotrxbytes  int    `json:"ipsectotrxbytes,omitempty"`
	Ipsectotrxpkts   int    `json:"ipsectotrxpkts,omitempty"`
	Ipsectottxbytes  int    `json:"ipsectottxbytes,omitempty"`
	Ipsectottxpkts   int    `json:"ipsectottxpkts,omitempty"`
	Ipsectxbytesrate int    `json:"ipsectxbytesrate,omitempty"`
	Ipsectxpktsrate  int    `json:"ipsectxpktsrate,omitempty"`
}
