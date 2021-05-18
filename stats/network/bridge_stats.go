package network

type Bridgestats struct {
	Bdgmbitsrate         int    `json:"bdgmbitsrate,omitempty"`
	Bdgpktsrate          int    `json:"bdgpktsrate,omitempty"`
	Clearstats           string `json:"clearstats,omitempty"`
	Tcpbdgcollisionsrate int    `json:"tcpbdgcollisionsrate,omitempty"`
	Tcpbdgmacmovedrate   int    `json:"tcpbdgmacmovedrate,omitempty"`
	Tcperrbdgmuted       int    `json:"tcperrbdgmuted,omitempty"`
	Tcperrbdgmutedrate   int    `json:"tcperrbdgmutedrate,omitempty"`
	Tcptotbdgcollisions  int    `json:"tcptotbdgcollisions,omitempty"`
	Tcptotbdgmacmoved    int    `json:"tcptotbdgmacmoved,omitempty"`
	Totbdgmbits          int    `json:"totbdgmbits,omitempty"`
	Totbdgpkts           int    `json:"totbdgpkts,omitempty"`
}
