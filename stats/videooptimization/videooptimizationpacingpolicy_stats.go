package videooptimization

type Videooptimizationpacingpolicystats struct {
	Clearstats            string `json:"clearstats,omitempty"`
	Name                  string `json:"name,omitempty"`
	Pipolicyhits          int    `json:"pipolicyhits,omitempty"`
	Pipolicyhitsrate      int    `json:"pipolicyhitsrate,omitempty"`
	Pipolicyundefhits     int    `json:"pipolicyundefhits,omitempty"`
	Pipolicyundefhitsrate int    `json:"pipolicyundefhitsrate,omitempty"`
}
