package pcp

type Pcpserverstats struct {
	Clearstats                 string `json:"clearstats,omitempty"`
	Name                       string `json:"name,omitempty"`
	Pcpaddrmismatchrate        int    `json:"pcpaddrmismatchrate,omitempty"`
	Pcperrinresrate            int    `json:"pcperrinresrate,omitempty"`
	Pcperrinrquestrate         int    `json:"pcperrinrquestrate,omitempty"`
	Pcpexcesspeersrate         int    `json:"pcpexcesspeersrate,omitempty"`
	Pcpmalformedoptionrate     int    `json:"pcpmalformedoptionrate,omitempty"`
	Pcpmalformedreqrate        int    `json:"pcpmalformedreqrate,omitempty"`
	Pcpmaprequestsrate         int    `json:"pcpmaprequestsrate,omitempty"`
	Pcpnetfailurerate          int    `json:"pcpnetfailurerate,omitempty"`
	Pcpnoexternalrate          int    `json:"pcpnoexternalrate,omitempty"`
	Pcpnoresourcesrate         int    `json:"pcpnoresourcesrate,omitempty"`
	Pcppeerrequestsrate        int    `json:"pcppeerrequestsrate,omitempty"`
	Pcprequestsrate            int    `json:"pcprequestsrate,omitempty"`
	Pcpresponsesrate           int    `json:"pcpresponsesrate,omitempty"`
	Pcptotaddrmismatch         int    `json:"pcptotaddrmismatch,omitempty"`
	Pcptoterrinres             int    `json:"pcptoterrinres,omitempty"`
	Pcptoterrinrquest          int    `json:"pcptoterrinrquest,omitempty"`
	Pcptotexcesspeers          int    `json:"pcptotexcesspeers,omitempty"`
	Pcptotmalformedoption      int    `json:"pcptotmalformedoption,omitempty"`
	Pcptotmalformedreq         int    `json:"pcptotmalformedreq,omitempty"`
	Pcptotmaprequests          int    `json:"pcptotmaprequests,omitempty"`
	Pcptotnetfailure           int    `json:"pcptotnetfailure,omitempty"`
	Pcptotnoexternal           int    `json:"pcptotnoexternal,omitempty"`
	Pcptotnoresources          int    `json:"pcptotnoresources,omitempty"`
	Pcptotpeerrequests         int    `json:"pcptotpeerrequests,omitempty"`
	Pcptotrequests             int    `json:"pcptotrequests,omitempty"`
	Pcptotresponses            int    `json:"pcptotresponses,omitempty"`
	Pcptotunsuppopcode         int    `json:"pcptotunsuppopcode,omitempty"`
	Pcptotunsuppoption         int    `json:"pcptotunsuppoption,omitempty"`
	Pcptotunsupportedprotocol  int    `json:"pcptotunsupportedprotocol,omitempty"`
	Pcptotunsuppvers           int    `json:"pcptotunsuppvers,omitempty"`
	Pcptotuserexqouta          int    `json:"pcptotuserexqouta,omitempty"`
	Pcpunsuppopcoderate        int    `json:"pcpunsuppopcoderate,omitempty"`
	Pcpunsuppoptionrate        int    `json:"pcpunsuppoptionrate,omitempty"`
	Pcpunsupportedprotocolrate int    `json:"pcpunsupportedprotocolrate,omitempty"`
	Pcpunsuppversrate          int    `json:"pcpunsuppversrate,omitempty"`
	Pcpuserexqoutarate         int    `json:"pcpuserexqoutarate,omitempty"`
}
