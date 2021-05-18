package ns

type Nstrafficdomainstats struct {
	Clearstats          string      `json:"clearstats,omitempty"`
	Nstddroppedpktsrate int         `json:"nstddroppedpktsrate,omitempty"`
	Nstdrxpktsrate      int         `json:"nstdrxpktsrate,omitempty"`
	Nstdtotdroppedpkts  int         `json:"nstdtotdroppedpkts,omitempty"`
	Nstdtotrxpkts       int         `json:"nstdtotrxpkts,omitempty"`
	Nstdtottxpkts       int         `json:"nstdtottxpkts,omitempty"`
	Nstdtxpktsrate      int         `json:"nstdtxpktsrate,omitempty"`
	Td                  int         `json:"td,omitempty"`
	Vlan                interface{} `json:"vlan,omitempty"`
	Vxlan               interface{} `json:"vxlan,omitempty"`
}
