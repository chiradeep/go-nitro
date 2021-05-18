package cache

type Cachecontentgroupstats struct {
	Clearstats     string `json:"clearstats,omitempty"`
	Group304hit    int    `json:"group304hit,omitempty"`
	Groupnon304hit int    `json:"groupnon304hit,omitempty"`
	Maxmemory      int    `json:"maxmemory,omitempty"`
	Name           string `json:"name,omitempty"`
	Timesflushed   int    `json:"timesflushed,omitempty"`
	Totcell        int    `json:"totcell,omitempty"`
	Totmarkercell  int    `json:"totmarkercell,omitempty"`
	Totmemory      int    `json:"totmemory,omitempty"`
}
