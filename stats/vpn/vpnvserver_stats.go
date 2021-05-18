package vpn

type Vpnvserverstats struct {
	Authenticationloginschemapolicy interface{} `json:"authenticationloginschemapolicy,omitempty"`
	Authenticationoauthidppolicy    interface{} `json:"authenticationoauthidppolicy,omitempty"`
	Authenticationpolicy            interface{} `json:"authenticationpolicy,omitempty"`
	Authenticationsamlidppolicy     interface{} `json:"authenticationsamlidppolicy,omitempty"`
	Cachepolicy                     interface{} `json:"cachepolicy,omitempty"`
	Clearstats                      string      `json:"clearstats,omitempty"`
	Icapolicy                       interface{} `json:"icapolicy,omitempty"`
	Name                            string      `json:"name,omitempty"`
	Primaryipaddress                string      `json:"primaryipaddress,omitempty"`
	Primaryport                     int         `json:"primaryport,omitempty"`
	Requestbytesrate                int         `json:"requestbytesrate,omitempty"`
	Requestsrate                    int         `json:"requestsrate,omitempty"`
	Responderpolicy                 interface{} `json:"responderpolicy,omitempty"`
	Responsebytesrate               int         `json:"responsebytesrate,omitempty"`
	Responsesrate                   int         `json:"responsesrate,omitempty"`
	Rewritepolicy                   interface{} `json:"rewritepolicy,omitempty"`
	State                           string      `json:"state,omitempty"`
	Totalrequestbytes               int         `json:"totalrequestbytes,omitempty"`
	Totalrequests                   int         `json:"totalrequests,omitempty"`
	Totalresponsebytes              int         `json:"totalresponsebytes,omitempty"`
	Totalresponses                  int         `json:"totalresponses,omitempty"`
	Type                            string      `json:"type,omitempty"`
	Vpnurlpolicy                    interface{} `json:"vpnurlpolicy,omitempty"`
}
