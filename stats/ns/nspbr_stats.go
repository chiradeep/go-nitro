package ns

type Nspbrstats struct {
	Clearstats         string `json:"clearstats,omitempty"`
	Name               string `json:"name,omitempty"`
	Pbrhitsrate        int    `json:"pbrhitsrate,omitempty"`
	Pbrmissesrate      int    `json:"pbrmissesrate,omitempty"`
	Pbrnulldroprate    int    `json:"pbrnulldroprate,omitempty"`
	Pbrperhits         int    `json:"pbrperhits,omitempty"`
	Pbrperhitsrate     int    `json:"pbrperhitsrate,omitempty"`
	Pbrpktsallowedrate int    `json:"pbrpktsallowedrate,omitempty"`
	Pbrpktsdeniedrate  int    `json:"pbrpktsdeniedrate,omitempty"`
	Pbrtothits         int    `json:"pbrtothits,omitempty"`
	Pbrtotmisses       int    `json:"pbrtotmisses,omitempty"`
	Pbrtotnulldrop     int    `json:"pbrtotnulldrop,omitempty"`
	Pbrtotpktsallowed  int    `json:"pbrtotpktsallowed,omitempty"`
	Pbrtotpktsdenied   int    `json:"pbrtotpktsdenied,omitempty"`
}
