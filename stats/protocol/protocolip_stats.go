package protocol

type Protocolipstats struct {
	Clearstats               string `json:"clearstats,omitempty"`
	Iproutedmbitsrate        int    `json:"iproutedmbitsrate,omitempty"`
	Iproutedpktsrate         int    `json:"iproutedpktsrate,omitempty"`
	Iprxbytesrate            int    `json:"iprxbytesrate,omitempty"`
	Iprxmbitsrate            int    `json:"iprxmbitsrate,omitempty"`
	Iprxpktsrate             int    `json:"iprxpktsrate,omitempty"`
	Iptotaddrlookup          int    `json:"iptotaddrlookup,omitempty"`
	Iptotaddrlookupfail      int    `json:"iptotaddrlookupfail,omitempty"`
	Iptotbadchecksums        int    `json:"iptotbadchecksums,omitempty"`
	Iptotbadlens             int    `json:"iptotbadlens,omitempty"`
	Iptotbadmacaddrs         int    `json:"iptotbadmacaddrs,omitempty"`
	Iptotbadtransport        int    `json:"iptotbadtransport,omitempty"`
	Iptotdupfragments        int    `json:"iptotdupfragments,omitempty"`
	Iptotfixheaderfail       int    `json:"iptotfixheaderfail,omitempty"`
	Iptotfragments           int    `json:"iptotfragments,omitempty"`
	Iptotfragpktsgen         int    `json:"iptotfragpktsgen,omitempty"`
	Iptotinvalidheadersz     int    `json:"iptotinvalidheadersz,omitempty"`
	Iptotinvalidpacketsize   int    `json:"iptotinvalidpacketsize,omitempty"`
	Iptotlandattacks         int    `json:"iptotlandattacks,omitempty"`
	Iptotmaxclients          int    `json:"iptotmaxclients,omitempty"`
	Iptotoutoforderfrag      int    `json:"iptotoutoforderfrag,omitempty"`
	Iptotreassemblyattempt   int    `json:"iptotreassemblyattempt,omitempty"`
	Iptotroutedmbits         int    `json:"iptotroutedmbits,omitempty"`
	Iptotroutedpkts          int    `json:"iptotroutedpkts,omitempty"`
	Iptotrxbytes             int    `json:"iptotrxbytes,omitempty"`
	Iptotrxmbits             int    `json:"iptotrxmbits,omitempty"`
	Iptotrxpkts              int    `json:"iptotrxpkts,omitempty"`
	Iptotsuccreassembly      int    `json:"iptotsuccreassembly,omitempty"`
	Iptottcpfragmentsfwd     int    `json:"iptottcpfragmentsfwd,omitempty"`
	Iptottoobig              int    `json:"iptottoobig,omitempty"`
	Iptottruncatedpackets    int    `json:"iptottruncatedpackets,omitempty"`
	Iptotttlexpired          int    `json:"iptotttlexpired,omitempty"`
	Iptottxbytes             int    `json:"iptottxbytes,omitempty"`
	Iptottxmbits             int    `json:"iptottxmbits,omitempty"`
	Iptottxpkts              int    `json:"iptottxpkts,omitempty"`
	Iptotudpfragmentsfwd     int    `json:"iptotudpfragmentsfwd,omitempty"`
	Iptotunknowndstrcvd      int    `json:"iptotunknowndstrcvd,omitempty"`
	Iptotunknownsvcs         int    `json:"iptotunknownsvcs,omitempty"`
	Iptotunsuccreassembly    int    `json:"iptotunsuccreassembly,omitempty"`
	Iptotvipdown             int    `json:"iptotvipdown,omitempty"`
	Iptotzerofragmentlen     int    `json:"iptotzerofragmentlen,omitempty"`
	Iptotzeronexthop         int    `json:"iptotzeronexthop,omitempty"`
	Iptxbytesrate            int    `json:"iptxbytesrate,omitempty"`
	Iptxmbitsrate            int    `json:"iptxmbitsrate,omitempty"`
	Iptxpktsrate             int    `json:"iptxpktsrate,omitempty"`
	Noniptottruncatedpackets int    `json:"noniptottruncatedpackets,omitempty"`
}
