package network

type Mapbmrstats struct {
	Bmrtotv4rxpkts      int    `json:"bmrtotv4rxpkts,omitempty"`
	Bmrtotv4rxpktsicmp  int    `json:"bmrtotv4rxpktsicmp,omitempty"`
	Bmrtotv4rxpktstcp   int    `json:"bmrtotv4rxpktstcp,omitempty"`
	Bmrtotv4rxpktsudp   int    `json:"bmrtotv4rxpktsudp,omitempty"`
	Bmrtotv4txpkts      int    `json:"bmrtotv4txpkts,omitempty"`
	Bmrtotv4txpktsicmp  int    `json:"bmrtotv4txpktsicmp,omitempty"`
	Bmrtotv4txpktstcp   int    `json:"bmrtotv4txpktstcp,omitempty"`
	Bmrtotv4txpktsudp   int    `json:"bmrtotv4txpktsudp,omitempty"`
	Bmrtotv6rxpkts      int    `json:"bmrtotv6rxpkts,omitempty"`
	Bmrtotv6rxpktsicmp  int    `json:"bmrtotv6rxpktsicmp,omitempty"`
	Bmrtotv6rxpktstcp   int    `json:"bmrtotv6rxpktstcp,omitempty"`
	Bmrtotv6rxpktsudp   int    `json:"bmrtotv6rxpktsudp,omitempty"`
	Bmrtotv6txpkts      int    `json:"bmrtotv6txpkts,omitempty"`
	Bmrtotv6txpktsicmp  int    `json:"bmrtotv6txpktsicmp,omitempty"`
	Bmrtotv6txpktstcp   int    `json:"bmrtotv6txpktstcp,omitempty"`
	Bmrtotv6txpktsudp   int    `json:"bmrtotv6txpktsudp,omitempty"`
	Bmrv4rxpktsicmprate int    `json:"bmrv4rxpktsicmprate,omitempty"`
	Bmrv4rxpktsrate     int    `json:"bmrv4rxpktsrate,omitempty"`
	Bmrv4rxpktstcprate  int    `json:"bmrv4rxpktstcprate,omitempty"`
	Bmrv4rxpktsudprate  int    `json:"bmrv4rxpktsudprate,omitempty"`
	Bmrv4txpktsicmprate int    `json:"bmrv4txpktsicmprate,omitempty"`
	Bmrv4txpktsrate     int    `json:"bmrv4txpktsrate,omitempty"`
	Bmrv4txpktstcprate  int    `json:"bmrv4txpktstcprate,omitempty"`
	Bmrv4txpktsudprate  int    `json:"bmrv4txpktsudprate,omitempty"`
	Bmrv6rxpktsicmprate int    `json:"bmrv6rxpktsicmprate,omitempty"`
	Bmrv6rxpktsrate     int    `json:"bmrv6rxpktsrate,omitempty"`
	Bmrv6rxpktstcprate  int    `json:"bmrv6rxpktstcprate,omitempty"`
	Bmrv6rxpktsudprate  int    `json:"bmrv6rxpktsudprate,omitempty"`
	Bmrv6txpktsicmprate int    `json:"bmrv6txpktsicmprate,omitempty"`
	Bmrv6txpktsrate     int    `json:"bmrv6txpktsrate,omitempty"`
	Bmrv6txpktstcprate  int    `json:"bmrv6txpktstcprate,omitempty"`
	Bmrv6txpktsudprate  int    `json:"bmrv6txpktsudprate,omitempty"`
	Clearstats          string `json:"clearstats,omitempty"`
	Name                string `json:"name,omitempty"`
}
