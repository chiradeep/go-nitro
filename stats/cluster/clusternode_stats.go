package cluster

type Clusternodestats struct {
	Clearstats            string `json:"clearstats,omitempty"`
	Clmasterstate         string `json:"clmasterstate,omitempty"`
	Clnodeeffectivehealth string `json:"clnodeeffectivehealth,omitempty"`
	Clnodeip              string `json:"clnodeip,omitempty"`
	Clptprx               int    `json:"clptprx,omitempty"`
	Clptpstate            string `json:"clptpstate,omitempty"`
	Clptptx               int    `json:"clptptx,omitempty"`
	Clsyncstate           string `json:"clsyncstate,omitempty"`
	Cltothbrx             int    `json:"cltothbrx,omitempty"`
	Cltothbtx             int    `json:"cltothbtx,omitempty"`
	Nnmcurconn            int    `json:"nnmcurconn,omitempty"`
	Nnmerrmsend           int    `json:"nnmerrmsend,omitempty"`
	Nnmtotconnrx          int    `json:"nnmtotconnrx,omitempty"`
	Nnmtotconntx          int    `json:"nnmtotconntx,omitempty"`
	Nodeid                int    `json:"nodeid,omitempty"`
}
