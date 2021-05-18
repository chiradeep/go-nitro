package cluster

type Clusterinstancestats struct {
	Clbkplanerx           int    `json:"clbkplanerx,omitempty"`
	Clbkplanerxrate       int    `json:"clbkplanerxrate,omitempty"`
	Clbkplanetx           int    `json:"clbkplanetx,omitempty"`
	Clbkplanetxrate       int    `json:"clbkplanetxrate,omitempty"`
	Clcurstatus           string `json:"clcurstatus,omitempty"`
	Clearstats            string `json:"clearstats,omitempty"`
	Clid                  int    `json:"clid,omitempty"`
	Clnumnodes            int    `json:"clnumnodes,omitempty"`
	Clviewleader          string `json:"clviewleader,omitempty"`
	Numdfddroppkts        int    `json:"numdfddroppkts,omitempty"`
	Totpropagationtimeout int    `json:"totpropagationtimeout,omitempty"`
	Totsteeredpkts        int    `json:"totsteeredpkts,omitempty"`
}
