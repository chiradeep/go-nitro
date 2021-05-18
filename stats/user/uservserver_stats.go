package user

type Uservserverstats struct {
	Actsvcs                       int    `json:"actsvcs,omitempty"`
	Clearstats                    string `json:"clearstats,omitempty"`
	Curclntconnections            int    `json:"curclntconnections,omitempty"`
	Cursrvrconnections            int    `json:"cursrvrconnections,omitempty"`
	Establishedconn               int    `json:"establishedconn,omitempty"`
	Hitsrate                      int    `json:"hitsrate,omitempty"`
	Inactsvcs                     int    `json:"inactsvcs,omitempty"`
	Invalidrequestresponse        int    `json:"invalidrequestresponse,omitempty"`
	Invalidrequestresponsedropped int    `json:"invalidrequestresponsedropped,omitempty"`
	Name                          string `json:"name,omitempty"`
	Pktsrecvdrate                 int    `json:"pktsrecvdrate,omitempty"`
	Pktssentrate                  int    `json:"pktssentrate,omitempty"`
	Primaryipaddress              string `json:"primaryipaddress,omitempty"`
	Primaryport                   int    `json:"primaryport,omitempty"`
	Protocolname                  string `json:"protocolname,omitempty"`
	Requestbytesrate              int    `json:"requestbytesrate,omitempty"`
	Requestsrate                  int    `json:"requestsrate,omitempty"`
	Responsebytesrate             int    `json:"responsebytesrate,omitempty"`
	Responsesrate                 int    `json:"responsesrate,omitempty"`
	Sortby                        string `json:"sortby,omitempty"`
	Sortorder                     string `json:"sortorder,omitempty"`
	State                         string `json:"state,omitempty"`
	Totalpktsrecvd                int    `json:"totalpktsrecvd,omitempty"`
	Totalpktssent                 int    `json:"totalpktssent,omitempty"`
	Totalrequestbytes             int    `json:"totalrequestbytes,omitempty"`
	Totalrequests                 int    `json:"totalrequests,omitempty"`
	Totalresponsebytes            int    `json:"totalresponsebytes,omitempty"`
	Totalresponses                int    `json:"totalresponses,omitempty"`
	Tothits                       int    `json:"tothits,omitempty"`
	Vslbhealth                    int    `json:"vslbhealth,omitempty"`
}
