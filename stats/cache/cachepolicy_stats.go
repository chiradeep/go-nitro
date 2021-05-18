package cache

type Cachepolicystats struct {
	Clearstats            string `json:"clearstats,omitempty"`
	Pipolicyhits          int    `json:"pipolicyhits,omitempty"`
	Pipolicyhitsrate      int    `json:"pipolicyhitsrate,omitempty"`
	Pipolicyundefhits     int    `json:"pipolicyundefhits,omitempty"`
	Pipolicyundefhitsrate int    `json:"pipolicyundefhitsrate,omitempty"`
	Policyname            string `json:"policyname,omitempty"`
}
