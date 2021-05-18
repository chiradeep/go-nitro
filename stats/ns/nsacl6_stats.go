package ns

type Nsacl6stats struct {
	Acl6hitsrate        int    `json:"acl6hitsrate,omitempty"`
	Acl6missesrate      int    `json:"acl6missesrate,omitempty"`
	Acl6name            string `json:"acl6name,omitempty"`
	Acl6perhits         int    `json:"acl6perhits,omitempty"`
	Acl6perhitsrate     int    `json:"acl6perhitsrate,omitempty"`
	Acl6pktsallowedrate int    `json:"acl6pktsallowedrate,omitempty"`
	Acl6pktsbridgedrate int    `json:"acl6pktsbridgedrate,omitempty"`
	Acl6pktsdeniedrate  int    `json:"acl6pktsdeniedrate,omitempty"`
	Acl6pktsnat64rate   int    `json:"acl6pktsnat64rate,omitempty"`
	Acl6pktsnatrate     int    `json:"acl6pktsnatrate,omitempty"`
	Acl6totcount        int    `json:"acl6totcount,omitempty"`
	Acl6tothits         int    `json:"acl6tothits,omitempty"`
	Acl6totmisses       int    `json:"acl6totmisses,omitempty"`
	Acl6totpktsallowed  int    `json:"acl6totpktsallowed,omitempty"`
	Acl6totpktsbridged  int    `json:"acl6totpktsbridged,omitempty"`
	Acl6totpktsdenied   int    `json:"acl6totpktsdenied,omitempty"`
	Acl6totpktsnat      int    `json:"acl6totpktsnat,omitempty"`
	Acl6totpktsnat64    int    `json:"acl6totpktsnat64,omitempty"`
	Clearstats          string `json:"clearstats,omitempty"`
	Dfdacl6hitsrate     int    `json:"dfdacl6hitsrate,omitempty"`
	Dfdacl6missesrate   int    `json:"dfdacl6missesrate,omitempty"`
	Dfdacl6totcount     int    `json:"dfdacl6totcount,omitempty"`
	Dfdacl6tothits      int    `json:"dfdacl6tothits,omitempty"`
	Dfdacl6totmisses    int    `json:"dfdacl6totmisses,omitempty"`
}
