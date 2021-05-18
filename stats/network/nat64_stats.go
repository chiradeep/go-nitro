package network

type Nat64stats struct {
	Clearstats            string `json:"clearstats,omitempty"`
	Nat64icmpsessionsrate int    `json:"nat64icmpsessionsrate,omitempty"`
	Nat64sessionsrate     int    `json:"nat64sessionsrate,omitempty"`
	Nat64tcpsessionsrate  int    `json:"nat64tcpsessionsrate,omitempty"`
	Nat64toticmpsessions  int    `json:"nat64toticmpsessions,omitempty"`
	Nat64totsessions      int    `json:"nat64totsessions,omitempty"`
	Nat64tottcpsessions   int    `json:"nat64tottcpsessions,omitempty"`
	Nat64totudpsessions   int    `json:"nat64totudpsessions,omitempty"`
	Nat64udpsessionsrate  int    `json:"nat64udpsessionsrate,omitempty"`
}
