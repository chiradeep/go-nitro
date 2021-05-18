package network

type Tunnelipstats struct {
	Clearstats     string `json:"clearstats,omitempty"`
	Tnlrxbytesrate int    `json:"tnlrxbytesrate,omitempty"`
	Tnlrxpktsrate  int    `json:"tnlrxpktsrate,omitempty"`
	Tnltotrxbytes  int    `json:"tnltotrxbytes,omitempty"`
	Tnltotrxpkts   int    `json:"tnltotrxpkts,omitempty"`
	Tnltottxbytes  int    `json:"tnltottxbytes,omitempty"`
	Tnltottxpkts   int    `json:"tnltottxpkts,omitempty"`
	Tnltxbytesrate int    `json:"tnltxbytesrate,omitempty"`
	Tnltxpktsrate  int    `json:"tnltxpktsrate,omitempty"`
	Tunnelip       string `json:"tunnelip,omitempty"`
}
