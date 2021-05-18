package vpn

type Vpnurlpolicystats struct {
	Clearstats       string `json:"clearstats,omitempty"`
	Name             string `json:"name,omitempty"`
	Pipolicyhits     int    `json:"pipolicyhits,omitempty"`
	Pipolicyhitsrate int    `json:"pipolicyhitsrate,omitempty"`
}
