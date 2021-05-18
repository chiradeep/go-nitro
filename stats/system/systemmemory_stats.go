package system

type Systemmemorystats struct {
	Cacmemmaxmemlimit     int     `json:"cacmemmaxmemlimit,omitempty"`
	Cacmemmaxmemlimitpcnt float64 `json:"cacmemmaxmemlimitpcnt,omitempty"`
	Cacmemmaxsyslimitmb   int     `json:"cacmemmaxsyslimitmb,omitempty"`
	Clearstats            string  `json:"clearstats,omitempty"`
	Memtotallocfailed     int     `json:"memtotallocfailed,omitempty"`
	Memtotallocmb         int     `json:"memtotallocmb,omitempty"`
	Memtotallocpcnt       float64 `json:"memtotallocpcnt,omitempty"`
	Memtotavail           int     `json:"memtotavail,omitempty"`
	Memtotfree            int     `json:"memtotfree,omitempty"`
	Memtotinmb            int     `json:"memtotinmb,omitempty"`
	Memtotuseinmb         int     `json:"memtotuseinmb,omitempty"`
	Memusagepcnt          float64 `json:"memusagepcnt,omitempty"`
	Shmemallocinmb        int     `json:"shmemallocinmb,omitempty"`
	Shmemallocpcnt        float64 `json:"shmemallocpcnt,omitempty"`
	Shmemerrallocfailed   int     `json:"shmemerrallocfailed,omitempty"`
	Shmemtotinmb          int     `json:"shmemtotinmb,omitempty"`
}
