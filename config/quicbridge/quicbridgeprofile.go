package quicbridge

type Quicbridgeprofile struct {
	Name             string `json:"name,omitempty"`
	Refcnt           int    `json:"refcnt,omitempty"`
	Routingalgorithm string `json:"routingalgorithm,omitempty"`
	Serveridlength   int    `json:"serveridlength,omitempty"`
}
