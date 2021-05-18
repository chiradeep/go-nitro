package basic

type Servicegroupmemberstats struct {
	Avgsvrttfb                    int    `json:"avgsvrttfb,omitempty"`
	Clearstats                    string `json:"clearstats,omitempty"`
	Curclntconnections            int    `json:"curclntconnections,omitempty"`
	Curload                       int    `json:"curload,omitempty"`
	Curreusepool                  int    `json:"curreusepool,omitempty"`
	Cursrvrconnections            int    `json:"cursrvrconnections,omitempty"`
	Frustratingttlbtransactions   int    `json:"frustratingttlbtransactions,omitempty"`
	Httpmaxhdrszpkts              int    `json:"httpmaxhdrszpkts,omitempty"`
	Ip                            string `json:"ip,omitempty"`
	Maxclients                    int    `json:"maxclients,omitempty"`
	Port                          int    `json:"port,omitempty"`
	Primaryipaddress              string `json:"primaryipaddress,omitempty"`
	Primaryport                   int    `json:"primaryport,omitempty"`
	Requestbytesrate              int    `json:"requestbytesrate,omitempty"`
	Requestsrate                  int    `json:"requestsrate,omitempty"`
	Responsebytesrate             int    `json:"responsebytesrate,omitempty"`
	Responsesrate                 int    `json:"responsesrate,omitempty"`
	Servername                    string `json:"servername,omitempty"`
	Servicegroupname              string `json:"servicegroupname,omitempty"`
	Servicetype                   string `json:"servicetype,omitempty"`
	State                         string `json:"state,omitempty"`
	Surgecount                    int    `json:"surgecount,omitempty"`
	Svrestablishedconn            int    `json:"svrestablishedconn,omitempty"`
	Tcpmaxooopkts                 int    `json:"tcpmaxooopkts,omitempty"`
	Toleratingttlbtransactions    int    `json:"toleratingttlbtransactions,omitempty"`
	Totalconnreassemblyqueue75    int    `json:"totalconnreassemblyqueue75,omitempty"`
	Totalconnreassemblyqueueflush int    `json:"totalconnreassemblyqueueflush,omitempty"`
	Totalrequestbytes             int    `json:"totalrequestbytes,omitempty"`
	Totalrequests                 int    `json:"totalrequests,omitempty"`
	Totalresponsebytes            int    `json:"totalresponsebytes,omitempty"`
	Totalresponses                int    `json:"totalresponses,omitempty"`
	Totsvrttlbtransactions        int    `json:"totsvrttlbtransactions,omitempty"`
}
