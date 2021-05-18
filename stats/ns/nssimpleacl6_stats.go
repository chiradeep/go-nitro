package ns

type Nssimpleacl6stats struct {
	Clearstats           string `json:"clearstats,omitempty"`
	Sacl6hitsrate        int    `json:"sacl6hitsrate,omitempty"`
	Sacl6missesrate      int    `json:"sacl6missesrate,omitempty"`
	Sacl6pktsallowedrate int    `json:"sacl6pktsallowedrate,omitempty"`
	Sacl6pktsbridgedrate int    `json:"sacl6pktsbridgedrate,omitempty"`
	Sacl6pktsdeniedrate  int    `json:"sacl6pktsdeniedrate,omitempty"`
	Sacl6scount          int    `json:"sacl6scount,omitempty"`
	Sacl6tothits         int    `json:"sacl6tothits,omitempty"`
	Sacl6totmisses       int    `json:"sacl6totmisses,omitempty"`
	Sacl6totpktsallowed  int    `json:"sacl6totpktsallowed,omitempty"`
	Sacl6totpktsbridged  int    `json:"sacl6totpktsbridged,omitempty"`
	Sacl6totpktsdenied   int    `json:"sacl6totpktsdenied,omitempty"`
}
