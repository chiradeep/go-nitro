package protocol

type Protocoludpstats struct {
	Clearstats                 string `json:"clearstats,omitempty"`
	Udpbadchecksum             int    `json:"udpbadchecksum,omitempty"`
	Udpcurratethreshold        int    `json:"udpcurratethreshold,omitempty"`
	Udpcurratethresholdexceeds int    `json:"udpcurratethresholdexceeds,omitempty"`
	Udprxbytesrate             int    `json:"udprxbytesrate,omitempty"`
	Udprxpktsrate              int    `json:"udprxpktsrate,omitempty"`
	Udptotrxbytes              int    `json:"udptotrxbytes,omitempty"`
	Udptotrxpkts               int    `json:"udptotrxpkts,omitempty"`
	Udptottxbytes              int    `json:"udptottxbytes,omitempty"`
	Udptottxpkts               int    `json:"udptottxpkts,omitempty"`
	Udptotunknownsvcpkts       int    `json:"udptotunknownsvcpkts,omitempty"`
	Udptxbytesrate             int    `json:"udptxbytesrate,omitempty"`
	Udptxpktsrate              int    `json:"udptxpktsrate,omitempty"`
}
