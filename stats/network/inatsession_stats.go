package network

type Inatsessionstats struct {
	Clearstats                 string `json:"clearstats,omitempty"`
	Globalinatcursessions      int    `json:"globalinatcursessions,omitempty"`
	Globalinatcursessionsrate  int    `json:"globalinatcursessionsrate,omitempty"`
	Globalinathits             int    `json:"globalinathits,omitempty"`
	Globalinathitsrate         int    `json:"globalinathitsrate,omitempty"`
	Globalinatpktreceived      int    `json:"globalinatpktreceived,omitempty"`
	Globalinatpktreceivedrate  int    `json:"globalinatpktreceivedrate,omitempty"`
	Globalinatpktsent          int    `json:"globalinatpktsent,omitempty"`
	Globalinatpktsentrate      int    `json:"globalinatpktsentrate,omitempty"`
	Globalinatreceivebytes     int    `json:"globalinatreceivebytes,omitempty"`
	Globalinatreceivebytesrate int    `json:"globalinatreceivebytesrate,omitempty"`
	Globalinatsentbytesrate    int    `json:"globalinatsentbytesrate,omitempty"`
	Globalinattotsentbytes     int    `json:"globalinattotsentbytes,omitempty"`
	Inatcursessions            int    `json:"inatcursessions,omitempty"`
	Inatcursessionsrate        int    `json:"inatcursessionsrate,omitempty"`
	Inathitsrate               int    `json:"inathitsrate,omitempty"`
	Inatpktreceivedrate        int    `json:"inatpktreceivedrate,omitempty"`
	Inatpktsentrate            int    `json:"inatpktsentrate,omitempty"`
	Inatreceivebytesrate       int    `json:"inatreceivebytesrate,omitempty"`
	Inatsentbytesrate          int    `json:"inatsentbytesrate,omitempty"`
	Inattothits                int    `json:"inattothits,omitempty"`
	Inattotpktreceived         int    `json:"inattotpktreceived,omitempty"`
	Inattotpktsent             int    `json:"inattotpktsent,omitempty"`
	Inattotreceivebytes        int    `json:"inattotreceivebytes,omitempty"`
	Inattotsentbytes           int    `json:"inattotsentbytes,omitempty"`
	Name                       string `json:"name,omitempty"`
}
