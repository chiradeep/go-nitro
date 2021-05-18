package snmp

type Snmpstats struct {
	Clearstats                   string `json:"clearstats,omitempty"`
	Snmpdecryptionerrors         int    `json:"snmpdecryptionerrors,omitempty"`
	Snmpgetbulkreqsrate          int    `json:"snmpgetbulkreqsrate,omitempty"`
	Snmpgetnextreqsrate          int    `json:"snmpgetnextreqsrate,omitempty"`
	Snmpgetreqsrate              int    `json:"snmpgetreqsrate,omitempty"`
	Snmpnotintimewindow          int    `json:"snmpnotintimewindow,omitempty"`
	Snmpresponsesrate            int    `json:"snmpresponsesrate,omitempty"`
	Snmprxpktsrate               int    `json:"snmprxpktsrate,omitempty"`
	Snmptotbadcommname           int    `json:"snmptotbadcommname,omitempty"`
	Snmptotbadcommuse            int    `json:"snmptotbadcommuse,omitempty"`
	Snmptotbadversions           int    `json:"snmptotbadversions,omitempty"`
	Snmptoterrreqdropped         int    `json:"snmptoterrreqdropped,omitempty"`
	Snmptotgetbulkreqs           int    `json:"snmptotgetbulkreqs,omitempty"`
	Snmptotgetnextreqs           int    `json:"snmptotgetnextreqs,omitempty"`
	Snmptotgetreqs               int    `json:"snmptotgetreqs,omitempty"`
	Snmptotparseerrs             int    `json:"snmptotparseerrs,omitempty"`
	Snmptotresponses             int    `json:"snmptotresponses,omitempty"`
	Snmptotrxpkts                int    `json:"snmptotrxpkts,omitempty"`
	Snmptottraps                 int    `json:"snmptottraps,omitempty"`
	Snmptottxpkts                int    `json:"snmptottxpkts,omitempty"`
	Snmptxpktsrate               int    `json:"snmptxpktsrate,omitempty"`
	Snmpunknownengineids         int    `json:"snmpunknownengineids,omitempty"`
	Snmpunknownusername          int    `json:"snmpunknownusername,omitempty"`
	Snmpunsupportedsecuritylevel int    `json:"snmpunsupportedsecuritylevel,omitempty"`
	Snmpwrongdigests             int    `json:"snmpwrongdigests,omitempty"`
}
