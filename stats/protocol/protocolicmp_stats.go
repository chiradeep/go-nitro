package protocol

type Protocolicmpstats struct {
	Clearstats                   string `json:"clearstats,omitempty"`
	Icmpcurratethreshold         int    `json:"icmpcurratethreshold,omitempty"`
	Icmprxbytesrate              int    `json:"icmprxbytesrate,omitempty"`
	Icmprxechorate               int    `json:"icmprxechorate,omitempty"`
	Icmprxechoreplyrate          int    `json:"icmprxechoreplyrate,omitempty"`
	Icmprxpktsrate               int    `json:"icmprxpktsrate,omitempty"`
	Icmptotbadchecksum           int    `json:"icmptotbadchecksum,omitempty"`
	Icmptotbadpmtuipchecksum     int    `json:"icmptotbadpmtuipchecksum,omitempty"`
	Icmptotbignextmtu            int    `json:"icmptotbignextmtu,omitempty"`
	Icmptotdstiplookup           int    `json:"icmptotdstiplookup,omitempty"`
	Icmptotinvalidbodylen        int    `json:"icmptotinvalidbodylen,omitempty"`
	Icmptotinvalidnextmtuval     int    `json:"icmptotinvalidnextmtuval,omitempty"`
	Icmptotinvalidprotocol       int    `json:"icmptotinvalidprotocol,omitempty"`
	Icmptotinvalidtcpseqno       int    `json:"icmptotinvalidtcpseqno,omitempty"`
	Icmptotneedfragrx            int    `json:"icmptotneedfragrx,omitempty"`
	Icmptotnonfirstipfrag        int    `json:"icmptotnonfirstipfrag,omitempty"`
	Icmptotnotcpconn             int    `json:"icmptotnotcpconn,omitempty"`
	Icmptotnoudpconn             int    `json:"icmptotnoudpconn,omitempty"`
	Icmptotpktsdropped           int    `json:"icmptotpktsdropped,omitempty"`
	Icmptotpmtudiscoverydisabled int    `json:"icmptotpmtudiscoverydisabled,omitempty"`
	Icmptotpmtunolink            int    `json:"icmptotpmtunolink,omitempty"`
	Icmptotportunreachablerx     int    `json:"icmptotportunreachablerx,omitempty"`
	Icmptotportunreachabletx     int    `json:"icmptotportunreachabletx,omitempty"`
	Icmptotrxbytes               int    `json:"icmptotrxbytes,omitempty"`
	Icmptotrxecho                int    `json:"icmptotrxecho,omitempty"`
	Icmptotrxechoreply           int    `json:"icmptotrxechoreply,omitempty"`
	Icmptotrxpkts                int    `json:"icmptotrxpkts,omitempty"`
	Icmptotthresholdexceeds      int    `json:"icmptotthresholdexceeds,omitempty"`
	Icmptottxbytes               int    `json:"icmptottxbytes,omitempty"`
	Icmptottxechoreply           int    `json:"icmptottxechoreply,omitempty"`
	Icmptottxpkts                int    `json:"icmptottxpkts,omitempty"`
	Icmptxbytesrate              int    `json:"icmptxbytesrate,omitempty"`
	Icmptxechoreplyrate          int    `json:"icmptxechoreplyrate,omitempty"`
	Icmptxpktsrate               int    `json:"icmptxpktsrate,omitempty"`
}
