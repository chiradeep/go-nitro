package network

type Rnat6stats struct {
	Clearstats       string `json:"clearstats,omitempty"`
	Rnat6cursessions int    `json:"rnat6cursessions,omitempty"`
	Rnat6rxbytesrate int    `json:"rnat6rxbytesrate,omitempty"`
	Rnat6rxpktsrate  int    `json:"rnat6rxpktsrate,omitempty"`
	Rnat6totrxbytes  int    `json:"rnat6totrxbytes,omitempty"`
	Rnat6totrxpkts   int    `json:"rnat6totrxpkts,omitempty"`
	Rnat6tottxbytes  int    `json:"rnat6tottxbytes,omitempty"`
	Rnat6tottxpkts   int    `json:"rnat6tottxpkts,omitempty"`
	Rnat6tottxsyn    int    `json:"rnat6tottxsyn,omitempty"`
	Rnat6txbytesrate int    `json:"rnat6txbytesrate,omitempty"`
	Rnat6txpktsrate  int    `json:"rnat6txpktsrate,omitempty"`
	Rnat6txsynrate   int    `json:"rnat6txsynrate,omitempty"`
}
