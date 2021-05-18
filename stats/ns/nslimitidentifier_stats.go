package ns

type Nslimitidentifierstats struct {
	Clearstats            string      `json:"clearstats,omitempty"`
	Name                  string      `json:"name,omitempty"`
	Pattern               interface{} `json:"pattern,omitempty"`
	Ratelmtobjdrops       int         `json:"ratelmtobjdrops,omitempty"`
	Ratelmtobjhits        int         `json:"ratelmtobjhits,omitempty"`
	Ratelmtsessionobjhits int         `json:"ratelmtsessionobjhits,omitempty"`
	Sortby                string      `json:"sortby,omitempty"`
	Sortorder             string      `json:"sortorder,omitempty"`
}
