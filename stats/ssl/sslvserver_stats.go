package ssl

type Sslvserverstats struct {
	Actsvcs                  int         `json:"actsvcs,omitempty"`
	Appfwpolicy              interface{} `json:"appfwpolicy,omitempty"`
	Cachepolicy              interface{} `json:"cachepolicy,omitempty"`
	Clearstats               string      `json:"clearstats,omitempty"`
	Cmppolicy                interface{} `json:"cmppolicy,omitempty"`
	Pqpolicy                 interface{} `json:"pqpolicy,omitempty"`
	Primaryipaddress         string      `json:"primaryipaddress,omitempty"`
	Primaryport              int         `json:"primaryport,omitempty"`
	Responderpolicy          interface{} `json:"responderpolicy,omitempty"`
	Rewritepolicy            interface{} `json:"rewritepolicy,omitempty"`
	Scpolicy                 interface{} `json:"scpolicy,omitempty"`
	Sslclientauthfailurerate int         `json:"sslclientauthfailurerate,omitempty"`
	Sslclientauthsuccessrate int         `json:"sslclientauthsuccessrate,omitempty"`
	Sslctxdecbytesrate       int         `json:"sslctxdecbytesrate,omitempty"`
	Sslctxencbytesrate       int         `json:"sslctxencbytesrate,omitempty"`
	Sslctxhwdecbytesrate     int         `json:"sslctxhwdec_bytesrate,omitempty"`
	Sslctxhwencbytesrate     int         `json:"sslctxhwencbytesrate,omitempty"`
	Sslctxsessionhitsrate    int         `json:"sslctxsessionhitsrate,omitempty"`
	Sslctxsessionnewrate     int         `json:"sslctxsessionnewrate,omitempty"`
	Sslctxtotdecbytes        int         `json:"sslctxtotdecbytes,omitempty"`
	Sslctxtotencbytes        int         `json:"sslctxtotencbytes,omitempty"`
	Sslctxtothwdecbytes      int         `json:"sslctxtothwdec_bytes,omitempty"`
	Sslctxtothwencbytes      int         `json:"sslctxtothwencbytes,omitempty"`
	Sslctxtotsessionhits     int         `json:"sslctxtotsessionhits,omitempty"`
	Sslctxtotsessionnew      int         `json:"sslctxtotsessionnew,omitempty"`
	Ssltotclientauthfailure  int         `json:"ssltotclientauthfailure,omitempty"`
	Ssltotclientauthsuccess  int         `json:"ssltotclientauthsuccess,omitempty"`
	State                    string      `json:"state,omitempty"`
	Type                     string      `json:"type,omitempty"`
	Vservername              string      `json:"vservername,omitempty"`
	Vslbhealth               int         `json:"vslbhealth,omitempty"`
}
