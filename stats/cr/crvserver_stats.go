package cr

type Crvserverstats struct {
	Appfwpolicy                   interface{} `json:"appfwpolicy,omitempty"`
	Appqoepolicy                  interface{} `json:"appqoepolicy,omitempty"`
	Cachepolicy                   interface{} `json:"cachepolicy,omitempty"`
	Clearstats                    string      `json:"clearstats,omitempty"`
	Cmppolicy                     interface{} `json:"cmppolicy,omitempty"`
	Hitsrate                      int         `json:"hitsrate,omitempty"`
	Icapolicy                     interface{} `json:"icapolicy,omitempty"`
	Invalidrequestresponse        int         `json:"invalidrequestresponse,omitempty"`
	Invalidrequestresponsedropped int         `json:"invalidrequestresponsedropped,omitempty"`
	Lbvserver                     interface{} `json:"lbvserver,omitempty"`
	Name                          string      `json:"name,omitempty"`
	Pktsrecvdrate                 int         `json:"pktsrecvdrate,omitempty"`
	Pktssentrate                  int         `json:"pktssentrate,omitempty"`
	Primaryipaddress              string      `json:"primaryipaddress,omitempty"`
	Primaryport                   int         `json:"primaryport,omitempty"`
	Requestbytesrate              int         `json:"requestbytesrate,omitempty"`
	Requestsrate                  int         `json:"requestsrate,omitempty"`
	Responderpolicy               interface{} `json:"responderpolicy,omitempty"`
	Responsebytesrate             int         `json:"responsebytesrate,omitempty"`
	Responsesrate                 int         `json:"responsesrate,omitempty"`
	Rewritepolicy                 interface{} `json:"rewritepolicy,omitempty"`
	Spilloverpolicy               interface{} `json:"spilloverpolicy,omitempty"`
	State                         string      `json:"state,omitempty"`
	Totalpktsrecvd                int         `json:"totalpktsrecvd,omitempty"`
	Totalpktssent                 int         `json:"totalpktssent,omitempty"`
	Totalrequestbytes             int         `json:"totalrequestbytes,omitempty"`
	Totalrequests                 int         `json:"totalrequests,omitempty"`
	Totalresponsebytes            int         `json:"totalresponsebytes,omitempty"`
	Totalresponses                int         `json:"totalresponses,omitempty"`
	Tothits                       int         `json:"tothits,omitempty"`
	Type                          string      `json:"type,omitempty"`
}
