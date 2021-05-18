package lsn

type Lsnpoolstats struct {
	Clearstats                   string  `json:"clearstats,omitempty"`
	Lsnpoolcurportallocother     int     `json:"lsnpoolcurportallocother,omitempty"`
	Lsnpoolcurportallocotherrate int     `json:"lsnpoolcurportallocotherrate,omitempty"`
	Lsnpoolcurportalloctcp       int     `json:"lsnpoolcurportalloctcp,omitempty"`
	Lsnpoolcurportalloctcprate   int     `json:"lsnpoolcurportalloctcprate,omitempty"`
	Lsnpoolotherportusagepercen  float64 `json:"lsnpoolotherportusagepercen,omitempty"`
	Lsnpooltcpportusagepercen    float64 `json:"lsnpooltcpportusagepercen,omitempty"`
	Lsnpooltotips                int     `json:"lsnpooltotips,omitempty"`
	Poolname                     string  `json:"poolname,omitempty"`
}
