package gslb

type Gslbdomainstats struct {
	Clearstats      string `json:"clearstats,omitempty"`
	Dnsqueriesrate  int    `json:"dnsqueriesrate,omitempty"`
	Dnsrecordtype   string `json:"dnsrecordtype,omitempty"`
	Dnstotalqueries int    `json:"dnstotalqueries,omitempty"`
	Gslbdnsrectype  string `json:"gslbdnsrectype,omitempty"`
	Name            string `json:"name,omitempty"`
}
