package ns

type Nsmemorystats struct {
	Allocf         float64 `json:"allocf,omitempty"`
	Clearstats     string  `json:"clearstats,omitempty"`
	Memcurallocper float64 `json:"memcurallocper,omitempty"`
	Memcurinkb     int     `json:"memcurinkb,omitempty"`
	Pool           string  `json:"pool,omitempty"`
}
