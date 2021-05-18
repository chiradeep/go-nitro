package ns

type Nspbr6stats struct {
	Clearstats          string `json:"clearstats,omitempty"`
	Name                string `json:"name,omitempty"`
	Pbr6hitsrate        int    `json:"pbr6hitsrate,omitempty"`
	Pbr6missesrate      int    `json:"pbr6missesrate,omitempty"`
	Pbr6nulldroprate    int    `json:"pbr6nulldroprate,omitempty"`
	Pbr6perhits         int    `json:"pbr6perhits,omitempty"`
	Pbr6perhitsrate     int    `json:"pbr6perhitsrate,omitempty"`
	Pbr6pktsallowedrate int    `json:"pbr6pktsallowedrate,omitempty"`
	Pbr6pktsdeniedrate  int    `json:"pbr6pktsdeniedrate,omitempty"`
	Pbr6tothits         int    `json:"pbr6tothits,omitempty"`
	Pbr6totmisses       int    `json:"pbr6totmisses,omitempty"`
	Pbr6totnulldrop     int    `json:"pbr6totnulldrop,omitempty"`
	Pbr6totpktsallowed  int    `json:"pbr6totpktsallowed,omitempty"`
	Pbr6totpktsdenied   int    `json:"pbr6totpktsdenied,omitempty"`
}
