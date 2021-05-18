package network

type Rnatstats struct {
	Clearstats      string `json:"clearstats,omitempty"`
	Rnatcursessions int    `json:"rnatcursessions,omitempty"`
	Rnatrxbytesrate int    `json:"rnatrxbytesrate,omitempty"`
	Rnatrxpktsrate  int    `json:"rnatrxpktsrate,omitempty"`
	Rnattotrxbytes  int    `json:"rnattotrxbytes,omitempty"`
	Rnattotrxpkts   int    `json:"rnattotrxpkts,omitempty"`
	Rnattottxbytes  int    `json:"rnattottxbytes,omitempty"`
	Rnattottxpkts   int    `json:"rnattottxpkts,omitempty"`
	Rnattottxsyn    int    `json:"rnattottxsyn,omitempty"`
	Rnattxbytesrate int    `json:"rnattxbytesrate,omitempty"`
	Rnattxpktsrate  int    `json:"rnattxpktsrate,omitempty"`
	Rnattxsynrate   int    `json:"rnattxsynrate,omitempty"`
}
