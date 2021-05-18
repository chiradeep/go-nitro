package gslb

type Gslbvserverstats struct {
	Actsvcs                  int         `json:"actsvcs,omitempty"`
	Clearstats               string      `json:"clearstats,omitempty"`
	Curclntconnections       int         `json:"curclntconnections,omitempty"`
	Curpersistencesessions   int         `json:"curpersistencesessions,omitempty"`
	Cursrvrconnections       int         `json:"cursrvrconnections,omitempty"`
	Establishedconn          int         `json:"establishedconn,omitempty"`
	Gslbservice              interface{} `json:"gslbservice,omitempty"`
	Gslbservicegroup         interface{} `json:"gslbservicegroup,omitempty"`
	Hitsrate                 int         `json:"hitsrate,omitempty"`
	Inactsvcs                int         `json:"inactsvcs,omitempty"`
	Name                     string      `json:"name,omitempty"`
	Requestbytesrate         int         `json:"requestbytesrate,omitempty"`
	Requestsrate             int         `json:"requestsrate,omitempty"`
	Responsebytesrate        int         `json:"responsebytesrate,omitempty"`
	Responsesrate            int         `json:"responsesrate,omitempty"`
	Sothreshold              int         `json:"sothreshold,omitempty"`
	Spilloverpolicy          interface{} `json:"spilloverpolicy,omitempty"`
	State                    string      `json:"state,omitempty"`
	Totalrequestbytes        int         `json:"totalrequestbytes,omitempty"`
	Totalrequests            int         `json:"totalrequests,omitempty"`
	Totalresponsebytes       int         `json:"totalresponsebytes,omitempty"`
	Totalresponses           int         `json:"totalresponses,omitempty"`
	Tothits                  int         `json:"tothits,omitempty"`
	Totspillovers            int         `json:"totspillovers,omitempty"`
	Totvserverdownbackuphits int         `json:"totvserverdownbackuphits,omitempty"`
	Type                     string      `json:"type,omitempty"`
	Vslbhealth               int         `json:"vslbhealth,omitempty"`
	Vsvrtotbkplbfail         int         `json:"vsvrtotbkplbfail,omitempty"`
	Vsvrtotbkplbhits         int         `json:"vsvrtotbkplbhits,omitempty"`
	Vsvrtotpersistencehits   int         `json:"vsvrtotpersistencehits,omitempty"`
}
