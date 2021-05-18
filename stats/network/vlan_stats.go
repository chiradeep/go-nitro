package network

type Vlanstats struct {
	Clearstats           string      `json:"clearstats,omitempty"`
	Id                   int         `json:"id,omitempty"`
	Interface            interface{} `json:"Interface,omitempty"`
	Vlanrxbytesrate      int         `json:"vlanrxbytesrate,omitempty"`
	Vlanrxpktsrate       int         `json:"vlanrxpktsrate,omitempty"`
	Vlantotbroadcastpkts int         `json:"vlantotbroadcastpkts,omitempty"`
	Vlantotdroppedpkts   int         `json:"vlantotdroppedpkts,omitempty"`
	Vlantotrxbytes       int         `json:"vlantotrxbytes,omitempty"`
	Vlantotrxpkts        int         `json:"vlantotrxpkts,omitempty"`
	Vlantottxbytes       int         `json:"vlantottxbytes,omitempty"`
	Vlantottxpkts        int         `json:"vlantottxpkts,omitempty"`
	Vlantxbytesrate      int         `json:"vlantxbytesrate,omitempty"`
	Vlantxpktsrate       int         `json:"vlantxpktsrate,omitempty"`
}
