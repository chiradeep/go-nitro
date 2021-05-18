package ns

type Nssimpleaclstats struct {
	Clearstats          string `json:"clearstats,omitempty"`
	Saclhitsrate        int    `json:"saclhitsrate,omitempty"`
	Saclmissesrate      int    `json:"saclmissesrate,omitempty"`
	Saclpktsallowedrate int    `json:"saclpktsallowedrate,omitempty"`
	Saclpktsbridgedrate int    `json:"saclpktsbridgedrate,omitempty"`
	Saclpktsdeniedrate  int    `json:"saclpktsdeniedrate,omitempty"`
	Saclscount          int    `json:"saclscount,omitempty"`
	Sacltothits         int    `json:"sacltothits,omitempty"`
	Sacltotmisses       int    `json:"sacltotmisses,omitempty"`
	Sacltotpktsallowed  int    `json:"sacltotpktsallowed,omitempty"`
	Sacltotpktsbridged  int    `json:"sacltotpktsbridged,omitempty"`
	Sacltotpktsdenied   int    `json:"sacltotpktsdenied,omitempty"`
}
