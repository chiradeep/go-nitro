package ns

type Nspartitionstats struct {
	Clearstats           string      `json:"clearstats,omitempty"`
	Connectiondropsrate  int         `json:"connectiondropsrate,omitempty"`
	Currentbandwidth     int         `json:"currentbandwidth,omitempty"`
	Currentconnections   int         `json:"currentconnections,omitempty"`
	Dropsrate            int         `json:"dropsrate,omitempty"`
	Maxbandwidth         int         `json:"maxbandwidth,omitempty"`
	Maxconnection        int         `json:"maxconnection,omitempty"`
	Maxmemory            int         `json:"maxmemory,omitempty"`
	Memoryusagepcnt      int         `json:"memoryusagepcnt,omitempty"`
	Partitionname        string      `json:"partitionname,omitempty"`
	Tokendropsrate       int         `json:"tokendropsrate,omitempty"`
	Totalconnectiondrops int         `json:"totalconnectiondrops,omitempty"`
	Totaldrops           int         `json:"totaldrops,omitempty"`
	Totaltokendrops      int         `json:"totaltokendrops,omitempty"`
	Vlan                 interface{} `json:"vlan,omitempty"`
	Vxlan                interface{} `json:"vxlan,omitempty"`
}
