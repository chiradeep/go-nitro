package basic

type Servicestats struct {
	Activetransactions            int         `json:"activetransactions,omitempty"`
	Avgsvrttfb                    int         `json:"avgsvrttfb,omitempty"`
	Clearstats                    string      `json:"clearstats,omitempty"`
	Curclntconnections            int         `json:"curclntconnections,omitempty"`
	Curload                       int         `json:"curload,omitempty"`
	Curreusepool                  int         `json:"curreusepool,omitempty"`
	Cursrvrconnections            int         `json:"cursrvrconnections,omitempty"`
	Curtflags                     int         `json:"curtflags,omitempty"`
	Dospolicy                     interface{} `json:"dospolicy,omitempty"`
	Frustratingttlbtransactions   int         `json:"frustratingttlbtransactions,omitempty"`
	Httpmaxhdrszpkts              int         `json:"httpmaxhdrszpkts,omitempty"`
	Maxclients                    int         `json:"maxclients,omitempty"`
	Name                          string      `json:"name,omitempty"`
	Primaryipaddress              string      `json:"primaryipaddress,omitempty"`
	Primaryport                   int         `json:"primaryport,omitempty"`
	Requestbytesrate              int         `json:"requestbytesrate,omitempty"`
	Requestsrate                  int         `json:"requestsrate,omitempty"`
	Responsebytesrate             int         `json:"responsebytesrate,omitempty"`
	Responsesrate                 int         `json:"responsesrate,omitempty"`
	Scpolicy                      interface{} `json:"scpolicy,omitempty"`
	Servicetype                   string      `json:"servicetype,omitempty"`
	State                         string      `json:"state,omitempty"`
	Surgecount                    int         `json:"surgecount,omitempty"`
	Svrestablishedconn            int         `json:"svrestablishedconn,omitempty"`
	Tcpmaxooopkts                 int         `json:"tcpmaxooopkts,omitempty"`
	Throughput                    int         `json:"throughput,omitempty"`
	Throughputrate                int         `json:"throughputrate,omitempty"`
	Toleratingttlbtransactions    int         `json:"toleratingttlbtransactions,omitempty"`
	Totalconnreassemblyqueue75    int         `json:"totalconnreassemblyqueue75,omitempty"`
	Totalconnreassemblyqueueflush int         `json:"totalconnreassemblyqueueflush,omitempty"`
	Totalrequestbytes             int         `json:"totalrequestbytes,omitempty"`
	Totalrequests                 int         `json:"totalrequests,omitempty"`
	Totalresponsebytes            int         `json:"totalresponsebytes,omitempty"`
	Totalresponses                int         `json:"totalresponses,omitempty"`
	Totsvrttlbtransactions        int         `json:"totsvrttlbtransactions,omitempty"`
	Vsvrservicehits               int         `json:"vsvrservicehits,omitempty"`
	Vsvrservicehitsrate           int         `json:"vsvrservicehitsrate,omitempty"`
}
