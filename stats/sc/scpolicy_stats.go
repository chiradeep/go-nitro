package sc

type Scpolicystats struct {
	Avgservertransactiontime       int    `json:"avgservertransactiontime,omitempty"`
	Avgservertransactiontimerate   int    `json:"avgservertransactiontimerate,omitempty"`
	Clearstats                     string `json:"clearstats,omitempty"`
	Clienttransactionrate          int    `json:"clienttransactionrate,omitempty"`
	Name                           string `json:"name,omitempty"`
	Openconnrate                   int    `json:"openconnrate,omitempty"`
	Scaverageclientttlb            int    `json:"scaverageclientttlb,omitempty"`
	Scaverageclientttlbrate        int    `json:"scaverageclientttlbrate,omitempty"`
	Scclientconnectionsrate        int    `json:"scclientconnectionsrate,omitempty"`
	Sccurrentclientconnections     int    `json:"sccurrentclientconnections,omitempty"`
	Sccurrentclientconnectionsrate int    `json:"sccurrentclientconnectionsrate,omitempty"`
	Sccurrentwaitingclients        int    `json:"sccurrentwaitingclients,omitempty"`
	Sccurrentwaitingclientsrate    int    `json:"sccurrentwaitingclientsrate,omitempty"`
	Sccurrentwaitingtime           int    `json:"sccurrentwaitingtime,omitempty"`
	Sccurrentwaitingtimerate       int    `json:"sccurrentwaitingtimerate,omitempty"`
	Scphysicalserviceip            string `json:"scphysicalserviceip,omitempty"`
	Scphysicalserviceport          int    `json:"scphysicalserviceport,omitempty"`
	Screquestbytesrate             int    `json:"screquestbytesrate,omitempty"`
	Screquestsreceivedrate         int    `json:"screquestsreceivedrate,omitempty"`
	Scresponsebytesrate            int    `json:"scresponsebytesrate,omitempty"`
	Scresponsesreceivedrate        int    `json:"scresponsesreceivedrate,omitempty"`
	Scserverconnectionsrate        int    `json:"scserverconnectionsrate,omitempty"`
	Scservertransactionsrate       int    `json:"scservertransactionsrate,omitempty"`
	Sctotalclientconnections       int    `json:"sctotalclientconnections,omitempty"`
	Sctotalrequestbytes            int    `json:"sctotalrequestbytes,omitempty"`
	Sctotalrequestsreceived        int    `json:"sctotalrequestsreceived,omitempty"`
	Sctotalresponsebytes           int    `json:"sctotalresponsebytes,omitempty"`
	Sctotalresponsesreceived       int    `json:"sctotalresponsesreceived,omitempty"`
	Sctotalserverconnections       int    `json:"sctotalserverconnections,omitempty"`
	Sctotalservertransactions      int    `json:"sctotalservertransactions,omitempty"`
	Totclienttransaction           int    `json:"totclienttransaction,omitempty"`
	Totopenconn                    int    `json:"totopenconn,omitempty"`
}
