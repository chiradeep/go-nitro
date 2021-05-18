package ipsecalg

type Ipsecalgcountersstats struct {
	Clearstats                 string `json:"clearstats,omitempty"`
	Ipsecalgcuractsessions     int    `json:"ipsecalgcuractsessions,omitempty"`
	Ipsecalgcuractsessionsrate int    `json:"ipsecalgcuractsessionsrate,omitempty"`
	Ipsecalgcurblksessions     int    `json:"ipsecalgcurblksessions,omitempty"`
	Ipsecalgcurblksessionsrate int    `json:"ipsecalgcurblksessionsrate,omitempty"`
	Ipsecalgdrsessionsrate     int    `json:"ipsecalgdrsessionsrate,omitempty"`
	Ipsecalgrxpktsrate         int    `json:"ipsecalgrxpktsrate,omitempty"`
	Ipsecalgsessionsrate       int    `json:"ipsecalgsessionsrate,omitempty"`
	Ipsecalgtotdrsessions      int    `json:"ipsecalgtotdrsessions,omitempty"`
	Ipsecalgtotrxpkts          int    `json:"ipsecalgtotrxpkts,omitempty"`
	Ipsecalgtotsessions        int    `json:"ipsecalgtotsessions,omitempty"`
	Ipsecalgtottxpkts          int    `json:"ipsecalgtottxpkts,omitempty"`
	Ipsecalgtxpktsrate         int    `json:"ipsecalgtxpktsrate,omitempty"`
	Name                       string `json:"name,omitempty"`
}
